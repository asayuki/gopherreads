package stores

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func comparePassword(hashed string, plain []byte) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), plain) == nil
}
