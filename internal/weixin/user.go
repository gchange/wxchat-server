package weixin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetOpenId(code string) (openid string, err error) {
	// 构建请求URL
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", c.AppId, c.Secret, code)

	// 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 解析JSON响应
	var result struct {
		OpenID     string `json:"openid"`
		SessionKey string `json:"session_key"`
		UnionID    string `json:"unionid"`
		ErrCode    int    `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	// 检查错误
	if result.ErrCode != 0 {
		return "", fmt.Errorf("WeChat API error: %d - %s", result.ErrCode, result.ErrMsg)
	}

	return result.OpenID, nil

}
