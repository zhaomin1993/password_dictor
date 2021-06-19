package password_dictor

import (
	"context"
	"fmt"
)

type Dictor struct {
	ctx    context.Context
	cancel context.CancelFunc
	dict   [][]rune
	out    chan string
}

func NewDictor(ctx context.Context, dict [][]rune) *Dictor {
	newCtx, cancel := context.WithCancel(ctx)
	return &Dictor{
		ctx:    newCtx,
		cancel: cancel,
		dict:   dict,
		out:    make(chan string),
	}
}

func (obj *Dictor) Run() (<-chan string, error) {
	//相当于最内层循环执行的次数
	var totalTimes = 1
	for i := range obj.dict {
		totalTimes *= len(obj.dict[i])
	}
	if totalTimes == 0 || len(obj.dict) == 0 {
		return nil, fmt.Errorf("bad dict")
	}
	go func() {
		n := 0
	LOOP:
		for n < totalTimes {
			div := n
			temRes := make([]rune, 0, len(obj.dict))
			for i := 0; i < len(obj.dict); i++ {
				var mod int
				//余数就是参与计算的元素的下标,商用于计算下一个列表参与元素的下标.
				div, mod = divMod(div, len(obj.dict[i]))
				temRes = append(temRes, obj.dict[i][mod])
			}
			select {
			case obj.out <- string(temRes):
			case <-obj.ctx.Done():
				break LOOP
			}
			n += 1
		}
		close(obj.out)
	}()
	return obj.out, nil
}

func (obj *Dictor) Stop() {
	select {
	case <-obj.ctx.Done():
	default:
		obj.cancel()
	}
}

func divMod(a, b int) (c, d int) {
	c = a / b
	d = a % b
	return
}
