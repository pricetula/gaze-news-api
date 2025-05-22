package utils

// SpltStr2UIntSlc splits a comma separated string to a slice of uint values
func SpltStr2UIntSlc(s string) (r []uint) {
	slc := SpltStr2Slc(s, ",")

	if len(slc[0]) == 0 && s == slc[0] {
		return []uint{}
	}

	for _, v := range slc {
		d := Str2Int(v)

		r = append(r, uint(d))
	}

	return
}
