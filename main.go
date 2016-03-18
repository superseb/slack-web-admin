package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/superseb/slack"
)

type Specification struct {
	SlackToken string
}

func showCurrentUsers(w http.ResponseWriter, r *http.Request) {
	var s Specification
	err := envconfig.Process("slackwebadmin", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	api := slack.New(s.SlackToken)
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// api.SetDebug(true)
	users, err := api.GetUsers()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, user := range users {
		fmt.Fprintf(w, "ID: %s, Name: %s\n", user.ID, user.Name)
	}
}
func inviteUsers(w http.ResponseWriter, r *http.Request) {
	var s Specification
	err := envconfig.Process("slackwebadmin", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	api := slack.New(s.SlackToken)
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// api.SetDebug(true)
	email := r.URL.Path[len("/invite/"):]
	err = api.SendInvite(email)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

}

func main() {
	http.HandleFunc("/", showCurrentUsers)   // set router
	http.HandleFunc("/invite/", inviteUsers) // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
