package middleware

import "net/http"

var sem chan struct{} = nil

func ConcurrentLimiter(maxConcurrentRequests int, next http.Handler) http.Handler {
	if sem == nil {
		sem = make(chan struct{}, maxConcurrentRequests)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sem <- struct{}{} // Acquire a semaphore slot
		defer func() {
			<-sem // Release the semaphore slot
		}()

		next.ServeHTTP(w, r)
	})
}
