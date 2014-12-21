package app

import (
	"net/http"
)

// Middleware type
type Middleware func(*C, http.Handler) http.Handler

// Controller creator type
type Ctrl func(*C) http.Handler

// Request context
type C struct {
	CurrentUser *User
}

// Create a context, create the controller using the context, and
// create a http.Handler that combines all the middleware.
func Mid(h Ctrl, mid ...Middleware) http.Handler {
	ctx := new(C)
	f := h(ctx)
	for i := len(mid) - 1; i >= 0; i-- {
		f = mid[i](ctx, f)
	}
	return f
}
