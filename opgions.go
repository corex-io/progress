package progress

type Options struct {
	Start   int64
	End     int64
	Graph   string
	Content string
}

// Option func
type Option func(*Options)

// NewOptions new request
func newOptions(opts ...Option) Options {
	opt := Options{
		Start: 0,
		End:   100,
		Graph: "█", // 这里设置进度条的样式
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Start set method
func Start(start int64) Option {
	return func(o *Options) {
		o.Start = start
	}
}

// End set method
func End(end int64) Option {
	return func(o *Options) {
		o.End = end
	}
}

// Graph set method
func Graph(graph string) Option {
	return func(o *Options) {
		o.Graph = graph
	}
}

func Content(content string) Option {
	return func(o *Options) {
		o.Content = content
	}
}
