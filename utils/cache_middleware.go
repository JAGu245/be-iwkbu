package utils

import (
	"net/http"
)

// ResponseRecorder captures the output written by the next handler.
type ResponseRecorder struct {
	http.ResponseWriter
	Body []byte
	Status int
}

func (w *ResponseRecorder) Write(b []byte) (int, error) {
	w.Body = append(w.Body, b...)
	return w.ResponseWriter.Write(b)
}

func (w *ResponseRecorder) WriteHeader(statusCode int) {
	w.Status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// CacheMiddleware intercepts the request to serve cached data if available, 
// or records the output of the next handler to populate the cache.
func CacheMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Only cache GET requests or standard requests
		if r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}

		key := r.URL.Path

		// Attempt to fetch from cache
		cachedData := GetCache(key)
		if cachedData != nil {
			// Serve from cache
			w.Header().Set("Content-Type", "application/json")
			w.Write(cachedData)
			return
		}

		// Not in cache, we need to capture the output
		recorder := &ResponseRecorder{
			ResponseWriter: w,
			Status:         http.StatusOK, // default status
		}

		next.ServeHTTP(recorder, r)

		// Only cache successful requests
		if recorder.Status >= 200 && recorder.Status < 300 && len(recorder.Body) > 0 {
			SetCache(key, recorder.Body)
		}
	}
}
