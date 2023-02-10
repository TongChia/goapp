package goapp

type Packages struct {
	Inject   string   `long:"inject" choice:"uber/fx" choice:"samber/do" choice:"google/wire"`
	Logger   string   `long:"logger" choice:"uber/zap" choice:"logrus"`
	Helper   []string `long:"helper" choice:"samber/lo"`
	Configs  []string `long:"config" choice:"dockerfile" choice:"docker-compose" choice:"k8s" choice:"envoy" choice:"nginx"`
	Database []string `long:"database" choice:"entgo"`
}

type Options struct {
	Verbose  []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Name     string `long:"name" description:"Project name"`
	Style    string `long:"style" description:"Project styles"`
	Packages `group:"packages"`
}
