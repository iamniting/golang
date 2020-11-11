package main

func minOperations(n int) int {

	/*`
	if n is odd, suppose n=5.
	The array is :
	1 3 5 7 9.
	Here, we will have the middle element as 5.
	We take 2 from 7 and add to 3 to make each one 5.
	We take 4 from 9 and add to 1 to make each one 5.
	Total steps: 2+4=6. (sum of first n/2 even numbers)
	Sum of first k EVEN numbers = k*(k+1)

	if n is even, suppose n=6.
	The array is :
	1 3 5 7 9 11.
	Here, we will have the middle element as (5+7)/2=6.
	We take 1 from 7 and add to 5 to make each one 6.
	We take 3 from 9 and add to 3 to make each one 6.
	We take 5 from 11 and add to 1 to make each one 6.
	Total steps: 1+3+5=9. (sum of first n/2 odd numbers)
	Sum of first k ODD numbers = k * k.
	`*/

	count := n / 2

	return count * (count + n&1)
}
