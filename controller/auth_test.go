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
	"net/http"
	"testing"
)

// RegisterSuite ...
type RegisterSuite struct {
	suite.Suite
	e *echo.Echo
}

// SetupSuite ...
func (suite *RegisterSuite) SetupSuite() {
	suite.e = test_helper.InitServer()

}

// TestRegisterSuccess ...
func (suite *RegisterSuite) TestRegister_Success() {
	var (
		registerBody = model.Player{
			Name:     "quang",
			Email:    "thaingocquang147@gmail.com",
			Password: "123456",
		}
		response util.Response
	)

	bodyMarshal, _ := json.Marshal(registerBody)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(bodyMarshal))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := test_helper.RunAndAssertHTTPOk(suite.e, req, suite.T())

	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
}

// TestEmailRequired ...
func (suite *RegisterSuite) TestRegister_Fail_EmailRequired() {
	var (
		registerBody = model.Player{
			Name:     "quang",
			Email:    "",
			Password: "123456",
		}
		response util.Response
	)

	bodyMarshal, _ := json.Marshal(registerBody)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(bodyMarshal))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := test_helper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
}

// TestEmailMissingAtSign ...
func (suite *RegisterSuite) TestRegister_Fail_EmailMissingAtSign() {
	var (
		registerBody = model.Player{
			Name:     "quang",
			Email:    "thaingocquang147gmail.com",
			Password: "123456",
		}
		response util.Response
	)

	bodyMarshal, _ := json.Marshal(registerBody)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(bodyMarshal))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := test_helper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
}

// TestPasswordRequired ...
func (suite *RegisterSuite) TestRegister_Fail_PasswordRequired() {
	var (
		registerBody = model.Player{
			Name:     "quang",
			Email:    "thaingocquang147@gmail.com",
			Password: "",
		}
		response util.Response
	)

	bodyMarshal, _ := json.Marshal(registerBody)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewReader(bodyMarshal))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := test_helper.RunAndAssertHTTPBadRequest(suite.e, req, suite.T())

	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(suite.T(), nil, response["data"])
}

func (suite *RegisterSuite) TestRegister_Fail_PasswordAtLeast8Chars() {
	if testing.Short() {
		suite.T().Skip()
	}
}

func (suite *RegisterSuite) TestRegister_Fail_PasswordMaximum256Chars() {
	if testing.Short() {
		suite.T().Skip()
	}
}

func (suite *RegisterSuite) TestRegister_Fail_PasswordWithoutNumber() {
	if testing.Short() {
		suite.T().Skip()
	}
}

func (suite *RegisterSuite) TestRegister_Fail_PasswordSpecialChar() {
	if testing.Short() {
		suite.T().Skip()
	}
}

// LoginSuite ...
type LoginSuite struct {
	suite.Suite
	e *echo.Echo
}

// SetupSuite ...
func (suite *LoginSuite) SetupSuite() {
	suite.e = test_helper.InitServer()
	test_helper.CreateFakePlayer()
}

// TestRegisterSuccess ...
func (suite *LoginSuite) TestLoginSuccess() {
	var (
		registerBody = model.Player{
			Email:    "fake@gmail.com",
			Password: "123456",
		}
		response util.Response
	)

	bodyMarshal, _ := json.Marshal(registerBody)

	// request
	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(bodyMarshal))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// run
	rec := test_helper.RunAndAssertHTTPOk(suite.e, req, suite.T())

	_ = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NotEqual(suite.T(), nil, response["data"])
}

// TestAuth ...
func TestAuth(t *testing.T) {
	suite.Run(t, new(RegisterSuite))
	suite.Run(t, new(LoginSuite))
}
