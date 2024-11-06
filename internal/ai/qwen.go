package ai

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"path"
	"strings"

	"github.com/levigross/grequests"
)

type QwenRole string
type QwenFormat string

const (
	QwenRoleUser      QwenRole = "user"
	QwenRoleSystem    QwenRole = "system"
	QwenRoleAssistant QwenRole = "assistant"

	QwenFormatJson QwenFormat = "json_object"
	QwenFormatText QwenFormat = "text"
)

type QwenConfig struct {
	Server string `json:"server" yaml:"server"`
	Model  string `json:"model" yaml:"model"`
	Key    string `json:"key" yaml:"key"`
}

type QWen struct {
	server *url.URL
	model  string
	key    string
}

func NewQwen(c *QwenConfig) (qwen *QWen, err error) {
	s, err := url.Parse(c.Server)
	if err != nil {
		return nil, err
	}
	return &QWen{
		server: s,
		model:  c.Model,
		key:    c.Key,
	}, nil
}

type QwenStreamOptions struct {
	IncludeUsage bool `json:"include_usage"`
}

type QwenResponseFormat struct {
	Type QwenFormat `json:"type,omitempty"`
}

type QwenChatMessage struct {
	Role    QwenRole `json:"role"`
	Content string   `json:"content"`
}

type QwenChatRequest struct {
	Model          string              `json:"model"`
	Messages       []QwenChatMessage   `json:"messages"`
	Stream         bool                `json:"stream,omitempty"`
	StreamOption   *QwenStreamOptions  `json:"stream_options,omitempty"`
	ResponseFormat *QwenResponseFormat `json:"response_format,omitempty"`
}

type QwenChatResponseChoiceDelta struct {
	Content string `json:"content"`
}

type QwenChatResponseChoiceMessage struct {
	Role         QwenRole `json:"role"`
	Content      string   `json:"content"`
	FinishReason string   `json:"finish_reason"`
	Index        int      `json:"index"`
	Logprobs     any      `json:"logprobs"`
}

type QWenChatResponseChoice struct {
	Delta   QwenChatResponseChoiceDelta   `json:"delta"`
	Message QwenChatResponseChoiceMessage `json:"message"`
}

type QwenChatResponse struct {
	Choices           []QWenChatResponseChoice `json:"choices"`
	FinishReason      any                      `json:"finish_reason"`
	Index             int                      `json:"index"`
	Logprobs          any                      `json:"logprobs"`
	Object            string                   `json:"object"`
	Usage             any                      `json:"usage"`
	Created           int                      `json:"created"`
	SystemFingerprint any                      `json:"system_fingerprint"`
	Model             string                   `json:"model"`
	ID                string                   `json:"id"`
}

func (qwen *QWen) decode(buf []byte) (r *QwenChatResponse) {
	data := []byte(strings.TrimSpace(strings.TrimPrefix(string(buf), "data:")))
	if !json.Valid(data) {
		return nil
	}
	r = &QwenChatResponse{}
	err := json.Unmarshal(data, &r)
	if err != nil {
		return nil
	}
	return
}

func (qwen *QWen) Chat(text string) (content string, err error) {
	p := path.Join(qwen.server.Path, "compatible-mode/v1/chat/completions")
	u, err := qwen.server.Parse(p)
	if err != nil {
		return "", err
	}
	ro := &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": qwen.key,
			"Content-Type":  "application/json",
		},
		JSON: QwenChatRequest{
			Model: qwen.model,
			Messages: []QwenChatMessage{
				{
					Role:    QwenRoleUser,
					Content: text,
				},
			},
		},
	}
	resp, err := grequests.Post(u.String(), ro)
	if err != nil {
		return "", err
	}
	r := qwen.decode(resp.Bytes())
	if r == nil {
		return "", fmt.Errorf("response is not valid json")
	}
	return r.Choices[0].Message.Content, nil
}

func (qwen *QWen) StreamChat(text string) (c <-chan string, err error) {
	p := path.Join(qwen.server.Path, "compatible-mode/v1/chat/completions")
	u, err := qwen.server.Parse(p)
	if err != nil {
		return nil, err
	}
	ro := &grequests.RequestOptions{
		Headers: map[string]string{
			"Authorization": qwen.key,
			"Content-Type":  "application/json",
		},
		JSON: QwenChatRequest{
			Model: qwen.model,
			Messages: []QwenChatMessage{
				{
					Role:    QwenRoleUser,
					Content: text,
				},
			},
			Stream: true,
			StreamOption: &QwenStreamOptions{
				IncludeUsage: true,
			},
		},
	}
	resp, err := grequests.Post(u.String(), ro)
	if err != nil {
		return nil, err
	}
	br := bufio.NewReader(resp.RawResponse.Body)
	ch := make(chan string, 1)
	go func() {
		defer close(ch)
		for {
			line, _, err := br.ReadLine()
			if err != nil {
				break
			}
			if len(line) == 0 {
				continue
			}
			r := qwen.decode(line)
			if r == nil {
				log.Printf("deocde response error. line=%s", string(line))
				break
			}
			if len(r.Choices) == 0 {
				continue
			}
			choice := r.Choices[0]
			if len(choice.Delta.Content) == 0 {
				continue
			}
			ch <- choice.Delta.Content
		}
	}()
	return ch, nil
}
