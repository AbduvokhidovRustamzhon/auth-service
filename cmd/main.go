package main

import (
	"auth/cmd/app"
	"auth/pkg/crud/services/users"
	"context"
	"flag"
	"github.com/AbduvokhidovRustamzhon/mux2/pkg/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	hostF = flag.String("host", "", "Server host")
	portF = flag.String("port", "", "Server port")
	dsnF  = flag.String("dsn", "", "Postgres DSN")
	eHost = "HOST"
	ePort = "PORT"
	EDSN  = "DATABASE_URL"
)

func main() {
	flag.Parse()

	host, ok := FlagOrEnv(*hostF, eHost)
	if !ok {
		log.Panic("can't get host")
	}

	port, ok := FlagOrEnv(*portF, ePort)
	if !ok {
		log.Panic("can't get port")
	}

	log.Println("set address to connect")
	dsn, ok := FlagOrEnv(*dsnF, EDSN)
	if !ok {
		log.Panic("can't get dsn")
	}

	addr := net.JoinHostPort(host, port)
	log.Println(host,port)
	log.Printf("try start server on: %s, dbUrl: %s", addr, dsn)
	start(addr, dsn)
	log.Printf("server success on: %s, dbUrl: %s", addr, dsn)
}

func start (addr string, dsn string) {
	router := mux.NewExactMux()

	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	userSvc := users.NewUserSvc([]byte("secret"), pool)
	server := app.NewServer(router,pool,userSvc)
    log.Print("Setting deadline for context")
	ctx, _ := context.WithTimeout(context.Background(),time.Second)
	conn, err := pool.Acquire(ctx)
	if err != nil {
		panic(err)
	}

	log.Print("Recreating users-table")
	_, err = conn.Exec(ctx,createTable)
	if err != nil {
		panic(err)
	}

	log.Print("Starting routes")
	server.InitRoutes()
	panic(http.ListenAndServe(addr, server))
}

func FlagOrEnv(flag string, envKey string) (string, bool) {
	if flag != "" {
		return flag, true
	}
	return os.LookupEnv(envKey)
}