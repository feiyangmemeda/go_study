package main

import (
	"net/http"
)

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
}

func (s *sdkHttpServer) Start(address string) error {
	return nil
}

type adminHttpServer struct {
	Name string
}

func (a *adminHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	//TODO implement me
	panic("implement me")
}

func (a *adminHttpServer) Start(address string) error {
	//TODO implement me
	panic("implement me")
}
