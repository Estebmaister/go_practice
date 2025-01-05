package concurrency

/*
 * The `sema` struct is defining a semaphore. A semaphore is a synchronization
 * primitive that limits the number of concurrent operations or access to a
 * shared resource. In this case, the `sema` struct has two fields: `ch` which
 * is a channel of `struct{}`, and `size` which represents the maximum number
 * of semaphores that can be acquired. The `Enter()` method is used to acquire
 * a semaphore by adding a value to the `ch` channel, and the `Exit()` method
 * is used to release a semaphore by removing a value from the `ch` channel.
 */
type sema struct {
	ch   chan struct{}
	size int
}

const (
	/*
	 * `DefaultSemaCounter` is a constant that represents the default size of
	 * the semaphore. If the size of the semaphore is not specified when
	 * creating a new semaphore using the `NewSema` function, `DefaultSemaCounter`
	 * will be used as the size.
	 */
	DefaultSemaCounter = 10
)

// The NewSema function creates a new semaphore with the specified size.
func NewSema(size int) *sema {
	if size == 0 {
		size = DefaultSemaCounter
	}
	return &sema{
		ch:   make(chan struct{}, size),
		size: size,
	}
}

/*
 *The `Enter()` method of the `sema` struct is used to acquire a semaphore.
 * It adds a struct{} value to the `ch` channel, effectively blocking if
 * the channel is already full (i.e., if the number of acquired semaphores
 * equals the size of the channel). Once a semaphore is acquired, it can be
 * used to control access to a shared resource or limit the number of
 *concurrent operations.
 */
func (s *sema) Enter() {
	s.ch <- struct{}{}
}

/*
 * The `Exit()` method of the `sema` struct is used to release a semaphore.
 * It removes a struct{} value from the `ch` channel, effectively allowing
 * another goroutine to acquire the semaphore. This is typically used when
 * a goroutine is finished using a shared resource or has completed its
 * concurrent operation.
 */
func (s *sema) Exit() {
	<-s.ch
}
