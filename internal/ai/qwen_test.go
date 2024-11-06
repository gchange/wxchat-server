package ai

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQwen(t *testing.T) {
	c := &QwenConfig{
		Server: "https://dashscope.aliyuncs.com",
		Model:  "qwen-turbo",
		Key:    "sk-0435ba7650234838bae51fd38fafc628",
	}
	qwen, err := NewQwen(c)
	assert.Nil(t, err)

	r, err := qwen.Chat("who are you?")
	assert.Nil(t, err)

	fmt.Println(r)
}
