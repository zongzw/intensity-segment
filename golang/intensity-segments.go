package main

import (
	"fmt"
	"sort"
)

// IntensitySegments
//
//	Using a map (because of it dscreteness) to store: [segemnt] -> intensity
type IntensitySegments struct {
	it map[int]int
}

// NewIntensitySegments new a IntensitySegments Object
// return a pointer to the new object
func NewIntensitySegments() *IntensitySegments {
	return &IntensitySegments{
		it: map[int]int{},
	}
}

// set set a new range(from <-> to, to not included) of intensity based on existing one.
func (is *IntensitySegments) set(from, to int, amount int) {
	keys := is.orderedKeys()
	if len(keys) == 0 { // if no segments, in the intial case
		is.it[to] = 0
		is.it[from] = amount
		return
	}

	if to < keys[0] || to > keys[len(keys)-1] { // the end of set segment is outside of the existing one
		is.it[to] = 0
	} else if _, f := is.it[to]; !f { // 'to' is not found in existing segment numbers
		l := is.leftKey(to)
		is.it[to] = is.it[l]
	}

	is.it[from] = amount // directly set the segment beginning

	for _, k := range is.orderedKeys() { // all keys between from and to is invalid, delete them
		if k > from && k < to {
			delete(is.it, k)
		}
	}

	is.merge()
}

// add add a new range(from <-> to, to not included) of intensity to existing one.
func (is *IntensitySegments) add(from, to int, amount int) {
	keys := is.orderedKeys()
	if len(keys) == 0 { // if no segments, in the intial case
		is.set(from, to, amount)
		return
	}

	if to < keys[0] || from > keys[len(keys)-1] { // no overlap case
		is.set(from, to, amount)
		return
	}

	// handle the 'to'
	if _, f := is.it[to]; !f { // if not found to, add it
		l := is.leftKey(to)
		is.it[to] = is.it[l]
	}

	// handle the 'from'
	if from < keys[0] {
		is.it[from] = amount
	} else if from == keys[0] {
		is.it[from] += amount
	} else {
		if _, f := is.it[from]; !f {
			is.it[from] = amount
			if l := is.leftKey(from); l != -1 {
				is.it[from] += is.it[l]
			}
		} else {
			is.it[from] += amount
		}
	}

	// handle the keys between from and 'to'
	for _, k := range keys {
		if k > from && k < to {
			is.it[k] += amount
		}
	}

	// merge segments
	is.merge()
}

// merge merge continous segments to together if they have same intensity
func (is *IntensitySegments) merge() {
	keys := is.orderedKeys()
	if len(keys) == 0 {
		return
	}

	// clear prefix 0
	i := 0
	for ; i < len(keys) && is.it[keys[i]] == 0; i++ {
		delete(is.it, keys[i])
	}
	keys = keys[i:]

	for i = len(keys) - 1; i >= 0; i-- {
		if is.it[keys[i]] == 0 {
			if i-1 >= 0 && is.it[keys[i-1]] == 0 {
				delete(is.it, keys[i])
				keys = append(keys[0:i], keys[i+1:]...)
			}
		}

	}

	// clear the surffix segment with same intensity <> 0
	lastIntensity := is.it[keys[0]]
	for i := 1; i < len(keys); i++ {
		if is.it[keys[i]] == lastIntensity && lastIntensity != 0 {
			delete(is.it, keys[i])
		} else {
			lastIntensity = is.it[keys[i]]
		}
	}

	// clear the surffix segment with intensity == 0
	lastIntensity = 0
	for i := len(keys) - 2; i >= 0; i-- {
		if is.it[keys[i]] == 0 && lastIntensity == 0 {
			delete(is.it, keys[i])
		} else {
			lastIntensity = is.it[keys[i]]
		}
	}
}

// orderedKeys orderedKeys return a sorted keys of is.it
func (is *IntensitySegments) orderedKeys() []int {
	keys := intSlice{}
	for k := range is.it {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return keys
}

// dumps dumps return a string of simple format, i.e. [[20 1] [30 2] [40 0]]
func (is *IntensitySegments) dumps() string {
	keys := is.orderedKeys()
	rlt := [][2]int{}
	for _, k := range keys {
		rlt = append(rlt, [2]int{k, is.it[k]})
	}
	return fmt.Sprintf("%v", rlt)
}

// leftKey return the left segment number of the given.
func (is *IntensitySegments) leftKey(x int) int {
	l := -1
	for _, k := range is.orderedKeys() {
		if k < x {
			l = k
		}
	}
	return l
}

// toString print the dumped string simply
func (is *IntensitySegments) toString() {
	fmt.Printf("%v\n", is.dumps())
}
