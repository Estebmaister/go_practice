package interviews

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Movies(t *testing.T) {
	movieA := NewMovie(1, 1.2)
	movieB := NewMovie(2, 2.4)
	movieC := NewMovie(3, 3.6)
	movieD := NewMovie(4, 4.8)

	movieA.AddSimilarMovie(nil)
	assert.Nil(t, movieA.GetSimilarMovies())

	movieA.AddSimilarMovie(movieB)
	movieA.AddSimilarMovie(movieC)
	movieB.AddSimilarMovie(movieD)
	movieC.AddSimilarMovie(movieD)

	assert.NotNil(t, movieA)
	assert.NotNil(t, movieB)
	assert.NotNil(t, movieC)
	assert.NotNil(t, movieD)
	assert.Equal(t, movieA.Rating, float32(1.2))
	assert.Equal(t, movieB.Rating, float32(2.4))
	assert.Equal(t, movieC.Rating, float32(3.6))
	assert.Equal(t, movieD.Rating, float32(4.8))
	for _, m := range movieA.GetSimilarMovies() {
		if m.ID != movieB.ID && m.ID != movieC.ID {
			assert.Fail(t, "Movie A should have only movie B and C as similar movies")
		}
	}
	for _, m := range movieB.GetSimilarMovies() {
		if m.ID != movieA.ID && m.ID != movieD.ID {
			assert.Fail(t, "Movie B should have only movie A and D as similar movies")
		}
	}
	for _, m := range movieC.GetSimilarMovies() {
		if m.ID != movieA.ID && m.ID != movieD.ID {
			assert.Fail(t, "Movie C should have only movie A and D as similar movies")
		}
	}
	for _, m := range movieD.GetSimilarMovies() {
		if m.ID != movieB.ID && m.ID != movieC.ID {
			assert.Fail(t, "Movie D should have only movie B and C as similar movies")
		}
	}

	// similar movies are not ordered because map is not ordered,
	// so the following assertion can fail sometimes:
	// assert.Equal(t, movieA.GetSimilarMovies(), []*Movie{movieB, movieC})
}

func Test_GetMovieRecommendations(t *testing.T) {

	// Example usage
	movieA := NewMovie(1, 1.2)
	movieB := NewMovie(2, 2.4)
	movieC := NewMovie(3, 3.6)
	movieD := NewMovie(4, 4.8)

	movieA.AddSimilarMovie(movieB)
	movieA.AddSimilarMovie(movieC)
	movieB.AddSimilarMovie(movieD)
	movieC.AddSimilarMovie(movieD)

	// Get top 2 recommendations
	recommendations := GetMovieRecommendations(movieA, 2)
	assert.NotNil(t, recommendations)
	assert.Equal(t, len(recommendations), 2)

	for _, m := range recommendations {
		assert.NotNil(t, m)
	}

	assert.Equal(t, recommendations[0].Rating, float32(4.8))
	assert.Equal(t, recommendations[1].Rating, float32(3.6))

	// Get top 4 recommendations
	recommendations = GetMovieRecommendations(movieA, 4)
	assert.NotNil(t, recommendations)
	assert.Equal(t, len(recommendations), 4)

	for _, m := range recommendations {
		assert.NotNil(t, m)
	}

	assert.Equal(t, recommendations[0].Rating, float32(4.8))
	assert.Equal(t, recommendations[1].Rating, float32(3.6))
	assert.Equal(t, recommendations[2].Rating, float32(2.4))
	assert.Equal(t, recommendations[3].Rating, float32(1.2))

	// Get top 1 recommendation
	recommendations = GetMovieRecommendations(movieA, 1)
	assert.NotNil(t, recommendations)
	assert.Equal(t, len(recommendations), 1)

	for _, m := range recommendations {
		assert.NotNil(t, m)
	}

	assert.Equal(t, recommendations[0].Rating, float32(4.8))

	// Get top 0 recommendations
	recommendations = GetMovieRecommendations(movieA, 0)
	assert.NotNil(t, recommendations)
	assert.Equal(t, len(recommendations), 0)

	// Get top 5 recommendations
	recommendations = GetMovieRecommendations(movieA, 5)
	assert.NotNil(t, recommendations)
	assert.Equal(t, len(recommendations), 4)

	for _, m := range recommendations {
		assert.NotNil(t, m)
	}

	assert.Equal(t, recommendations[0].Rating, float32(4.8))
	assert.Equal(t, recommendations[1].Rating, float32(3.6))
	assert.Equal(t, recommendations[2].Rating, float32(2.4))
	assert.Equal(t, recommendations[3].Rating, float32(1.2))
}

/* A class Movie like below
*     public class Movie
*     {
*          private int movieId;
*          private float rating;
*          private List<Movie> similarMovies; // Similarity is bidirectional
*          //private HashSet<Movie> similarMovies; // Similarity is bidirectional
*          ...
*     }
*
* Implement a function to return top rated movies in the network of movies
* reachable (directly or indirectly) from the current movie:
*
* //public HashSet<Movie> getMovieRecommendations(Movie movie, int numTopRatedSimilarMovies)
* public IList<Movie> getMovieRecommendations(Movie movie, int numTopRatedSimilarMovies)
*
* eg:            A(Rating 1.2)
*               /   \
*            B(2.4)  C(3.6)
*              \     /
*               D(4.8)
* In the above example edges represent similarity and the number is rating.
* getMovieRecommendations(A,2)should return C and D (sorting order doesn't matter so it can also return D and C)
* getMovieRecommendations(A,4) should return A, B, C, D (it can also return these in any order eg: B,C,D,A)
* getMovieRecommendations(A,1) should return D. Note distance from A to D doesn't matter,
*                            return the highest  rated.
*
*     @param movie
*     @param numTopRatedSimilarMovies
*                      number of movies we want to return
*     @return List of top rated similar movies
*     //@return HashSet of top rated similar movies
 */
