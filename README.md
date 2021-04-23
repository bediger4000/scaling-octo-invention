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

1. increments a from zero to M/2
2. For each a, calculates b = M - a
3. if x xor b == N, prints a, b

You only have to increment up to M/2. 
For a > M/2 you'll find pairs b,a that
mirror pairs a,b you've already found.
Xor is commutative.

This is an O(M/2) algorithm, for whatever the input M is.

There's got to be a better algorithm,
because only a limited number of values of a xor b
occur for any given M.

## Around the web 

An extremely clever, inobvious and subtle bitwise
[solution](https://initcodes.blogspot.com/2019/06/daily-coding-problem-332-relationship.html)
