package main

import "casino/games"

func main() {
	bj := games.NewBlackjack()
	_ = bj.Play()
}
