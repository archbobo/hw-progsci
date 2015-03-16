package main

import (
    "fmt"
    "math"
)

/*==============================================================================
 * 1. Gauss' formula. (Lecture 2)
 *
 * Implement a function that takes a positive integer n, and returns
 * the sum of the integers from 1 to n.
 *
 *============================================================================*/

func SumOfFirstNIntegers(n int) int {
    // WRITE YOUR CODE HERE!!!
    return n * (n + 1) / 2
}


/*==============================================================================
 * 2. Speed of a Marathoner (lecture 2)
 *
 * Write a function TimeToRun(marathonHours, marathonMinutes, miles) that
 * takes: the time a runner ran a marathon in possibly fractional hours and
 * possibly fractional minutes and a possibly fractional number of miles and
 * return the time in DAYS it should take the runner to run miles if he or she
 * runs at the same pace as they did in the marathon.
 *
 * For example: TimeToRun(3.1, 23.2, 107.1) should return 0.5938.
 *
 * Your function should also print out the answer in the format:
 *
 *     You could run 107.1 miles in 0.5938 days.
 *
 * Recall that there are 26.2 miles in a marathon.
 *============================================================================*/

func TimeToRun(marathonHours, marathonMinutes, miles float64) float64 {
    // WRITE YOUR CODE HERE
    var velocity float64 = 26.2 / (marathonHours * 60 + marathonMinutes)
    var days float64 = miles / velocity / 60 /24
    return days
}

/*==============================================================================
 * 3. Generalized Fibonacci sequences (Lecture 3)
 *
 * Implement a function GenFib(a0,a1,n) that takes: two positive integers a0,
 * a1 and a positive integer n, and returns the nth item in the sequence
 * defined by the rule:
 *      a_n = a_{n-1} + a_{n-2}.
 *============================================================================*/

// WRITE YOUR CODE HERE
func GenFib(a0, a1, n int) int {
  switch {
    case n == 0:
      return a0
    case n == 1:
      return a1
    default:
      return GenFib(a0, a1, n - 2) + GenFib(a0, a1, n - 1)
  }
}

/*==============================================================================
 * 4. Reversing Integers (Lecture 3)
 *
 * Write a function ReverseInteger(n) that takes an integer, and returns
 * the integer  formed by reversing the decimal digits of n. For example:
 *      1234 -> 4321
 *      20000 -> 2
 *      1331  -> 1331
 *      -60 -> -6
 *===========================================================================*/

// WRITE YOUR CODE HERE
func ReverseInteger(n int) int {
  var reverseValue int
  for n != 0 {
    reverseValue = n % 10 + reverseValue * 10
    n = n / 10
  }
  return reverseValue
}

/*==============================================================================
 * 5. Growth of a Population (Lecture 3)
 *
 * The size at time t of a population with a birth rate r can be modeled as:
 *
 *      x_t = r*x_{t-1}(1 - x_{t-1})
 *
 * Write a function PopSize(r, x_0, max_t) that prints out the size of the
 * population (x_t) for t=0...max_t, where x_0 is the initial population size.
 * Assume population size and the birth rate parameter r are real numbers; t is
 * an integer.
 *
 * Your function should also return the final population size.
 *============================================================================*/

// WRITE YOUR CODE HERE
func PopSize(r, x0 float64, max_t int) float64 {
  var x float64 = x0
  for i := 0; i < max_t; i++ {
    if x < 0 {
      x = 0
    }
    x = r * x * (1 - x)
    fmt.Println(x)
  }
  return x
}

/*==============================================================================
 * 6. Hailstone function (Lecture 3)
 *
 * The Hailstone function h(n) is defined as n/2 if n is even or 3n+1 if n is
 * odd.  The Hailstone sequence for n is defined by repeatedly applying this
 * function:
 *
 *      h(n),  h(h(n)),  h(h(h(n))), ...
 *
 * It's conjectured that for all n, this sequence eventually returns to 1.
 * Write a function HailstoneReturnsTo1(n) to compute the smallest number of times h
 * must be applied to n before the sequence returns to 1.
 *============================================================================*/

// WRITE YOUR CODE HERE
func h(n int) int {
  if n % 2 == 0 {
    return n / 2
  } else {
    return 3 * n + 1
  }
}

func HailstoneReturnsTo1(n int) int {
  var iterNum int
  for n != 1 {
    n = h(n)
    iterNum++
  }
  return iterNum
}

/*==============================================================================
 * 7. Hailstone function maximum (Lecture 3)
 *
 * Write a function MaxHailstoneValue(n) that takes an integer, and returns the
 * maximum value that the Hailstone sequence:
 *
 *      h(n),  h(h(n)),  h(h(h(n))), ...
 *
 * achieves before it returns to 1.
 *============================================================================*/

// WRITE YOUR CODE HERE
func MaxHailstoneValue(n int) int {
  var maxNum int
  for n != 1 {
    n = h(n)
    if n > maxNum {
      maxNum = n
    }
  }
  return maxNum
}

/*==============================================================================
 * 8. Find the kth digit of an integer n (Lecture 4)
 *
 * Implement a function that takes an integer n, and a positive integer k and
 * returns the k-th decimal digit of n, with digit 1 being the rightmost (least
 * significant) digit.
 *
 *============================================================================*/

// WRITE YOUR CODE HERE
func KthDigit(n, k int) int {
  n = n / int(math.Pow10(k - 1))
  if (n % 10) < 0 {
    return -(n % 10)
  }
  return n % 10
}

/*===========================================================================
 * 9. Hypergeometric distribution (Lecture 4)
 *
 * Write a function Hypergeometric(M,N,n,k) that takes 4 integers and returns
 * a float64 which is the value of the hypergeometric distribution
 *   Pr[red = k] = {M choose k}{N choose n-k} / {M+N choose n}
 *
 * Be careful about overflow: Your funciton should be able to compute
 *      Hypergeometric(5000, 5000, 25, 15)
 *      Hypergeometric(5000, 5000, 50, 15)
 * but not necessarily:
 *      Hypergeometric(5000, 5000, 100, 15)
 *===========================================================================*/

// WRITE YOUR CODE HERE
func continueMultiple(n0, ni int) float64 {
  var out float64 = 1.0
  for i := n0; i >= ni; i-- {
    out *= float64(i)
  }
  return out
}

func factorial(n int) float64 {
  var out float64 = 1.0
  for i := 1; i <= n; i++ {
    out = out * float64(i)
  }
  return out
}

func Hypergeometric(M, N, n, k int) float64 {
  var nton_k_1 float64 = continueMultiple(n, n - k + 1)
  var MtoM_k_1 float64 = continueMultiple(M, M - k + 1)
  var NtoN_n_k_1 float64 = continueMultiple(N, N - n + k + 1)
  var fac_k float64 = factorial(k)
  var M_NtoM_N_n_1 float64 = continueMultiple(M + N, M + N - n + 1)
  return (nton_k_1 * MtoM_k_1 * NtoN_n_k_1) / (fac_k * M_NtoM_N_n_1)
}
