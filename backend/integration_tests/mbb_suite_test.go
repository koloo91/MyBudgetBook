package integration_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/koloo91/controller"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	connectionString = "postgres://postgres:@localhost/postgres?sslmode=disable"
	appUser          = "kolo"
	appUserPassword  = "Pass00"

	uuidRegExp = `[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}`
	timeRegExp = `[\d]{4}-[\d]{2}-[\d]{2}T[\d]{2}:[\d]{2}:[\d]{2}.*`
)

type MbbTestSuite struct {
	suite.Suite
	db     *gorm.DB
	router *gin.Engine
}

func (suite *MbbTestSuite) SetupSuite() {
	log.Println("Setup suite")

	//db, err := sql.Open("postgres", connectionString)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	driver, _ := postgres.WithInstance(db.DB(), &postgres.Config{})
	migrator, _ := migrate.NewWithDatabaseInstance("file://../migrations", "postgres", driver)
	if err := migrator.Up(); err != nil {
		log.Println(err.Error())
	}

	suite.db = db
	suite.router = controller.SetupRoutes(db, appUser, appUserPassword)
}

func (suite *MbbTestSuite) SetupTest() {
	log.Println("Tear down test")

	_, _ = suite.db.DB().Exec("DELETE FROM bookings;")
	_, _ = suite.db.DB().Exec("DELETE FROM accounts;")
	_, _ = suite.db.DB().Exec("DELETE FROM categories;")
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(MbbTestSuite))
}

func (suite *MbbTestSuite) TestBasicAuthentication() {

	request, _ := http.NewRequest("GET", "/api/ping", nil)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, request)

	suite.Equal(http.StatusUnauthorized, recorder.Code)
}

func (suite *MbbTestSuite) createAccount(name string, startingBalance float64) string {
	body := []byte(fmt.Sprintf(`{"name": "%s", "startingBalance": %f}`, name, startingBalance))
	request, _ := http.NewRequest("POST", "/api/accounts", bytes.NewBuffer(body))
	request.SetBasicAuth(appUser, appUserPassword)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, request)

	suite.Equal(http.StatusCreated, recorder.Code)

	var response map[string]interface{}
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	return response["id"].(string)
}

func (suite *MbbTestSuite) createCategory(name string) string {
	body := []byte(fmt.Sprintf(`{"name": "%s"}`, name))
	request, _ := http.NewRequest("POST", "/api/categories", bytes.NewBuffer(body))
	request.SetBasicAuth(appUser, appUserPassword)
	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, request)

	suite.Equal(http.StatusCreated, recorder.Code)

	var response map[string]interface{}
	_ = json.Unmarshal(recorder.Body.Bytes(), &response)

	return response["id"].(string)
}
