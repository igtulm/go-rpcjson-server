package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
)

func initConfig() *Config {
	cfg := NewConfig()
	cfg.Init()
	return cfg
}

func initDb(cfg *Config) IDbConnector {
	db := NewDbConnector(cfg.DB.Host, cfg.DB.Port, cfg.DB.Database, cfg.DB.User, cfg.DB.Password)
	db.Init()
	return db
}

func main() {
	cfg := initConfig()
	db := initDb(cfg)

	user := NewUserRepository(db)
	userRpc := NewUserService(user)

	rpcServer := rpc.NewServer()

	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	rpcServer.RegisterService(userRpc, "user")

	router := mux.NewRouter()
	router.Handle("/rpc", rpcServer)

	log.Printf("RPC-JSON Server is started on %s\n", cfg.Server.Conn)
	http.ListenAndServe(cfg.Server.Conn, router)
}
