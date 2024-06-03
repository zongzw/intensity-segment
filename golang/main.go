package main

func main() {
	is := NewIntensitySegments()
	is.set(10, 50, 1)
	is.toString() // [[10 1] [50 0]]

	is.set(5, 8, 2)
	is.toString() // [[5 2] [8 0] [10 1] [50 0]]

	is.add(30, 40, 3)
	is.toString() // [[5 2] [8 0] [10 1] [30 4] [40 1] [50 0]]

	is.add(20, 40, -2)
	is.toString() // [[5 2] [8 0] [10 1] [20 -1] [30 2] [40 1] [50 0]]

	is.add(20, 40, -1)
	is.toString() // [[5 2] [8 0] [10 1] [20 -2] [30 1] [50 0]]

}
