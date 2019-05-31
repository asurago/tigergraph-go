package tigergraph_go

import (
    "net/http"
    "time"
)

type Client struct {
    HttpClient *http.Client
}

func NewClient() (*Client, error) {
    tr := &http.Transport{
        MaxIdleConns:       10,
        IdleConnTimeout:    5 * time.Second,
        DisableCompression: true,
    }
    client := &http.Client{Transport: tr}
    return &Client{HttpClient: client}, nil
}

func (c *Client) Post(urlString string) {
}
