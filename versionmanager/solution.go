package template

type version struct {
	number           int
	isPrevCompatible bool
}

type serviceManager struct {
	versions map[int]version
}

func NewVersionManager() serviceManager {
	return serviceManager{versions: map[int]version{}}
}

func (s *serviceManager) addVersion(number int, isPrevCompatible bool) {
	s.versions[number] = version{number: number, isPrevCompatible: isPrevCompatible}
}

func (s *serviceManager) isCompatible(src int, target int) bool {
	v, ok := s.versions[target]
	if !ok {
		return false
	}
	if !v.isPrevCompatible {
		return false
	}
	prev := target - 1
	if prev == src {
		return true
	}
	return s.isCompatible(src, prev)
}
