package integration_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func (suite *MbbTestSuite) TestCreateAccounts() {

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
	suite.NotEmpty(timeRegExp, response["updated"])
}
