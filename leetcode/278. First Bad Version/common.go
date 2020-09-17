package main

import "math"

var threshold = math.MinInt64

func isBadVersion(version int) bool {
	return version >= threshold
}
