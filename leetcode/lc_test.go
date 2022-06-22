package leetcode

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func removeOuterParentheses(s string) string {
	bytes := []byte(s)
	leftNums := 0
	rightNums := 0
	var res []byte
	for _, item := range bytes {
		switch item {
		case '(':
			if leftNums != 0 {
				res = append(res, item)
			}
			leftNums++
		case ')':
			rightNums++
			if rightNums < leftNums {
				res = append(res, item)
			} else {
				leftNums = 0
				rightNums = 0
			}
		}
	}
	return string(res)
}

func TestRemoveOuterParentheses(t *testing.T) {
	assert.Equal(t, "()()()", removeOuterParentheses("(()())(())"))
	assert.Equal(t, "()()()()(())", removeOuterParentheses("(()())(())(()(()))"))
	assert.Equal(t, "", removeOuterParentheses("()()"))
}

func totalSteps(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	maxTimes := 0
	for index := 0; index < len(nums); index++ {
		temp := 0
		for i := index + 1; i < len(nums); i++ {
			if nums[i] < nums[index] {
				temp++
			} else {
				if temp > maxTimes {
					maxTimes = temp
				}
				break
			}
		}
	}
	return maxTimes
}

func TestTotalSteps(t *testing.T) {
	fmt.Println(totalSteps([]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}))
	fmt.Println(totalSteps([]int{4, 5, 7, 7, 13}))
	fmt.Println(totalSteps([]int{10, 1, 2, 3, 4, 5, 6, 1, 2, 3}))
}
