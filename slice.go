package common

import "sort"

//In - возвращает true,  если в проверяемом срезе есть искомое число
func In(source []int, digit int) bool {
	result := make(map[int]bool)
	for _, e := range source {
		result[e] = true
	}
	return result[digit]
}

//MaxInt - возвращает максимальное число из среза
func MaxInt(source []int) (max int) {
	for _, e := range source {
		if max < e {
			max = e
		}
	}
	return max
}

//MinInt64 - возвращает минимальное число из среза
func MinInt64(source []int64) (min int64) {
	sort.Slice(source, func(i, j int) bool {
		return source[i] < source[j]
	})
	return source[0]
}

//AvgInt64 - возвращает среднее число из среза
func AvgInt64(source []int64) (avg int64) {
	var sum int64
	for _, e := range source {
		sum += e
	}
	return sum / int64(len(source))
}

//IsOnlyIncrease - возвращает true, только если каждое последующее число в срезе не меньше предыдущего
func IsOnlyIncrease(source []int) bool {
	for i := 0; i < len(source)-1; i++ {
		if source[i] > source[i+1] {
			return false
		}
	}
	return true
}

//IsTwin - возвращает true, если срез является спаркой
//т.е. все последующие числа в срезе равны или выше первого, а последний строго больше первого
func IsTwin(source []int) bool {
	if len(source) == 1 { //если срез представлен всего одним элементом, его можно назвать спаркой по рангу 1
		return true
	}
	for i := 0; i < len(source)-2; i++ {
		if source[0] > source[i+1] { //если какой либо последующий но не последний меньше первого, вернем ложь
			return false
		}
	}
	if source[0] >= source[len(source)-1] { //если последний не строго больше первого, вернем ложь
		return false
	}
	return true
}

//CountEntriesUpTo - возвращает, сколько раз встречаются числа от 0 до max в выборке source
func CountEntriesUpTo(max int, source []int) map[int]int64 {
	results := make(map[int]int64)
	for _, e := range source {
		results[e]++
	}
	return results
}
