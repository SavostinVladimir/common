package common

import (
	"sort"
)

//IntKeysWithMaxValue - возвращает ключи с максимальным значением,  в количестве amount
func IntKeysWithMaxValue(source map[int]int64, amount int) (digits []int) {
	type kv struct {
		Key   int
		Value int64
	}
	var pairs []kv
	for k, v := range source {
		pairs = append(pairs, kv{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	for i := 0; i < amount; i++ {
		digits = append(digits, pairs[i].Key)
	}
	return digits
}

//StringKeysWithMaxValue - возвращает ключи с максимальным значением,  в количестве amount
func StringKeysWithMaxValue(source map[string]int, amount int) (keys []string) {
	type kv struct {
		Key   string
		Value int
	}
	var pairs []kv
	for k, v := range source {
		pairs = append(pairs, kv{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})
	for i := 0; i < amount; i++ {
		keys = append(keys, pairs[i].Key)
	}
	return keys
}


//MinKeys - возвращает ключи с минимальным значением,  в количестве amount
func MinKeys(source map[int]int64, amount int) (digits []int) {
	type kv struct {
		Key   int
		Value int64
	}
	var pairs []kv
	for k, v := range source {
		pairs = append(pairs, kv{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	for i := 0; i < amount; i++ {
		digits = append(digits, pairs[i].Key)
	}
	return digits
}

//MaxKeys - возвращает максимальные ключи,  в количестве amount
func MaxKeys(source map[int]int64, amount int) (digits []int) {
	type kv struct {
		Key   int
		Value int64
	}
	var pairs []kv
	for k, v := range source {
		pairs = append(pairs, kv{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key > pairs[j].Key
	})

	for i := 0; i < amount; i++ {
		digits = append(digits, pairs[i].Key)
	}
	return digits
}


