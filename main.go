package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	argsWithoutProg := strings.Join(os.Args[1:], "")
	str := strings.Fields(argsWithoutProg)
	fmt.Println(len(str))

}
