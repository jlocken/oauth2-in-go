package main

import (
		"fmt"
		"log"
		"strings"
		"net/http"
		"io/ioutil"
		"golang.org/x/oauth2"
		"github.com/gorilla/mux"		
)

//HandleRoutes handles all the app routes
func initializeRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/", handleHome)
	router.HandleFunc("/login", handleLogin)
	router.HandleFunc("/callback", handleCallback)
	log.Fatal(http.ListenAndServe(":8080", router))
}


func handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != randomState {
		fmt.Fprintf(w, "Invalid state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		log.Printf("could not get response code: %s\n", err.Error())
		http.Redirect(w,r,"/", http.StatusTemporaryRedirect)
		return
	}
	
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken)

	if err !=nil {
		log.Printf("could not create get request: %s\n",err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return

	}

	defer resp.Body.Close()
	
	content, err := ioutil.ReadAll(resp.Body)

	if err !=nil {
		fmt.Printf("could not parse response: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Response: %s", content)
	
}

func handleLogin(w http.ResponseWriter, r *http.Request){
	url:= googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}


func handleHome(w http.ResponseWriter, r *http.Request){
	log.Println(strings.Split())

	var html =`
		<html>
			<body>
				<div>
					<a href="/login"> Google Log In</a>
				</div>
			</body>
		</html>
	`
	fmt.Fprintf(w, html)
}
