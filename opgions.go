package progress

import (
	"io"
	"os"
	"time"
)

type Options struct {
	Start   int64
	End     int64
	Div     float64
	Graph   string
	Content string
	Refresh time.Duration
	Writer  io.Writer
	Tty     bool
}

// Option func
type Option func(*Options)

// NewOptions new request
func newOptions(opts ...Option) Options {
	opt := Options{
		Start:   0,
		End:     100,
		Div:     1.0,
		Graph:   "█",                    // 设置进度条的样式
		Refresh: 500 * time.Millisecond, // 0.5s
		Writer:  os.Stdout,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func Tty(tty bool) Option {
	return func(o *Options) {
		o.Tty = tty
	}
}

// Writer to
func Writer(w io.Writer) Option {
	return func(o *Options) {
		o.Writer = w
	}
}

func Div(div float64) Option {
	return func(o *Options) {
		o.Div = div
	}
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

func Refresh(refresh time.Duration) Option {
	return func(o *Options) {
		o.Refresh = refresh
	}
}
