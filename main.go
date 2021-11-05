package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var (
		authorArg  = flag.String("a", "", "auther name and Japanese author name")
		titleArg   = flag.String("t", "", "title name and Japanese title name")
		titleIdArg = flag.Int("i", 0, "title ID")
	)
	flag.Parse()

	author, authorJp := splitArg(*authorArg)
	title, titleJp := splitArg(*titleArg)
	id := *titleIdArg

	fmt.Println(author, authorJp, title, titleJp, id)
}

func splitArg(arg string) (string, string) {
	args := strings.Split(arg, ":")
	arg0 := args[0]
	arg1 := arg0
	if len(args) > 1 {
		arg1 = args[1]
	}
	return arg0, arg1
}
