package progress

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
	"sync"
	"time"
)

// Progress 进图结构体
type Progress struct {
	now     time.Time
	options Options
	current int64 // 开始
	mutex   sync.Mutex
	once    sync.Once
	stop    bool
}

// New 初始化方法
func New(opts ...Option) *Progress {
	options := newOptions(opts...)
	return &Progress{
		now:     time.Now(),
		options: options,
		current: options.Start,
	}
}

// Add 增加进度
func (progress *Progress) Add(i int64) *Progress {
	progress.mutex.Lock()
	defer progress.mutex.Unlock()
	if progress.current += i; progress.current > progress.options.End {
		progress.current = progress.options.End
	}
	return progress
}

func (progress *Progress) Print() {
	progress.mutex.Lock()
	defer progress.mutex.Unlock()
	if progress.stop {
		return
	}
	display := []string{progress.options.Content, "[", "", "]", "0.00%", "0/0", "(0s)"}
	// 0: content  1: 开始符号  2: 进度显示  3: 结束符号  4: 进度  5: 当前/总量  6: 时间

	width := 80
	width, _, _ = term.GetSize(int(os.Stdout.Fd()))

	percent := float64(progress.current) / float64(progress.options.End)

	display[4] = fmt.Sprintf("%7.3f%%", percent*100)
	display[5] = fmt.Sprintf("%10d/%-10d", progress.current, progress.options.End)
	display[6] = fmt.Sprintf("%11s", Human(time.Since(progress.now)))
	width -= len(display[0]) + len(display[1]) + len(display[3]) + len(display[4]) + len(display[5]) + len(display[6])
	width = width * int(progress.options.Div*100) / 100
	display[2] = fmt.Sprintf("%-"+fmt.Sprintf("%d", width)+"s", strings.Repeat(progress.options.Graph, int(percent*float64(width))))
	_, _ = fmt.Fprintf(os.Stderr, "\r%s", strings.Join(display, ""))
	if progress.IsFinish() {
		progress.Finish()
	}
}

// IsFinish return progress is finish or not.
func (progress *Progress) IsFinish() bool {
	return progress.current >= progress.options.End
}

func (progress *Progress) Finish() {
	progress.once.Do(func() {
		progress.stop = true
		_, _ = fmt.Fprintf(os.Stderr, "\n")
	})
}

func Human(d time.Duration) string {
	ss := []struct {
		dis string
		time.Duration
	}{
		{"y", time.Hour * 24 * 365},
		{"d", time.Hour * 24},
		{"h", time.Hour},
		{"m", time.Minute},
		{"s", time.Second},
		{"ms", time.Millisecond},
	}
	var b strings.Builder
	i := 0
	for _, post := range ss {
		if d >= post.Duration {
			t := d / post.Duration
			fmt.Fprintf(&b, "%d%s", t, post.dis)
			d -= t * post.Duration
			i += 1
		}
		if i == 3 {
			return b.String()
		}
	}
	return b.String()
}
