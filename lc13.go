package main

func roman(s string) int {
	Map := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	};
	res := 0;
	for i := 0; i < len(s)-1; i++ {
		if (Map[string(s[i])] < Map[string(s[i+1])]) {
			res -= Map[string(s[i])];
		} else {
			res += Map[string(s[i])];
		}
	}
	res += Map[string(s[len(s)-1])];
	return res;
}
