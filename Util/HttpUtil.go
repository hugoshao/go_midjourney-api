package Util

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

// HTTPGet 发起一个HTTP GET请求，并返回响应体的字节切片和错误（如果有）。
// 如果指定了proxyURL，则使用代理发起请求。
func HTTPGet(urlStr string, proxyURL *url.URL) ([]byte, error) {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 如果设置了代理，创建一个Transport来使用代理
	if proxyURL != nil {
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		client.Transport = transport
	}

	// 发起GET请求
	resp, err := client.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// HTTPPost 发起一个HTTP POST请求，并返回响应体的字节切片和错误（如果有）。
// 如果指定了proxyURL，则使用代理发起请求。
func HTTPPost(urlStr string, body []byte, proxyURL *url.URL, userToken string) ([]byte, error) {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 如果设置了代理，创建一个Transport来使用代理
	if proxyURL != nil {
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		client.Transport = transport
	}

	// 构建请求
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	if len(userToken) > 0 {
		req.Header.Set("Authorization", userToken)
	}
	// 发起POST请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
