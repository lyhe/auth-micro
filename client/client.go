package client

import (
	auth_srv_user "github.com/lyhe/auth-micro/proto/proto_files/user"
	"github.com/micro/go-micro/client"
)

// user 服务客户端
func UserClient() auth_srv_user.UserService {
	return auth_srv_user.NewUserService("go.micro.api.lyhe", client.DefaultClient)
}
