package hash

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	out := make([][]string, 0)

	for _, value := range strs {
		count := [26]int{}
		for _, v := range value {
			count[v-'a']++
		}
		//取出对应的hash
		m[count] = append(m[count], value)
	}

	for _, group := range m {
		out = append(out, group)
	}

	return out
}
