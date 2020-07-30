package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

const foo = "foobar"

func main() {
	fmt.Println(foo)
	fmt.Println(common.BytesToHash([]byte(foo)).Hex())
}
