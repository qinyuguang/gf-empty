package requestid

const defaultTTL = 3600

type (
	// OptionFunc uses a constructor pattern to customize middleware.
	OptionFunc func(*options) *options

	options struct {
		scope string
		ttl   int
	}
)

func newOptions(opts ...OptionFunc) *options {
	o := &options{
		ttl: defaultTTL,
	}

	for i := range opts {
		o = opts[i](o)
	}
	return o
}

// WithScope set scope of requestid cachekey.
func WithScope(scope string) OptionFunc {
	return func(o *options) *options {
		o.scope = scope
		return o
	}
}

// WithTTL set expired time of requestid cachekey.
func WithTTL(ttl int) OptionFunc {
	return func(o *options) *options {
		o.ttl = ttl
		return o
	}
}
