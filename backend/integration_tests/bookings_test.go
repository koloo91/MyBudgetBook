package integration_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func (suite *MbbTestSuite) TestCreateBooking() {
	accountId := suite.createAccount("account", 10.0)
	categoryId := suite.createCategory("category")

	body := []byte(fmt.Sprintf(`{"title": "booking", "amount": -12, "date": "2019-11-08T22:17:20Z", "accountId": "%s", "categoryId": "%s"}`, accountId, categoryId))
	request, _ := http.NewRequest("POST", "/api/bookings", bytes.NewBuffer(body))
	request.SetBasicAuth(appUser, appUserPassword)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, request)

	suite.Equal(http.StatusCreated, recorder.Code)

	var response map[string]interface{}
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	suite.Regexp(uuidRegExp, response["id"])
	suite.Equal("booking", response["title"])
	suite.Equal(-12.0, response["amount"])
	suite.Equal("2019-11-08T22:17:20Z", response["date"])
	suite.Equal(accountId, response["accountId"])
	suite.Equal(categoryId, response["categoryId"])
	suite.Regexp(timeRegExp, response["created"])
	suite.NotEmpty(timeRegExp, response["updated"])
}

//func (suite *MbbTestSuite) TestGetBookings() {
//
//	body := []byte(`{"name": "category_name"}`)
//	postRequest, _ := http.NewRequest("POST", "/api/categories", bytes.NewBuffer(body))
//	postRequest.SetBasicAuth(appUser, appUserPassword)
//	postRecorder := httptest.NewRecorder()
//	suite.router.ServeHTTP(postRecorder, postRequest)
//
//	getRequest, _ := http.NewRequest("GET", "/api/categories", nil)
//	getRequest.SetBasicAuth(appUser, appUserPassword)
//	getRecorder := httptest.NewRecorder()
//	suite.router.ServeHTTP(getRecorder, getRequest)
//
//	suite.Equal(http.StatusCreated, postRecorder.Code)
//
//	var response map[string]interface{}
//	_ = json.Unmarshal(getRecorder.Body.Bytes(), &response)
//
//	accounts := response["content"].([]interface{})
//	suite.Equal(1, len(accounts))
//
//	account := accounts[0].(map[string]interface{})
//	suite.Regexp(uuidRegExp, account["id"])
//	suite.Equal("category_name", account["name"])
//	suite.Regexp(timeRegExp, account["created"])
//	suite.NotEmpty(timeRegExp, account["updated"])
//}
//
//func (suite *MbbTestSuite) TestUpdateBooking() {
//
//	postBody := []byte(`{"name": "category_name"}`)
//	postRequest, _ := http.NewRequest("POST", "/api/categories", bytes.NewBuffer(postBody))
//	postRequest.SetBasicAuth(appUser, appUserPassword)
//	postRecorder := httptest.NewRecorder()
//	suite.router.ServeHTTP(postRecorder, postRequest)
//
//	var postResponse map[string]interface{}
//	_ = json.Unmarshal(postRecorder.Body.Bytes(), &postResponse)
//
//	putBody := []byte(`{"name": "category_name updated"}`)
//	putRequest, _ := http.NewRequest("PUT", fmt.Sprintf("/api/categories/%s", postResponse["id"]), bytes.NewBuffer(putBody))
//	putRequest.SetBasicAuth(appUser, appUserPassword)
//	putRecorder := httptest.NewRecorder()
//	suite.router.ServeHTTP(putRecorder, putRequest)
//
//	suite.Equal(http.StatusOK, putRecorder.Code)
//
//	getRequest, _ := http.NewRequest("GET", "/api/categories", nil)
//	getRequest.SetBasicAuth(appUser, appUserPassword)
//	getRecorder := httptest.NewRecorder()
//	suite.router.ServeHTTP(getRecorder, getRequest)
//
//	suite.Equal(http.StatusCreated, postRecorder.Code)
//
//	var response map[string]interface{}
//	_ = json.Unmarshal(getRecorder.Body.Bytes(), &response)
//
//	accounts := response["content"].([]interface{})
//	suite.Equal(1, len(accounts))
//
//	account := accounts[0].(map[string]interface{})
//	suite.Regexp(uuidRegExp, account["id"])
//	suite.Equal("category_name updated", account["name"])
//	suite.Regexp(timeRegExp, account["created"])
//	suite.NotEmpty(timeRegExp, account["updated"])
//}
