package tempconv

import "fmt"

type Celsius float64    // 摄氏度
type Fahrenheit float64 // 华氏度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 冰点
	BoilingC      Celsius = 100     // 沸点
)

/* Celsius类型的参数c出现在了函数名的前面，表示声明的是Celsius类型的一个名叫String的方法，该方法返回该类型对象c带着°C温度单位的字符串 */
func (c Celsius) String() string { return fmt.Sprintf("%g °C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g °F", f) }
