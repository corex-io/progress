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
	now      time.Time
	options  Options
	current  int64 // 开始
	isFinish bool
	mutex    sync.Mutex
	once     sync.Once
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
func (progress *Progress) Add(i int64) {
	progress.mutex.Lock()
	defer progress.mutex.Unlock()
	progress.current += i
}

func (progress *Progress) Print() {
	progress.mutex.Lock()
	defer progress.mutex.Unlock()
	if progress.isFinish {
		return
	}
	width := 80
	width, _, _ = term.GetSize(int(os.Stdout.Fd()))
	width = width*1/3 - 30 - len(progress.options.Content)

	percent := float64(progress.current) / float64(progress.options.End)

	rate := strings.Repeat(progress.options.Graph, int(percent*100*float64(width)/float64(progress.options.End)))
	_, _ = fmt.Fprintf(os.Stderr, "\r%s[%-"+fmt.Sprintf("%d", width)+"s] %.2f%% %8d/%d (%.2fs)", progress.options.Content, rate, percent*100, progress.current, progress.options.End, time.Since(progress.now).Seconds())
	// %-50s 左对齐, 占50个字符位置, 打印string
	// %3d   右对齐, 占3个字符位置 打印int
	if progress.options.End <= progress.current {
		progress.Finish()
	}
}

// IsFinish return progress is finish or not.
func (progress *Progress) IsFinish() bool {
	return progress.isFinish
}

func (progress *Progress) Finish() {
	progress.once.Do(func() {
		progress.isFinish = true
		_, _ = fmt.Fprintf(os.Stderr, "\n")
	})
}
