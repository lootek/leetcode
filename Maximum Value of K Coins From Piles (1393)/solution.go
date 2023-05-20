func maxValueOfCoins(piles [][]int, k int) int {
    P := len(piles)

    knapsack := make([][]int, P+1, P+1)
    for p := 0; p <= P; p++ {
        knapsack[p] = make([]int, k+1, k+1)
    }

    // printPiles := func() {
    //     for p := 0; p < P; p++ {
    //         for i := 0; i < len(piles[p]) && i < k; i++ {
    //             fmt.Printf(" %4d", piles[p][i])
    //         }
    //         fmt.Println()
    //     }
    //     fmt.Println()
    // }

    // printKnapsack := func() {
    //     for i := 1; i < len(knapsack); i++ {
    //         for j := 1; j < len(knapsack[i]); j++ {
    //             fmt.Printf("%4d", knapsack[i][j])
    //         }
    //         fmt.Println()
    //     }
    //     fmt.Println()
    // }

    // printPiles()

    for p := 0; p < P; p++ {
        for i := 1; i < len(piles[p]) && i < k; i++ {
            if piles[p][i] == 0 {
                break
            }

            piles[p][i] += piles[p][i-1]
        }
    }

    // printPiles()

    for p := 1; p <= P; p++ {
        for i := 0; i <= k; i++ {
            knapsack[p][i] = knapsack[p-1][i]

            for j := 1; j <= len(piles[p-1]) && j <= i; j++ {
                if knapsack[p][i] < knapsack[p-1][i-j] + piles[p-1][j-1] {
                    knapsack[p][i] = knapsack[p-1][i-j] + piles[p-1][j-1]
                }
            }
        }

        // printKnapsack()
    }

    // printKnapsack()

    return knapsack[P][k]
}
