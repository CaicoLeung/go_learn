package tempconv

// C2F converts a Celsius temperature to Fahrenheit.
func C2F(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// F2C converts a Fahrenheit temperature to Celsius.
func F2C(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
