package helpers

func Set[T comparable](values ...T) map[T]struct{} {
	set := make(map[T]struct{})
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

// HasIntersection 判断两个集合是否有交集
func HasIntersection[T comparable](set1, set2 map[T]struct{}) bool {
	for value := range set1 {
		if _, ok := set2[value]; ok {
			return true
		}
	}
	return false
}

// HasUnion 判断两个集合是否有并集
func HasUnion[T comparable](set1, set2 map[T]struct{}) bool {
	for value := range set1 {
		if _, ok := set2[value]; !ok {
			return false
		}
	}
	return true
}

// HasDifference 判断两个集合是否有差集
func HasDifference[T comparable](set1, set2 map[T]struct{}) bool {
	for value := range set1 {
		if _, ok := set2[value]; !ok {
			return true
		}
	}
	return false
}
