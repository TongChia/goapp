package goapp

import (
	"fmt"
	"os"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

// clone code
func clone(path, url string) error {
	fmt.Println("git clone " + url)
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		panic(err)
	}
	return nil
}

// fix project name
func fixup(name string) {
	fmt.Println("fixup project name")
}

// print ok
func final(name string) {
	fmt.Printf("OK\n"+
		"next step: \n"+
		"cd %s\n"+
		"make init\n"+
		"make api\n"+
		"go run cmd/main.go\n",
		name,
	)
}

func Generate(opt *Options) {
	name := strings.Replace(strings.ToLower(opt.Name), " ", "-", -1)
	clone(name, "https://github.com/tongchia/goapp-layout-simple-grpc.git")
	fixup(name)
	final(name)
}
