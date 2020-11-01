package main

import (
	"log"
	"os"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/joho/godotenv"	
)

var (
	googleOauthConfig *oauth2.Config
	//TODO: randomize
	randomState = "random"
)

//Helper function to read an environment or return a default value
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}


func main(){
	
	godotenv.Load()

	log.Println("Initiallizing auth app....")
	initApp()
	log.Println("starting auth app....")
	initializeRoutes()

}




func initApp(){

	googleOauthConfig = &oauth2.Config{
		RedirectURL:"http://localhost:8080/callback",
		ClientID: getEnv("GOOGLE_CLIENT_ID","default GOOGLE_CLIENT_ID"),
		ClientSecret: getEnv("GOOGLE_CLIENT_SECRET", "default GOOGLE_CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email","https://www.googleapis.com/auth/userinfo.profile" },
		Endpoint: google.Endpoint,
	}

}



