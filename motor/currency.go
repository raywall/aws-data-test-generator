package motor

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomCurrencyValue(limit int, minValue int) (string, error) {
	source := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(source)

	if limit < minValue {
		return "", errors.New("O valor máximo é menor do que o valor mínimo!")
	}

	// Gere um número inteiro aleatório no intervalo de 0 a limit (inclusive)
	intValue := randGenerator.Intn(limit+1) + minValue

	return fmt.Sprintf("%.2f", float64(intValue)/100.0), nil
}
