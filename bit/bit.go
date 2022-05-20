package bit

type Bit int

func Add(a, b Bit) Bit {
	return (a | b)
}

func Remove(a, b Bit) Bit {
	return (a ^ b)
}

func Has(a, b Bit) bool {
	return ((a & b) == b)
}
