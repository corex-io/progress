package progress

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	p := New(context.Background(), Content("测试1111111111: "), Div(1), Tty(true))

	for {
		time.Sleep(1000 * time.Millisecond)
		p.Add(int64(20)).Print()
		if p.IsFinish() {
			break
		}
	}

	p2 := New(context.Background(), Content("Hello:"), Refresh(100*time.Millisecond), Graph("#"))
	p2.Start()
	for {
		time.Sleep(1000 * time.Millisecond)
		p2.Add(int64(20))
		if p2.IsFinish() {
			break
		}
	}

}

func TestLen(t *testing.T) {
	c := fmt.Sprintf("%7.3f", 82.3)
	fmt.Println(c, len(c))
	d := fmt.Sprintf("%10d/%d", 1, 2)
	fmt.Println(d, len(d))
	e := time.Since(time.Now().Add(-4 * time.Second))
	ee := fmt.Sprintf("%6s", e)
	fmt.Println(ee, len(ee))
	day := time.Minute * 60 * 24
	year := 365 * day
	fmt.Println(HumanTime(12*year + 365*day + 5*day + 5*time.Hour + 1998*time.Millisecond))

}

func TestChan(t *testing.T) {
	c := make([]string, 0, 5)
	for i, v := range []string{"a", "b", "c"} {
		if i%2 == 0 {
			c = append(c, v)
		}
	}
	fmt.Println(c, len(c), cap(c))
}
