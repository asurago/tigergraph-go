package tigergraph_go

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Request struct {
    HttpRequest *http.Request
}

func NewRequest(urlString string, requestBody string, c *Client) ([]byte, error) {
    response := []byte{}
    reader := bytes.NewReader(bytes.NewBufferString(requestBody).Bytes())
    request, err := http.NewRequest("POST", urlString, reader)
    if err != nil {
        fmt.Println(err.Error())
        return response, err
    }
    request.Header.Set("Content-Type", "application/json;charset=UTF-8")
    resp, err := c.HttpClient.Do(request)
    if err != nil {
        fmt.Println(err.Error())
        return response, err
    }
    respBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err.Error())
    }
    return respBytes, err
}
