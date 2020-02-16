package simplemath

import (
	"errors"
	"math"
	"strconv"
)

func main() {
	return
}

func negExponent(x float64, y float64) float64 {
	result := 1.0
	y = y * -1
	for i := 1.0; i <= y; i++ {
		result = result * x
	}

	return result

}

// Strconvert converts any int to a string
func Strconvert(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

// Add adds x and y, and returns the sum
func Add(x float64, y float64) float64 {
	result := 0.0
	result = x + y
	return result
}

// Subtract subtracts y from x and returns the difference
func Subtract(x float64, y float64) float64 {
	result := 0.0
	result = x - y
	return result
}

// Multiply multiplies x by y and returns the result
func Multiply(x float64, y float64) float64 {
	result := 0.0
	result = x * y
	return result
}

// Divide divides x by y and returns the result as a float64
func Divide(x float64, y float64) float64 {
	result := 0.0
	result = x / y
	return result
}

// Sqrt returns square root of x as a float64
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Nonreal answer")
	}

	return math.Sqrt(x), nil
}

// Exponent returns x^y as a float64
func Exponent(x float64, y float64) float64 {
	if y > 0.0 {
		result := 1.0

		for i := 1.0; i <= y; i++ {
			result = result * x
		}

		return result
	}
	if y < 0.0 {
		result := 1.0
		result = 1 / negExponent(x, y)

		return result
	}

	return 0
}
