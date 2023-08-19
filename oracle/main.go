package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	// 导入godror驱动
	go_ora "github.com/sijms/go-ora/v2"
)

// 用户结构，用于存储用户名和原始密码
type User struct {
	Username         string
	OriginalPassword string
}

func main() {
	// 连接参数
	adminUser := flag.String("user", "", "Admin user")
	adminPass := flag.String("password", "", "Admin password")
	host := flag.String("host", "170.106.195.164", "Host")
	port := flag.String("port", "1521", "Port")
	newPass := flag.String("newPass", "", "New password for all users")
	pdbs := flag.String("pdbs", "pdb1,pdb2,pdb3", "Comma separated list of PDBs")
	// 操作类似，改密或验证
	action := flag.String("action", "change", "Action to perform (change|validate)")
	flag.Parse()

	// 连接到CDB
	// 分割PDB名称并迭代
	pdbList := strings.Split(*pdbs, ",")
	// 已经改密的pdb列表
	var changedPDBs []string
	// Remove the unused variable declaration
	// var err error
	for _, pdb := range pdbList {
		db, err := login(*host, *port, *adminUser, *adminPass, pdb)
		if err != nil {
			log.Fatalf("Failed to login: %v", err)
		}
		if *action == "validate" {
			continue // 跳过验证
		}
		if err := changePassword(db, *adminUser, *newPass); err != nil {
			log.Printf("Failed to change password for %s: %v", *adminUser, err)
			if len(changedPDBs) > 0 {
				for _, pdb := range changedPDBs {
					db, err := login(*host, *port, *adminUser, *adminPass, pdb)
					if err != nil {
						log.Fatalf("Failed to login: %v", err)
					}
					err = changePassword(db, *adminUser, *adminUser)
					if err != nil {
						log.Printf("Failed to rollback password for user %s: %v", *adminUser, err)
					}
				}
			}
			return
		}
		changedPDBs = append(changedPDBs, pdb)
	}
	// 退出
	if *action == "validate" {
		fmt.Println("Validation complete")
		return
	}
	fmt.Println("Password update complete")
}

// login
func login(host, port, username, password, serviceName string) (*sql.DB, error) {
	iPort, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	connStr := go_ora.BuildUrl(host, iPort, serviceName, username, password, nil)
	conn, err := sql.Open("oracle", connStr)
	if err != nil {
		return nil, err
	}
	// check for error
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// 切换到指定的PDB
func switchToPDB(db *sql.DB, pdbName string) error {
	_, err := db.Exec(fmt.Sprintf("ALTER SESSION SET CONTAINER = %s", pdbName))
	return err
}

// 获取所有用户
func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT USERNAME FROM DBA_USERS")
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
func changePassword(db *sql.DB, username, newPassword string) error {
	_, err := db.Exec(fmt.Sprintf("ALTER USER %s IDENTIFIED BY %s", username, newPassword))
	return err
}
