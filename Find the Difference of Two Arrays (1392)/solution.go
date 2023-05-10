func findDifference(nums1 []int, nums2 []int) [][]int {
    map1Ch := asMap(nums1)
    map2Ch := asMap(nums2)

    map1 := <- map1Ch
    map2 := <- map2Ch

    return [][]int{
        <- diff(map1, map2),
        <- diff(map2, map1),
    }
}

func asMap(nums []int) chan map[int]struct{} {
    res := make(chan map[int]struct{})

    go func() {
        mapped := make(map[int]struct{}, len(nums))
        for _, v := range nums {
            mapped[v] = struct{}{}
        }

        res <- mapped
    }()

    return res
}

func diff(map1, map2 map[int]struct{}) chan []int {
    res := make(chan []int)

    go func() {
        only1 := make([]int, 0, len(map1))
        for k := range map1 {
            if _, ok := map2[k]; !ok {
                only1 = append(only1, k)
            }
        }

        res <- only1
    }()

    return res
}
