package progress

import (
	"fmt"
	"strings"
	"time"
)

var tt = []struct {
	unit string
	time.Duration
}{
	{"y", time.Hour * 24 * 365},
	{"d", time.Hour * 24},
	{"h", time.Hour},
	{"m", time.Minute},
	{"s", time.Second},
	{"ms", time.Millisecond},
}

func HumanTime(d time.Duration) string {
	var b strings.Builder
	i, t := 0, time.Duration(0)
	for _, post := range tt {
		if d >= post.Duration {
			i, t = i+1, d/post.Duration
			if _, _ = fmt.Fprintf(&b, "%d%s", t, post.unit); i == 3 {
				return b.String()
			}
			d = d % post.Duration
		}
	}
	return b.String()
}
