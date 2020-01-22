package integration_tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func (suite *MbbTestSuite) TestGetBalances() {
	accountId := suite.createAccount("account", 10.0)
	categoryId := suite.createCategory("category")

	suite.createBooking("my title", -12.0, "2019-11-08T22:17:20Z", accountId, categoryId)

	request, _ := http.NewRequest("GET", "/api/balances", nil)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, request)

	suite.Equal(http.StatusOK, recorder.Code)

	var response map[string]interface{}
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	balances := response["content"].([]interface{})
	suite.Equal(1, len(balances))

	balance := balances[0].(map[string]interface{})

	suite.Equal(accountId, balance["accountId"])
	suite.Equal("account", balance["name"])
	suite.Equal(-2.0, balance["balance"])
}
