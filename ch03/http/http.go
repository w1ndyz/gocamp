package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type HttpServer struct {
	s       *http.Server
	handler http.Handler
	ctx     context.Context
}

func NewServer(address string, handler http.Handler) *HttpServer {
	hs := &HttpServer{ctx: context.Background()}
	hs.s = &http.Server{
		Addr:         address,
		Handler:      handler,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	}
	return hs
}

func (hs *HttpServer) Start() error {
	fmt.Println("Start HttpServer...")
	return hs.s.ListenAndServe()
}

func (hs *HttpServer) Stop() error {
	_ = hs.s.Shutdown(hs.ctx)
	return fmt.Errorf("Stop HttpServer")
}
