package main

import "strings"

func main() {
	/*
			输入: ["flower","flow","flight"]
			输出: "fl"


			输入: ["dog","racecar","car"]
			输出: ""

		    所有输入只包含小写字母 a-z

	*/

	longestCommonPrefix([]string{"rose", "rothing", "koro"})

}

// GO
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, k := range strs {
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
