package router

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

type muxEntry struct {
	h       http.Handler
	pattern string
}

type Router struct {
	mu sync.RWMutex
	m  map[string]muxEntry
	es []muxEntry
	http.ServeMux
}

// Handle registers the handler for the given pattern.
// If a handler already exists for pattern, Handle panics.
func (mux *Router) Handle(pattern string, handler http.Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if pattern == "" || pattern == "//" {
		panic("router: invalid pattern")
	}
	if handler == nil {
		panic("router: nil handler")
	}
	if _, exist := mux.m[pattern]; exist {
		panic("router: multiple registrations for " + pattern)
	}
	if pattern[len(pattern)-1] != '/' || pattern[0] != '/' {
		panic("router: first and end of every route has to contain `/`")
	}

	pattern = fmt.Sprintf("^%s$", pattern)
	if mux.m == nil {
		mux.m = make(map[string]muxEntry)
	}
	e := muxEntry{h: handler, pattern: pattern}
	mux.m[pattern] = e
	mux.es = append(mux.es, e)
}

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (mux *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}

func (mux *Router) Handler(r *http.Request) (h http.Handler, pattern string) {
	path := r.URL.Path
	if path != "/" {
		if path[len(path)-1] != '/' {
			path = fmt.Sprintf("%s/", path)
		}
	}

	return mux.handler(path)
}

// handler is the main implementation of Handler.
// The path is known to be in canonical form, except for CONNECT methods.
func (mux *Router) handler(path string) (h http.Handler, pattern string) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()

	if h == nil {
		for i := 0; i < len(mux.es); i++ {
			element := mux.es[i]
			r, _ := regexp.Compile(element.pattern)
			if r.MatchString(path) {
				h = element.h
				pattern = element.pattern
				break
			}
		}
	}
	if h == nil {
		h, pattern = http.NotFoundHandler(), ""
	}
	return
}
