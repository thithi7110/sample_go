package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Hoge struct {
	aaa int
	bbb int
}

func (h *Hoge) add(num int) error {
	if num == 0 {
		return errors.New("ゼロは無効")
	}

	h.aaa += num
	h.bbb += num

	return nil
}

func (h *Hoge) text() string {
	return strconv.Itoa(h.aaa) + strconv.Itoa(h.bbb)
}

func main() {
	fmt.Println("start")

	hoge := Hoge{1, 2}

	fmt.Println(hoge.aaa, hoge.bbb)
	err := hoge.add(0)
	if err != nil {
		fmt.Println("エラー：", err)
		defer fmt.Println("えらーはっせい")
	}
	fmt.Println(hoge.text())
}
