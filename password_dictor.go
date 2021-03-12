package password_dictor

import (
  "context"
  "fmt"
)

func Dictor(ctx context.Context, li [][]rune) (ch chan string, err error) {
	//相当于最内层循环执行的次数
	var totalTimes = 1
	for i := range li {
		totalTimes *= len(li[i])
	}
	if totalTimes == 0 || len(li) == 0 {
		return nil, fmt.Errorf("bad arg")
	}
	ch = make(chan string)
	go func() {
		n := 0
	LOOP:
		for n < totalTimes {
			tmp := n
			temRes := make([]rune, 0, len(li))
			for i := 0; i < len(li); i++ {
				var cur int
				//余数就是参与计算的元素的下标,商用于计算下一个列表参与元素的下标.
				tmp, cur = divmod(tmp, len(li[i]))
				temRes = append(temRes, li[i][cur])
			}
			select {
			case ch <- string(temRes):
			case <-ctx.Done():
				break LOOP
			}
			n += 1
		}
		close(ch)
	}()
	return
}

func divmod(a, b int) (int, int) {
	c := a / b
	d := a % b
	return c, d
}
