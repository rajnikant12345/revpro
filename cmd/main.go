package main

import (
	"github.com/rajnikant12345/revpro"
	handler2 "github.com/rajnikant12345/revpro/revphandler"
	"github.com/sirupsen/logrus"
)

func getOptions( ) *revpro.Config {
	cfg := &revpro.Config{}
	op := []revpro.ConfigOpt{revpro.WithAddressOpts(":8765"),}
	handler,err := handler2.NewProxyHandler()
	if err != nil {
		panic(err)
	}
	op = append(op,revpro.WithHandler(handler))
	for _,v := range op {
		v(cfg)
	}
	return cfg
}

func main() {
	revpro.Logger = logrus.New()
	proxy := revpro.NewProxyServer(getOptions())
	proxy.Serve()
}
