## Go Card Poker Game REST API Example
A RESTful API example for simple card poker game with Go

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

## API

#### /auth/register `POST` : Register player

#### /auth/login `POST` : Player login

#### /auth/admin-login `POST` : Admin login

#### /bot/
* `POST` : Create bot
* `PUT` : Update bot

#### /bot/:id `GET` : Get bot by id

#### /games/ `POST` : Play game

#### /players/me
* `GET` : Update my profile
* `PUT` : Get my profile