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

var DefaultSlotSetPayout = map[SlotSet]Payout{
	SlotSet{Slot0, Slot0, Slot0}: {Sum: 1, Text: "Beautiful, but all zeros. Keep it up!"},
	SlotSet{Slot1, Slot1, Slot1}: {Sum: 11, Text: "Win! But try bigger"},
	SlotSet{Slot2, Slot2, Slot2}: {Sum: 12, Text: "Not bad! Winner!"},
	SlotSet{Slot3, Slot3, Slot3}: {Sum: 200, Text: "Nice bet!"},
	SlotSet{Slot4, Slot4, Slot4}: {Sum: 14, Text: "Good result!"},
	SlotSet{Slot5, Slot5, Slot5}: {Sum: 15, Text: "Beautiful!"},
	SlotSet{Slot6, Slot6, Slot6}: {Sum: 16, Text: "Great result!"},
	SlotSet{Slot7, Slot7, Slot7}: {Sum: 600, Text: "Jackpot!!!"},
	SlotSet{Slot8, Slot8, Slot8}: {Sum: 18, Text: "Look at you, winner!"},
	SlotSet{Slot9, Slot9, Slot9}: {Sum: 19, Text: "Biggest win aside of Jackpot!"},
}

type SlotMachine struct {
	betSize     int64
	stats       *Stats
	max         *big.Int
	slotValues  []SlotValue
	payoutTable map[SlotSet]Payout
}

func NewSlotMachine(betSize int64, slotValues []SlotValue, payoutTable map[SlotSet]Payout) SlotMachine {
	m := SlotMachine{
		betSize:     betSize,
		stats:       &Stats{},
		slotValues:  slotValues,
		max:         big.NewInt(int64(len(slotValues))),
		payoutTable: payoutTable,
	}

	return m
}

func (m *SlotMachine) GetStats() Stats {
	return *m.stats
}

func (m *SlotMachine) GetPayoutRate() PayoutRate {
	wc := int64(0)
	s := int64(0)
	maxSum := int64(0)
	for _, v := range m.payoutTable {
		if maxSum < v.Sum {
			maxSum = v.Sum
		}
		s += v.Sum
		wc += 1
	}
	return PayoutRate{
		TotalCombinations:   int64(math.Pow(float64(len(m.slotValues)), float64(SlotSetSize))),
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

	pay := m.payoutTable[set]

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
}

func (m *SlotMachine) GetBetSize() int64 {
	return m.betSize
}
