// This solution is a modified/adapted solution of spiral-matrix problem (54.)

func generateMatrix(n int) [][]int {
    matrix := make([][]int, n, n)
    for i := 0; i < n; i++ {
        matrix[i] = make([]int, n, n)
    }

    left, top, right, bottom := 0, 0, n, n
    spiralLen := n*n

    var i, j int
    for spiralPos := 0; spiralPos < spiralLen; {
        // go right
        i = top
        for j = left; j <= right-1; j++ {
            matrix[i][j] = spiralPos+1
            spiralPos++

            if spiralPos >= spiralLen {
                return matrix
            }
        }
        top++

        // go down
        j = right-1
        for i = top; i <= bottom-1; i++ {
            matrix[i][j] = spiralPos+1
            spiralPos++

            if spiralPos >= spiralLen {
                return matrix
            }
        }
        right--

        // go left
        i = bottom-1
        for j = right-1; j >= left; j-- {
            matrix[i][j] = spiralPos+1
            spiralPos++

            if spiralPos >= spiralLen {
                return matrix
            }
        }
        bottom--

        // go up
        j = left
        for i = bottom-1; i >= top; i-- {
            matrix[i][j] = spiralPos+1
            spiralPos++

            if spiralPos >= spiralLen {
                return matrix
            }
        }
        left++
    }

    return matrix
}
