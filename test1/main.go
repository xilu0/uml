package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	// 导入godror驱动
	go_ora "github.com/sijms/go-ora/v2"
)

// 用户结构，用于存储用户名和原始密码
type User struct {
	Username         string
	OriginalPassword string
}

// 数据库连接类
type DBConnection struct {
	db *sql.DB
}

// 连接到数据库
func (conn *DBConnection) Connect(host string, port int, serviceName string, username string, password string) error {
	connStr := go_ora.BuildUrl(host, port, serviceName, username, password, nil)
	db, err := sql.Open("oracle", connStr)
	if err != nil {
		return err
	}
	// check for error
	err = db.Ping()
	if err != nil {
		return err
	}
	conn.db = db
	return nil
}

// 切换到指定的PDB
func (conn *DBConnection) SwitchToPDB(pdbName string) error {
	_, err := conn.db.Exec(fmt.Sprintf("ALTER SESSION SET CONTAINER = %s", pdbName))
	return err
}

// 获取所有用户
func (conn *DBConnection) GetUsers() ([]User, error) {
	rows, err := conn.db.Query("SELECT USERNAME FROM DBA_USERS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			return nil, err
		}

		// 这里可以添加逻辑来获取原始密码，如果有访问权限
		users = append(users, User{Username: username})
	}

	return users, nil
}

// 更改指定用户的密码
func (conn *DBConnection) ChangePassword(username string, newPassword string) error {
	_, err := conn.db.Exec(fmt.Sprintf("ALTER USER %s IDENTIFIED BY %s", username, newPassword))
	return err
}

// 密码修改类
type PasswordChanger struct {
	conn *DBConnection
}

// 修改指定PDB中所有用户的密码
func (changer *PasswordChanger) ChangePasswords(pdbName string, adminUser string, adminPass string, newPass string) error {
	err := changer.conn.Connect("170.106.195.164", 1521, pdbName, adminUser, adminPass)
	if err != nil {
		return err
	}
	users, err := changer.conn.GetUsers()
	if err != nil {
		return err
	}
	for _, user := range users {
		err = changer.conn.ChangePassword(user.Username, newPass)
		if err != nil {
			return err
		}
	}
	return nil
}

// 密码验证类
type PasswordValidator struct {
	conn *DBConnection
}

// 验证指定PDB中所有用户的密码是否一致
func (validator *PasswordValidator) ValidatePasswords(pdbName string, adminUser string, adminPass string, expectedPass string) (bool, error) {
	err := validator.conn.Connect("170.106.195.164", 1521, pdbName, adminUser, adminPass)
	if err != nil {
		return false, err
	}
	users, err := validator.conn.GetUsers()
	if err != nil {
		return false, err
	}
	for _, user := range users {
		if user.OriginalPassword != expectedPass {
			return false, nil
		}
	}
	return true, nil
}

// 密码修改器类
type PasswordUpdater struct {
	changer     *PasswordChanger
	validator   *PasswordValidator
	adminUser   string
	adminPass   string
	newPass     string
	pdbList     []string
	changedPDBs []string
}

// 修改所有PDB中所有用户的密码
func (updater *PasswordUpdater) UpdatePasswords() error {
	for _, pdb := range updater.pdbList {
		err := updater.changer.ChangePasswords(pdb, updater.adminUser, updater.adminPass, updater.newPass)
		if err != nil {
			log.Printf("Failed to change password for PDB %s: %v", pdb, err)
			if len(updater.changedPDBs) > 0 {
				for _, pdb := range updater.changedPDBs {
					err = updater.changer.ChangePasswords(pdb, updater.adminUser, updater.adminPass, updater.adminUser)
					if err != nil {
						log.Printf("Failed to rollback password for PDB %s: %v", pdb, err)
					}
				}
			}
			return err
		}
		updater.changedPDBs = append(updater.changedPDBs, pdb)
	}
	return nil
}

// 验证所有PDB中所有用户的密码是否一致
func (updater *PasswordUpdater) ValidatePasswords() (bool, error) {
	for _, pdb := range updater.pdbList {
		valid, err := updater.validator.ValidatePasswords(pdb, updater.adminUser, updater.adminPass, updater.newPass)
		if err != nil {
			return false, err
		}
		if !valid {
			return false, nil
		}
	}
	return true, nil
}

// 主程序类
type MainProgram struct {
	updater *PasswordUpdater
}

// 主程序入口
func (program *MainProgram) Run() {
	// 连接参数
	adminUser := "admin"
	adminPass := "password"
	newPass := "newpassword"
	pdbs := "pdb1,pdb2,pdb3"
	// 操作类似，改密或验证
	action := "change"

	// 连接到CDB
	// 分割PDB名称并迭代
	pdbList := strings.Split(pdbs, ",")
	program.updater.adminUser = adminUser
	program.updater.adminPass = adminPass
	program.updater.newPass = newPass
	program.updater.pdbList = pdbList
	if action == "validate" {
		valid, err := program.updater.ValidatePasswords()
		if err != nil {
			log.Fatalf("Failed to validate passwords: %v", err)
		}
		if valid {
			fmt.Println("Validation complete")
		} else {
			fmt.Println("Validation failed")
		}
	} else {
		err := program.updater.UpdatePasswords()
		if err != nil {
			log.Fatalf("Failed to update passwords: %v", err)
		}
		fmt.Println("Password update complete")
	}
}

func main() {
	conn := &DBConnection{}
	changer := &PasswordChanger{conn: conn}
	validator := &PasswordValidator{conn: conn}
	updater := &PasswordUpdater{changer: changer, validator: validator}
	program := &MainProgram{updater: updater}
	program.Run()
}
