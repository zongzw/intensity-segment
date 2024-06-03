package main

type intSlice []int

func (sl intSlice) Len() int {
	return len(sl)
}
func (sl intSlice) Less(i, j int) bool {
	return sl[i] < sl[j]
}
func (sl intSlice) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}
