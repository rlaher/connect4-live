package main

import (
	"connect4/game"
	"connect4/player"
	"fmt"
)

func main() {
	var mygame game.Game
	player1 := player.Player{
		Name:   "Bob",
		Marker: "O",
		Game:   &mygame,
	}
	player2 := player.Player{
		Name:   "Joe",
		Marker: "X",
		Game:   &mygame,
	}
	for !mygame.IsComplete {
		//player one gets to go
		if !mygame.IsComplete {
			err := player1.MakeMove()
			for err != nil {
				fmt.Println(err.Error())
				err = player1.MakeMove()
			}
			mygame.IsComplete = mygame.CheckWinner()
		}
		//player2 gets to go
		if !mygame.IsComplete {
			err := player2.MakeMove()
			for err != nil {
				fmt.Println(err.Error())
				err = player2.MakeMove()
			}
			mygame.IsComplete = mygame.CheckWinner()
		}
	}
	fmt.Println("Game over!")
}
