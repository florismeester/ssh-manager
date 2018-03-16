package main

/* Json rest API ssh-manager
   @author Floris Meester floris@rawpacket.com */

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"gopkg.in/validator.v2"
//	"io"
//	"io/ioutil"
)


// Restapi context routes

func  restapi() {
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/keygen/{username}", keygen)
	if configuration.Tls {
		log.Fatal(http.ListenAndServeTLS(configuration.Listen, configuration.Certpath, configuration.Keypath, router))
	} else {
		log.Fatal(http.ListenAndServe(configuration.Listen, router))
	}

}

// Rest api functions

func index(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintln(w, "ssh-manager  api, usage: /keygen/<username>")
}

func keygen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
        pw := genpw()
        message := createKeyPair(username + "-id_rsa.pub", username + "-id_rsa.priv", pw)
        if err := validator.Validate(message); err != nil {
                printerr(err)
                t := err
                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(422)
                err := json.NewEncoder(w).Encode(t)
                printerr(err)

        } else {
                //t := "succes"
		message.Status = "succes"
		message.Username = username
                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(message)
        }

}

/*
func postdata(w http.ResponseWriter, r *http.Request) {

	var m Message

	// Read the body
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	printerr(err)
	
	if err := r.Body.Close(); err != nil {
		printerr(err)
	}

	// Unmarshal json in to a message struct
	fmt.Println(body)
	err = json.Unmarshal(body, &m) 
	
	if err := validator.Validate(m); err != nil {
		printerr(err)
		
		
		t := err
    		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    		w.WriteHeader(422) 
   		err := json.NewEncoder(w).Encode(t)
       		printerr(err)

	} else {
		t := "succes"
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)
    		printerr(err)

		message, err := json.Marshal(m)
		printerr(err)
		fmt.Println("Rest received:", m)
		producerchannel <- message
	}
}	
*/
