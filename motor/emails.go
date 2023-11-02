package motor

import "fmt"

// Gera um endereço de e-mail aleatório
func GenerateRandomEmail() string {
	domains := []string{"gmail.com", "yahoo.com", "hotmail.com", "example.com", "testmail.com"}
	username := GenerateRandomName() + GenerateRandomNumericCode(3)
	domain := domains[randGenerator.Intn(len(domains))]
	return fmt.Sprintf("%s@%s", username, domain)
}
