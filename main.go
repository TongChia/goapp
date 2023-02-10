package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
	goapp "github.com/tongchia/goapp/src"
)

func main() {
	opt := goapp.Options{}
	parser := flags.NewParser(&opt, flags.HelpFlag|flags.IgnoreUnknown|flags.PassDoubleDash)
	_, err := parser.Parse()
	if err == nil {
		ok := goapp.Inquirer(&opt)
		if ok {
			goapp.Generate(&opt)
		}
	} else {
		fmt.Println(err)
	}
}
