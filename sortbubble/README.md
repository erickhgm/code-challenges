# Bubble sort

## Instructions

Sort list of numbers from lowest number to greatest number using
[bubble sort](https://en.wikipedia.org/wiki/Bubble_sort).

## Algorithm

Starting from the beginning of the list, compare every adjacent pair, swap their position if they are not in the right
order (the latter one is smaller than the former one). After each iteration, one less element (the last one)
is needed to be compared until there are no more elements left to be compared.

[challenge](solution_test.go) | [solution](solution.go)

`O(n^2)`