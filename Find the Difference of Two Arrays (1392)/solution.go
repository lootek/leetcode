func findDifference(nums1 []int, nums2 []int) [][]int {
    map1 := asMap(nums1)
    map2 := asMap(nums2)

    return [][]int{
        diff(map1, map2),
        diff(map2, map1),
    }
}

func asMap(nums []int) map[int]struct{} {
    res := map[int]struct{}{}
    for _, v := range nums {
        res[v] = struct{}{}
    }

    return res
}

func diff(map1, map2 map[int]struct{}) []int {
    only1 := []int{}
    for k := range map1 {
        if _, ok := map2[k]; !ok {
            only1 = append(only1, k)
        }
    }

    return only1
}
