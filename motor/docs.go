package motor

// Gera um número de CPF aleatório
func GenerateRandomCPF() string {
	return GenerateRandomNumericCode(3) + "." + GenerateRandomNumericCode(3) + "." + GenerateRandomNumericCode(3) + "-" + GenerateRandomNumericCode(2)
}

// Gera um número de CNPJ aleatório
func GenerateRandomCNPJ() string {
	return GenerateRandomNumericCode(2) + "." + GenerateRandomNumericCode(3) + "." + GenerateRandomNumericCode(3) + "/0001-" + GenerateRandomNumericCode(2)
}
