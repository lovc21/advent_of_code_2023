package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Bid struct {
	Value    int
	Strength int
	Hand     string
}

func Day7task1() {
	var input string = `2345A 2
	Q2Q2Q 13
	2345J 5
	J345A 3
	32T3K 7
	T55J5 19
	KK677 11
	KTJJT 29
	QQQJA 23
	JJJJJ 31
	JAAAA 43
	AAAAJ 53
	AAAAA 59
	2AAAA 17
	2JJJJ 47
	JJJJ2 34`

	//cardsarr := []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

	lines := strings.Split(input, "\n")

	handMultiplier := len(lines)
	fmt.Println(handMultiplier)
	fiveOfAKindBids := []Bid{}
	fourOfAKindBids := []Bid{}
	fullHouseBids := []Bid{}
	threeOfAKindBids := []Bid{}
	twoPairBids := []Bid{}
	onePairBids := []Bid{}
	highCardBids := []Bid{}
	bidssorted := []Bid{}

	for _, line := range lines {
		bidandhand := strings.Fields(strings.TrimSpace(line))
		hand := bidandhand[0]
		cardCount := make(map[string]int)
		for _, card := range hand {
			cardCount[string(card)]++
		}

		if isFiveOfAKind(cardCount) {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			bidStrength := handStrength(hand)
			fiveOfAKindBids = append(fiveOfAKindBids, Bid{Value: bidValue, Strength: bidStrength, Hand: hand})
		} else if isFourOfAKind(cardCount) {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			bidStrength := handStrength(hand)
			fourOfAKindBids = append(fourOfAKindBids, Bid{Value: bidValue, Strength: bidStrength, Hand: hand})
		} else if isFullHouse(cardCount) {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			bidStrength := handStrength(hand)
			fullHouseBids = append(fullHouseBids, Bid{Value: bidValue, Strength: bidStrength, Hand: hand})
		} else if isThreeOfAKind(cardCount) {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			bidStrength := handStrength(hand)
			threeOfAKindBids = append(threeOfAKindBids, Bid{Value: bidValue, Strength: bidStrength, Hand: hand})
		} else if isTwoPair(cardCount) {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			bidStrength := handStrength(hand)
			twoPairBids = append(twoPairBids, Bid{Value: bidValue, Strength: bidStrength, Hand: hand})
		} else if isOnePair(cardCount) {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			bidStrength := handStrength(hand)
			onePairBids = append(onePairBids, Bid{Value: bidValue, Strength: bidStrength, Hand: hand})
		} else {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			bidStrength := handStrength(hand)
			highCardBids = append(highCardBids, Bid{Value: bidValue, Strength: bidStrength, Hand: hand})
		}
	}

	// sort bids by strength for fiveOfAKindBids
	sort.Slice(fiveOfAKindBids, func(i, j int) bool {
		return fiveOfAKindBids[i].Strength > fiveOfAKindBids[j].Strength
	})
	bidssorted = append(bidssorted, fiveOfAKindBids...)

	// sort bids by strength for fourOfAKindBids
	sort.Slice(fourOfAKindBids, func(i, j int) bool {
		return fourOfAKindBids[i].Strength > fourOfAKindBids[j].Strength
	})
	bidssorted = append(bidssorted, fourOfAKindBids...)

	// sort bids by strength for fullHouseBids
	sort.Slice(fullHouseBids, func(i, j int) bool {
		return fullHouseBids[i].Strength > fullHouseBids[j].Strength
	})
	bidssorted = append(bidssorted, fullHouseBids...)

	// sort bids by strength for threeOfAKindBids
	sort.Slice(threeOfAKindBids, func(i, j int) bool {
		return threeOfAKindBids[i].Strength > threeOfAKindBids[j].Strength
	})
	bidssorted = append(bidssorted, threeOfAKindBids...)

	// sort bids by strength for twoPairBids
	sort.Slice(twoPairBids, func(i, j int) bool {
		return twoPairBids[i].Strength > twoPairBids[j].Strength
	})

	bidssorted = append(bidssorted, twoPairBids...)

	// sort bids by strength for onePairBids
	sort.Slice(onePairBids, func(i, j int) bool {
		return onePairBids[i].Strength > onePairBids[j].Strength
	})
	bidssorted = append(bidssorted, onePairBids...)

	// sort bids by strength for highCardBids
	sort.Slice(highCardBids, func(i, j int) bool {
		return highCardBids[i].Strength > highCardBids[j].Strength
	})
	bidssorted = append(bidssorted, highCardBids...)

	for i, bid := range bidssorted {
		fmt.Printf("index %s, Hand: %s, Strength: %d, Value: %d\n", i, bid.Hand, bid.Strength, bid.Value)
	}

	//fmt.Println(bidssorted)
	//fmt.Println("Hello, playground day7 is working ")

}

func handStrength(hand string) int {
	cardValues := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	cards := make([]string, 0, len(hand))
	for _, card := range hand {
		cards = append(cards, string(card))
	}

	// Sort the cards by their values (highest to lowest)
	sort.Slice(cards, func(i, j int) bool {
		return cardValues[cards[i]] > cardValues[cards[j]]
	})

	strength := 0
	for i, card := range cards {
		strength += cardValues[card] * (len(cards) - i)
	}
	return strength
}

func isFiveOfAKind(cardCount map[string]int) bool {
	for _, count := range cardCount {
		if count == 5 {
			return true
		}
	}
	return false
}

func isFourOfAKind(cardCount map[string]int) bool {
	for _, count := range cardCount {
		if count == 4 {
			return true
		}
	}
	return false
}

func isFullHouse(cardCount map[string]int) bool {
	hasThree, hasTwo := false, false
	for _, count := range cardCount {
		if count == 3 {
			hasThree = true
		} else if count == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}

func isThreeOfAKind(cardCount map[string]int) bool {
	threeCount := 0
	pairCount := 0
	for _, count := range cardCount {
		if count == 3 {
			threeCount++
		} else if count == 2 {
			pairCount++
		}
	}
	return threeCount == 1 && pairCount == 0
}

func isTwoPair(cardCount map[string]int) bool {
	pairs := 0
	for _, count := range cardCount {
		if count == 2 {
			pairs++
		}
	}
	return pairs == 2
}

func isOnePair(cardCount map[string]int) bool {
	pairCount := 0
	threeCount := 0
	for _, count := range cardCount {
		if count == 2 {
			pairCount++
		} else if count == 3 {
			threeCount++
		}
	}
	return pairCount == 1 && threeCount == 0
}
