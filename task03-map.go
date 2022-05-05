package homework

import (
	"fmt"
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {

	type kv struct {
		k int
		v string
	}
	kvs := make([]kv, 0, len(input))
	for k, v := range input {
		kvs = append(kvs, kv{k, v})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].k < kvs[j].k
	})
	n := make([]string, 0, len(input))
	for k, v := range kvs {
		fmt.Println(kvs[k].v, v)
		nn := kvs[k].v
		n = append(n, string(nn))
	}

	return n
}
