package fizzbuzz

const (
	fizz     = 3
	buzz     = 5
)

func Fizzbuzz(angka int) string {
	if angka%fizz == 0 {
		return "FIZZ"
	}
	if angka%buzz == 0 {
		return "BUZZ"
	}
	if angka%fizz == 0 && angka%buzz == 0 {
		return "FIZZBUZZ"
	}
	return string(rune(angka))
}

