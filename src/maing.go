package main

import "time"

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

type User struct {
	id int
}

var GetUserById = func(id int) (User, error) {
	time.Sleep(3 * time.Second)
	return User{}, nil
}
