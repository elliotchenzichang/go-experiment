package leetcode

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func checkDistances(s string, distance []int) bool {
	mark := make(map[byte]int)
	bytes := []byte(s)
	for i, b := range bytes {
		if index, exist := mark[b]; exist {
			if i-index-1 != distance[b-'a'] {
				return false
			}
		} else {
			mark[b] = i
		}
	}
	return true
}

func TestCheckDistances(t *testing.T) {
	assert.Equal(t, true, checkDistances("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, false, checkDistances("aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}

func numberOfWays(startPos int, endPos int, k int) int {
	marker := []int{0}
	for i := 1; i <= k; i++ {
		marker = append([]int{1}, marker...)
		marker = append(marker, 1)
		for i := 1; i < len(marker)-1; i++ {
			marker[i] = marker[i-1] + marker[i+1]
		}
	}
	middle := len(marker) / 2
	effetice := endPos - startPos - 1
	if effetice > middle {
		return 0
	}
	return marker[effetice+middle]
}

func TestNumberOfWays(t *testing.T) {
	assert.Equal(t, 3, numberOfWays(1, 2, 3))
	assert.Equal(t, 0, numberOfWays(2, 5, 10))
	assert.Equal(t, 0, numberOfWays(989, 1000, 99))
}

func longestNiceSubarray(nums []int) int {
	res := 0
	left, right := 0, 0
	for i := 1; i < len(nums); i++ {
		for j := right; j >= left; j-- {
			if nums[j]&nums[i] != 0 {
				if right-left+1 > res {
					res = right - left + 1
				}
				left = j + 1
				break
			}
		}
		right = i
	}
	if right-left+1 >= res {
		res = right - left + 1
	}
	return res
}

func TestLongestNiceSubarray(t *testing.T) {
	assert.Equal(t, 3, longestNiceSubarray([]int{1, 3, 8, 48, 10}))
	assert.Equal(t, 1, longestNiceSubarray([]int{3, 1, 5, 11, 13}))
}
