package sumoftwo

func v1(a []int, b []int, v int) bool {
	for i := range a {
		for j := range b {
			if (a[i] + b[j]) == v {
				return true
			}
		}
	}
	return false
}

var keyExists = struct{}{}

func v2(a []int, b []int, v int) bool {
	set := make(map[int]struct{})
	for _, vb := range b {
		set[vb] = keyExists
	}
	for _, va := range a {
		target := v - va
		if _, ok := set[target]; ok {
			return true
		}
	}
	return false
}

func v3(a []int, b []int, v int) bool {
	set := make(map[int]struct{}, len(b))
	for _, vb := range b {
		set[vb] = keyExists
	}
	for _, va := range a {
		target := v - va
		if _, ok := set[target]; ok {
			return true
		}
	}
	return false
}
