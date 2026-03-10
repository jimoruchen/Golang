package main

func backspaceCompare1(s string, t string) bool {
	var ch1 []byte
	for i := range s {
		if s[i] != '#' {
			ch1 = append(ch1, s[i])
		} else if len(ch1) > 0 {
			ch1 = ch1[:len(ch1)-1]
		}
	}
	var ch2 []byte
	for i := range t {
		if t[i] != '#' {
			ch2 = append(ch2, t[i])
		} else if len(ch2) > 0 {
			ch2 = ch2[:len(ch2)-1]
		}
	}
	return string(ch1) == string(ch2)
}

func backspaceCompare2(s string, t string) bool {
	return helper(s) == helper(t)
}
func helper(s string) string {
	c := []byte(s)
	left := 0
	for right := 0; right < len(c); right++ {
		if c[right] != '#' {
			c[left] = c[right]
			left++
		} else {
			if left > 0 {
				left--
			}
		}
	}
	return string(c[:left])
}
