package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := atoi(os.Args[1])
	op := os.Args[2]
	b, ok2 := atoi(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	if len(op) != 1 {
		return
	}

	switch op[0] {
	case '+', '-', '*', '/', '%':
		// Valid operator
	default:
		return
	}

	if b == 0 {
		if op[0] == '/' {
			os.Stdout.WriteString("No division by 0\n")
			return
		} else if op[0] == '%' {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
	}

	var result int
	switch op[0] {
	case '+':
		result = a + b
		if (a > 0 && b > 0 && result < a) || (a < 0 && b < 0 && result > a) {
			return
		}
	case '-':
		result = a - b
	case '*':
		result = a * b
		if a != 0 && result/a != b {
			return
		}
	case '/':
		result = a / b
	case '%':
		result = a % b
	}

	// Print result followed by newline (will show as $ on new line with cat -e)
	itoa(result)
	os.Stdout.WriteString("\n")
}

func atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	sign := 1
	i := 0
	if s[0] == '-' {
		sign = -1
		i++
	}

	var n int
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		digit := int(s[i] - '0')
		if n > (1<<63-1-digit)/10 {
			return 0, false
		}
		n = n*10 + digit
	}

	n *= sign
	if n < (-1<<63) || n > (1<<63-1) {
		return 0, false
	}

	return n, true
}

func itoa(n int) {
	if n == 0 {
		os.Stdout.WriteString("0")
		return
	}

	if n < 0 {
		os.Stdout.WriteString("-")
		n = -n
	}

	var digits []byte
	for n > 0 {
		digits = append(digits, byte(n%10+'0'))
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		os.Stdout.Write([]byte{digits[i]})
	}
}
