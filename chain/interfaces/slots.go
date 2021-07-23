package interfaces

type Slot uint64

func (s Slot) Mul(x uint64) Slot {
	res, err := s.SafeMul(x)
	if err != nil {
		panic(err.Error())
	}
	return res
}

// SafeMul multiplies slot by x.
// In case of arithmetic issues (overflow/underflow/div by zero) error is returned.
func (s Slot) SafeMul(x uint64) (Slot, error) {
	res, err := Mul64(uint64(s), x)
	return Slot(res), err
}

// MulSlot multiplies slot by another slot.
// In case of arithmetic issues (overflow/underflow/div by zero) panic is thrown.
func (s Slot) MulSlot(x Slot) Slot {
	return s.Mul(uint64(x))
}

// SafeMulSlot multiplies slot by another slot.
// In case of arithmetic issues (overflow/underflow/div by zero) error is returned.
func (s Slot) SafeMulSlot(x Slot) (Slot, error) {
	return s.SafeMul(uint64(x))
}

// Div divides slot by x.
// In case of arithmetic issues (overflow/underflow/div by zero) panic is thrown.
func (s Slot) Div(x uint64) Slot {
	res, err := s.SafeDiv(x)
	if err != nil {
		panic(err.Error())
	}
	return res
}

// SafeDiv divides slot by x.
// In case of arithmetic issues (overflow/underflow/div by zero) error is returned.
func (s Slot) SafeDiv(x uint64) (Slot, error) {
	res, err := Div64(uint64(s), x)
	return Slot(res), err
}

// DivSlot divides slot by another slot.
// In case of arithmetic issues (overflow/underflow/div by zero) panic is thrown.
func (s Slot) DivSlot(x Slot) Slot {
	return s.Div(uint64(x))
}

// SafeDivSlot divides slot by another slot.
// In case of arithmetic issues (overflow/underflow/div by zero) error is returned.
func (s Slot) SafeDivSlot(x Slot) (Slot, error) {
	return s.SafeDiv(uint64(x))
}

// Add increases slot by x.
// In case of arithmetic issues (overflow/underflow/div by zero) panic is thrown.
func (s Slot) Add(x uint64) Slot {
	res, err := s.SafeAdd(x)
	if err != nil {
		panic(err.Error())
	}
	return res
}

// SafeAdd increases slot by x.
// In case of arithmetic issues (overflow/underflow/div by zero) error is returned.
func (s Slot) SafeAdd(x uint64) (Slot, error) {
	res, err := Add64(uint64(s), x)
	return Slot(res), err
}



// AddSlot increases slot by another slot.
// In case of arithmetic issues (overflow/underflow/div by zero) panic is thrown.
func (s Slot) AddSlot(x Slot) Slot {
	return s.Add(uint64(x))
}

// SafeAddSlot increases slot by another slot.
// In case of arithmetic issues (overflow/underflow/div by zero) error is returned.
func (s Slot) SafeAddSlot(x Slot) (Slot, error) {
	return s.SafeAdd(uint64(x))
}



// MaxSlot returns the larger of the two slots.
func MaxSlot(a, b Slot) Slot {
	if a > b {
		return a
	}
	return b
}

// MinSlot returns the smaller of the two slots.
func MinSlot(a, b Slot) Slot {
	if a < b {
		return a
	}
	return b
}