package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	str := string(strs[0])
	for i := 1; i < len(strs); i++ {
		instr := ""
		for j := 0; j < len(str) && j < len(strs[i]); j++ {
			if str[j] != strs[i][j] {
				break
			}
			instr = instr + string(str[j])
		}

		str = instr
	}
	return str
}

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
