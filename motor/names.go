package motor

// "github.com/openai/openai-go/v2"

// Gera um nome aleatório
func GenerateRandomName() string {
	names := []string{"Alice", "Bob", "Charlie", "David", "Emma", "Frank", "Grace", "Hannah"}
	return names[randGenerator.Intn(len(names))]
}

// Gera um sobrenome aleatório
func GenerateRandomSurname() string {
	surnames := []string{"Smith", "Johnson", "Brown", "Davis", "Wilson", "Lee", "Hall", "Moore"}
	return surnames[randGenerator.Intn(len(surnames))]
}

// // Gera nome utilizando a OpenAI API
// func GenerateRandmNameWithOpenAI() (string, error) {
// 	apiKey := "SUA_CHAVE_DE_API_DO_OPENAI" // Substitua pela sua chave de API do OpenAI
// 	client := openai.NewClient(apiKey)

// 	prompt := "Gere um nome aleatório."
// 	options := openai.ChatCompletionRequest{
// 		Messages: []openai.Message{
// 			{Role: "system", Content: "Você é um gerador de nomes."},
// 			{Role: "user", Content: prompt},
// 		},
// 	}

// 	response, err := client.ChatCompletion.Create(options)
// 	if err != nil {
// 		return string{}, err
// 	}

// 	return response.Choices[0].Message.Content, nil
// }
