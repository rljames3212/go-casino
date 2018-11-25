package games

import (
	"bufio"
	"casino/player"
	"fmt"
	"os"
)

func createPlayers() ([]*player.Player, error) {
	players := make([]*player.Player, 0)
	atLeastOnePlayer := false
	r := bufio.NewReader(os.Stdin)
	for i := 0; ; i++ {
		fmt.Printf("Enter Player %d's name (Press Enter if done): ", i+1)
		name, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		// If name is empty and at least 1 player, exit loop
		if name == "\n" && atLeastOnePlayer {
			break
		}
		if name != "\n" {
			players = append(players, player.NewPlayer(name[:len(name)-1]))
			atLeastOnePlayer = true
		}
	}
	return players, nil
}

type Game interface {
	Play() error
	GetName() string
}

func Execute(g Game) {
	fmt.Printf("Starting %s game", g.GetName())
	g.Play()
}
