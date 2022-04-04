package controller_test

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"testing"
)

// GamePlaySuite ...
type GamePlaySuite struct {
	suite.Suite
	e *echo.Echo
}

func (suite *GamePlaySuite) GamePlay_Success() {
	if testing.Short() {
		suite.T().Skip()
	}
}

// ...

// TestAuth ...
func TestGame(t *testing.T) {
	suite.Run(t, new(GamePlaySuite))
}
