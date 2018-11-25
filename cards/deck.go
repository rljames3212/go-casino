package cards

import (
	"math/rand"
	"time"
)

type Face int

const (
	ACE Face = iota + 1
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

func (f Face) String() string {
	return [...]string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}[f-1]
}

type Suit int

const (
	HEARTS Suit = iota
	CLUBS
	DIAMONDS
	SPADES
)

func (s Suit) String() string {
	return [...]string{
		"Hearts",
		"Clubs",
		"Diamonds",
		"Spades",
	}[s]
}

type Card struct {
	suit  Suit
	value int
	face  Face
}

func NewCard(s Suit, f Face) *Card {
	n := int(f)
	if n > 10 {
		n = 10
	}
	return &Card{s, n, f}
}

func (c *Card) GetSuit() Suit {
	return c.suit
}

func (c *Card) GetValue() int {
	return c.value
}

func (c *Card) GetFace() Face {
	return c.face
}

func (c *Card) String() string {
	return c.face.String() + " of " + c.suit.String()
}

type Deck []*Card

func NewDeck() *Deck {
	d := new(Deck)
	for s := 0; s < 4; s++ {
		for v := 1; v <= 13; v++ {
			c := NewCard(Suit(s), Face(v))
			*d = append(*d, c)
		}
	}
	return d
}

func (d *Deck) Shuffle() {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for i, _ := range *d {
		j := r.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

func (d *Deck) Pop() *Card {
	var c *Card
	c, *d = (*d)[len(*d)-1], (*d)[:len(*d)-1]
	return c
}

func (d *Deck) Push(c *Card) {
	*d = append(*d, c)
}
