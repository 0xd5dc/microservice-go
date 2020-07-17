package service

import (
	model "awesomeProject/models"
	"encoding/json"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("SESSION_SECRET"))
var users = map[string]string{"naren": "passme", "admin": "password"}

// Response is a representation of JSON response for JWT
type Response struct {
	Token  string      `json:"token"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
type Data struct {
	Id primitive.ObjectID `faker:"-" json:"_id,omitempty" bson:"_id,omitempty"`
}

// HealthcheckHandler returns the date and time
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	tokenString, err := request.HeaderExtractor{"access_token"}.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secretKey, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Access Denied; Please check the access token"))
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// If token is valid
		response := make(map[string]string)
		// response["user"] = claims["username"]
		response["time"] = time.Now().String()
		response["user"] = claims["username"].(string)
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
	}
}

// LoginHandler validates the user credentials
func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	if originalPassword, ok := users[username]; ok {
		if password == originalPassword {
			// Create a claims map
			claims := jwt.MapClaims{
				"username":  username,
				"ExpiresAt": 15000,
				"IssuedAt":  time.Now().Unix(),
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(secretKey)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				w.Write([]byte(err.Error()))
			}
			response := Response{Token: tokenString, Status: "success"}
			responseJSON, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)
		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "User is not found", http.StatusNotFound)
		return
	}
}
func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/getToken", getTokenHandler)
	//r.HandleFunc("/healthcheck", HealthcheckHandler)
	r.HandleFunc("/user", UserHandler)
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func UserHandler(writer http.ResponseWriter, h *http.Request) {
	//validate authentication
	switch h.Method {
	case http.MethodGet:
		//200 (OK), single customer. 404 (Not Found), if ID not found or invalid.

		//get bson from request
		//convert bson to User
		//find user and respond
		//user:=bsonToModel(user_bson)
		//GetUser(&user)

		writer.WriteHeader(http.StatusOK)
		writer.WriteHeader(http.StatusNotFound)
		errors.New("Not implemented")
	case http.MethodPost:
		//404 (Not Found), 409 (Conflict) if resource already exists..
		user := model.User{}
		if err := json.NewDecoder(h.Body).Decode(&user); err != nil {
			writer.WriteHeader(http.StatusNotFound)
		} else {
			insertedId, err := CreateUser(user)
			if err != nil {
				writer.WriteHeader(http.StatusConflict)
			} else {
				response := Response{Data: Data{Id: insertedId}}
				res, _ := json.Marshal(response)
				writer.Write(res)
				writer.WriteHeader(http.StatusCreated)
			}
		}
	case http.MethodPatch:
		//200 (OK) or 204 (No Content). 404 (Not Found), if ID not found or invalid.
		errors.New("Not implemented")
		writer.WriteHeader(http.StatusNoContent)
		writer.WriteHeader(http.StatusNotFound)
		writer.WriteHeader(http.StatusOK)

	case http.MethodPut:
		//200 (OK) or 204 (No Content). 404 (Not Found), if ID not found or invalid.
		errors.New("Not implemented")
		writer.WriteHeader(http.StatusOK)
		writer.WriteHeader(http.StatusNoContent)
		writer.WriteHeader(http.StatusNotFound)
		//res, _ := UpdateUser(filter, update)
	case http.MethodDelete:
		//200 (OK). 404 (Not Found), if ID not found or invalid.
		errors.New("Not implemented")
		writer.WriteHeader(http.StatusOK)
		writer.WriteHeader(http.StatusNotFound)
	default:
		//respond with 405 (Method Not Allowed),
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
