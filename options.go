package revpro

import (
	"crypto/tls"
	"net/http"
)

type Config struct {
	address     string
	httpHandler http.Handler
	tlsCfg      *tls.Config
}

// ConfigOpt configures a specific feature of fixture instance creation
type ConfigOpt func(*Config)

// WithCfgSTFConfigOpts sets config opts for server address
func WithAddressOpts(opt string) ConfigOpt {
	return func(co *Config) {
		co.address = opt
	}
}

// WithTLSOpts sets config opts for TLS
func WithTLSOpts(opt *tls.Config) ConfigOpt {
	return func(co *Config) {
		co.tlsCfg = opt
	}
}

// WithTLSOpts sets config opts for TLS
func WithHandler(opt http.Handler) ConfigOpt {
	return func(co *Config) {
		co.httpHandler = opt
	}
}
