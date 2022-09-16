package progress

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	p := New(Content("测试1: "))

	for {
		time.Sleep(100 * time.Millisecond)
		p.Add(int64(2))
		p.Print()
		if p.IsFinish() {
			return
		}
	}
}
