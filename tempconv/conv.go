package tempconv

func FToC(f Farenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToF(c Celsius) Farenheit {
	return Farenheit(c*9/5 + 32)
}
