package interviews

import (
	"container/heap"
)

var _ heap.Interface = (*MovieHeap)(nil)

// Movie represents a movie with its ID, rating, and similar movies
type Movie struct {
	ID            int
	Rating        float32
	similarMovies map[int]*Movie // Using map for O(1) lookups
}

// Helper function to create a new movie
func NewMovie(id int, rating float32) *Movie {
	return &Movie{
		ID:            id,
		Rating:        rating,
		similarMovies: make(map[int]*Movie),
	}
}

// AddSimilarMovie adds a bidirectional similarity between two movies
func (m *Movie) AddSimilarMovie(movie *Movie) {
	// Check for nil or self-similarity
	if m == nil || movie == nil || m == movie {
		return
	}

	// Add to current movie's similar movies if not already present
	if _, exists := m.similarMovies[movie.ID]; !exists {
		m.similarMovies[movie.ID] = movie
	}

	// Add the reverse relationship if not already present
	if _, exists := movie.similarMovies[m.ID]; !exists {
		movie.similarMovies[m.ID] = m
	}
}

// GetSimilarMovies returns a slice of similar movies
func (m *Movie) GetSimilarMovies() []*Movie {
	if m == nil || len(m.similarMovies) == 0 {
		return nil
	}

	movies := make([]*Movie, len(m.similarMovies))
	i := 0
	for _, movie := range m.similarMovies {
		movies[i] = movie
		i++
	}
	return movies
}

// MovieHeap is a min-heap of movies based on their rating
type MovieHeap []*Movie

func (h MovieHeap) Len() int           { return len(h) }
func (h MovieHeap) Less(i, j int) bool { return h[i].Rating < h[j].Rating }
func (h MovieHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MovieHeap) Push(x any) { *h = append(*h, x.(*Movie)) }
func (h *MovieHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// GetMovieRecommendations returns the top rated movies in the network
// using BFS to traverse the movie graph and a min-heap to track top rated movies
func GetMovieRecommendations(movie *Movie, numTopRated int) []*Movie {
	if movie == nil || numTopRated <= 0 {
		return []*Movie{}
	}

	// Use a min-heap to keep track of top rated movies
	h := &MovieHeap{}
	heap.Init(h)
	visited := make(map[int]bool)
	queue := []*Movie{movie} // Initialize FIFO queue with the starting movie (BFS)
	visited[movie.ID] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Add current movie to heap
		heap.Push(h, current)
		if h.Len() > numTopRated {
			heap.Pop(h) // Remove the movie with smallest rating
		}

		// Add unvisited similar movies to queue
		for _, similar := range current.similarMovies {
			if !visited[similar.ID] {
				visited[similar.ID] = true
				queue = append(queue, similar) // Add to queue for BFS
			}
		}
	}

	// Convert heap to result slice and reverse to get highest rated first
	result := make([]*Movie, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(*Movie)
	}

	return result
}
