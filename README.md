# Word Break (DP)

> Striver SDE Sheet 182 · Day 26 DP II — dp over prefixes: s[:i] breaks iff some breakable shorter prefix plus one dictionary word reaches i. Go, stdlib only, tested.

## Problem

Can a string be segmented entirely into dictionary words (each usable repeatedly)? The decision twin of #060, which enumerated the actual sentences — comparing the two shows what dropping the output does to complexity.

## Approach

`dp[i]`: the prefix of length i is breakable. It is, exactly when some j < i has `dp[j]` true and `s[j:i]` in the dictionary — seeded by the trivially-breakable empty prefix. Two pragmatic bounds keep the inner loop tight: scan j backwards from i and stop beyond the longest dictionary word (nothing longer can match), and break on the first success. Where #060's output could be exponential (every segmentation), the boolean collapses to polynomial — the standard lesson about decision vs enumeration versions of the same recurrence.

**Complexity:** O(n · maxWordLen) time with O(1) substring hashing via the map, O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
