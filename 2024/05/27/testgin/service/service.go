// service.go
package service

type User struct {
	ID   int
	Name string
}
type UserService interface {
	GetUser(id int) (User, error)
}

type UserServiceImpl struct{}

func (s *UserServiceImpl) GetUser(id int) (User, error) {
	// 模拟获取用户信息
	return User{ID: id, Name: "John Doe"}, nil
}
