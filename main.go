package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/burl/inquire"
	"github.com/burl/inquire/widget"
	"github.com/jessevdk/go-flags"
	git "gopkg.in/src-d/go-git.v4"
)

type packages struct {
	Inject   string   `long:"inject" choice:"uber/fx" choice:"uber/dig" choice:"samber/do" choice:"google/wire"`
	Logger   string   `long:"logger" choice:"uber/zap" choice:"logrus"`
	Helper   []string `long:"helper" choice:"samber/lo"`
	Configs  []string `long:"config" choice:"dockerfile" choice:"docker-compose" choice:"k8s" choice:"envoy" choice:"nginx"`
	Database []string `long:"database" choice:"entgo"`
}

type options struct {
	Verbose  []bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	Name     string   `long:"name" description:"Project name"`
	Style    string   `long:"style" description:"Project styles"`
	Packages packages `group:"packages"`
}

func generate(opt options) {
	var ok bool

	q := inquire.Query()
	if opt.Name == "" {
		q = q.Input(&opt.Name, "What is your project name", nil)
	}
	if opt.Style == "" {
		q = q.Menu(&opt.Style, "What is your project styles", func(w *widget.Menu) {
			w.Hint("use arrow keys, pick one")
			w.Item("simple-grpc", "Simple gRPC service")
			w.Item("grpc-http", "gRPC with HTTP service")
			w.Item("grpc-micro", "gRPC micro services")
		})
	}
	if opt.Packages.Inject == "" {
		q = q.Menu(&opt.Packages.Inject, "what is your favorite dependency injection package", func(w *widget.Menu) {
			w.Hint("use arrow keys, pick one")
			w.Item("uber/fx", "uber/fx")
			w.Item("uber/dig", "uber/dig")
			w.Item("samber/do", "samber/do")
			w.Item("google/wire", "google/wire")
		})
	}
	if opt.Packages.Helper == nil {
		opts := map[string]bool{
			"samber/lo": false,
		}
		q = q.Select("what are your favorite helper packages", func(w *widget.Select) {
			w.Hint("use arrow/space, select multiple")
			for k, v := range opts {
				w.Item(&v, k)
			}
		})
	}
	if opt.Packages.Database == nil {
		opts := map[string]bool{
			"entgo":    false,
			"go-redis": false,
		}
		q = q.Select("what database driver / ORM do you need", func(w *widget.Select) {
			w.Hint("use arrow/space, select multiple")
			for k, v := range opts {
				w.Item(&v, k)
			}
		})
	}
	if opt.Packages.Configs == nil {
		opts := map[string]bool{
			"dockerfile":     false,
			"docker-compose": false,
			"k8s":            false,
			"envoy":          false,
			"nginx":          false,
		}
		q = q.Select("what config do you need", func(w *widget.Select) {
			w.Hint("use arrow/space, select multiple")
			for k, v := range opts {
				w.Item(&v, k)
			}
		})
	}

	q.YesNo(&ok, "Continue").Exec()

	if !ok {
		fmt.Println("aborted.")
		os.Exit(1)
	}

	fmt.Printf("\nReady to generate project `%s`\n---------------------\n", opt.Name)
	fmt.Println("git clone https://github.com/go-kratos/kratos-layout")
	_, err := git.PlainClone(strings.Replace(strings.ToLower(opt.Name), " ", "-", 0), false, &git.CloneOptions{
		URL:      "https://github.com/go-kratos/kratos-layout.git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OK")
}

func main() {
	opt := options{}
	parser := flags.NewParser(&opt, flags.Default)
	_, err := parser.Parse()
	if err == nil {
		generate(opt)
	} else {
		if e, ok := err.(*flags.Error); ok {
			if e.Type == flags.ErrHelp {
				return
			}
		}
		fmt.Println(err)
	}
}
