## Go Card Poker Game REST API Example
A RESTful API example for simple card poker game with Go

## API

#### /auth/register `POST` : Register player

#### /auth/login `POST` : Player login

#### /auth/admin-login `POST` : Admin login

#### /bots
* `POST` : Create bot (admin)
* `PUT` : Update bot (admin)

#### /bots/configs `POST` : Reset bot point when low point

#### /bots/:id `GET` : Get bot by id (admin)

#### /bots `GET` : Get list all bot (admin, player)

#### /games 
* `POST` : Play game
* `GET` : Get recent games

#### /players/me
* `GET` : Update my profile
* `PUT` : Get my profile

#### /players/password `PATCH` : update player password

#### /players
* `DELETE` : Delete player (admin)
* `GET` : Get all player (admin)

#### /players/:id `GET` : Get player by ID (admin)

#### / players/page=&per_page= `GET` : Get all player, 10 player per page (admin)

## Structure
```
├── config
|   ├── config.go
|   ├── define.go
|   └── env_keys.go
|
├── controller
│   ├── auth.go
|   ├── auth_test.go
|   ├── bot.go
|   ├── game.go
|   ├── player.go
|   └── player_test.go
|
├── dao
│   ├── admin.go
|   ├── bot.go
|   ├── game.go
|   ├── player.go
|   └── stats.go
|
├── middleware
│   ├── admin.go
|   ├── bot.go
|   ├── game.go
|   └── validator.go
|
├── model
│   ├── admin.go
|   ├── bot.go
|   ├── card.go
|   ├── game.go
|   ├── player.go
|   └── stats.go
|
├── module
|   └── database
|       ├── collection.go
|       └── database.go
|
├── route
│   ├── auth.go
|   ├── bot.go
|   ├── game.go
|   ├── player.go
|   └── route.go
|
├── service
│   ├── auth.go
|   ├── bot.go
|   ├── game.go
|   └── player.go
|
├── test_helper
│   ├── database.go
|   ├── init.go
|   └── run_http.go
|
├── util
│   ├── jwt.go
|   └── response.go
|
└── main.go
```