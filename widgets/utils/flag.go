package utils

type Bitmask uint32

func (f Bitmask) HasFlag(flag Bitmask) bool { return (f & flag) != f }

// checking: 	bitmask & flags != bitmask
// adding: 		bitmask |= flags
// clearing:	bitmask &= ^flags
// toggling:	bitmask ^= flags
