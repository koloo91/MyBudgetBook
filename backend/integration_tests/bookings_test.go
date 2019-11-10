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

func (suite *MbbTestSuite) TestGetBookings() {
	accountId := suite.createAccount("account", 10.0)
	categoryId := suite.createCategory("category")

	for i := 0; i < 10; i++ {
		suite.createBooking(fmt.Sprintf("%d", i), -12.0, "2019-11-08T22:17:20Z", accountId, categoryId)
	}

	getRequest, _ := http.NewRequest("GET", "/api/bookings?startDate=2019-11-01T00:00:00Z", nil)
	getRequest.SetBasicAuth(appUser, appUserPassword)
	getRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(getRecorder, getRequest)

	suite.Equal(http.StatusOK, getRecorder.Code)

	var response map[string]interface{}
	_ = json.Unmarshal(getRecorder.Body.Bytes(), &response)

	bookings := response["content"].([]interface{})
	suite.Equal(10, len(bookings))

	booking := bookings[0].(map[string]interface{})
	suite.Regexp(uuidRegExp, booking["id"])
	suite.Equal("0", booking["title"])
	suite.Equal(-12.0, booking["amount"])
	suite.Equal("2019-11-08T22:17:20Z", booking["date"])
	suite.Equal(accountId, booking["accountId"])
	suite.Equal(categoryId, booking["categoryId"])
	suite.Regexp(timeRegExp, booking["created"])
	suite.NotEmpty(timeRegExp, booking["updated"])
}

func (suite *MbbTestSuite) TestUpdateBooking() {

	accountId := suite.createAccount("account", 10.0)
	categoryId := suite.createCategory("category")
	bookingId := suite.createBooking("my title", -12.0, "2019-11-08T22:17:20Z", accountId, categoryId)

	newAccountId := suite.createAccount("account 2", 10.0)
	newCategoryId := suite.createCategory("category 2")

	putBody := []byte(fmt.Sprintf(`{"title": "booking", "amount": -12.1, "date": "2019-11-09T22:17:20Z", "accountId": "%s", "categoryId": "%s"}`, newAccountId, newCategoryId))
	putRequest, _ := http.NewRequest("PUT", fmt.Sprintf("/api/bookings/%s", bookingId), bytes.NewBuffer(putBody))
	putRequest.SetBasicAuth(appUser, appUserPassword)
	putRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(putRecorder, putRequest)

	suite.Equal(http.StatusOK, putRecorder.Code)

	getRequest, _ := http.NewRequest("GET", "/api/bookings?startDate=2019-11-01T00:00:00Z", nil)
	getRequest.SetBasicAuth(appUser, appUserPassword)
	getRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(getRecorder, getRequest)

	var response map[string]interface{}
	_ = json.Unmarshal(getRecorder.Body.Bytes(), &response)

	bookings := response["content"].([]interface{})
	suite.Equal(1, len(bookings))

	booking := bookings[0].(map[string]interface{})
	suite.Regexp(uuidRegExp, booking["id"])
	suite.Equal("booking", booking["title"])
	suite.Equal(-12.1, booking["amount"])
	suite.Equal("2019-11-09T22:17:20Z", booking["date"])
	suite.Equal(newAccountId, booking["accountId"])
	suite.Equal(newCategoryId, booking["categoryId"])
	suite.Regexp(timeRegExp, booking["created"])
	suite.NotEmpty(timeRegExp, booking["updated"])
}

func (suite *MbbTestSuite) TestDeleteBooking() {

	accountId := suite.createAccount("account", 10.0)
	categoryId := suite.createCategory("category")
	bookingId := suite.createBooking("my title", -12.0, "2019-11-08T22:17:20Z", accountId, categoryId)

	deleteRequest, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/bookings/%s", bookingId), nil)
	deleteRequest.SetBasicAuth(appUser, appUserPassword)
	deleteRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(deleteRecorder, deleteRequest)

	suite.Equal(http.StatusNoContent, deleteRecorder.Code)

	getRequest, _ := http.NewRequest("GET", "/api/bookings?startDate=2019-11-01T00:00:00Z", nil)
	getRequest.SetBasicAuth(appUser, appUserPassword)
	getRecorder := httptest.NewRecorder()
	suite.router.ServeHTTP(getRecorder, getRequest)

	suite.Equal(http.StatusOK, getRecorder.Code)

	var response map[string]interface{}
	_ = json.Unmarshal(getRecorder.Body.Bytes(), &response)

	bookings := response["content"].([]interface{})
	suite.Equal(0, len(bookings))

}
