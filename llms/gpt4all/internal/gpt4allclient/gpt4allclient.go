package gpt4allclient

import (
	"context"
	"sync"

	gpt4all "github.com/nomic-ai/gpt4all/gpt4all-bindings/golang"
)

// Client is a client for a gpt4all LLM.
type Client struct {
	model *gpt4all.Model
}

// New returns a new gpt4all client.
func New(model string, modelType, threads int) (*Client, error) {
	mt := gpt4all.ModelType(modelType)
	m, err := gpt4all.New(model, gpt4all.SetModelType(mt), gpt4all.SetThreads(threads))
	if err != nil {
		return nil, err
	}
	c := &Client{model: m}
	return c, nil
}

// CompletionRequest is a request to create a completion.
type CompletionRequest struct {
	Prompt            string  `json:"prompt"`
	MaxTokens         int     `json:"max_tokens"`
	Temperature       float64 `json:"temperature"`
	TopK              int     `json:"top_k"`
	TopP              float64 `json:"top_p"`
	RepetitionPenalty float64 `json:"repetition_penalty"`
}

// Completion is a completion.
type Completion struct {
	Text string `json:"text"`
}

func (c *Client) CreateCompletion(ctx context.Context, r *CompletionRequest) (*Completion, error) {
	completion := &Completion{}
	// Predict in a waitgroup to avoid concurrent calls to the model.
	var (
		wg         sync.WaitGroup
		predictErr error
	)
	c.model.SetTokenCallback(func(token string) bool {
		//fmt.Print(token)
		return true
	})
	wg.Add(1)
	go func() {
		defer wg.Done()
		text, err := c.model.Predict(r.Prompt,
			gpt4all.SetTokens(r.MaxTokens),
			gpt4all.SetTemperature(r.Temperature),
			gpt4all.SetTopK(r.TopK),
			gpt4all.SetTopP(r.TopP),
			gpt4all.SetRepeatPenalty(r.RepetitionPenalty))
		if err != nil {
			predictErr = err
			return
		}
		completion.Text = text
	}()
	wg.Wait()
	return completion, predictErr
}
