func arraySign(nums []int) int {
    countNegative := 0
    for _, n := range nums {
        if n < 0 {
            countNegative++
        }

        if n == 0 {
            return 0
        }
    }

    if countNegative%2 == 0 {
        return 1
    }

    return -1
}
