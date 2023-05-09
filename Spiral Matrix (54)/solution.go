func spiralOrder(matrix [][]int) []int {
    height := len(matrix)
    width := len(matrix[0])
    left, top, right, bottom := 0, 0, width, height

    spiralLen := height*width
    spiral := make([]int, spiralLen, spiralLen)

    var i, j int
    for spiralPos := 0; spiralPos < spiralLen; {
        // go right
        i = top
        for j = left; j <= right-1; j++ {
            spiral[spiralPos] = matrix[i][j]
            spiralPos++

            if spiralPos >= spiralLen {
                return spiral
            }
        }
        top++

        // go down
        j = right-1
        for i = top; i <= bottom-1; i++ {
            spiral[spiralPos] = matrix[i][j]
            spiralPos++

            if spiralPos >= spiralLen {
                return spiral
            }
        }
        right--

        // go left
        i = bottom-1
        for j = right-1; j >= left; j-- {
            spiral[spiralPos] = matrix[i][j]
            spiralPos++

            if spiralPos >= spiralLen {
                return spiral
            }
        }
        bottom--

        // go up
        j = left
        for i = bottom-1; i >= top; i-- {
            spiral[spiralPos] = matrix[i][j]
            spiralPos++

            if spiralPos >= spiralLen {
                return spiral
            }
        }
        left++
    }

    return spiral
}
