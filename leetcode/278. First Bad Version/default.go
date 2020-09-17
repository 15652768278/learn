package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {
	if isBadVersion(1) {
		return 1
	}
	left, right := 1, n
	for right-left > 1 {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
