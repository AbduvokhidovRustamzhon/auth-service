package app

import (
	"auth/pkg/crud/services/users"
	"errors"
	_ "github.com/AbduvokhidovRustamzhon/mux2/pkg/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

type server struct {
	pool          *pgxpool.Pool
	router        http.Handler
	usersSvc      *users.UsersSvc
}


func NewServer(router http.Handler, pool *pgxpool.Pool, usersSvc *users.UsersSvc) *server {
	if router == nil {
		panic(errors.New("router can't be nil"))
	}
	if pool == nil {
		panic(errors.New("pool can't be nil"))
	}
	if usersSvc == nil {
		panic(errors.New("booksSvc can't be nil"))
	}

	return &server{
		router:        router,
		pool:          pool,
		usersSvc:      usersSvc,
	}
}

func (receiver *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	receiver.router.ServeHTTP(writer, request)
}
