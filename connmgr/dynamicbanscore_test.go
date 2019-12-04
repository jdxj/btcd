// Copyright (c) 2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package connmgr

import (
	"fmt"
	"math"
	"testing"
	"time"
)

// TestDynamicBanScoreDecay tests the exponential decay implemented in
// DynamicBanScore.
func TestDynamicBanScoreDecay(t *testing.T) {
	var bs DynamicBanScore
	base := time.Now()

	r := bs.increase(100, 50, base)
	if r != 150 {
		t.Errorf("Unexpected result %d after ban score increase.", r)
	}

	r = bs.int(base.Add(time.Minute))
	if r != 125 {
		t.Errorf("Halflife check failed - %d instead of 125", r)
	}

	r = bs.int(base.Add(7 * time.Minute))
	if r != 100 {
		t.Errorf("Decay after 7m - %d instead of 100", r)
	}
}

// TestDynamicBanScoreLifetime tests that DynamicBanScore properly yields zero
// once the maximum age is reached.
func TestDynamicBanScoreLifetime(t *testing.T) {
	var bs DynamicBanScore
	base := time.Now()

	r := bs.increase(0, math.MaxUint32, base)
	r = bs.int(base.Add(Lifetime * time.Second))
	if r != 3 { // 3, not 4 due to precision loss and truncating 3.999...
		t.Errorf("Pre max age check with MaxUint32 failed - %d", r)
	}
	r = bs.int(base.Add((Lifetime + 1) * time.Second))
	if r != 0 {
		t.Errorf("Zero after max age check failed - %d instead of 0", r)
	}
}

// TestDynamicBanScore tests exported functions of DynamicBanScore. Exponential
// decay or other time based behavior is tested by other functions.
func TestDynamicBanScoreReset(t *testing.T) {
	var bs DynamicBanScore
	if bs.Int() != 0 {
		t.Errorf("Initial state is not zero.")
	}
	bs.Increase(100, 0)
	r := bs.Int()
	if r != 100 {
		t.Errorf("Unexpected result %d after ban score increase.", r)
	}
	bs.Reset()
	if bs.Int() != 0 {
		t.Errorf("Failed to reset ban score.")
	}
}

// TestPrintPrecomputedFactor 用于显示 precomputedFactor 中的值.
func TestPrintPrecomputedFactor(t *testing.T) {
	for i, v := range precomputedFactor {
		fmt.Printf("%d: %f\n", i, v)
	}
}

func TestInt(t *testing.T) {
	const count = 200
	var (
		dur0010m = 10 * time.Minute
		dur0001m = 1 * time.Minute

		dur0010s = 10 * time.Second
		dur0005s = 5 * time.Second
		dur0001s = 1 * time.Second

		dur100ms = 100 * time.Millisecond
		dur010ms = 10 * time.Millisecond
		dur001ms = 1 * time.Millisecond
	)

	printScore(count, dur0010m)
	printScore(count, dur0001m)

	printScore(count, dur0010s)
	printScore(count, dur0005s)
	printScore(count, dur0001s)

	printScore(count, dur100ms)
	printScore(count, dur010ms)
	printScore(count, dur001ms)
}

func printScore(count int, dur time.Duration) {
	var bs DynamicBanScore

	fmt.Printf("--- %s internal\n", dur)
	bs.Reset()
	base := time.Now()

	score := bs.increase(100, 50, base)
	fmt.Printf("init score: %d\n", score)
	for i := 0; i < count; i++ {
		base = base.Add(dur)
		score = bs.int(base)
		fmt.Printf("%d, ", score)
	}
	fmt.Println()
}

func TestDurSecond(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(time.Nanosecond)

	fmt.Println(t2.Sub(t1))
	fmt.Println(t2.Unix() - t1.Unix())
}
