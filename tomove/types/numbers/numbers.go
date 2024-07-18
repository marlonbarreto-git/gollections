package numbers

type (
	// Natural represents a natural number.
	// A natural number is a non-negative integer.
	// {0, 1, 2, 3, ...}
	Natural interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}

	// Integer represents an integer number.
	// An integer is a number that can be written without a fractional component.
	// {..., -2, -1, 0, 1, 2, ...}
	Integer interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}

	// Float represents a floating-point number.
	// A floating-point number is a number that has a fractional part.
	// {..., -0.2, -0.1, 0.0, 0.1, 0.2, ...}
	Float interface {
		~float32 | ~float64
	}

	// Number represents a number.
	// A number is a mathematical object used to count, measure, and label.
	Number interface {
		Integer | Natural | Float
	}
)
