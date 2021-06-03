package revpro

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

var Logger *logrus.Logger

// ReverseProxyServer this is the server struct in reverse proxy
type ReverseProxyServer struct {
	server *http.Server
	config *Config
}

//NewDpgProxy creating proxy inst
func NewProxyServer(cfg *Config) (ret *ReverseProxyServer) {
	// we can manipulate all the Server parameters via configuration here
	ret = &ReverseProxyServer{}
	ret.server = &http.Server{
		Addr:      cfg.address,
		TLSConfig: cfg.tlsCfg,
		Handler:   cfg.httpHandler,
	}
	ret.config = cfg
	return
}

//Serve Proxy
func (d *ReverseProxyServer) Serve() {
	Logger.Info("Server Listening on :" + d.config.address)
	d.server.ListenAndServe()
}
