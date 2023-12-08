package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Bid struct {
	Value int
	Hand  string
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

	lines := strings.Split(input, "\n")

	handMultiplier := len(lines)
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
			fiveOfAKindBids = append(fiveOfAKindBids, Bid{Value: bidValue, Hand: hand})
		} else if isFourOfAKind(cardCount) {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			fourOfAKindBids = append(fourOfAKindBids, Bid{Value: bidValue, Hand: hand})
		} else if isFullHouse(cardCount) {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			fullHouseBids = append(fullHouseBids, Bid{Value: bidValue, Hand: hand})
		} else if isThreeOfAKind(cardCount) {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			threeOfAKindBids = append(threeOfAKindBids, Bid{Value: bidValue, Hand: hand})
		} else if isTwoPair(cardCount) {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			twoPairBids = append(twoPairBids, Bid{Value: bidValue, Hand: hand})
		} else if isOnePair(cardCount) {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			onePairBids = append(onePairBids, Bid{Value: bidValue, Hand: hand})
		} else {
			bidValue, _ := strconv.Atoi(bidandhand[1])
			highCardBids = append(highCardBids, Bid{Value: bidValue, Hand: hand})
		}
	}

	// sort the bids
	//funcion that sorts the bids by this So, 33332 and 2AAAA are both four of a kind hands, but 33332 is stronger because its first card is stronger. Similarly, 77888 and 77788 are both a full house, but 77888 is stronger because its third card is stronger (and both hands have the same first and second card).

	sort.Slice(fiveOfAKindBids, func(i, j int) bool {
		return cardValues[string(fiveOfAKindBids[i].Hand[0])] > cardValues[string(fiveOfAKindBids[j].Hand[0])]
	})
	sort.Slice(fourOfAKindBids, func(i, j int) bool {
		return cardValues[string(fourOfAKindBids[i].Hand[0])] > cardValues[string(fourOfAKindBids[j].Hand[0])]
	})
	sort.Slice(fullHouseBids, func(i, j int) bool {
		return cardValues[string(fullHouseBids[i].Hand[2])] > cardValues[string(fullHouseBids[j].Hand[2])]
	})
	sort.Slice(threeOfAKindBids, func(i, j int) bool {
		return cardValues[string(threeOfAKindBids[i].Hand[2])] > cardValues[string(threeOfAKindBids[j].Hand[2])]
	})
	sort.Slice(twoPairBids, func(i, j int) bool {
		return cardValues[string(twoPairBids[i].Hand[3])] > cardValues[string(twoPairBids[j].Hand[3])]
	})
	sort.Slice(onePairBids, func(i, j int) bool {
		return cardValues[string(onePairBids[i].Hand[2])] > cardValues[string(onePairBids[j].Hand[2])]
	})
	sort.Slice(highCardBids, func(i, j int) bool {
		return cardValues[string(highCardBids[i].Hand[4])] > cardValues[string(highCardBids[j].Hand[4])]
	})

	// add the bids to the bidssorted slice
	bidssorted = append(bidssorted, fiveOfAKindBids...)
	bidssorted = append(bidssorted, fourOfAKindBids...)
	bidssorted = append(bidssorted, fullHouseBids...)
	bidssorted = append(bidssorted, threeOfAKindBids...)
	bidssorted = append(bidssorted, twoPairBids...)
	bidssorted = append(bidssorted, onePairBids...)
	bidssorted = append(bidssorted, highCardBids...)

	fmt.Println(bidssorted)
	// add up the total value of the bidskind
	totalValue := 0
	for i, bid := range bidssorted {
		totalValue += bid.Value * (handMultiplier - i)
	}
	fmt.Println(totalValue)

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
