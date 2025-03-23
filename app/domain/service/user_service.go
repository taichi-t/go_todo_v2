package service

// 複数のエンティティにまたがるドメインロジック
type UserService struct{}

// ユーザー名の重複チェックなど、エンティティ単体では判断できないロジック
func (s *UserService) IsUsernameDuplicate(username string, users []entity.User) bool {
	for _, user := range users {
		if user.Name == username {
			return true
		}
	}
	return false
}
