package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/go-querystring/query"
)

type Client struct {
	apiKey string
}

// NewClient creates a new client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

// Post makes a post request
func (c *Client) Post(url string, input any) (response []byte, err error) {
	response = make([]byte, 0)

	rJson, err := json.Marshal(input)
	if err != nil {
		return response, err
	}

	resp, err := c.Call(http.MethodPost, url, bytes.NewReader(rJson))
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	response, err = io.ReadAll(resp.Body)
	return response, err
}

// Get makes a get request
func (c *Client) Get(url string, input any) (response []byte, err error) {
	if input != nil {
		vals, _ := query.Values(input)
		query := vals.Encode()

		if query != "" {
			url += "?" + query
		}
	}

	resp, err := c.Call(http.MethodGet, url, nil)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	response, err = io.ReadAll(resp.Body)
	return response, err
}

// Call makes a request
func (c *Client) Call(method string, url string, body io.Reader) (response *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return response, err
	}

	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}

type CreateCompletionsRequest struct {
	Model            string            `json:"model,omitempty"`
	Messages         []Message         `json:"messages,omitempty"`
	Prompt           StrArray          `json:"prompt,omitempty"`
	Suffix           string            `json:"suffix,omitempty"`
	MaxTokens        int               `json:"max_tokens,omitempty"`
	Temperature      float64           `json:"temperature,omitempty"`
	TopP             float64           `json:"top_p,omitempty"`
	N                int               `json:"n,omitempty"`
	Stream           bool              `json:"stream,omitempty"`
	LogProbs         int               `json:"logprobs,omitempty"`
	Echo             bool              `json:"echo,omitempty"`
	Stop             StrArray          `json:"stop,omitempty"`
	PresencePenalty  float64           `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64           `json:"frequency_penalty,omitempty"`
	BestOf           int               `json:"best_of,omitempty"`
	LogitBias        map[string]string `json:"logit_bias,omitempty"`
	User             string            `json:"user,omitempty"`
}

func (c *Client) CreateCompletionsRaw(r CreateCompletionsRequest) ([]byte, error) {
	return c.Post(os.Getenv("CHATGPT_URL"), r)
}

func (c *Client) CreateCompletions(r CreateCompletionsRequest) (response CreateCompletionsResponse, err error) {
	raw, err := c.CreateCompletionsRaw(r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}

type CreateCompletionsResponse struct {
	ID      string `json:"id,omitempty"`
	Object  string `json:"object,omitempty"`
	Created int    `json:"created,omitempty"`
	Model   string `json:"model,omitempty"`
	Choices []struct {
		Message struct {
			Role    string `json:"role,omitempty"`
			Content string `json:"content,omitempty"`
		} `json:"message"`
		Text         string      `json:"text,omitempty"`
		Index        int         `json:"index,omitempty"`
		Logprobs     interface{} `json:"logprobs,omitempty"`
		FinishReason string      `json:"finish_reason,omitempty"`
	} `json:"choices,omitempty"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens,omitempty"`
		CompletionTokens int `json:"completion_tokens,omitempty"`
		TotalTokens      int `json:"total_tokens,omitempty"`
	} `json:"usage,omitempty"`

	Error *Error `json:"error,omitempty"`
}

// Error is the error standard response from the API
type Error struct {
	Message string      `json:"message,omitempty"`
	Type    string      `json:"type,omitempty"`
	Param   interface{} `json:"param,omitempty"`
	Code    interface{} `json:"code,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}
