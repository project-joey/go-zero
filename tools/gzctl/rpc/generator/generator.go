package generator

import (
	"log"

	conf "github.com/project-joey/go-zero/tools/gzctl/config"
	"github.com/project-joey/go-zero/tools/gzctl/env"
	"github.com/project-joey/go-zero/tools/gzctl/util/console"
)

// Generator defines the environment needs of rpc service generation
type Generator struct {
	log     console.Console
	cfg     *conf.Config
	verbose bool
}

// NewGenerator returns an instance of Generator
func NewGenerator(style string, verbose bool) *Generator {
	cfg, err := conf.NewConfig(style)
	if err != nil {
		log.Fatalln(err)
	}

	colorLogger := console.NewColorConsole(verbose)

	return &Generator{
		log:     colorLogger,
		cfg:     cfg,
		verbose: verbose,
	}
}

// Prepare provides environment detection generated by rpc service,
// including go environment, protoc, whether protoc-gen-go is installed or not
func (g *Generator) Prepare() error {
	return env.Prepare(true, true, g.verbose)
}
