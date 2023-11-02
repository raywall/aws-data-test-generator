package motor

import "fmt"

// Gera um número de telefone aleatório no formato (XX) XXXX-XXXX
func GenerateRandomPhoneNumber() string {
	areaCode := randGenerator.Intn(90) + 10 // Gera um código de área entre 10 e 99
	numberPart1 := randGenerator.Intn(9000) + 1000
	numberPart2 := randGenerator.Intn(9000) + 1000
	return fmt.Sprintf("(%02d) %04d-%04d", areaCode, numberPart1, numberPart2)
}
