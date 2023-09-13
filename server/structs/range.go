package structs

type Range struct {
	val [2]uint
}

func NewRange(s, e uint) Range {
	var r [2]uint = [2]uint{s, e}

	return Range {
		val: r,
	}
}

func (rng Range) Start() uint  {
	return rng.val[0]
}

func (rng Range) End() uint  {
	return rng.val[1]
}

