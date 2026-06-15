# Tug of War Team Selection

This project solves a competitive programming problem about optimal team selection and bribery strategy.

---

## Problem Summary

There are **n participants**, each with a unique weight.

Team selection happens as follows:

- **Romina picks 1 person**
- **Ali picks 2 people**
- Then they **alternate picking 2 people each**
- If fewer people remain at the end, the **current player picks the remaining ones**

Both players pick **optimally to maximize their team strength**.

After teams are formed, one player may **bribe some members of the opposing team** so that they do not contribute to the tug-of-war.

The goal is to **find the minimum number of people to bribe so that the chosen player wins**.

---

## Key Idea

1. Sort participants by weight in **descending order**
2. Simulate the **team selection process**
3. Compute the **total weight of each team**
4. If the target player is **losing or tied**:
    - Bribe the **heaviest members of the opponent’s team first**
    - Continue until the target player’s total weight becomes greater

This greedy strategy guarantees the **minimum number of bribes**.

---

## Time Complexity

- **Sorting:** `O(n log n)`
- **Simulation + bribery:** `O(n)`

**Total complexity**
O(n log n)

---

## How to Run
```bash
go run main.go

#Then provide input in this format:
n
weights
target_player

#Example
## in
4
5 7 3 1
ali
##out
1
1






