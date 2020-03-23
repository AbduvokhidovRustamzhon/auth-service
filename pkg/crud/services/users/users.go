package users

import (
	"auth/pkg/crud/models"
	"context"
	"errors"
	"fmt"
	jwt "github.com/AbduvokhidovRustamzhon/jwt/pkg/cmd"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)


type UsersSvc struct {
	secret jwt.Secret
	pool *pgxpool.Pool
}


func NewUserSvc(secret jwt.Secret,pool *pgxpool.Pool) *UsersSvc {
	if pool == nil {
		panic(errors.New("pool can't be nil")) // <- be accurate
	}
	return &UsersSvc{secret: secret, pool: pool}
}

var ErrInvalidLoginOrPassword = errors.New("login or password is wrong")

type Payload struct {
	Id    int64    `json:"id"`
	Exp   int64    `json:"exp"`
	Roles []string `json:"roles"`
}

type RequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseDTO struct {
	Id    int64    `json:"id"`
	Token string `json:"token"`
}

type ErrorDTO struct {
	Error string `json:"error"`
}




func (service *UsersSvc) AddNewUser(ctx context.Context, model models.User) (err error) {
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't execute pool: %v",err)
		return errors.New(fmt.Sprintf("can't execute pool: %v", err))
	}
	defer conn.Release()

	log.Print("Checking login for not repeating")
	rows, err := conn.Query(ctx, "SELECT id FROM users WHERE login = $1;",&model.Login)
	if err != nil {
		return errors.New(fmt.Sprintf("can't execute a querry: %v", err))
	}
	defer rows.Close()
	log.Print("Login is repeating")
	if rows.Next() {
		return errors.New("login is repeating")
	}


	log.Printf("register a  new user ")
	password, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	_, err = conn.Exec(ctx, "INSERT INTO users(name, login, password, role) VALUES ($1, $2, $3, $4);",model.Name, model.Login,password, model.Role)
	//---------------------Id, 		 Name,       Login,       Password,       Role,        Removed


	if err != nil {
		log.Printf("can't register a  new user ")
		return errors.New(fmt.Sprintf("can't save a new user: %v ", err))
	}
	log.Printf("new user successufuly added")
	return nil
}

func (service *UsersSvc)Login(ctx context.Context, model models.User) (response ResponseDTO, err error) {

	var pass string
	var id int64
	err = service.pool.QueryRow(ctx, `SELECT id,password FROM users WHERE login = $1;`, &model.Login).Scan(&id,&pass)
	if err != nil {
		return ResponseDTO{}, ErrInvalidLoginOrPassword
	}

	err = bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(pass))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		log.Printf("Error password %s %s",model.Password,pass)
		return ResponseDTO{}, ErrInvalidLoginOrPassword
	}

	response.Token, err = jwt.Encode(Payload{Id: id,  Exp:   time.Now().Add(time.Hour).Unix(), Roles: []string{"ROLE_USER"}, 	}, service.secret)
    response.Id = id
	return response,nil
}

