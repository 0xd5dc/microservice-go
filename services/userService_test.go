package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

//register a user return error if the returned object id is nil
func TestUserHandler(t *testing.T) {

	client := &http.Client{}
	server := httptest.NewServer(
		http.HandlerFunc(UserHandler))
	defer server.Close()

	body := []byte("")
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

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if _, ok := res.Header["User"]; !ok {
		t.Error("User header is not set")
	}
	if _, ok := res.Header["Token"]; !ok {
		t.Error("Token header is not set")
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected response status 201, received %s",
			res.Status)
	}

	//fmt.Printf("Payload: %s", string(payload))
}

//login a user return error if the returned token is nil
func TestLoginHandler(t *testing.T) {
	//	generate access token
	t.Errorf("error")
}

//logout a user return error if the returned token is not nil
func TestLogoutHandler(t *testing.T) {
	//	delete access token
	t.Errorf("error")

}

//delete a user return error if the returned deleteAt is nil
func TestDeRegisterHandler(t *testing.T) {
	//	touch deleteAt
	t.Errorf("error")
}
