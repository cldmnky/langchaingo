package gpt4all

import (
	"os"
	"runtime"
	"strings"
	"sync"
)

const (
	LLaMAType ModelType = iota
	GPTJType
	MPTType

	modelEnvVarName     = "GPT4ALL_MODEL"      //nolint:gosec
	modelTypeEnvVarName = "GPT4ALL_MODEL_TYPE" //nolint:gosec
)

type ModelType int

var (
	// nolint: gochecknoglobals
	initOptions sync.Once

	// nolint: gochecknoglobals
	defaultOptions *options
)

type options struct {
	model     string
	modelType ModelType
	threads   int
}

type Option func(*options)

// initOpts initializes defaultOptions with the environment variables.
func initOpts() {
	var mt ModelType
	switch strings.ToLower(os.Getenv(modelTypeEnvVarName)) {
	case "gptj":
		mt = ModelType(GPTJType)
	case "mpt":
		mt = ModelType(MPTType)
	default:
		mt = ModelType(LLaMAType)
	}
	defaultOptions = &options{
		model:     os.Getenv(modelEnvVarName),
		modelType: mt,
		threads:   runtime.NumCPU(),
	}
}

func WithModelType(modelType ModelType) Option {
	return func(opts *options) {
		mt := ModelType(modelType)
		opts.modelType = mt
	}
}

func WithModel(model string) Option {
	return func(opts *options) {
		opts.model = model
	}
}

func WithThreads(threads int) Option {
	return func(opts *options) {
		opts.threads = threads
	}
}
