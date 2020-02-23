package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	CharsetUtf8        = "utf8"
	ContentTypeJSON    = "application/json"
	defaultWithoutTime = 5 * time.Second
)

// Client client
type Client struct {
	contentType string
	httpClient  *http.Client
	withoutTime int
}

// NewClient new client
func NewClient(httpClient *http.Client, withoutTime int) Client {
	return Client{
		httpClient:  httpClient,
		withoutTime: withoutTime,
	}
}

// GetHttpClient get http_client
func (c Client) GetHttpClient() *http.Client {
	if c.httpClient == nil {
		return &http.Client{}
	}
	return c.httpClient
}

// GetWithoutTime get without time
func (c Client) GetWithoutTime() time.Duration {
	if c.withoutTime <= 0 {
		return defaultWithoutTime
	}
	return time.Duration(c.withoutTime) * time.Second
}

// Do do transport
func (c Client) Do(req *http.Request) ([]byte, error) {
	var (
		res *http.Response
		err error
	)

	ctx, cancel := context.WithTimeout(context.Background(), c.GetWithoutTime())
	defer cancel()

	req = req.WithContext(ctx)

	client := c.GetHttpClient()

	// 出错重试机制
	tryCount := 0
tryAgain:
	res, err = client.Do(req)
	if err != nil {
		fmt.Println("do post err:", err, tryCount)
		select {
		case <-ctx.Done():
			return nil, err
		default:
		}
		tryCount++
		if tryCount < 3 {
			goto tryAgain
		}
		return nil, err
	}

	if res.Body == nil {
		return nil, fmt.Errorf("do post response is nil")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code error: %v", res.StatusCode)
	}

	return ioutil.ReadAll(res.Body)
}

func (c Client) JSONGet(turl string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, turl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Charset", CharsetUtf8)
	req.Header.Add("Content-Type", ContentTypeJSON)

	body, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c Client) JSONPost(turl string, param interface{}) ([]byte, error) {
	p, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, turl, bytes.NewReader(p))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Charset", CharsetUtf8)
	req.Header.Add("Content-Type", ContentTypeJSON)

	body, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}
