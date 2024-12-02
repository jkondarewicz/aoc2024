package solutions

import (
	"fmt"
	"strconv"

	aocmath "github.comjkondarewicz/aoc2024/pkg/math"
)

type Report struct {
	Nums []int
}
type reportDiff struct {
	nums []int
	diffs []int
}
func (r Report) toReportDiff() reportDiff {
	diffs := make([]int, len(r.Nums) - 1)
	for i := 1; i < len(r.Nums); i++ {
		diff := aocmath.Vertex { A: r.Nums[i - 1], B: r.Nums[i] }.Diff()
		diffs[i - 1] = diff
	}
	return reportDiff { diffs: diffs, nums: r.Nums }
}

func checkSafety(diffs []int, maxDiff int) bool {
	checks := map[int]int { -1: 0, 1: 0}
	for _, diff := range diffs {
		diffNormalized := aocmath.Normalize(diff)
		if diffNormalized == 0 || aocmath.Abs(diff) > maxDiff {
			return false
		}
		checks[diffNormalized] = checks[diffNormalized] + 1
	}
	return checks[1] == len(diffs) || checks[-1] == len(diffs)
}

func (rep Report) isSafe(maxDiff int, canRemove bool) bool {
	if canRemove {
		safe := checkSafety(rep.toReportDiff().diffs, maxDiff)
		if safe {
			return true
		}
		for i := 0; i < len(rep.Nums); i++ {
			nNums := make([]int, 0)
			nNums = append(nNums, rep.Nums[0:i]...)
			nNums = append(nNums, rep.Nums[i + 1 : len(rep.Nums)]...)
			nRep := Report { Nums: nNums }
			safe = checkSafety(nRep.toReportDiff().diffs, maxDiff)
			if safe {
				return true
			}
		}
		return false

	} else {
		return checkSafety(rep.toReportDiff().diffs, maxDiff)
	}
}


type Day02Part01 struct {
	Reports []Report
}

type Day02Part02 struct {
	Reports []Report
}

func (data *Day02Part01) Exec() (string, error) {
	safeReports := 0
	for i := 0; i < len(data.Reports); i++ {
		if data.Reports[i].isSafe(3, false) {
			safeReports++
		}
	}
	return strconv.Itoa(safeReports), nil
}

func (data *Day02Part02) Exec() (string, error) {
	safeReports := 0
	for i := 0; i < len(data.Reports); i++ {
		if data.Reports[i].isSafe(3, true) {
			safeReports++
		}
	}
	fmt.Println()
	return strconv.Itoa(safeReports), nil
}
