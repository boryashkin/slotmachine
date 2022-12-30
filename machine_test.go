package slotmachine

import (
	"testing"
)

func TestPayotRate(t *testing.T) {
	m := NewSlotMachine()
	tries := 1000
	for i := 0; i < tries; i++ {
		br, err := m.BetResult()
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		m.ApplyBetResultToStats(br)
	}
	s := m.GetStats()
	balance := s.Revenue - s.Payouts

	t.Logf("[%d games] balance: %d, revenue: %d, payouts: %d", tries, balance, s.Revenue, s.Payouts)
}

func TestWinRate(t *testing.T) {
	m := NewSlotMachine()
	s := GetPayoutRate()
	totalInputs := s.TotalCombinations * m.GetBetSize()
	totalOutputs := s.SumOfWinningAmounts
	if totalInputs < totalOutputs {
		t.Errorf("Total bet sum should not be less than total payout sum, now: %d / %d", totalInputs, totalOutputs)
		t.FailNow()
	}
}
