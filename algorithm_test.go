package sternbrocot

import (
    "fmt"
    . "github.com/djwelch/go-matchers"
    . "testing"
)

func IgnoreTestAlgorithm1(t *T) {
    num, den := search(0.1)
    AssertThat(t, num, Equals(1))
    AssertThat(t, den, Equals(10))
}

func IgnoreTestAlgorithm2(t *T) {
    num, den := search(0.2)
    AssertThat(t, num, Equals(1))
    AssertThat(t, den, Equals(5))
}

func IgnoreTestAlgorithm3(t *T) {
    num, den := search(0.3)
    AssertThat(t, num, Equals(3))
    AssertThat(t, den, Equals(10))
}

func IgnoreTestAlgorithm4(t *T) {
    num, den := search(0.5)
    AssertThat(t, num, Equals(1))
    AssertThat(t, den, Equals(2))
}

func IgnoreTestAlgorithm5(t *T) {
    num, den := search(2.25)
    AssertThat(t, num, Equals(9))
    AssertThat(t, den, Equals(4))
}

func IgnoreTestAlgorithm6(t *T) {
    num, den := search(2.125)
    AssertThat(t, num, Equals(17))
    AssertThat(t, den, Equals(8))
}

func IgnoreTestAlgorithm7(t *T) {
    num, den := search(0.1111)
    AssertThat(t, num, Equals(1))
    AssertThat(t, den, Equals(9))
}

func TestPrint(t *T) {
    //cases := []float64 { 1.11, 1.15,1.16,1.20,1.22,1.25,1.36,1.50,1.57,1.66,1.72,1.80,1.90,10.00,101.00,11.00,12.00,13.00,15.00,16.00,19.00,2.00,2.10,2.20,2.25,2.75,2.87,201.00,29.00,3.25,3.50,3.75,34.00,4.00,4.50,41.00,5.50,6.50,67.00,7.00,7.50,81.00,9.00,1.00,1.18,1.28,1.30,1.33,1.40,1.44,1.53,1.61,1.83,126.00,151.00,17.00,2.05,2.37,2.50,2.60,2.62,21.00,23.00,251.00,26.00,3.00,3.12,4.33,5.00,51.00,6.00,8.00,8.50,9.50 }
    cases := []float64 { 1.9, 1.8, 1.83, 3, 2.87, 2.5, 2.37, 3.25, 5.0 }

    for _, dec := range cases {
        fmt.Printf("%2f = ", dec)
        num, den := search(dec-1)
        fmt.Printf("%d / %d\n", num, den)
    }
}
