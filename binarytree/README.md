# Binary search tree

## Instructions

Implement [binary search tree (BST)](https://en.wikipedia.org/wiki/Binary_search_tree).

[challenge](solution_test.go) | [solution](solution.go)

## Complecity
Balanced Binary tree: `O(log n)` (insert and find)
Unbalanced Binary tree: `O(n)` (insert and find)

## Traversing

### In Order 
`left` then `root` then `right`

It is the most used, you can print the elemets in order.

### Pre Order 
First the `root` then `bottom` (leaf)

`root` then `left` then `right`

### Post Order
First the `bottom` (leaf) then `root`

So: `left` then `right` then `root`


Ref.: https://www.youtube.com/watch?v=oSWTXtMglKE