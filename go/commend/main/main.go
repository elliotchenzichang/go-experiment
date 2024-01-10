package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	flag.StringVar(&user, "u", "root", "账号，默认为root")
	fmt.Println(user)
}
