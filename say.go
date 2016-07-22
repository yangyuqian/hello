package hello

import (
	"fmt"
	_ "github.com/RangelReale/osin"
)

func SayHello(msg string) {
	fmt.Println("SayHello =>", msg)
}
