package app

import (
	"auth/pkg/crud/models"
	"auth/pkg/crud/services/token"
	"auth/pkg/crud/services/users"
	_ "auth/pkg/crud/services/users"
	"context"
	"fmt"
	"github.com/AbduvokhidovRustamzhon/rest/pkg/rest"
	_ "github.com/AbduvokhidovRustamzhon/rest/pkg/rest"
	"log"
	"net/http"
	"time"
)


func (receiver *server) handleNewUser() func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		userJSON := models.User{}
		err := rest.ReadJSONBody(request, &userJSON)
		if err != nil {
			log.Print("Bad JSON file for reading from body")
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}

		log.Print(userJSON)
		ctx, _ := context.WithTimeout(request.Context(),time.Second)
		log.Print("Setting context deadline")
		err = receiver.usersSvc.AddNewUser(ctx,userJSON)
		if err != nil {
			if fmt.Sprintf("%v",err) == "login is repeating" {
				writer.WriteHeader(http.StatusBadRequest)
				err := rest.WriteJSONBody(writer, &token.ErrorDTO{Error:"login is exists"})
				log.Print(err)
				return
			}
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

			writer.Header().Set("Content-Type", "application/json")
			err = rest.WriteJSONBody(writer, userJSON)
			if err != nil {
				log.Print(err)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

		http.Redirect(writer, request, "/", http.StatusFound)

	}
}


func (receiver *server) handleLogin() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		userJSON := models.User{}
		err := rest.ReadJSONBody(request, &userJSON)
		if err != nil {
			log.Print("Bad JSON file for reading from body")
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}

		log.Print(userJSON)
		ctx, _ := context.WithTimeout(request.Context(),time.Second)
		log.Print("Setting context deadline")
		tokenJSON := users.ResponseDTO{}
		tokenJSON,err = receiver.usersSvc.Login(ctx,userJSON)

		if err != nil {
			if  fmt.Sprintf("%v",err) == "login or password is wrong" {
				writer.WriteHeader(http.StatusBadRequest)
				err := rest.WriteJSONBody(writer, &token.ErrorDTO{Error:"login or password is wrong"})
				log.Print(err)
				return
			}
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		err = rest.WriteJSONBody(writer, tokenJSON)
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/", http.StatusFound)
	}
}



/*
func (receiver *server) handleUsersRemove() func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}


		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}

		user := models.User{}
		err = json.Unmarshal(body,&user)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}

		log.Print(user.Id)

		err =  receiver.usersSvc.RemoveById(user.Id)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}

		http.Redirect(writer, request, "/admin/books", http.StatusFound)
		return
	}
}






*/















/*
func (receiver *server) handleUsersList() func(http.ResponseWriter, *http.Request) {
		return func(writer http.ResponseWriter, request *http.Request) {

		list, err := receiver.usersSvc.Users(request.Context())
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		urlsJSON, err := json.Marshal(list)
		if err != nil {
			log.Print(err)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(urlsJSON)
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

func (receiver *server) handleUsersRemove() func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}


		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}

		user := models.User{}
		err = json.Unmarshal(body,&user)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}

		log.Print(user.Id)

		err =  receiver.usersSvc.RemoveById(user.Id)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}

		http.Redirect(writer, request, "/admin/books", http.StatusFound)
		return
	}
}

func (receiver *server) handleUserShow() func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}

		bookJSON := models.Book{}
		err = json.Unmarshal(body,&bookJSON)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
			return
		}
		log.Print((bookJSON))
		book, err := receiver.booksSvc.ShowById(request.Context(),int(bookJSON.Id))
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		urlsJSON, err := json.Marshal(book)
		if err != nil {
			log.Print(err)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(urlsJSON)
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

*/