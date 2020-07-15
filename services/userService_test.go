package service

import "testing"

//register a user return error if the returned object id is nil
func TestRegisterHandler(t *testing.T) {
	t.Errorf("error")
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
