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
func Play(gameBody model.GameBody, playerID string, botID string) (model.Game, error) {
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

	// to objectId
	playerObjID, _ := primitive.ObjectIDFromHex(playerID)
	botObjID, _ := primitive.ObjectIDFromHex(botID)

	// get bot by name
	bot, err := dao.GetBotByID(botObjID)
	if err != nil {
		return model.Game{}, errors.New("get bot by name error")
	}

	fmt.Println(bot)

	// get player stats
	stats, err := dao.GetStatsByPlayerID(playerObjID)
	if err != nil {
		return model.Game{}, errors.New("get player stats error")
	}

	fmt.Println(stats)

	// compare hand
	if botHand.CompareHandIsHigher(playerHand) {
		// add bot point
		botRemainPoint := bot.RemainPoints + gameBody.BetValue
		if botRemainPoint > bot.TotalPoints {
			bot.RemainPoints = bot.TotalPoints
		} else {
			bot.RemainPoints = botRemainPoint
		}

		// minus player point
		playerPoint := stats.Point - gameBody.BetValue
		if playerPoint < 0 {
			stats.Point = 0
		} else {
			stats.Point = playerPoint
		}

		fmt.Println("BOT WIN!!")

	} else {
		// minus bot point
		botRemainPoint := bot.RemainPoints - gameBody.BetValue
		//if botRemainPoint > bot.TotalPoints {
		if botRemainPoint < 0 {
			bot.RemainPoints = 0
		} else {
			bot.RemainPoints = botRemainPoint
		}

		stats.Point += gameBody.BetValue

		fmt.Println("PLAYER WIN!!")

	}

	fmt.Println("AFTER")
	fmt.Println(bot)
	fmt.Println(stats)

	err = dao.UpdateStatsByPlayerID(playerObjID, stats)
	if err != nil {
		return model.Game{}, errors.New("update stats failed")
	}

	err = dao.UpdateBotByID(botObjID, bot)
	if err != nil {
		return model.Game{}, errors.New("update stats failed")
	}

	// init game
	game := model.Game{
		ID:         primitive.NewObjectID(),
		GameNo:     dao.CountAllGame(),
		PlayerID:   playerObjID,
		BotID:      botObjID,
		WinnerID:   botObjID,
		PlayerHand: playerHand,
		BotHand:    botHand,
		BetValue:   gameBody.BetValue,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err = dao.RecordGame(game)
	if err != nil {
		return model.Game{}, errors.New("record game failed")
	}

	return game, err
}
