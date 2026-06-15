# Maximum Bitwise OR After Removing One Element

## Problem

You are given an array of integers. Exactly **one element must be removed**, and the remaining elements will be combined using the **bitwise OR** operation.

Your task is to determine the **maximum possible OR value** that can be obtained after removing exactly one element.

The process for each scenario is:

1. Remove one element from the array.
2. Compute the **bitwise OR** of all remaining elements.
3. Choose the removal that produces the **maximum result**.

---

## Input

The input consists of multiple test cases.

- The first line contains an integer `t` — the number of test cases.
- For each test case:
    - An integer `n` — the number of elements in the array.
    - A line containing `n` integers `a1, a2, ..., an`.