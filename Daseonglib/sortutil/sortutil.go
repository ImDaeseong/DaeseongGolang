package sortutil

import (
	"fmt"
	"sort"
)

type game struct {
	sname  string
	svalue string
}

var (
	gMap map[int]game
	sMap map[int]string
)

func sort_sMap(m map[int]string, reverse bool) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}

	if reverse {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	} else {
		sort.Sort(sort.IntSlice(keys))
	}
	return keys
}

func sort_gMap(m map[int]game, reverse bool) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}

	if reverse {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	} else {
		sort.Sort(sort.IntSlice(keys))
	}

	return keys
}

func Getgame_gMap() {

	gMap = make(map[int]game)

	gMap[1] = game{"sname1", "svalue1"}
	gMap[3] = game{"sname3", "svalue3"}
	gMap[2] = game{"sname2", "svalue2"}
	gMap[7] = game{"sname7", "svalue7"}
	gMap[5] = game{"sname5", "svalue5"}

	for _, i := range sort_gMap(gMap, false) {
		fmt.Println(gMap[i])
	}
}

func Getgame_sMap() {

	sMap = make(map[int]string)

	sMap[1] = "sMap1"
	sMap[3] = "sMap3"
	sMap[2] = "sMap2"
	sMap[7] = "sMap7"
	sMap[5] = "sMap5"

	for _, i := range sort_sMap(sMap, true) {
		fmt.Println(sMap[i])
	}
}
