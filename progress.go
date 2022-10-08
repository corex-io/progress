package progress

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/term"
)

// Progress 进图结构体
type Progress struct {
	ctx     context.Context
	cancel  context.CancelFunc
	start   time.Time
	options Options
	current int64
	mutex   sync.Mutex
	once    sync.Once
	stop    bool
}

// New 初始化方法
func New(ctx context.Context, opts ...Option) *Progress {
	ctx, cancel := context.WithCancel(ctx)
	options := newOptions(opts...)
	return &Progress{
		ctx:     ctx,
		cancel:  cancel,
		start:   time.Now(),
		options: options,
		current: options.Start,
	}
}

func (progress *Progress) Start() {
	go func() {
		timer := time.NewTimer(0 * time.Second)
		defer timer.Stop()
		var exit bool
		for {
			if progress.Print(); progress.IsFinish() || exit {
				progress.Finish()
				return
			}
			select {
			case <-progress.ctx.Done():
				exit = true
			case <-timer.C:
				timer.Reset(progress.options.Refresh)
			}
		}
	}()
}

func (progress *Progress) Wait() {
	select {
	case <-progress.ctx.Done():
		return
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

func (progress *Progress) display() (string, error) {
	display := []string{progress.options.Content, "[", "", "]", "0.00%", "0/0", "(0s)"}
	// 0: content  1: 开始符号  2: 进度显示  3: 结束符号  4: 进度  5: 当前/总量  6: 时间
	percent := float64(progress.current) / float64(progress.options.End)
	display[4] = fmt.Sprintf("%7.3f%%", percent*100)
	display[5] = fmt.Sprintf("%10d/%-10d", progress.current, progress.options.End)
	display[6] = fmt.Sprintf("%11s", HumanTime(time.Since(progress.start)))

	width, _, err := term.GetSize(int(os.Stdout.Fd()))

	width -= len(display[0]) + len(display[1]) + len(display[3]) + len(display[4]) + len(display[5]) + len(display[6])
	width = width * int(progress.options.Div*100) / 100
	if err != nil || width <= 0 {
		return strings.Join(display, ""), err
	}

	display[2] = fmt.Sprintf("%-"+fmt.Sprintf("%d", width)+"s", strings.Repeat(progress.options.Graph, int(percent*float64(width))))
	return strings.Join(display, ""), nil
}

func (progress *Progress) Print() {
	progress.mutex.Lock()
	defer progress.mutex.Unlock()
	if progress.stop {
		return
	}
	display, err := progress.display()
	if err == nil && progress.options.Tty {
		_, _ = fmt.Fprintf(progress.options.Writer, "\r%s", display)
	} else {
		_, _ = fmt.Fprintf(progress.options.Writer, "%s\n", display)
	}
	if progress.IsFinish() {
		progress.Finish()
	}
}

// IsFinish return progress is finish or not.
func (progress *Progress) IsFinish() bool {
	return progress.current >= progress.options.End || progress.stop
}

func (progress *Progress) Finish() {
	progress.once.Do(func() {
		defer progress.cancel()
		progress.stop = true
		_, _ = fmt.Fprintf(progress.options.Writer, "\n")
	})
}
