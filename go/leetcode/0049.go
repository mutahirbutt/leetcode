package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {

	byteMap := map[string][]string{}

	res := [][]string{}

	for _, s := range strs {
		runeS := []rune(s)
		sort.Slice(runeS, func(i int, j int) bool {
			return runeS[i] < runeS[j]
		})
		str := string(runeS)

		if _, ok := byteMap[str]; !ok {
			byteMap[str] = []string{s}
		} else {
			byteMap[str] = append(byteMap[str], s)

		}

	}

	for _, v := range byteMap {
		res = append(res, v)
	}

	return res

}
