package piscine

func StrRev(s string) string {
	rev := ""
	for _, v := range s {
		rev = string(v) + rev
	}
	return rev
}
