package weight

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string    { return fmt.Sprintf("%gpounds", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }

func PoundToKilogram(p Pound) Kilogram {
	return Kilogram(0.45359237 * p)
}

func KilogramToPound(kg Kilogram) Pound {
	return Pound(kg / 0.45359237)
}
