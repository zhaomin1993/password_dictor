package password_dictor

import (
	"context"
	"testing"
)

func TestDictor(t *testing.T) {
	var ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	var li1 = []rune("abcdefghijklmnopqrstuvwxyz")
	var li2 = []rune("abcdefghijklmnopqrstuvwxyz")
	var li3 = []rune("12")
	var li4 = []rune("0189")
	var li5 = []rune("0123456789")
	var li6 = []rune("0123456789")
	ch, err := NewDictor(ctx, [][]rune{li1, li2, li3, li4, li5, li6}).Run()
	if err != nil {
		t.Fatal(err)
	}

	for v := range ch {
		t.Log(v)
		if v == "cs2009" {
			cancel()
		}
	}
}
