package gpt4all

import (
	"context"
	"errors"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/gpt4all/internal/gpt4allclient"
)

var (
	ErrMissingModel  = errors.New("missing the gpt4all model, set it in the GPT4ALL_MODEL environment variable")
	ErrEmptyResponse = errors.New("no response")
)

// LLM is a gpt4all LLM implementation.
type LLM struct {
	client *gpt4allclient.Client
}

// _ ensures that LLM implements the llms.LLM interface.
var _ llms.LLM = (*LLM)(nil)

// Call requests a completion for the given prompt.
func (o *LLM) Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	r, err := o.Generate(ctx, []string{prompt}, options...)
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", ErrEmptyResponse
	}
	return r[0].Text, nil
}

func (o *LLM) Generate(ctx context.Context, prompts []string, options ...llms.CallOption) ([]*llms.Generation, error) {
	opts := llms.CallOptions{}

	// Set default options
	opts.MaxTokens = 512
	opts.TopK = 90
	opts.Temperature = 0.1
	opts.TopP = 0.86
	opts.RepetitionPenalty = 1.1

	for _, opt := range options {
		opt(&opts)
	}
	result, err := o.client.CreateCompletion(ctx, &gpt4allclient.CompletionRequest{
		Prompt:            prompts[0],
		TopK:              opts.TopK,
		TopP:              opts.TopP,
		Temperature:       opts.Temperature,
		MaxTokens:         opts.MaxTokens,
		RepetitionPenalty: opts.RepetitionPenalty,
	})
	if err != nil {
		return nil, err
	}
	return []*llms.Generation{
		{Text: result.Text},
	}, nil
}

// New returns a new gpt4all LLM.
func New(opts ...Option) (*LLM, error) {
	// Ensure options are initialized only once.
	initOptions.Do(initOpts)
	options := &options{}
	*options = *defaultOptions // Copy default options.

	for _, opt := range opts {
		opt(options)
	}

	if len(options.model) == 0 {
		return nil, ErrMissingModel
	}
	c, err := gpt4allclient.New(options.model, int(options.modelType), options.threads)
	return &LLM{
		client: c,
	}, err
}
