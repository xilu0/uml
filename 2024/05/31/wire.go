// wire.go

//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

// 初始化依赖项
func InitializeServer(dsn string) *Server {
	wire.Build(NewDatabase, NewServer)
	return &Server{}
}
