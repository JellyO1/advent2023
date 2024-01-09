package day7

import "testing"

func TestNewHand_FiveOfAKind(t *testing.T) {
	hand := NewHand([]int{5, 5, 5, 5, 5}, 100)

	if hand.Type != FiveOfAKind {
		t.Fatalf("Hand type isn't a five of a kind")
	}
}

func TestNewHand_FourOfAKind(t *testing.T) {
	hand := NewHand([]int{5, 5, 5, 5, 2}, 100)

	if hand.Type != FourOfAKind {
		t.Fatalf("Hand type isn't a four of a kind")
	}
}

func TestNewHand_OrderedFullHouse(t *testing.T) {
	hand := NewHand([]int{5, 5, 5, 2, 2}, 100)

	if hand.Type != FullHouse {
		t.Fatalf("Hand type isn't a fullhouse")
	}
}
func TestNewHand_UnorderedFullHouse(t *testing.T) {
	hand := NewHand([]int{5, 2, 5, 5, 2}, 100)

	if hand.Type != FullHouse {
		t.Fatalf("Hand type isn't a fullhouse")
	}
}

func TestNewHand_ThreeOfAKind(t *testing.T) {
	hand := NewHand([]int{5, 1, 5, 5, 2}, 100)

	if hand.Type != ThreeOfAKind {
		t.Fatalf("Hand type isn't a three of a kind")
	}
}

func TestNewHand_TwoPair(t *testing.T) {
	hand := NewHand([]int{int(King), int(King), int(Six), int(Seven), int(Seven)}, 100)

	if hand.Type != TwoPair {
		t.Fatalf("Hand type isn't a two pair")
	}
}

func TestNewHand_OnePair(t *testing.T) {
	hand := NewHand([]int{1, 2, 3, 4, 2}, 100)

	if hand.Type != OnePair {
		t.Fatalf("Hand type isn't a pair")
	}
}

func TestNewHand_HighCard(t *testing.T) {
	hand := NewHand([]int{5, 4, 3, 2, 1}, 100)

	if hand.Type != HighCard {
		t.Fatalf("Hand type isn't a highcard")
	}
}

func TestOrderedByWeakest(t *testing.T) {
	hands := []Hand{
		*NewHand([]int{int(Ten), int(Five), int(Five), int(Five), int(Five)}, 0),
		*NewHand([]int{int(Three), int(Two), int(Ten), int(Three), int(King)}, 0),
		*NewHand([]int{int(Ten), int(Five), int(Five), int(Jack), int(Five)}, 0),
	}

	handsByWeakest := OrderedByWeakest(hands)

	if handsByWeakest[0].Type != OnePair && handsByWeakest[1].Type != ThreeOfAKind && handsByWeakest[2].Type != FourOfAKind {
		t.Fatalf("Hands aren't ordered properly")
	}
}

func TestOrderedByWeakest_SameKind_WinsFirstHighCard(t *testing.T) {
	hands := []Hand{
		*NewHand([]int{int(Ten), int(Five), int(Five), int(Five), int(Five)}, 0),
		*NewHand([]int{int(Four), int(Three), int(Four), int(Four), int(Four)}, 0),
	}

	handsByWeakest := OrderedByWeakest(hands)

	if handsByWeakest[0].Cards[0] != Four {
		t.Fatalf("Ordering by highcard when same type isn't correct")
	}
}

func TestRun(t *testing.T) {
	totalWinnings := Run("../input/07_test.txt")

	if totalWinnings != 6440 {
		t.Fatalf("The total winnings is %d when it should be 6440", totalWinnings)
	}
}
