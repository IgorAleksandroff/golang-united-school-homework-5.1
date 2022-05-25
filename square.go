package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	// implement me
	s.start.y = s.start.y + int(s.a)
	s.start.x = s.start.x + int(s.a)
	return s.start
}

func (s *Square) Area() uint {
	// implement me
	return s.a * s.a
}

func (s *Square) Perimeter() uint {
	// implement me
	return s.a * 4
}
