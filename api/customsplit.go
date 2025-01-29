package api

func Split(s string) []string {
	var temp string
	var fin []string
	for i, c := range s {
		if c == '-' {
			fin = append(fin, temp)
			temp = ""
		} else {
			temp += string(c)
		}
		if i == len(s)-1 {
			fin = append(fin, temp)
		}
	}
	return fin
}
