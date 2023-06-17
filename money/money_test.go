package money

import (
	"errors"

	"testing"
)

func TestParseDecimal(t *testing.T) {
	tt := map[string]struct {
		decimal  string
		expected Decimal
		err      error
	}{
		"2 decimal digits": {
			decimal: "1.52",
			expected: Decimal{
				subunits:  152,
				precision: 2,
			},
			err: nil,
		},
		"no decimal digits": {
			decimal: "10",
			expected: Decimal{
				subunits:  10, // Assuming the subunit is in cents (e.g., dollars to cents)
				precision: 0,
			},
			err: nil,
		},
		"suffix 0 as decimal digits": {
			decimal: "15.50",
			expected: Decimal{
				subunits:  1550,
				precision: 2,
			},
			err: nil,
		},
		"prefix 0 as decimal digits": {
			decimal: "0.75",
			expected: Decimal{
				subunits:  75,
				precision: 2,
			},
			err: nil,
		},
		"multiple of 10": {
			decimal: "100.00",
			expected: Decimal{
				subunits:  10000,
				precision: 2,
			},
			err: nil,
		},
		"invalid decimal part": {
			decimal:  "1.abc",
			expected: Decimal{},
			err:      ErrInvalidDecimal,
		},
		"Not a number": {
			decimal:  "NaN",
			expected: Decimal{},
			err:      ErrInvalidDecimal,
		},
		"empty string": {
			decimal:  "",
			expected: Decimal{},
			err:      ErrInvalidDecimal,
		},
		"too large": {
			decimal:  "1234567890123",
			expected: Decimal{},
			err:      ErrTooLarge,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ParseDecimal(tc.decimal)

			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}

			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
