// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[2,2,1,1,5,3,3,5]`, 
			`7`,
		},
		{
			`[1,1,1,2,2,2,3,3,3,4,4,4,5]`, 
			`13`,
		},
		{
			`[1,1,1,2,2,2]`, 
			`5`,
		},
		{
			`[10,2,8,9,3,8,1,5,2,3,7,6]`, 
			`8`,
		},
		
	}
	targetCaseNum :=3
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxEqualFreq, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-158/problems/maximum-equal-frequency/