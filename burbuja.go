package main

func main() {
	s := []int64{54, 26, 93, 17, 77, 31, 44, 55, 20}
	Burbuja(s)
}

func Burbuja(s []int64) {
	var auxiliar int64
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] < s[j] {
				auxiliar = s[i]
				s[i] = s[j]
				s[j] = auxiliar
			}
		}
	}
}
