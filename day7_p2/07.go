package day7_p2

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card int

const (
	Joker Card = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Queen
	King
	Ace
)

func (c Card) String() string {
	switch c {
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "T"
	case Joker:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		panic("unhandled default case")
	}
}

func ToCard(card rune) (Card, error) {
	switch card {
	case '2':
		return Two, nil
	case '3':
		return Three, nil
	case '4':
		return Four, nil
	case '5':
		return Five, nil
	case '6':
		return Six, nil
	case '7':
		return Seven, nil
	case '8':
		return Eight, nil
	case '9':
		return Nine, nil
	case 'T':
		return Ten, nil
	case 'J':
		return Joker, nil
	case 'Q':
		return Queen, nil
	case 'K':
		return King, nil
	case 'A':
		return Ace, nil
	}

	return 0, fmt.Errorf(`failed to convert rune "%c" into card`, card)
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	// Unordered list
	Cards [5]Card
	Bid   int
	Type  HandType
}

func NewHand(cards []int, bid int) *Hand {
	handType := HighCard

	var (
		cardTimes = make(map[int]int)
		enumCards [5]Card
	)

	highestCardIndex := int(Two)
	for i, card := range cards[:5] {
		cardTimes[card]++

		// keep the highest most repeated card that isn't the wildcard
		if card != int(Joker) && (cardTimes[card] > cardTimes[highestCardIndex] ||
			(cardTimes[card] == cardTimes[highestCardIndex] && card > highestCardIndex)) {
			highestCardIndex = card
		}

		enumCards[i] = Card(card)
	}

	// there's jokers but it's not a five of a kind of jokers
	if cardTimes[int(Joker)] > 0 && cardTimes[int(Joker)] < 5 {
		cardTimes[highestCardIndex] += cardTimes[int(Joker)]
		cardTimes[int(Joker)] = 0
	}

	for _, times := range cardTimes {
		if times == 2 && handType == ThreeOfAKind || times == 3 && handType == OnePair {
			handType = FullHouse
			break
		}

		if times == 2 && handType == OnePair {
			handType = TwoPair
			break
		}

		if times == 5 {
			handType = FiveOfAKind
			break
		}

		if times == 4 {
			handType = FourOfAKind
			break
		}

		if times == 3 {
			handType = ThreeOfAKind
			continue
		}

		if times == 2 {
			handType = OnePair
			continue
		}
	}

	return &Hand{
		Cards: enumCards,
		Bid:   bid,
		Type:  handType,
	}
}

func OrderedByWeakest(hands []Hand) (orderedHands []Hand) {
	orderedHands = make([]Hand, len(hands))
	copy(orderedHands, hands)
	slices.SortFunc(orderedHands, func(a, b Hand) int {
		v := cmp.Compare(a.Type, b.Type)
		if v == 0 {
			for i := 0; i < len(a.Cards); i++ {
				v = cmp.Compare(a.Cards[i], b.Cards[i])
				if v != 0 {
					// found strongest
					break
				}
			}
		}

		return v
	})

	return orderedHands
}

func inputFromFile(path string) ([]Hand, error) {
	fs, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	hands := make([]Hand, 0)
	scanner := bufio.NewScanner(fs)

	for scanner.Scan() {
		ln := scanner.Text()
		splitLine := strings.Split(ln, " ")
		cardsStr := splitLine[0]
		betStr := splitLine[1]

		cards := make([]int, 5)
		for i, cardRune := range cardsStr {
			card, err := ToCard(cardRune)
			if err != nil {
				return nil, err
			}

			cards[i] = int(card)
		}

		bet, err := strconv.Atoi(strings.TrimSpace(betStr))
		if err != nil {
			return nil, err
		}

		hands = append(hands, *NewHand(cards, bet))
	}

	return hands, nil
}

func Run(inputFilePath string) int {
	hands, err := inputFromFile(inputFilePath)
	if err != nil {
		panic(err)
	}

	handsByWeakest := OrderedByWeakest(hands)

	//fmt.Printf("%v\n", handsByWeakest)

	total := 0
	for i := 1; i <= len(handsByWeakest); i++ {
		hand := handsByWeakest[i-1]
		total += hand.Bid * i
	}

	return total
}
