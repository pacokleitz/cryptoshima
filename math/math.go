package math

func Mod(a, b uint16) uint16 {
	return (a%b + b) % b
}
