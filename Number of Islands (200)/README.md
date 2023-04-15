# Number of Islands (200)

## The Problem

[Number of Islands (200)](https://leetcode.com/problems/number-of-islands)

### Level

Medium

### Tags

 [Array](https://leetcode.com/tag/array), [Depth-First Search](https://leetcode.com/tag/depth-first-search), [Breadth-First Search](https://leetcode.com/tag/breadth-first-search), [Union Find](https://leetcode.com/tag/union-find), [Matrix](https://leetcode.com/tag/matrix)

## Solution

### Code

[solution.go](solution.go)

[LeetCode submissions section](https://leetcode.com/problems/number-of-islands/submissions/934124999/)

### Performance

[LeetCode submission details](https://leetcode.com/submissions/detail/934124999/)

#### CPU

Took 0 ms (beats 100.00%)

#### Memory

Used 3.9 MB (beats 90.66%)

## Problem statement

Given an `m x n` 2D binary grid `grid` which represents a map of `'1'`s (land) and `'0'`s (water), return _the number of islands_.

An **island** is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

**Example 1:**


**Input:** grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
**Output:** 1

**Example 2:**


**Input:** grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
**Output:** 3

**Constraints:**

* `m == grid.length`
* `n == grid[i].length`
* `1 <= m, n <= 300`
* `grid[i][j]` is `'0'` or `'1'`.

## Social rank

:thumbsup: 19713 / :thumbsdown: 435
