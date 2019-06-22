package set

type Set struct {
	data []interface{}
	less func(a, b interface{}) bool // a < b == true
}

func New(less func(a, b interface{}) bool) *Set {
	s := &Set{nil, less}
	return s
}

func (s *Set) find(v interface{}) int {
	i, j := 0, len(s.data)
	mid := (i + j) / 2
	for i != j {
		mid = (i + j) / 2
		if s.less(v, s.data[mid]) {
			j = mid - 1
		} else if s.less(s.data[mid], v) {
			i = mid + 1
		} else {
			return mid
		}
	}

	mid = (i + j) / 2
	if mid < len(s.data) && s.data[mid] == v {
		return mid
	}

	return -1
}

func (s *Set) insert(i int, v interface{}) {
	s.data = append(s.data, v)
	copy(s.data[i+1:], s.data[i:])
	s.data[i] = v
}

func (s *Set) Exists(v interface{}) bool {
	return s.find(v) != -1
}

func (s *Set) Empty() bool {
	return len(s.data) == 0
}

func (s *Set) Insert(v interface{}) {
	i, j := 0, len(s.data)
	index := 0
	for j-i > 1 {

		mid := (i + j) / 2

		if s.less(v, s.data[mid]) {
			j = mid
		} else {
			i = mid
		}

		index = mid
	}

	if j-i == 1 {
		if s.less(v, s.data[i]) {
			index = i
		} else {
			index = j
		}
	}

	s.insert(index, v)
}

func (s *Set) erase(i int) {
	copy(s.data[i:], s.data[i+1:])
	s.data[len(s.data)-1] = nil
	s.data = s.data[:len(s.data)-1]
}

func (s *Set) Erase(v interface{}) {
	i := s.find(v)
	if i >= 0 {
		s.erase(i)
	}
}
