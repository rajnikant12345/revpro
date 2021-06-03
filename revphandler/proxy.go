package revphandler

import (
	"github.com/rajnikant12345/revpro"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type MyProxy struct {
	ReverseProxy *httputil.ReverseProxy
	ProxyURL string
}

type ProxyHandler struct {
	Proxies map[string]*MyProxy
}

//ServeHTTP to serve http request
func (h *ProxyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	revpro.Logger.Info("ServeHTTP")
	p := h.Proxies[req.URL.Path]
	req.URL,_ = url.Parse(p.ProxyURL)
	req.Host = req.URL.Host
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.URL.Path = ""
	req.URL.RawQuery = ""
	p.ReverseProxy.ServeHTTP(res,req)
}

func NewProxyHandler() (http.Handler,error) {

	p := &ProxyHandler{}

	p.Proxies = make(map[string]*MyProxy)

	googleURL := "https://google.com"

	yahooURL := "https://yahoo.com"

	///////// forward to google ///////////////////////
	url, _ := url.Parse(googleURL)
	p.Proxies["/path1"] = &MyProxy {
		httputil.NewSingleHostReverseProxy(url),
		googleURL,
	}
	////////////////////////////////////////////

	///////// forward to yahoo ///////////////////////
	url, _ = url.Parse(yahooURL)
	p.Proxies["/path2"] = &MyProxy {
		httputil.NewSingleHostReverseProxy(url),
		yahooURL,
	}
	////////////////////////////////////////////



	return p,nil
}
