package _014_Longest_Common_Prefix

func longestCommonPrefix(strs []string) string {
	res := strs[0]
	for _, s := range strs[1:] {
		for j := 0; j < len(res); j++ {
			if j >= len(s) || res[j] != s[j] {
				res = res[0:j]
				break
			}
		}
	}
	return res
}
