package kibana

import (
	"crypto/tls"

	"github.com/ggsood/kibana-client/kbapi"
	"github.com/go-resty/resty/v2"
)

// Config contain the value to access on Kibana API
// The Kibana APIs support key (Basic Authentication) and token-based authentication.
type Config struct {
	Address          string
	Username         string
	Password         string
	DisableVerifySSL bool
	ApiKey           string
	CAs              []string
}

// Client contain the REST client and the API specification
type Client struct {
	*kbapi.API
	Client *resty.Client
}

// NewDefaultClient init client with empty config
func NewDefaultClient() (*Client, error) {
	return NewClient(Config{})
}

// NewClient init client with custom config
func NewClient(cfg Config) (*Client, error) {
	if cfg.Address == "" {
		cfg.Address = "http://localhost:5601"
	}

	restyClient := getKibanaClient(cfg)

	for _, path := range cfg.CAs {
		restyClient.SetRootCertificate(path)
	}

	client := &Client{
		Client: restyClient,
		API:    kbapi.New(restyClient),
	}

	if cfg.DisableVerifySSL {
		client.Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	return client, nil

}

// Get Kibana Client Based on Basic / Token Based Authentication
func getKibanaClient(cfg Config) *resty.Client {
	if cfg.ApiKey != "" {
		apiKey := "ApiKey " + cfg.ApiKey
		return resty.New().
			SetBaseURL(cfg.Address).
			SetHeader("Authorization", apiKey).
			SetHeader("kbn-xsrf", "true").
			SetHeader("Content-Type", "application/json")
	} else {
		return resty.New().
			SetBaseURL(cfg.Address).
			SetBasicAuth(cfg.Username, cfg.Password).
			SetHeader("kbn-xsrf", "true").
			SetHeader("Content-Type", "application/json")
	}
}
