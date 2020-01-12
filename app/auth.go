package app

import (
	"fmt"
	"net/http"
	u "PhoneBook/utils"
	"os"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"go-contacts/models"
	"context"
)
//MiddleWare Implementation - Middleware which will intercept every requqest to check the presence of JWT
//Checking for Valid JWT(JSON Web Token) and responding accordingly
//If the token is valid proceed to request
//If the token is InValid Send error

var JWTAuthentication = func(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		//Endpoints Which doesn't require Auth
		noAuth := []string{"api/user/new", "api/user/login"}
		requestPath := r.URL.Path //Current Url

		//Checking auth requirement
		for _, value:= range noAuth{
			if value == requestPath {
					next.ServeHTTP(w, r)
					return
			}
		}

		//Getting the JWT token
		response := map[string]interface{}{}
		tokenHeader := w.Header().Get("Authorization")

		//Error - If the JWT is missing
		if tokenHeader == ""{
			response = u.Message(false, "Token not found")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenSplit := strings.Split(tokenHeader, " ")
		//Check JWT Format
		if len(tokenSplit) != 2{
			response = u.Message(false, "Invalid Token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			u.Respond(w, response)
		}
		//Did not understand Part
		tokenPart := tokenSplit[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil { //Malformed token, returns with http code 403 as usual
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User %", tk.Username) //Useful for monitoring

		//context need to be understood
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!



	})
}