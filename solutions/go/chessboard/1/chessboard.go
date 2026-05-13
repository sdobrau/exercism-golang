package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var squares int = 0
	if cb[file] == nil {
		return 0
	}

	for _, fileValue := range cb[file] {
		if fileValue == true {
			squares += 1
		}
	}
	return squares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	
	if rank < 1 || rank > 8 {
		return 0
	}	
	var squares int = 0
	// for each file
	for _, file := range cb {
		// for rank nr. rank
		if file[rank - 1] == true {
			squares += 1
		}
	}
	return squares
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var squares int = 0
	// for all files
	for _, file := range cb {
		// for all ranks in file
		for _, _ = range file {
			squares += 1
		}
	}
	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var squares int = 0
	for _, file := range cb {
		// for all ranks in file
		for _, value := range file {
			if value == true {
				squares += 1
			}
		}
	}
	return squares
}
