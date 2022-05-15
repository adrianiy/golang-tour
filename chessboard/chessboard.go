package chessboard


// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool 

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
    var count int
    _, exists := cb[rank]

    if !exists {
        return 0
    }

    for _, x := range cb[rank] {
        if x {
            count++
        }
    }

    return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
    if (file < 0 || file > 8) {
        return 0
    }
    var count int

    for _, x := range cb {
        if x[file - 1] {
            count++
        }
    }
    return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
    var count int

    for _, x := range cb {
        count += len(x)
    }

    return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
    var count int

    for _, x := range cb {
        for _, y := range x {
            if y {
                count++
            }
        }
    }

    return count
}
