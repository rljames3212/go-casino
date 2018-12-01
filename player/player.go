package player

import (
	"casino/cards"
)

//Represents a hand of cards
type Hand []*cards.Card

func NewHand() *Hand {
	return new(Hand)
}

//Add a card to the hand that calls this method
func (h *Hand) AddCard(c *cards.Card) {
	*h = append(*h, c)
}

func (h *Hand) String() string {
	var s string
	for i, c := range *h {
		s += c.String()
		if i != len(*h)-1 {
			s += "\n"
		}
	}
	return s
}

//Represents a Player in a card game
type Player struct {
	name  string
	Hands []*Hand
}

//NewPlayer: Creates a new player with a given name and empty hand
func NewPlayer(name string) *Player {
	p := new(Player)
	p.name = name
	p.AddHand()
	return p
}

//GetName: Return the name of the Player
func (p *Player) GetName() string {
	return p.name
}

func (p *Player) AddHand() {
	p.Hands = append(p.Hands, NewHand())
}

func (p *Player) String() string {
	s := p.name + ":\n"
	for i, hand := range p.Hands {
		s += "[" + hand.String() + "]"
		if i != len(p.Hands)-1 {
			s += ",\n"
		}
	}
	return s
}
