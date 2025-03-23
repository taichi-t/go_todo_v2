type Money struct {
	amount   int64
	currency string
}

// 値オブジェクトは不変
func NewMoney(amount int64, currency string) (*Money, error) {
	// ビジネスルール：金額は0以上
	if amount < 0 {
		return nil, errors.ErrNegativeAmount
	}
	return &Money{amount: amount, currency: currency}, nil
}

// 演算は新しいインスタンスを返す
func (m *Money) Add(other *Money) (*Money, error) {
	if m.currency != other.currency {
		return nil, errors.ErrDifferentCurrency
	}
	return NewMoney(m.amount+other.amount, m.currency)
} 