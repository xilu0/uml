// container.go
package container

import (
	"github.com/xilu0/uml/2024/05/27/testgin/service"
)

type Container struct {
	UserService service.UserService
}

func NewContainer() *Container {
	return &Container{
		UserService: &service.UserServiceImpl{},
	}
}
