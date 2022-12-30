package slotmachine

import (
	"crypto/rand"
	"math"
	"math/big"
)

const SlotSetSize = 3

type SlotValue int
type SlotSet [SlotSetSize]SlotValue
type Payout struct {
	Sum  int64
	Text string
}
type BetResult struct {
	SlotSet *SlotSet
	Payout  *Payout
}

type Stats struct {
	Payouts int64
	Revenue int64
}

type PayoutRate struct {
	TotalCombinations   int64
	WinningCombinations int64 // the more winning combinations, the more interesting a game is
	SumOfWinningAmounts int64 // should be less than total amount of combinations * amount of a bet
	BiggestWin          int64
}

const Slot0 = SlotValue(0)
const Slot1 = SlotValue(1)
const Slot2 = SlotValue(2)
const Slot3 = SlotValue(3)
const Slot4 = SlotValue(4)
const Slot5 = SlotValue(5)
const Slot6 = SlotValue(6)
const Slot7 = SlotValue(7)
const Slot8 = SlotValue(8)
const Slot9 = SlotValue(9)

var SlotValues = []SlotValue{
	Slot0,
	Slot1,
	Slot2,
	Slot3,
	Slot4,
	Slot5,
	Slot6,
	Slot7,
	Slot8,
	Slot9,
}

