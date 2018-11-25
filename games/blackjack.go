package games

import (
	"bufio"
	"casino/cards"
	"casino/player"
	"fmt"
	"os"
	"strings"
)

type Blackjack struct {
	name    string
	deck    *cards.Deck
	dealer  *player.Player
	Players []*player.Player
}

func NewBlackjack() *Blackjack {
	bj := new(Blackjack)
	bj.name = "Blackjack"
	bj.deck = cards.NewDeck()
	bj.deck.Shuffle()
	bj.dealer = player.NewPlayer("Dealer")
	bj.dealer.AddHand()
	bj.dealer.Hands[0].AddCard(bj.deck.Pop())
	bj.dealer.Hands[0].AddCard(bj.deck.Pop())

	var err error
	bj.Players, err = createPlayers()
	if err != nil {
		fmt.Println("Could not create Blackjack game")
		panic(err)
	}
	for _, p := range bj.Players {
		p.Hands[0].AddCard(bj.deck.Pop())
		p.Hands[0].AddCard(bj.deck.Pop())
	}
	return bj
}

func (bj *Blackjack) GetName() string {
	return bj.name
}

func (bj *Blackjack) Play() error {
	err := bj.playerTurns()
	if err != nil {
		return err
	}
	dealerHand := bj.dealer.Hands[0]
	fmt.Println("Dealer Total:", calcTotal(dealerHand))
	for _, c := range *dealerHand {
		fmt.Println(c)
	}
	bj.dealerTurn()
	bj.printResults()
	return nil
}

func (bj *Blackjack) playerTurns() error {
	reader := bufio.NewReader(os.Stdin)
	for _, p := range bj.Players {
		for i, hand := range p.Hands {
			done := false
			for !done {
				total := calcTotal(hand)
				fmt.Printf("%[1]s's hand %[2]d: \n%[3]v\n%[1]s's total: %[4]d\n", p.GetName(), i, hand, total)
				fmt.Println("Dealer first card:", bj.getDealerFirstCardString())
				var err error
				done, err = bj.handleInput(reader, hand)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func calcTotal(h *player.Hand) int {
	total := 0
	highAce := false
	for _, c := range *h {
		total += c.GetValue()
		if c.GetValue() == 1 {
			// If having a high ace doesnt bust the player, give high ace
			if total+10 <= 21 {
				total += 10
				highAce = true
			}
		}
		if total > 21 && highAce {
			total -= 10
			highAce = false
		}
	}
	return total
}

func (bj *Blackjack) handleInput(r *bufio.Reader, hand *player.Hand) (bool, error) {
	fmt.Println("\n1. Hit")
	fmt.Println("2. Stand")
	fmt.Println("3. Split")
	for {
		input, err := r.ReadString('\n')
		if err != nil {
			return true, err
		}
		switch strings.TrimSuffix(input, "\n") {
		case "1":
			hand.AddCard(bj.deck.Pop())

			//If total >= 21, player busts or wins and turn is over
			if calcTotal(hand) >= 21 {
				return true, nil
			}
			return false, nil

		case "2":
			// Player chooses to stand, turn is over
			return true, nil

		case "3":
			return true, nil
		default:
			fmt.Println("Enter a valid input")
		}
	}
}

func (bj *Blackjack) dealerTurn() {
	dealerHand := bj.dealer.Hands[0]
	for total := calcTotal(dealerHand); total < 17; {
		c := bj.deck.Pop()
		dealerHand.AddCard(c)
		fmt.Println("\n", c)
		total = calcTotal(dealerHand)
		fmt.Println("Dealer Total:", total)
	}
}

func (bj *Blackjack) printResults() {
	dealerTotal := calcTotal(bj.dealer.Hands[0])
	for _, player := range bj.Players {
		for i, hand := range player.Hands {
			playerTotal := calcTotal(hand)
			if playerTotal > 21 {
				fmt.Printf("%s's hand %d Busted! %d - %d\n", player.GetName(), i+1, playerTotal, dealerTotal)
			} else if dealerTotal > 21 {
				fmt.Printf("Dealer Busted! %s's hand %d won! %d - %d\n", player.GetName(), i+1, playerTotal, dealerTotal)
			} else if playerTotal > dealerTotal {
				fmt.Printf("%s's hand %d won! %d - %d\n", player.GetName(), i+1, playerTotal, dealerTotal)
			} else if playerTotal < dealerTotal {
				fmt.Printf("%s's hand %d lost! %d - %d\n", player.GetName(), i+1, playerTotal, dealerTotal)
			} else {
				fmt.Printf("%s's hand %d Tied! %d - %d\n", player.GetName(), i+1, playerTotal, dealerTotal)
			}
		}
	}
}

func (bj *Blackjack) getDealerFirstCardString() string {
	hand := bj.dealer.Hands[0]
	firstCard := (*hand)[0]
	return firstCard.String()
}
