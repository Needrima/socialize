package controllers

import (
	"campmart/database"
	"campmart/helpers"
	"campmart/middlewares"
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl = helpers.LoadTemplate()

// redirects "/" to "/home"
func RedirectToHome() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func SubscribeToNewsLetter() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		subscriber, err := middlewares.CreateNewSubscriber(r)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		subscribersCollection := database.GetDatabaseCollection("subscribers")

		insertOneResult, err := subscribersCollection.InsertOne(context.TODO(), subscriber)
		if err != nil {
			fmt.Println("Error inserting subscriber to subscribers collection:", err)
			w.Write([]byte("something went wrong, try again later"))
			return
		}

		fmt.Println("successfull added new subscriber with id:", insertOneResult.InsertedID)
		w.Write([]byte("you have succesfully subscribed to our newsletter"))
	}
}
