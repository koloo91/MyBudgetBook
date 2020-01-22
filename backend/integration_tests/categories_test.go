package integration_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func (suite *MbbTestSuite) TestCreateCategory() {

	body := []byte(`{"name": "category_name"}`)
	request, _ := http.NewRequest("POST", "/api/categories", bytes.NewBuffer(body))
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, request)

	suite.Equal(http.StatusCreated, recorder.Code)

	var response map[string]interface{}
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	suite.Regexp(uuidRegExp, response["id"])
	suite.Equal("category_name", response["name"])
	suite.Regexp(timeRegExp, response["created"])
	suite.NotEmpty(timeRegExp, response["updated"])
}

func (suite *MbbTestSuite) TestGetCategories() {

	body := []byte(`{"name": "category_name"}`)
	postRequest, _ := http.NewRequest("POST", "/api/categories", bytes.NewBuffer(body))
	postRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	postRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(postRecorder, postRequest)

	getRequest, _ := http.NewRequest("GET", "/api/categories", nil)
	getRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	getRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(getRecorder, getRequest)

	suite.Equal(http.StatusCreated, postRecorder.Code)

	var response map[string]interface{}
	_ = json.Unmarshal(getRecorder.Body.Bytes(), &response)

	accounts := response["content"].([]interface{})
	suite.Equal(1, len(accounts))

	account := accounts[0].(map[string]interface{})
	suite.Regexp(uuidRegExp, account["id"])
	suite.Equal("category_name", account["name"])
	suite.Regexp(timeRegExp, account["created"])
	suite.NotEmpty(timeRegExp, account["updated"])
}

func (suite *MbbTestSuite) TestUpdateCategory() {

	postBody := []byte(`{"name": "category_name"}`)
	postRequest, _ := http.NewRequest("POST", "/api/categories", bytes.NewBuffer(postBody))
	postRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	postRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(postRecorder, postRequest)

	var postResponse map[string]interface{}
	_ = json.Unmarshal(postRecorder.Body.Bytes(), &postResponse)

	putBody := []byte(`{"name": "category_name updated"}`)
	putRequest, _ := http.NewRequest("PUT", fmt.Sprintf("/api/categories/%s", postResponse["id"]), bytes.NewBuffer(putBody))
	putRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	putRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(putRecorder, putRequest)

	suite.Equal(http.StatusOK, putRecorder.Code)

	getRequest, _ := http.NewRequest("GET", "/api/categories", nil)
	getRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	getRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(getRecorder, getRequest)

	suite.Equal(http.StatusCreated, postRecorder.Code)

	var response map[string]interface{}
	_ = json.Unmarshal(getRecorder.Body.Bytes(), &response)

	accounts := response["content"].([]interface{})
	suite.Equal(1, len(accounts))

	account := accounts[0].(map[string]interface{})
	suite.Regexp(uuidRegExp, account["id"])
	suite.Equal("category_name updated", account["name"])
	suite.Regexp(timeRegExp, account["created"])
	suite.NotEmpty(timeRegExp, account["updated"])
}
