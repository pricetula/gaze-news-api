package utils

// SpltStr2UIntSlc splits a comma separated string to a slice of uint values
func SpltStr2IntSlc(s string) (r []int) {
	slc := SpltStr2Slc(s, ",")

	if len(slc[0]) == 0 && s == slc[0] {
		return []int{}
	}

	for _, v := range slc {
		d := Str2Int(v)

		r = append(r, int(d))
	}

	return
}
