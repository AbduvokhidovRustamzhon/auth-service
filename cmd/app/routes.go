package app

import (
	"github.com/AbduvokhidovRustamzhon/mux2/pkg/mux"
)

func (receiver *server) InitRoutes() {
	mux := receiver.router.(*mux.ExactMux)
	mux.POST(
		"/newUser",
		receiver.handleNewUser())
		/*logger.Logger("Registration"),*/


	mux.POST(
		"/login",
		receiver.handleLogin())
		/*logger.Logger("Autorization"),*/

	/*
	mux.GET("/admin", receiver.handleBooksList())
	mux.POST("/user", receiver.handleBooksList())
	mux.GET("/user/books", receiver.handleBooksList())
	mux.POST("/user/books", receiver.handleBooksList())

	mux.GET("/admin/books", receiver.handleBooksList())
	mux.POST("/admin/books", receiver.handleBooksList())

     //http://localhost:9999/admin/books

	mux.POST("/admin/books/save", receiver.handleBooksSave())

	mux.POST("/admin/books/remove", receiver.handleBooksRemove())

	mux.POST("/admin/book/show", receiver.handleBookShow())
*/
}
