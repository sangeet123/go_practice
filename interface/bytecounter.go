package main

import (
	"fmt"
)

type ByteCounter int

// This function implements Writer interface with Write method
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // conver int to ByteCounter
	return len(p), nil
}

// Now ByteCounter can be passed to Fprintf since it implements Writer interface
// func Fprintf(w io.Writer, format string, args ...interface{})

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)

	c = 0 // reset

	var name = "Ryan"
	fmt.Fprintf(&c, "Hello, %s", name) // Fprintf internally calls c.Write method so the length will b 11
	fmt.Println(c)
}
