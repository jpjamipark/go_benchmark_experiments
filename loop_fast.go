package main

func countFor1(s string) int {
	result := 0
	for i := 0; i < len(s); {
		if s[i]&1 == 1 {
			result += 1
		}
		i++
	}

	return result
}

func main() {
  countFor1("a")
}
