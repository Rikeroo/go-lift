package set

type Set struct {
	SetNumber int
	Weight    float32
	Reps      int
}

func (s *Set) Update(selector string, value float32) {
	switch selector {
	case "weight":
		s.Weight = value
	case "reps":
		s.Reps = int(value)
	}
}

func (s Set) Get(setNum int) [3]float32 {
	return [3]float32{float32(s.SetNumber), s.Weight, float32(s.Reps)}
}
