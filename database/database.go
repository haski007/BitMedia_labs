package database

import (
	"log"

	"github.com/Haski007/BitMedia_labs/config"
	"github.com/globalsign/mgo"
)

// UsersCollection - users table.
var UsersCollection *mgo.Collection

// UserGamesCollection - user_games table.
var UserGamesCollection *mgo.Collection

// GamesCollection - games table.
var GamesCollection *mgo.Collection

// InitDB - initialises mongoDB with your configurations at config/config.go.
func InitDB() {
	session, err := mgo.Dial(config.DataBaseHost)
	if err != nil {
		log.Fatal(err)
	}

	UsersCollection = session.DB("BitMedia").C("users")
	GamesCollection = session.DB("BitMedia").C("games")
	UserGamesCollection = session.DB("BitMedia").C("user_games")

	if err = session.Ping(); err != nil {
		log.Fatal(err)
	}
}