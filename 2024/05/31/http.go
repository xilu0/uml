package main

// Database 定义数据库结构体
type Database struct {
	DSN string
}

// NewDatabase 创建一个新的数据库实例
func NewDatabase(dsn string) *Database {
	return &Database{DSN: dsn}
}

// Server 定义HTTP服务器结构体
type Server struct {
	DB *Database
}

// NewServer 创建一个新的服务器实例
func NewServer(db *Database) *Server {
	return &Server{DB: db}
}
