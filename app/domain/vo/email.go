package value

type Email string

func NewEmail(address string) (Email, error) {
	// メールアドレスのバリデーション
	if !isValidEmail(address) {
		return "", ErrInvalidEmail
	}
	return Email(address), nil
}

func (e Email) String() string {
	return string(e)
}
