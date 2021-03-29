# Daily Coding Problem: Problem #841 [Easy]

This problem was asked by Jane Street.

Given integers M and N,
write a program that counts how many positive integer pairs (a, b)
satisfy the following conditions:

* a + b = M
* a XOR b = N

## Analysis

The "[Easy]" rating suggests a brute force approach,
no tricky O(log N) inobvious and subtle algorithms.

I wrote a [brute force](bruteforce.go) program that

1. increments b from zero to M/2
2. chooses b = M - a
3. if x xor b == N, prints a, b

You only have to increment up to M/2, because you'll just
find pairs b,a of pairs a,b you've already found.
