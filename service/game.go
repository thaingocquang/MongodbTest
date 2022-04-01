package service

import (
	"MongodbTest/dao"
	"MongodbTest/model"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Play ...
func Play(gameBody model.GameBody, playerID string) error {
	// empty deck card
	var deckCard model.DeckCard

	// init deck card
	deckCard.Init()

	// shuffle deck card
	deckCard.Shuffle()

	// deal card
	var botHand, playerHand model.Hand
	botHand.Cards = deckCard[0:3]
	playerHand.Cards = deckCard[3:6]

	// find max card
	botHand.InitMaxCard()
	playerHand.InitMaxCard()

	// get bot by name
	bot, err := dao.GetBotByName(gameBody.BotName)
	if err != nil {
		return errors.New("get bot by name error")
	}

	// compare hand
	if botHand.CompareHandIsHigher(playerHand) {
		// add bot point
		updateBot := model.Bot{RemainPoints: bot.RemainPoints + gameBody.BetValue}
		err := dao.UpdateBotByName(bot.Name, updateBot)
		if err != nil {
			return errors.New("update bot failed")
		}
		// minus player point
		stats, err := dao.GetStatsByPlayerID(playerID)
		if err != nil {
			return errors.New("get stats failed")
		}

		updateStats := model.Stats{Point: stats.Point - gameBody.BetValue}
		err = dao.UpdateStatsByPlayerID(playerID, updateStats)
		if err != nil {
			return errors.New("update stats failed")
		}

	} else {
	}

	// init game
	var game model.Game

	game.ID = primitive.NewObjectID()
	game.GameNo = 1
	objID, _ := primitive.ObjectIDFromHex(playerID)
	game.PlayerID = objID
	game.BotID = bot.ID
	// winner ID
	game.WinnerID = bot.ID
	game.PlayerHand.Cards = deckCard[0:3]
	game.BotHand.Cards = deckCard[3:6]
	game.BetValue = gameBody.BetValue
	game.CreatedAt = time.Now()
	game.UpdatedAt = time.Now()

	if game.BotHand.CompareHandIsHigher(game.PlayerHand) {
		fmt.Println("BOT win!!")
	} else {
		fmt.Println("PLAYER win!!")
	}

	fmt.Println("BOT hand:", game.BotHand, " + SUM rank:", game.BotHand.SumRank())
	fmt.Println("PLAYER hand:", game.PlayerHand, " + SUM rank:", game.PlayerHand.SumRank())

	err = dao.RecordGame(game)
	if err != nil {
		return errors.New("record game failed")
	}
	// deal card

	// minus point if who win

	// record (point)

	return err
}
