package app

import (
	"fmt"
	"net/http"

	// local module
	"../controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	fmt.Printf("Inside app package.\n")

	var err error
	if err = http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
