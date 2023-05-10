# Merge k Sorted Lists (23)

## The Problem

[Merge k Sorted Lists (23)](https://leetcode.com/problems/merge-k-sorted-lists)

### Level

Hard

### Tags

 [Linked List](https://leetcode.com/tag/linked-list), [Divide and Conquer](https://leetcode.com/tag/divide-and-conquer), [Heap (Priority Queue)](https://leetcode.com/tag/heap-priority-queue), [Merge Sort](https://leetcode.com/tag/merge-sort)

## Solution

### Code

[solution.go](solution.go)

[LeetCode submissions section](https://leetcode.com/problems/merge-k-sorted-lists/submissions/947935820/)

### Performance

[LeetCode submission details](https://leetcode.com/submissions/detail/947935820/)

#### CPU

Took 212 ms (beats 20.50%)

#### Memory

Used 5.3 MB (beats 97.64%)

## Problem statement

You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order.

_Merge all the linked-lists into one sorted linked-list and return it._

**Example 1:**


**Input:** lists = [[1,4,5],[1,3,4],[2,6]]
**Output:** [1,1,2,3,4,4,5,6]
**Explanation:** The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

**Example 2:**


**Input:** lists = []
**Output:** []

**Example 3:**


**Input:** lists = [[]]
**Output:** []

**Constraints:**

* `k == lists.length`
* `0 <= k <= 104`
* `0 <= lists[i].length <= 500`
* `-104 <= lists[i][j] <= 104`
* `lists[i]` is sorted in **ascending order**.
* The sum of `lists[i].length` will not exceed `104`.

## Social rank

:thumbsup: 16945 / :thumbsdown: 617
