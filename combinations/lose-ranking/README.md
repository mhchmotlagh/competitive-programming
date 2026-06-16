# Recover Missing Contest Ranking

## Problem

The organizers of a competition archived the results of several contest categories.  
However, the ranking of **one category** was lost.

You are given:

- The rankings of all other categories
- The participants of the missing category (without order)
- The **final overall ranking**

Your task is to **reconstruct the missing ranking**.

If multiple valid rankings exist, any correct one is acceptable.

---

# Scoring Formula

The score of a participant in a category is computed as:

Score = S × (P - R) / (P - 1)

Where:

- **S** = total score of the category
- **P** = number of participants
- **R** = rank of the participant in that category

Special case:

If there is only one participant:

If a participant participates in multiple categories, their **final score** is the sum of their scores from all categories.

The final ranking is determined by sorting participants by their **total score**.

---

# Input

1. Integer `n` — number of known categories.

For each category:

- `S P`
- A comma-separated list of `P` usernames in ranking order.

Then the **missing category**:

- `S P`
- `P` usernames (unordered)

Then the **final ranking**:

- `P` (total participants)
- A comma-separated list of usernames in final ranking order.

---

# Output

Print the **recovered ranking** of the missing category.

Each username should be printed on a separate line.

If multiple answers are valid, any correct one is accepted.

---

# Key Observation

The missing category has at most:
