package cook

import (
	"strconv"
	"strings"
)

const (
	RepeatRangeOp      = "-"
	RepeatHorizontalOp = "*"
	RepeatVerticalOp   = "**"
)

func RepeatOp(value string, array *[]string) bool {

	getRange := func(looprange string) (int, int, bool) {
		t := strings.Split(looprange, RepeatRangeOp)

		start, err := strconv.Atoi(t[0])
		if err != nil {
			return 0, 0, false
		}

		stop, err := strconv.Atoi(t[1])
		if err != nil {
			return 0, 0, false
		}

		if start == stop {
			return 0, 0, false
		}

		return start, stop, true
	}

	if strings.Count(value, RepeatVerticalOp) >= 1 {
		s := strings.Split(value, RepeatVerticalOp)
		input := strings.Join(s[:len(s)-1], RepeatVerticalOp)
		last := s[len(s)-1]
		till, err := strconv.Atoi(last)
		if err == nil {
			for i := 0; i < till; i++ {
				*array = append(*array, input)
			}
			return true
		}
	}

	if strings.Count(value, RepeatHorizontalOp) >= 1 {
		s := strings.Split(value, RepeatHorizontalOp)

		input := strings.Join(s[:len(s)-1], RepeatHorizontalOp)
		last := s[len(s)-1]

		if strings.Count(last, RepeatRangeOp) == 1 {
			start, stop, chk := getRange(last)
			if !chk {
				return false
			}

			if start < stop {
				for i := start; i <= stop; i++ {
					*array = append(*array, strings.Repeat(input, i))
				}
			} else {
				for i := start; i >= stop; i-- {
					*array = append(*array, strings.Repeat(input, i))
				}
			}
			return true
		}

		times, err := strconv.Atoi(last)
		if err != nil {
			return false
		}

		*array = append(*array, strings.Repeat(input, times))

		return true
	}

	return false
}
