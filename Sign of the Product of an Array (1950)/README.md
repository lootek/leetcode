# Sign of the Product of an Array (1950)

## Category

### Level

Easy

### Tags

[Array](https://leetcode.com/tag/array),[Math](https://leetcode.com/tag/math)

## Solution performance

### CPU

Took 3 ms

Beats 86.92%

### Memory

Used 3.3 MB

Beats 98.81%

## Problem statement

[Sign of the Product of an Array (1950)](https://leetcode.com/problems/sign-of-the-product-of-an-array)

There is a function `signFunc(x)` that returns:

* `1` if `x` is positive.
* `-1` if `x` is negative.
* `0` if `x` is equal to `0`.

You are given an integer array `nums`. Let `product` be the product of all values in the array `nums`.

Return `signFunc(product)`.

**Example 1:**


**Input:** nums = [-1,-2,-3,-4,3,2,1]
**Output:** 1
**Explanation:** The product of all values in the array is 144, and signFunc(144) = 1

**Example 2:**


**Input:** nums = [1,5,0,2,-3]
**Output:** 0
**Explanation:** The product of all values in the array is 0, and signFunc(0) = 0

**Example 3:**


**Input:** nums = [-1,1,-1,1,-1]
**Output:** -1
**Explanation:** The product of all values in the array is -1, and signFunc(-1) = -1

**Constraints:**

* `1 <= nums.length <= 1000`
* `-100 <= nums[i] <= 100`

### Social

:thumbsup: 1820 / :thumbsdown: 182