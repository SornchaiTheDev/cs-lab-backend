package auth

import "github.com/google/uuid"

func generateState() (string, error) {
	gen, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	strGen := gen.String()

	return strGen, nil
}