var SlotSetPayout = map[SlotSet]Payout{
	SlotSet{Slot0, Slot0, Slot0}: {Sum: 1, Text: "Beautiful, but all zeros. Keep it up!"},
	SlotSet{Slot1, Slot1, Slot1}: {Sum: 11, Text: "Win! But try bigger"},
	SlotSet{Slot2, Slot2, Slot2}: {Sum: 12, Text: "Not bad! Winner!"},
	SlotSet{Slot3, Slot3, Slot3}: {Sum: 100, Text: "Nice bet!"},
	SlotSet{Slot4, Slot4, Slot4}: {Sum: 14, Text: "Good result!"},
	SlotSet{Slot5, Slot5, Slot5}: {Sum: 15, Text: "Beautiful!"},
	SlotSet{Slot6, Slot6, Slot6}: {Sum: 16, Text: "Great result!"},
	SlotSet{Slot7, Slot7, Slot7}: {Sum: 300, Text: "Jackpot!!!"},
	SlotSet{Slot8, Slot8, Slot8}: {Sum: 18, Text: "Look at you, winner!"},
	SlotSet{Slot9, Slot9, Slot9}: {Sum: 19, Text: "Biggest win aside of Jackpot!"},

	SlotSet{Slot0, Slot1, Slot2}: {Sum: 2, Text: "Nice!"},
	SlotSet{Slot1, Slot2, Slot3}: {Sum: 2, Text: "Nice!"},
	SlotSet{Slot3, Slot2, Slot1}: {Sum: 2, Text: "Nice!"},
	SlotSet{Slot3, Slot4, Slot5}: {Sum: 2, Text: "Nice!"},
	SlotSet{Slot5, Slot4, Slot3}: {Sum: 2, Text: "Good!"},
	SlotSet{Slot6, Slot7, Slot8}: {Sum: 2, Text: "Beautiful!"},
	SlotSet{Slot8, Slot7, Slot6}: {Sum: 2, Text: "Good!"},
	SlotSet{Slot7, Slot8, Slot9}: {Sum: 2, Text: "Nice!"},
	SlotSet{Slot9, Slot8, Slot7}: {Sum: 2, Text: "Not bad!"},

	SlotSet{Slot0, Slot0, Slot1}: {Sum: 1, Text: "Double!"},
	SlotSet{Slot0, Slot0, Slot2}: {Sum: 2, Text: "Double!"},
	SlotSet{Slot0, Slot0, Slot3}: {Sum: 3, Text: "Double!"},
	SlotSet{Slot0, Slot0, Slot4}: {Sum: 4, Text: "Double!"},
	SlotSet{Slot0, Slot0, Slot5}: {Sum: 5, Text: "Double!"},
	SlotSet{Slot0, Slot0, Slot6}: {Sum: 6, Text: "Double!"},
	SlotSet{Slot0, Slot0, Slot7}: {Sum: 7, Text: "Double!"},
	SlotSet{Slot0, Slot0, Slot8}: {Sum: 8, Text: "Double!"},
	SlotSet{Slot0, Slot0, Slot9}: {Sum: 9, Text: "Double!"},

	SlotSet{Slot1, Slot1, Slot2}: {Sum: 2, Text: "Double!"},
	SlotSet{Slot1, Slot1, Slot3}: {Sum: 3, Text: "Double!"},
	SlotSet{Slot1, Slot1, Slot4}: {Sum: 4, Text: "Double!"},
	SlotSet{Slot1, Slot1, Slot5}: {Sum: 5, Text: "Double!"},
	SlotSet{Slot1, Slot1, Slot6}: {Sum: 6, Text: "Double!"},
	SlotSet{Slot1, Slot1, Slot7}: {Sum: 7, Text: "Double!"},
	SlotSet{Slot1, Slot1, Slot8}: {Sum: 8, Text: "Double!"},
	SlotSet{Slot1, Slot1, Slot9}: {Sum: 9, Text: "Double!"},

	SlotSet{Slot2, Slot2, Slot1}: {Sum: 1, Text: "Double!"},
	SlotSet{Slot2, Slot2, Slot3}: {Sum: 3, Text: "Double!"},
	SlotSet{Slot2, Slot2, Slot4}: {Sum: 4, Text: "Double!"},
	SlotSet{Slot2, Slot2, Slot5}: {Sum: 5, Text: "Double!"},
	SlotSet{Slot2, Slot2, Slot6}: {Sum: 6, Text: "Double!"},
	SlotSet{Slot2, Slot2, Slot7}: {Sum: 7, Text: "Double!"},
	SlotSet{Slot2, Slot2, Slot8}: {Sum: 8, Text: "Double!"},
	SlotSet{Slot2, Slot2, Slot9}: {Sum: 9, Text: "Double!"},

	SlotSet{Slot3, Slot3, Slot1}: {Sum: 1, Text: "Double!"},
	SlotSet{Slot3, Slot3, Slot2}: {Sum: 2, Text: "Double!"},
	SlotSet{Slot3, Slot3, Slot4}: {Sum: 4, Text: "Double!"},
	SlotSet{Slot3, Slot3, Slot5}: {Sum: 5, Text: "Double!"},
	SlotSet{Slot3, Slot3, Slot6}: {Sum: 6, Text: "Double!"},
	SlotSet{Slot3, Slot3, Slot7}: {Sum: 7, Text: "Double!"},
	SlotSet{Slot3, Slot3, Slot8}: {Sum: 8, Text: "Double!"},
	SlotSet{Slot3, Slot3, Slot9}: {Sum: 9, Text: "Double!"},

	SlotSet{Slot4, Slot4, Slot1}: {Sum: 1, Text: "Double!"},
	SlotSet{Slot4, Slot4, Slot2}: {Sum: 2, Text: "Double!"},
	SlotSet{Slot4, Slot4, Slot3}: {Sum: 3, Text: "Double!"},
	SlotSet{Slot4, Slot4, Slot5}: {Sum: 5, Text: "Double!"},
	SlotSet{Slot4, Slot4, Slot6}: {Sum: 6, Text: "Double!"},
	SlotSet{Slot4, Slot4, Slot7}: {Sum: 7, Text: "Double!"},
	SlotSet{Slot4, Slot4, Slot8}: {Sum: 8, Text: "Double!"},
	SlotSet{Slot4, Slot4, Slot9}: {Sum: 9, Text: "Double!"},

	SlotSet{Slot5, Slot5, Slot1}: {Sum: 1, Text: "Double!"},
	SlotSet{Slot5, Slot5, Slot2}: {Sum: 2, Text: "Double!"},
	SlotSet{Slot5, Slot5, Slot3}: {Sum: 3, Text: "Double!"},
	SlotSet{Slot5, Slot5, Slot4}: {Sum: 4, Text: "Double!"},
	SlotSet{Slot5, Slot5, Slot6}: {Sum: 6, Text: "Double!"},
	SlotSet{Slot5, Slot5, Slot7}: {Sum: 7, Text: "Double!"},
	SlotSet{Slot5, Slot5, Slot8}: {Sum: 8, Text: "Double!"},
	SlotSet{Slot5, Slot5, Slot9}: {Sum: 9, Text: "Double!"},

	SlotSet{Slot6, Slot6, Slot1}: {Sum: 1, Text: "Double!"},
	SlotSet{Slot6, Slot6, Slot2}: {Sum: 2, Text: "Double!"},
	SlotSet{Slot6, Slot6, Slot3}: {Sum: 3, Text: "Double!"},
	SlotSet{Slot6, Slot6, Slot4}: {Sum: 4, Text: "Double!"},
	SlotSet{Slot6, Slot6, Slot5}: {Sum: 5, Text: "Double!"},
	SlotSet{Slot6, Slot6, Slot7}: {Sum: 7, Text: "Double!"},
	SlotSet{Slot6, Slot6, Slot8}: {Sum: 8, Text: "Double!"},
	SlotSet{Slot6, Slot6, Slot9}: {Sum: 9, Text: "Double!"},

	SlotSet{Slot7, Slot7, Slot1}: {Sum: 1, Text: "Double!"},
	SlotSet{Slot7, Slot7, Slot2}: {Sum: 2, Text: "Double!"},
	SlotSet{Slot7, Slot7, Slot3}: {Sum: 3, Text: "Double!"},
	SlotSet{Slot7, Slot7, Slot4}: {Sum: 4, Text: "Double!"},
	SlotSet{Slot7, Slot7, Slot5}: {Sum: 5, Text: "Double!"},
	SlotSet{Slot7, Slot7, Slot6}: {Sum: 6, Text: "Double!"},
	SlotSet{Slot7, Slot7, Slot8}: {Sum: 8, Text: "Double!"},
	SlotSet{Slot7, Slot7, Slot9}: {Sum: 9, Text: "Double!"},

	SlotSet{Slot8, Slot8, Slot1}: {Sum: 1, Text: "Double!"},
	SlotSet{Slot8, Slot8, Slot2}: {Sum: 2, Text: "Double!"},
	SlotSet{Slot8, Slot8, Slot3}: {Sum: 3, Text: "Double!"},
	SlotSet{Slot8, Slot8, Slot4}: {Sum: 4, Text: "Double!"},
	SlotSet{Slot8, Slot8, Slot5}: {Sum: 5, Text: "Double!"},
	SlotSet{Slot8, Slot8, Slot6}: {Sum: 6, Text: "Double!"},
	SlotSet{Slot8, Slot8, Slot7}: {Sum: 7, Text: "Double!"},
	SlotSet{Slot8, Slot8, Slot9}: {Sum: 9, Text: "Double!"},

	SlotSet{Slot9, Slot9, Slot1}: {Sum: 1, Text: "Double!"},
	SlotSet{Slot9, Slot9, Slot2}: {Sum: 2, Text: "Double!"},
	SlotSet{Slot9, Slot9, Slot3}: {Sum: 3, Text: "Double!"},
	SlotSet{Slot9, Slot9, Slot4}: {Sum: 4, Text: "Double!"},
	SlotSet{Slot9, Slot9, Slot5}: {Sum: 5, Text: "Double!"},
	SlotSet{Slot9, Slot9, Slot6}: {Sum: 6, Text: "Double!"},
	SlotSet{Slot9, Slot9, Slot7}: {Sum: 7, Text: "Double!"},
	SlotSet{Slot9, Slot9, Slot8}: {Sum: 8, Text: "Double!"},
}

