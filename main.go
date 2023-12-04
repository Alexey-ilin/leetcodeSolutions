package leetcodeSols

import (
	"fmt"
	"strconv"
)

func M() {
	if n, err := strconv.Atoi("123"); err != nil {
		fmt.Println(n)
	}
}
