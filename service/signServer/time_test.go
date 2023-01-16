package signServer

import (
	"fmt"
	"testing"
	"time"
)

func TestDayNum(t *testing.T) {
	t1 := time.Now().Unix()
	//diff := float64(1673332863)
	d := int64(t1 / OneDayTime)
	tm := t1 % OneDayTime

	if tm > 0 {
		d = d + 1
	}

	fmt.Println(d - 19300)

}
