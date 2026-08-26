package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var summ int
	for key, val := range cb {
		if key == file {
			for _, v := range val {
				if v == true {
					summ++
				}
			}
		}
	}
	return summ
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var summ int

	if rank < 1 && rank > 8 {
		return 0
	}

	res := rank - 1
	for _, val := range cb {
		for i, v := range val {
			if i == res {
				if v == true {
					summ++
				}
			}
		}
	}
	return summ
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var sum int

	for _, val := range cb {
		for range val {
			sum++
		}
	}
	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var sum int

	for _, val := range cb {
		for _, v := range val {
			if v == true {
				sum++
			}
		}
	}
	return sum
}
