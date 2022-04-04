package controller_test

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"testing"
)

// CreateBotSuite ...
type CreateBotSuite struct {
	suite.Suite
	e *echo.Echo
}

func (suite *CreateBotSuite) TestRegister_Success() {
	if testing.Short() {
		suite.T().Skip()
	}
}

// GetBotSuite ...
type GetBotSuite struct {
	suite.Suite
	e *echo.Echo
}

func (suite *GetBotSuite) GetBot_Success() {
	if testing.Short() {
		suite.T().Skip()
	}
}

// UpdateBotSuite ...
type UpdateBotSuite struct {
	suite.Suite
	e *echo.Echo
}

func (suite *UpdateBotSuite) UpdateBot_Success() {
	if testing.Short() {
		suite.T().Skip()
	}
}

// TestAuth ...
func TestBot(t *testing.T) {
	suite.Run(t, new(CreateBotSuite))
	suite.Run(t, new(GetBotSuite))
	suite.Run(t, new(UpdateBotSuite))
}
