//////////////////////////////////////////////////////////////////////
// config.go
//////////////////////////////////////////////////////////////////////
package v1

import (
    "net/url"
)

type Config struct {
    Cred string
    JsonData []byte
    TimeoutSec int
    Url string
    UrlQuery *url.Values
}


