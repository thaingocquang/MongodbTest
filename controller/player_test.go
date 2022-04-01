package controller_test

import (
	"MongodbTest/model"
	"MongodbTest/test_helper"
	"MongodbTest/util"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MyProfileSuite ...
type MyProfileSuite struct {
	suite.Suite
	e *echo.Echo
}

// SetupSuite ...
func (suite *MyProfileSuite) SetupSuite() {
	suite.e = test_helper.InitServer()

	// create fake player
	test_helper.CreateFakePlayer()
}

// TestGetMyProfile_Success ...
func (suite *MyProfileSuite) TestGetMyProfile_Success() {
	var (
		response util.Response
		token    string
	)

	data := map[string]interface{}{
		"id": test_helper.PlayerObjID,
	}

	token = util.GenerateUserToken(data)

	// request
	req := httptest.NewRequest(http.MethodGet, "/players/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// run
	rec := test_helper.RunAndAssertHTTPOk(suite.e, req, suite.T())

	// parse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NotEqual(suite.T(), nil, response["data"])
}

// TestGetMyProfile_Fail_MissingJWT ...
func (suite *MyProfileSuite) TestGetMyProfile_Fail_MissingJWT() {
	var (
		response util.Response
	)

	// request
	req := httptest.NewRequest(http.MethodGet, "/players/me", nil)

	// run
	rec := test_helper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// parse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
}

// TestGetMyProfile_Fail_MalformedJWT ...
func (suite *MyProfileSuite) TestGetMyProfile_Fail_MalformedJWT() {
	var (
		response util.Response
		token    string
	)

	// malformed  JWT
	token = "abcxyz"

	// request
	req := httptest.NewRequest(http.MethodGet, "/players/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// run
	rec := test_helper.RunAndAssertHTTPUnauthorized(suite.e, req, suite.T())

	// parse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
}

// TestGetMyProfile_Fail_InvalidJWT ...
func (suite *MyProfileSuite) TestGetMyProfile_Fail_InvalidJWT() {
	var (
		response util.Response
		token    string
		// formed JWT but invalid
		objID, _ = primitive.ObjectIDFromHex("5f24d45125ea51bc57a8285b")
	)

	data := map[string]interface{}{
		"id": objID,
	}

	token = util.GenerateUserToken(data)

	// request
	req := httptest.NewRequest(http.MethodGet, "/players/me", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// run
	rec := test_helper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	// parse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
}

func (suite *MyProfileSuite) TestGetMyProfileIdNotFound() {
	if testing.Short() {
		suite.T().Skip()
	}
}

// UpdateMyProfile ...
type UpdateMyProfile struct {
	suite.Suite
	e *echo.Echo
}

// SetupSuite ...
func (suite *UpdateMyProfile) SetupSuite() {
	suite.e = test_helper.InitServer()

	// create fake player
	test_helper.CreateFakePlayer()
}

// TestUpdateMyProfile_Success ...
func (suite *UpdateMyProfile) TestUpdateMyProfile_Success() {
	var (
		response util.Response
		body     = model.Player{
			Name: "name updated",
		}
		token string
	)

	data := map[string]interface{}{
		"id": test_helper.PlayerObjID,
	}

	token = util.GenerateUserToken(data)

	bodyMarshal, _ := json.Marshal(body)

	// request
	req := httptest.NewRequest(http.MethodGet, "/players/me", bytes.NewReader(bodyMarshal))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := test_helper.RunAndAssertHTTPOk(suite.e, req, suite.T())

	// parse
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NotEqual(suite.T(), nil, response["data"])
}

// TestPlayer ...
func TestPlayer(t *testing.T) {
	suite.Run(t, new(MyProfileSuite))
	suite.Run(t, new(UpdateMyProfile))
}
