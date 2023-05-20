package llms

// CallOption is a function that configures a CallOptions.
type CallOption func(*CallOptions)

// CallOptions is a set of options for LLM.Call.
type CallOptions struct {
	// Model is the model to use.
	Model string `json:"model"`
	// MaxTokens is the maximum number of tokens to generate.
	MaxTokens int `json:"max_tokens"`
	// Temperature is the temperature for sampling, between 0 and 1.
	Temperature float64 `json:"temperature"`
	// StopWords is a list of words to stop on.
	StopWords []string `json:"stop_words"`
	// TopK is the number of tokens to sample from.
	TopK int `json:"top_k"`
	// TopP is the cumulative probability for top-k sampling.
	TopP float64 `json:"top_p"`
	// RepetitionPenalty is the repetition penalty.
	RepetitionPenalty float64 `json:"repetition_penalty"`
}

// WithModel is an option for LLM.Call.
func WithModel(model string) CallOption {
	return func(o *CallOptions) {
		o.Model = model
	}
}

// WithMaxTokens is an option for LLM.Call.
func WithMaxTokens(maxTokens int) CallOption {
	return func(o *CallOptions) {
		o.MaxTokens = maxTokens
	}
}

// WithTemperature is an option for LLM.Call.
func WithTemperature(temperature float64) CallOption {
	return func(o *CallOptions) {
		o.Temperature = temperature
	}
}

// WithStopWords is an option for LLM.Call.
func WithStopWords(stopWords []string) CallOption {
	return func(o *CallOptions) {
		o.StopWords = stopWords
	}
}

// WithTopK is an option for LLM.Call.
func WithTopK(topK int) CallOption {
	return func(o *CallOptions) {
		o.TopK = topK
	}
}

// WithTopP is an option for LLM.Call.
func WithTopP(topP float64) CallOption {
	return func(o *CallOptions) {
		o.TopP = topP
	}
}

// WithRepetitionPenalty is an option for LLM.Call.
func WithRepetitionPenalty(repetitionPenalty float64) CallOption {
	return func(o *CallOptions) {
		o.RepetitionPenalty = repetitionPenalty
	}
}

// WithOptions is an option for LLM.Call.
func WithOptions(options CallOptions) CallOption {
	return func(o *CallOptions) {
		(*o) = options
	}
}
