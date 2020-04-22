package common

// BinaryMatrix A binary matrix means that all elements are 0 or 1.
// For each individual row of the matrix, this row is sorted in non-decreasing order.
type BinaryMatrix struct {
	data [][]int
	n, m int
}

// Get BinaryMatrix.Get(x, y) returns the element of the matrix at index (x, y) (0-indexed).
func (bm BinaryMatrix) Get(x, y int) int {
	return bm.data[x][y]
}

// Dimensions BinaryMatrix.Dimensions() returns a list of 2 elements [n, m], which means the matrix is n * m.
func (bm BinaryMatrix) Dimensions() []int {
	return []int{bm.n, bm.m}
}

// CreateBinaryMatrix Create binary matrix with a 2d array
func CreateBinaryMatrix(data [][]int) BinaryMatrix {
	n, m := 0, 0
	if len(data) > 0 {
		n, m = len(data), len(data[0])
	}
	return BinaryMatrix{data, n, m}
}
