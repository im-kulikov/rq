package client

import (
	"net/http"
	"net/http/cookiejar"
	"time"

	"gopkg.in/ddo/go-dlog.v2"

	"github.com/ddo/rq"
)

const (
	defaultTimeout = 3 * time.Minute
)

var log = dlog.New("rq:client", nil)

// Client contains stdlib http client and other custom client settings
type Client struct {
	httpClient *http.Client

	DefaultRq *rq.Rq // TODO
}

// Option contains client settings
type Option struct {
	Timeout   time.Duration
	NoTimeout bool // if NoTimeout is false Timeout will be set as default

	Jar      http.CookieJar
	NoCookie bool // if NoCookie is true Jar will be skip

	Transport http.RoundTripper
}

// New returns new client which init with provided options
// cookie management is enable by default
func New(opt *Option) *Client {
	if opt == nil {
		opt = &Option{}
	}

	timeout := opt.Timeout
	if opt.Timeout == 0 && !opt.NoTimeout {
		timeout = defaultTimeout
	}
	if opt.NoTimeout {
		timeout = 0
	}

	jar := opt.Jar
	if opt.Jar == nil && !opt.NoCookie {
		jar, _ = cookiejar.New(nil)
	}
	if opt.NoCookie {
		jar = nil
	}

	log.Info("opt:", opt)
	return &Client{
		httpClient: &http.Client{
			Timeout:   timeout,
			Jar:       jar,
			Transport: opt.Transport,
		},
	}
}

// DefaultClient has default timeout and stdlib default transport
// no cookie management
var DefaultClient = New(&Option{NoCookie: true})