package progress

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	p := New(Content("测试1111111111: "))

	for {
		time.Sleep(1000 * time.Millisecond)
		p.Add(int64(2))
		p.Print()
		if p.IsFinish() {
			return
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
	fmt.Println(Human(time.Duration(12*year + 365*day + 5*day + 5*time.Hour + 1998*time.Millisecond)))

}
