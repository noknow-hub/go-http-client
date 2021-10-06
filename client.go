//////////////////////////////////////////////////////////////////////
// client.go
//////////////////////////////////////////////////////////////////////
package v1

import (
    "io/ioutil"
    "net/http"
    "time"
)


type Client struct {
    Config *Config
}


//////////////////////////////////////////////////////////////////////
// New client.
//////////////////////////////////////////////////////////////////////
func NewClient(url string) *Client {
    return &Client{
        Config: &Config{
            Url: url,
        },
    }
}


//////////////////////////////////////////////////////////////////////
// HTTP GET request.
//////////////////////////////////////////////////////////////////////
func (c *Client) Get() (*Response, error) {
    req, err := http.NewRequest(HTTP_METHOD_GET, c.Config.Url, nil)
    if err != nil {
        return nil, err
    }

    client := &http.Client{}

    // Timeout
    if c.Config.TimeoutSec > 0 {
        client.Timeout = time.Second * time.Duration(c.Config.TimeoutSec)
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    response := &Response{
        Body: respBody,
        ContentLength: resp.ContentLength,
        Header: resp.Header,
        Proto: resp.Proto,
        ProtoMajor: resp.ProtoMajor,
        ProtoMinor: resp.ProtoMinor,
        Status: resp.Status,
        StatusCode: resp.StatusCode,
        TransferEncoding: resp.TransferEncoding,
        Uncompressed: resp.Uncompressed,
        Trailer: resp.Trailer,
        Request: resp.Request,
        TLS: resp.TLS,
    }

    return response, nil
}