type SlotMachine struct {
	betSize         int64
	stats           *Stats
	winDistribution map[SlotSet]int64
	max             *big.Int
	slotValues      [10]SlotValue
}

func NewSlotMachine() SlotMachine {
	sv := [10]SlotValue{
		Slot0,
		Slot1,
		Slot2,
		Slot3,
		Slot4,
		Slot5,
		Slot6,
		Slot7,
		Slot8,
		Slot9,
	}
	m := SlotMachine{
		betSize:         1,
		stats:           &Stats{},
		winDistribution: make(map[SlotSet]int64, 100),
		slotValues:      sv,
		max:             big.NewInt(int64(len(sv))),
	}

	return m
}

func (m *SlotMachine) GetStats() Stats {
	return *m.stats
}

func (m *SlotMachine) GetWinDistribution() map[SlotSet]int64 {
	return m.winDistribution
}

func GetPayoutRate() PayoutRate {
	wc := int64(0)
	s := int64(0)
	maxSum := int64(0)
	for _, v := range SlotSetPayout {
		if maxSum < v.Sum {
			maxSum = v.Sum
		}
		s += v.Sum
		wc += 1
	}
	return PayoutRate{
		TotalCombinations:   int64(math.Pow(float64(len(SlotValues)), float64(SlotSetSize))),
		WinningCombinations: wc,
		SumOfWinningAmounts: s,
		BiggestWin:          maxSum,
	}
}

func (m *SlotMachine) spinSlot() (int64, error) {
	s, err := rand.Int(rand.Reader, m.max)
	if err != nil {
		return -1, err
	}

	return s.Int64(), nil
}

func (m *SlotMachine) BetResult() (*BetResult, error) {
	firstSlot, err := m.spinSlot()
	if err != nil {
		return nil, err
	}
	secondSlot, err := m.spinSlot()
	if err != nil {
		return nil, err
	}
	thirdSlot, err := m.spinSlot()
	if err != nil {
		return nil, err
	}
	set := SlotSet{SlotValue(firstSlot), SlotValue(secondSlot), SlotValue(thirdSlot)}
	m.stats.Revenue += m.betSize

	pay := SlotSetPayout[set]

	return &BetResult{SlotSet: &set, Payout: &pay}, nil
}

// bet result can be discarded, this hack needed to keep stats relevant
func (m *SlotMachine) ApplyBetResultToStats(br *BetResult) {
	if br.Payout == nil {
		return
	}

	m.stats.Payouts += br.Payout.Sum
	if br.SlotSet == nil {
		return
	}
	m.winDistribution[*br.SlotSet] += 1
}

func (m *SlotMachine) GetBetSize() int64 {
	return m.betSize
}
