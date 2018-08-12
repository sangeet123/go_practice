package main

import(
"fmt"
)

func swap(p, q *int){
	*p = *p ^ *q
	*q = *p ^ *q
	*p = *p ^ *q 
}

func main(){
	a := 2
	b := 3

    fmt.Printf("Before swapping a=%d, b = %d\n", a, b)
	
	swap(&a, &b)

	fmt.Printf("After swapping a=%d, b = %d\n", a, b)
}