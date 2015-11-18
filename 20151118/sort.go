package main

import (
	"fmt"
	"sort"
)

// func (m Map) Len() int {
// 	return len(m)
// }

// func (m Map) Less(i, j int) bool {
// 	m[i], m[j] = m[j], m[i]
// }

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	// map[string]intのそーとしたい
	// var m map[string]int //これはnilのmapで、m["hoge"] = 1のように代入不可
	// Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that. To initialize a map, use the built in make function:
	// m = make(map[string]int) m["hoge"] = 1したい時はこっちを使う
	mp := make(map[string]int)
	mp["baz"] = 3
	mp["hoge"] = 1
	mp["car"] = 2

	var keys []string
	for k, _ := range mp {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// rangeするとsliceはmapなのかarrayなのかよくわからん動きするが、
	// どうもmapはkey,val
	// sliceはindex,val
	// arrayはvのみを持つらしい
	sorted_map := make(map[string]int)
	for _, k2 := range keys {
		sorted_map[k2] = mp[k2]
	}
	fmt.Println(sorted_map)

}
