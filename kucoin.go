package kugo

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SpotEndpoint Base URL
const SpotEndpoint = "https://api.kucoin.com"
const FutureEndpoint = "https://api-futures.kucoin.com"
const ApiKeyVersionV2 = "2"

type Kucoin struct {
	spotEndpoint   string
	futureEndpoint string
	secretKey      string
	accessKey      string
	passphrase     string
	signer         Signer

	client  *resty.Client
	reqLog  func(...interface{})
	respLog func(...interface{})
	debug   bool
}

type Option func(kc *Kucoin) error

// NewKucoin Create a Kucoin API instance
func NewKucoin(opts ...Option) (*Kucoin, error) {
	kc := &Kucoin{}
	kc.spotEndpoint = SpotEndpoint
	kc.futureEndpoint = FutureEndpoint
	kc.reqLog = defaultLog
	kc.respLog = defaultLog
	kc.client = resty.NewWithClient(
		&http.Client{Transport: &http.Transport{
			DisableKeepAlives: true},
			Timeout: 10 * time.Second})

	err := kc.Set(opts...)
	if err != nil {
		return nil, err
	}

	kc.registerMiddleware()
	return kc, nil
}

// Print the log to the console
func defaultLog(v ...interface{}) {
	log.Println(v...)
}

// Set optional parameters
func (kc *Kucoin) Set(options ...Option) error {
	for _, option := range options {
		if err := option(kc); err != nil {
			return err
		}
	}
	kc.signer = NewKcSignerV2(kc.accessKey, kc.secretKey, kc.passphrase)
	return nil
}

// SetApiKey Set the API key created in Kucoin
func SetApiKey(accessKey, secretKey, passphrase string) Option {
	return func(kc *Kucoin) error {
		if kc == nil {
			return errors.New("instance is nil")
		}
		kc.accessKey = accessKey
		kc.secretKey = secretKey
		kc.passphrase = passphrase
		return nil
	}
}

// SetSpotEndpoint Use the specified base URL to access the Kucoin spot API interface
func SetSpotEndpoint(spotEndpoint string) Option {
	return func(kc *Kucoin) error {
		if kc == nil {
			return errors.New("instance is nil")
		}
		kc.spotEndpoint = spotEndpoint
		return nil
	}
}

// SetFutureEndpoint Use the specified base URL to access the Kucoin future API interface
func SetFutureEndpoint(futureEndpoint string) Option {
	return func(kc *Kucoin) error {
		if kc == nil {
			return errors.New("instance is nil")
		}
		kc.futureEndpoint = futureEndpoint
		return nil
	}
}

// SetClient Set HTTP client
func SetClient(client *http.Client) Option {
	return func(kc *Kucoin) error {
		if kc == nil {
			return errors.New("instance is nil")
		}
		if client == nil {
			return errors.New("client is nil")
		}
		kc.client = resty.NewWithClient(client)
		return nil
	}
}

// SetRequestLog Set the HTTP request log printing method. Set debug to true for it to work
func SetRequestLog(l func(...interface{})) Option {
	return func(kc *Kucoin) error {
		if kc == nil {
			return errors.New("instance is nil")
		}
		if l == nil {
			return errors.New("logger is nil")
		}
		kc.reqLog = l
		return nil
	}
}

// SetResponseLog Set the HTTP response log printing method. Set debug to true for it to work
func SetResponseLog(l func(...interface{})) Option {
	return func(kc *Kucoin) error {
		if kc == nil {
			return errors.New("instance is nil")
		}
		if l == nil {
			return errors.New("logger is nil")
		}
		kc.reqLog = l
		return nil
	}
}

// SetDebug If set debug to ture, the http request and response logs will be output
func SetDebug(debug bool) Option {
	return func(kc *Kucoin) error {
		kc.debug = debug
		return nil
	}
}

// Register HTTP request middleware
func (kc *Kucoin) registerMiddleware() {
	// Registering Request Middleware
	kc.client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		// Now you have access to Client and current Request object
		// manipulate it as per your need
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36")
		req.Header.Add("Content-Type", "application/json")
		if kc.client != nil && kc.debug {
			body := ""
			if req.Body != nil {
				body = string(req.Body.([]byte))
			}
			kc.reqLog(fmt.Sprintf("info:request\turl:%s\tquery_param:%+v\tform_data:%+v\tbody:%+v\theader=%+v", req.URL, req.QueryParam, req.FormData, body, req.Header))
		}
		return nil
	})

	// Registering Response Middleware
	kc.client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		if kc.respLog != nil && kc.debug {
			body := ""
			if resp.Request.Body != nil {
				body = string(resp.Request.Body.([]byte))
			}
			kc.respLog(fmt.Sprintf("info:response\turl:%s\tquery_param:%+v\tform_data:%+v\tbody:%+v\tresponse:%s\theader=%+v", resp.Request.URL, resp.Request.QueryParam, resp.Request.FormData, body, string(resp.Body()), resp.Request.Header))
		}

		return nil
	})
}

// do Send http request to Kucoin.
// When method is GET or DELETE, the type of params is map[string]string{}.
// When method is POST, the type of params is []byte.
// Examples:
// do("https://www.kucoin.com", "GET", "/api/v1/accounts", map[string]string{"currency":"BTC", "type":"trade"})
// do("https://www.kucoin.com", "POST", "/api/v1/orders", []byte("{\"price\":\"100\",...}"))
func (kc *Kucoin) do(endpoint string, method string, uri string, params interface{}) (resp *resty.Response, err error) {
	us := fmt.Sprintf("%s%s", endpoint, uri)
	header := make(map[string]string)
	body := make([]byte, 0)
	if kc.signer != nil {
		var b bytes.Buffer
		b.WriteString(method)
		if method == http.MethodGet || method == http.MethodDelete {
			if params != nil {
				p, _ := params.(map[string]string)
				qs := make([]string, 0, len(p))
				for k, v := range p {
					qs = append(qs, fmt.Sprintf("%s=%s", k, v))
				}
				us = fmt.Sprintf("%s?%s", us, strings.Join(qs, "&"))

				u, e := url.Parse(us)
				if e != nil {
					err = e
					return
				}
				us = u.String()
				uri = u.RequestURI()
				b.WriteString(uri)
			}
		} else {
			b.WriteString(uri)
			if params != nil {
				body, _ = params.([]byte)
				fmt.Println(string(body))
				b.Write(body)
			}
		}
		h := kc.signer.(*KcSigner).Headers(b.String())
		for k, v := range h {
			header[k] = v
		}
	}

	req := kc.client.R().
		SetHeaders(header)
	switch method {
	case http.MethodGet:
		resp, err = req.Get(us)
		return
	case http.MethodPost:
		resp, err = req.SetBody(body).Post(us)
		return
	}
	return
}
