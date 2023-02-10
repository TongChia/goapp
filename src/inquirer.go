package goapp

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Inquirer(opt *Options) bool {
	if opt.Name == "" {
		survey.AskOne(&survey.Input{Message: "What is the project name?"}, &opt.Name, survey.WithValidator(survey.Required))
	}
	if opt.Style == "" {
		survey.AskOne(&survey.Select{
			Message: "What is the project styles?",
			Options: []string{"simple-grpc", "grpc-http", "grpc-micro"},
			Description: func(value string, _ int) string {
				switch value {
				case "simple-grpc":
					return "Simple gRPC service"
				case "grpc-http":
					return "gRPC with HTTP service"
				case "grpc-micro":
					return "gRPC micro services"
				default:
					return ""
				}
			},
		}, &opt.Style, survey.WithValidator(survey.Required))
	}
	if opt.Inject == "" {
		survey.AskOne(&survey.Select{
			Message: "What is your favorite dependency injection package?",
			Options: []string{"uber/fx", "samber/do", "google/wire"},
		}, &opt.Inject)
	}
	if opt.Helper == nil {
		survey.AskOne(&survey.MultiSelect{
			Message: "What helper package do you need?",
			Options: []string{"samber/lo"},
		}, &opt.Helper)
	}
	if opt.Database == nil {
		survey.AskOne(&survey.MultiSelect{
			Message: "What database driver or ORM do you need?",
			Options: []string{"entgo", "go-redis"},
		}, &opt.Database)
	}
	if opt.Configs == nil {
		survey.AskOne(&survey.MultiSelect{
			Message: "What config do you need?",
			Options: []string{"dockerfile", "docker-compose", "k8s", "envoy", "nginx"},
		}, &opt.Configs)
	}

	ok := false
	survey.AskOne(&survey.Confirm{Message: "Continue"}, &ok)

	if ok {
		fmt.Printf("%+v\n", opt)
	}
	return ok
}
