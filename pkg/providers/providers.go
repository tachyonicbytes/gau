package providers

import (
	"context"
	// "tinygo.org/x/drivers/net/http"
	"net/http"
)

const Version = `2.1.2`

// Provider is a generic interface for all archive fetchers
type Provider interface {
	Fetch(ctx context.Context, domain string, results chan string) error
	Name() string
}

type URLScan struct {
	Host   string
	APIKey string
}

type Config struct {
	Threads           uint
	Timeout           uint
	Verbose           bool
	MaxRetries        uint
	IncludeSubdomains bool
	RemoveParameters  bool
	Client            *http.Client
	Providers         []string
	Blacklist         map[string]struct{}
	Output            string
	JSON              bool
	URLScan           URLScan
	OTX               string
}
