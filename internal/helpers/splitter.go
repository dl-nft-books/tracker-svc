package helpers

type Interval struct {
	From, To uint64
}

func SplitIntoIntervals(from, to, interval uint64) []Interval {
	intervals := make([]Interval, 0)

	currentValue := from
	for currentValue+interval <= to {
		intervals = append(intervals, Interval{
			From: currentValue,
			To:   currentValue + interval,
		})

		currentValue += interval + 1
	}

	if currentValue <= to {
		intervals = append(intervals, Interval{
			From: currentValue,
			To:   to,
		})
	}

	return intervals
}
