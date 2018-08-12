package tempconv

import "fmt"

type Celsius float64
type Farenheit float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g0c", c)
}

func (f Farenheit) String() string {
	return fmt.Sprintf("%g0f", f)
}
