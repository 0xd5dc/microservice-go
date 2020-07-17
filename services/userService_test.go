package service

import (
	model "awesomeProject/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//register a user return error if the returned object id is nil
func TestUserHandler(t *testing.T) {

	client, server := helper(UserHandler)
	defer server.Close()

	user, _ := model.CreateObject(model.User{}).(model.User)
	body, err := json.Marshal(&user)
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest("POST",
		server.URL, bytes.NewBuffer(body))

	if err != nil {
		t.Errorf("Error in creating POST request for createMatchHandler: %v",
			err)
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)

	if err != nil {
		t.Errorf("Error in POST to createMatchHandler: %v", err)
	}

	defer res.Body.Close()

	//payload, err := ioutil.ReadAll(res.Body)
	result := Response{}
	err = json.NewDecoder(res.Body).Decode(&result)
	fmt.Printf("%+v\n", result)
	//user.Id = result.Data
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	//if _, ok := res.Header["User"]; !ok {
	//	t.Error("User header is not set")
	//}
	//if _, ok := res.Header["Token"]; !ok {
	//	t.Error("Token header is not set")
	//}
	//
	//if res.StatusCode != http.StatusCreated {
	//	t.Errorf("Expected response status 201, received %s",
	//		res.Status)
	//}

	//fmt.Printf("Payload: %s", string(payload))
}

func helper(handler http.HandlerFunc) (*http.Client, *httptest.Server) {
	client := &http.Client{}
	server := httptest.NewServer(
		http.HandlerFunc(handler))
	return client, server
}

//login a user return error if the returned token is nil
func TestLoginHandler(t *testing.T) {
	//	generate access token
	//oid, _ := primitive.ObjectIDFromHex("5f10c8e7f1a937281e8c9ac7")
	//user:=model.User{Id: oid}
	t.Errorf("error")
}

//logout a user return error if the returned token is not nil
func TestLogoutHandler(t *testing.T) {
	//	deleteTest access token
	t.Errorf("error")

}

//deleteTest a user return error if the returned deleteAt is nil
func TestDeRegisterHandler(t *testing.T) {
	//	touch deleteAt
	t.Errorf("error")
}
