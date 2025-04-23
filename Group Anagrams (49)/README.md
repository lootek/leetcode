# Group Anagrams (49)

## The Problem

[Group Anagrams (49)](https://leetcode.com/problems/group-anagrams)

### Level

Medium

### Tags

 [Array](https://leetcode.com/tag/array), [Hash Table](https://leetcode.com/tag/hash-table), [String](https://leetcode.com/tag/string), [Sorting](https://leetcode.com/tag/sorting)

## Solution

### Code

[solution.go](solution.go)

[LeetCode submissions section](https://leetcode.com/problems/group-anagrams/submissions/1615992382/)

### Performance

[LeetCode submission details](https://leetcode.com/submissions/detail/1615992382/)

#### CPU

Took 18 ms (beats 21.57%)

#### Memory

Used 10 MB (beats 41.11%)

## Problem statement

Given an array of strings `strs`, group the anagrams together. You can return the answer in **any order**.

**Example 1:**

**Input:** strs = \["eat","tea","tan","ate","nat","bat"\]

**Output:** \[\["bat"\],\["nat","tan"\],\["ate","eat","tea"\]\]

**Explanation:**

* There is no string in strs that can be rearranged to form `"bat"`.
* The strings `"nat"` and `"tan"` are anagrams as they can be rearranged to form each other.
* The strings `"ate"`, `"eat"`, and `"tea"` are anagrams as they can be rearranged to form each other.

**Example 2:**

**Input:** strs = \[""\]

**Output:** \[\[""\]\]

**Example 3:**

**Input:** strs = \["a"\]

**Output:** \[\["a"\]\]

**Constraints:**

* `1 <= strs.length <= 104`
* `0 <= strs[i].length <= 100`
* `strs[i]` consists of lowercase English letters.

## Social rank

:thumbsup: 20365 / :thumbsdown: 682
