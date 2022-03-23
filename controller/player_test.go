package controller_test

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MyProfileSuite struct {
	suite.Suite
	e *echo.Echo
}

func (suite *LoginSuite) TestGetMyProfileSuccess() {
	if testing.Short() {
		suite.T().Skip()
	}
}

func (suite *LoginSuite) TestGetMyProfileMissingJwt() {
	if testing.Short() {
		suite.T().Skip()
	}
}

func (suite *LoginSuite) TestGetMyProfileJwtExpire() {
	if testing.Short() {
		suite.T().Skip()
	}
}

func (suite *LoginSuite) TestGetMyProfileIdNotFound() {
	if testing.Short() {
		suite.T().Skip()
	}
}

type UpdateMyProfile struct {
	suite.Suite
	e *echo.Echo
}

func (suite *LoginSuite) TestUpdateMyProfileSuccess() {
	if testing.Short() {
		suite.T().Skip()
	}
}

func TestPlayer(t *testing.T) {
	suite.Run(t, new(MyProfileSuite))
	suite.Run(t, new(UpdateMyProfile))
}
