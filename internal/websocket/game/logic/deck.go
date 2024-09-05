package logic

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Suit int

// 任何类型只要实现了这个接口（即具有 String() string 方法），
// 它就满足 fmt.Stringer 接口。在使用 fmt.Println 或 fmt.Sprintf 打印该类型的值时，fmt 包会自动调用 String 方法。
func (suit Suit) String() string {
	switch suit {
	case Spades:
		return "Spades"
	case Harts:
		return "Harts"
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	default:
		panic("invalid card suit")
	}
}

const (
	Spades Suit = iota
	Harts
	Diamonds
	Clubs
)

// 卡牌
type Card struct {
	Suit  Suit
	Value int
}

func (card Card) String() string {
	value := strconv.Itoa(card.Value)
	if card.Value == 1 {
		value = "ACE"
	}
	return fmt.Sprintf("%s of %s %s", value, card.Suit, suitToUnicode(card.Suit))
}

func NewCard(suit Suit, value int) Card {
	if value > 13 {
		panic("the value of cannot be higher then 13")
	}
	return Card{
		Suit:  suit,
		Value: value,
	}
}

type Deck [52]Card

func (d *Deck) DealCard(numPlayers, cardsPerPlayer int) [][]Card {
	players := make([][]Card, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = make([]Card, cardsPerPlayer)
		for j := 0; j < cardsPerPlayer; j++ {
			players[i][j] = d[i*cardsPerPlayer+j]
		}
	}
	return players
}

func New() Deck {
	var (
		numSuits = 4
		numCards = 13
		deck     = [52]Card{}
	)
	k := 0
	for i := 0; i < numSuits; i++ {
		for j := 0; j < numCards; j++ {
			deck[k] = NewCard(Suit(i), j+1)
			k++
		}
	}
	return shuffle(deck)
}

// 洗牌
func shuffle(deck Deck) Deck {
	for i := 0; i < len(deck); i++ {
		p := rand.Intn(i + 1)
		deck[i], deck[p] = deck[p], deck[i]
	}
	return deck
}

func suitToUnicode(s Suit) string {
	switch s {
	case Spades:
		return "♠"
	case Harts:
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		panic("错误的 suit")
	}
}
