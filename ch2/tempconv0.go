package tempconv

import (
	"fmt"
)

type Celcius float64
type Farenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celcius) Farenheit{return Farenheit(c*9/5 +32)}

func FToC(f Farenheit) Celcius {return Celcius((f-32)*5/9)}