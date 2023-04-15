func rangeBitwiseAnd(left int, right int) int {
    if left == 0 {
        return 0
    }

    // initial2powValue := 0
    // for val := 1; val < right; val = val << 1 {
    //     initial2powValue = val
    // }

    // fmt.Println(initial2powValue)

    // if initial2powValue > left {
    //     left = initial2powValue
    // }

    res := left & right
    // fmt.Println(left, right, res)
    for res != 0 && left < right {
        left++
        right--

        res = res & left
        res = res & right
    }

    return res
}