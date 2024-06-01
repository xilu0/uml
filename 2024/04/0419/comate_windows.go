package main

import (
	"fmt"
	"os/user"
)

// get 用户 SID 和 所属主组 SID
// 参考: https://github.com/golang/go/blob/master/src/syscall/security_windows.go#L165
func GetSidAndPrimaryGroupSid(username string) (string, string, error) {
	var err error
	var u *user.User

	if username == "" {
		u, err = user.Current()
	} else {
		u, err = user.Lookup(username)
	}
	if err != nil {
		return "", "", err
	}

	sid := u.Uid // 修改此处，使用u.Uid代替u.Id()

	// 获取用户的主组名称
	GroupIds, err := u.GroupIds()
	if err != nil {
		return "", "", err
	}

	// 使用主组名称查找组ID
	primaryGroup, err := user.LookupGroup(GroupIds[0])
	if err != nil {
		return "", "", err
	}

	return string(sid), string(primaryGroup.Gid), nil
}

func main() {
	sid, primaryGroupSid, err := GetSidAndPrimaryGroupSid("username")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User SID: %s\n", sid)
	fmt.Printf("Primary Group SID: %s\n", primaryGroupSid)
}
