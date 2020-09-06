package moon

import (
	"fmt"
	"testing"
	"time"

	"github.com/starainrt/astro/basic"
)

func TestMoonI(t *testing.T) {
	fmt.Println(basic.MoonR(2465445.9755443))
}

func Test_NewCalc(t *testing.T) {
	fmt.Printf("%.14f", basic.MoonCalcNew(2, 2451546.0))
}

func TestDownTime(t *testing.T) {
	fmt.Println(DownTime(time.Now(), 114, 21, 8, false))
}

func TestRiseTime(t *testing.T) {
	fmt.Println(RiseTime(time.Now(), 114, 21, 8, false))
}
