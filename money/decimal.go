package money

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// ErrInvalidDecimal is returned if the decimal is malformed.
	ErrInvalidDecimal = Error("unable to convert the decimal")
	// ErrTooLarge is returned if the quantity is too large - this would cause floating point precision errors.
	ErrTooLarge = Error("quantity over 10^12 is too large")
)

// Decimal is capable of storing a floating-point number with fixed precision.
// example: 1.52 = 152 * 10^(-2) will be stored as {152, 2}
type Decimal struct {
	// subunits is the amount of subunits. Multiply it by the precision to get the real value
	subunits int64
	// Number of "subunits" in a unit, expressed as a power of 10
	precision byte
}

// ParseDecimal converts a string into its Decimal representation.
// It assumes  there is up to one decimal separator, and that the separator is "."(full stop character)
func ParseDecimal(value string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")

	// maxDecimal is the number of digigts in a thousand billion.
	const maxDecimal = 12

	if len(intPart) > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	precision := byte(len(fracPart))

	return Decimal{subunits: subunits, precision: precision}, nil

}

// simplify the decimal ex: 1.50 -> 1.5
func (d *Decimal) Simplify() {
	for d.subunits%10 == 0 && d.precision > 0 {
		d.precision--
		d.subunits /= 10
	}
}
