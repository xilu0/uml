package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func GetUidAndGid(u string) (int, int, error) {
	// 用用户名获取用户信息
	userInfo, err := user.Lookup(u)
	if err != nil {
		return 0, 0, err
	}
	uid, err := strconv.Atoi(userInfo.Uid)
	if err != nil {
		return 0, 0, err
	}
	gid, err := strconv.Atoi(userInfo.Gid)
	if err != nil {
		return 0, 0, err
	}
	return uid, gid, nil
}

func checkAndFixFilePermissions(filePath, expectedUser, expectedPermissions string) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("获取文件信息失败: %v", err)
	}

	uid, gid, err := GetUidAndGid(expectedUser)
	if err != nil {
		return fmt.Errorf("get user info error: %v", err)
	}

	// 使用syscall包替代unix包，因为unix包不是Go标准库的一部分
	sys := syscall.Stat_t{}
	err = syscall.Stat(filePath, &sys)
	if err != nil {
		return err
	}

	if sys.Uid != uint32(uid) || sys.Gid != uint32(gid) {
		if err := os.Chown(filePath, int(uid), int(gid)); err != nil {
			return fmt.Errorf("更改文件所有者失败: %v", err)
		}
	}

	expectedPerm, err := strconv.ParseUint(expectedPermissions, 8, 32)
	if err != nil {
		return fmt.Errorf("解析权限失败: %v", err)
	}

	if fileInfo.Mode().Perm() != os.FileMode(expectedPerm) {
		if err := os.Chmod(filePath, os.FileMode(expectedPerm)); err != nil {
			return fmt.Errorf("更改文件权限失败: %v", err)
		}
	}

	return nil
}
