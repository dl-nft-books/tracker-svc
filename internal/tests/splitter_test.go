package tests

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/nft-books/contract-tracker/'/helpers"
	"testing"
)

type SplitterTest struct {
	Input struct {
		From, To, Interval uint64
	}
	Expected []helpers.Interval
}

var splitterTests = []SplitterTest{
	{
		Input:    struct{ From, To, Interval uint64 }{From: 100, To: 120, Interval: 100},
		Expected: []helpers.Interval{{100, 120}},
	},
	{
		Input:    struct{ From, To, Interval uint64 }{From: 100, To: 220, Interval: 100},
		Expected: []helpers.Interval{{100, 200}, {201, 220}},
	},
	{
		Input:    struct{ From, To, Interval uint64 }{From: 100, To: 301, Interval: 100},
		Expected: []helpers.Interval{{100, 200}, {201, 301}},
	},
	{
		Input:    struct{ From, To, Interval uint64 }{From: 100, To: 302, Interval: 100},
		Expected: []helpers.Interval{{100, 200}, {201, 301}, {302, 302}},
	},
}

func TestSplitter(t *testing.T) {
	for i, test := range splitterTests {
		actualValue := helpers.SplitIntoIntervals(test.Input.From, test.Input.To, test.Input.Interval)
		assert.Equalf(t, test.Expected, actualValue, "test %d failed", i+1)
	}
}
