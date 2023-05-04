func isWater(i, j, m, n int, grid [][]byte) bool {
    return i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0'
}

func isNewLand(i, j, m, n int, grid [][]byte) bool {
    return i >= 0 && j >= 0 && i < m && j < n && grid[i][j] == '1'
}

func isVisitedLand(i, j, m, n int, grid [][]byte) bool {
    return i >= 0 && j >= 0 && i < m && j < n && grid[i][j] == '2'
}

func visitLand(i, j, m, n int, grid [][]byte) {
    if isWater(i, j, m, n, grid) || isVisitedLand(i, j, m, n, grid) {
        return
    }

    grid[i][j] = '2'

    visitLand(i, j-1, m, n, grid)
    visitLand(i+1, j, m, n, grid)
    visitLand(i, j+1, m, n, grid)
    visitLand(i-1, j, m, n, grid)
}

func numIslands(grid [][]byte) int {
    num := 0
    i := 0
    j := 0
    m := len(grid)
    n := len(grid[0])

    for {
        // fmt.Println()
        // for k := 0; k < m; k++ {
        //     for l := 0; l < n; l++ {
        //         fmt.Print(string(grid[k][l]))
        //     }
        //     fmt.Println()
        // }

        if isNewLand(i, j, m, n, grid) {
            // fmt.Println("+")
            num++

            visitLand(i, j, m, n, grid)
        }

        if j < n {
            j++
            continue
        }

        j = 0
        if i < m {
            i++
            continue
        }

        break
    }

    return num
}
