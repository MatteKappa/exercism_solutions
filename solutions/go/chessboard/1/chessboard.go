package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	counter := 0
	for k, v := range cb {
		if k == file {
			for _, x := range v {
				if x == true {
					counter++
				}
			} 
		}
	}
	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	counter :=0
	number := rank-1
        if rank < 1 || rank > 8 {
		return 0
	}
	for letter := range cb {
		if cb[letter][number] == true {
			counter++
		}
	}	
	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	counter := 0
	for k := range cb{
		for i:=0; i<len(cb[k]); i++{
			counter++
		} 
	} 
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	counter := 0
	for k := range cb {
		for v := range cb[k]{
			if cb[k][v] == true {
				counter++
			}
		}
	}
	return counter
}
