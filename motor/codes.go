package motor

import "fmt"

// Gera um código numérico aleatório com a quantidade de dígitos informados
func GenerateRandomNumericCode(digits int) string {
	code := ""
	for i := 0; i < digits; i++ {
		code += fmt.Sprintf("%d", randGenerator.Intn(10))
	}
	return code
}
