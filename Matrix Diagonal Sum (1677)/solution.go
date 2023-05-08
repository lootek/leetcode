func diagonalSum(mat [][]int) int {
    sum := 0

    size := len(mat)
    for i := 0; i < size; i++ {
        sum += mat[i][i]

        if i != size - 1 - i {
            sum += mat[i][size - 1 - i]
        }
    }

    return sum
}