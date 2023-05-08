# Matrix Diagonal Sum (1677)

## The Problem

[Matrix Diagonal Sum (1677)](https://leetcode.com/problems/matrix-diagonal-sum)

### Level

Easy

### Tags

 [Array](https://leetcode.com/tag/array), [Matrix](https://leetcode.com/tag/matrix)

## Solution

### Code

[solution.go](solution.go)

[LeetCode submissions section](https://leetcode.com/problems/matrix-diagonal-sum/submissions/946647360/)

### Performance

[LeetCode submission details](https://leetcode.com/submissions/detail/946647360/)

#### CPU

Took 13 ms (beats 42.98%)

#### Memory

Used 5.6 MB (beats 26.32%)

## Problem statement

Given a square matrix `mat`, return the sum of the matrix diagonals.

Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/08/14/sample_1911.png) 


**Input:** mat = [[**1**,2,**3**],
              [4,**5**,6],
              [**7**,8,**9**]]
**Output:** 25
**Explanation:** Diagonals sum: 1 + 5 + 9 + 3 + 7 = 25
Notice that element mat[1][1] = 5 is counted only once.

**Example 2:**


**Input:** mat = [[**1**,1,1,**1**],
              [1,**1**,**1**,1],
              [1,**1**,**1**,1],
              [**1**,1,1,**1**]]
**Output:** 8

**Example 3:**


**Input:** mat = [[**5**]]
**Output:** 5

**Constraints:**

* `n == mat.length == mat[i].length`
* `1 <= n <= 100`
* `1 <= mat[i][j] <= 100`

## Social rank

:thumbsup: 2856 / :thumbsdown: 37
