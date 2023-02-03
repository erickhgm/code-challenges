# Big O

## What is Big O?

Big O, also known as Big O notation, represents an algorithm's `worst-case complexity`. It uses algebraic terms to describe the complexity of an algorithm.

Big O defines the runtime required to execute an algorithm by identifying how the performance of your algorithm will change as the input size grows. But it does not tell you how fast your algorithm's runtime is.

Big O notation measures the efficiency and performance of your algorithm using `time` and `space` complexity.

An algorithm's `time` complexity specifies `how long it will take to execute` an algorithm as a function of its input size. Similarly, an algorithm's `space` complexity specifies the `total amount of space or memory` required to execute an algorithm as a function of the size of the input.

## Video tutorial series
[Big O For Software Engineering](https://www.youtube.com/playlist?list=PL7g1jYj15RUPVZDU9C276SZvlJjf4hzqV)

## Complexities
In Big O, there are six major types of complexities (time and space):

- Constant: `O(1)`
- Logarithmic: `O(log n)`
- Linear time: `O(n)`
- Logarithmic linear time: `O(n * log n)`
- Quadratic time: `O(n^2)`
- Cubic time: `O(n^3)`
- Exponential time: `O(2^n)`
- Factorial time: `O(n!)`

![Arquitectura](https://paper-attachments.dropbox.com/s_2D428973624E7FC84C7D69D11421DE762BEA6B6F3361231FCDCAE0425D14526F_1664885448372_Untitled.drawio+17.png)

- When your calculation is not dependent on the input size, it is a constant time complexity `O(1)`.
- When the input size is reduced by half, maybe when iterating, handling recursion, or whatsoever, it is a logarithmic time complexity `O(log n)`.
- When you have a single loop within your algorithm, it is linear time complexity `O(n)`.
- When you have nested loops within your algorithm, meaning a loop in a loop, it is quadratic time complexity `O(n^2)`.
- When the growth rate doubles with each addition to the input, it is exponential time complexity `O(2^n)`.

## Constant Time: O(1)
When your algorithm is `not dependent on the input size n`, it is said to have a constant time complexity with order O(1). This means that the run time will always be the same regardless of the input size.
```golang
func firstElement (array []int) int {
  return array[0];
};

score := []uint{12, 55, 67, 94, 22};
fmt.Println(firstElement(score)); // 12
```

## Linear Time: O(n)
You get linear time complexity when the running time of an algorithm `increases linearly with the size of the input`. This means that when a function has an iteration that iterates over an input size of n, it is said to have a time complexity of order `O(n)`. 

Eg. Linear search (unsorted array)

```golang
const calcFactorial = (n) => {
  factorial := 1;
  for (i := 2; i <= n; i++) {
    factorial = factorial * i;
  }
  return factorial;
};

fmt.Println(calcFactorial(5)); // 120
```

## Logarithm Time: O(log n)
This is similar to linear time complexity, except that the runtime does not depend on the input size but rather on `half the input size`. When the input size decreases on each iteration or step, an algorithm is said to have `logarithmic time` complexity.

```golang
func logN(n int) {
  if n == 0 {
    return
  }
  n = math.Floor(n / 2)
  return logN(n)
}
```

Eg. Binary search (already sorted array)

```golang
func binarySearch(elmts []int, elmt int) int {
	left := 0
	right := len(elmts) - 1

	for left <= right {
		mid := (right + left) / 2
		v := elmts[mid]

		if elmt == v {
			return mid
		}

		if elmt > v {
			left = mid + 1 // x greater than middle, so ignore left half
		} else {
			right = mid - 1 // x greater than middle, so ignore right half
		}
	}
	return -1
}

score := []int{12, 22, 45, 67, 96}
fmt.Println(binarySearch(score, 96))
```

## Logarithmic linear: O(n log n)

This is similar to logarithm Time complexity `O(log n)`, except that the runtime execute the `O(log n) n times`, so we have `O(n log n)`.

```golang
func nLogN(n int) {
	y := n
    for n > 1 {
        n = math.Floor(n / 2)
        for (i := 1; i <= y; i++) {
            fmt.Println(i)
        }
    }
}

fmt.Println(nLogN(4))
```

## Quadratic Time: O(n^2) - squared
When you perform nested iteration, meaning having a `loop in a loop`, the time complexity is `quadratic`, which is horrible.

```golang
func matchElements (array []int) string {
  for (i := 0; i < len(array); i++) {
    for (j := 0; j < len(array); j++) {
      if (i != j && array[i] == array[j]) {
        return fmt.Sprintf("Match found at %v and %v", i, j)
      }
    }
  }
  return "No matches found";
};

fruit := []string{
    "strawberry", 
    "pear", 
    "orange", 
    "banana", 
    "pineapple",
    "peach", 
    "apple", 
    "melon", 
    "orange", 
    "grape"}

console.log(matchElements(fruit)) // "Match found at 2 and 8"
```

## Cubic time: O(nˆ3) - cubed
Same than before, but having a `loop in a loop in a loop`.

## Exponential Time: O(2^n)
You get exponential time complexity when the growth rate doubles with each addition to the input (n), often iterating through all subsets of the input elements. `Any time an input unit increases by 1, the number of operations executed is doubled`.

You almost never use recursion for performance reasons. You use recursion to make the problem more simple.

It is possible that recursion will be more expensive, depending on if the recursive function is [tail recursive](https://en.wikipedia.org/wiki/Tail_call). Tail recursion should be recognized by the compiler and optimized to its iterative counterpart (while maintaining the concise, clear implementation you have in your code).

Eg. Recursive Fibonacci sequence.

```golang
func recursiveFibonacci(n int) int {
  if n < 2 {
    return n
  }
  return recursiveFibonacci(n - 1) + recursiveFibonacci(n - 2)
};

fmt.Println(recursiveFibonacci(6)) // 8
```

## Factorial Time: O(n!)
Factorial notation is the least efficient of them all, but there are some problems for which there is no easy solution.

These are what are known as `NP-complete problems`. NP-complete is a concept in complexity theory used to describe a category of problems for which there is no known correct and fast solution. In other words, the solution to an NP-complete problem can be quickly verified, but there is no known way to quickly find a solution.

Ref.: https://jarednielsen.com/big-o-factorial-time-complexity

Eg. [The Traveling Salesman](https://en.wikipedia.org/wiki/Travelling_salesman_problem)

Say you’re a traveling salesperson and you need to visit n cities. What is the shortest route that visits each and returns you to your start? To solve this, we need to calculate every possible route.

```
Austin > Boston > Chicago
Austin > Chicago > Boston
Boston > Austin > Chicago
Boston > Chicago > Austin
Chicago > Austin > Boston
Chicago > Boston > Austin
```

This is 3!, which is six permutations.

What about 4? That’s 120.

What about 10? That’s 3.628.800.