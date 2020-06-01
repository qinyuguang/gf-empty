package requestid

type ctxKey int

// RequestIDKey is the request context key used to store the request ID.
const RequestIDKey ctxKey = iota
