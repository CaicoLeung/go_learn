package tempconv

import (
	"fmt"
	"testing"
)

/* go test -v ./ch2/tempconv/ -test.run TestExampleOne */
func TestExampleOne(t *testing.T) {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := C2F(BoilingC)
	fmt.Printf("%g\n", boilingF-C2F(FreezingC))
	fmt.Printf("%g\n", C2F(AbsoluteZeroC))
}

func TestExampleTwo(t *testing.T)  {
	c := F2C(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c) // 不用显式调用String()
	fmt.Printf("%s\n", c) // 不用显式调用String()
	fmt.Printf("%g\n", c) // 不用显式调用String()
	fmt.Println(c)
	fmt.Println(float64(c))
}
