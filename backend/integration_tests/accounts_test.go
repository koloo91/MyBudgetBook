package integration_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func (suite *MbbTestSuite) TestCreateAccount() {

	body := []byte(`{"name": "account_name", "startingBalance": 1200}`)
	request, _ := http.NewRequest("POST", "/api/accounts", bytes.NewBuffer(body))
	request.SetBasicAuth(appUser, appUserPassword)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, request)

	suite.Equal(http.StatusCreated, recorder.Code)

	var response map[string]interface{}
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	suite.Regexp(uuidRegExp, response["id"])
	suite.Equal("account_name", response["name"])
	suite.Equal(1200.0, response["startingBalance"])
	suite.Regexp(timeRegExp, response["created"])
	suite.Regexp(timeRegExp, response["updated"])
}

func (suite *MbbTestSuite) TestGetAccounts() {

	body := []byte(`{"name": "account_name", "startingBalance": 1200}`)
	createRequest, _ := http.NewRequest("POST", "/api/accounts", bytes.NewBuffer(body))
	createRequest.SetBasicAuth(appUser, appUserPassword)
	createRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(createRecorder, createRequest)

	getRequest, _ := http.NewRequest("GET", "/api/accounts", nil)
	getRequest.SetBasicAuth(appUser, appUserPassword)
	getRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(getRecorder, getRequest)

	suite.Equal(http.StatusCreated, createRecorder.Code)

	var response map[string]interface{}
	_ = json.Unmarshal(getRecorder.Body.Bytes(), &response)

	accounts := response["content"].([]interface{})
	suite.Equal(1, len(accounts))

	account := accounts[0].(map[string]interface{})
	suite.Regexp(uuidRegExp, account["id"])
	suite.Equal("account_name", account["name"])
	suite.Equal(1200.0, account["startingBalance"])
	suite.Regexp(timeRegExp, account["created"])
	suite.Regexp(timeRegExp, account["updated"])
}
