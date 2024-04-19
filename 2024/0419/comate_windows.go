package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func GetSidAndPrimaryGroupSid(expectedUser string) (string, string, error) {
	// Retrieve the user's SID
	userSid, _, domain, err := windows.LookupSID("", expectedUser)
	if err != nil {
		return "", "", fmt.Errorf("failed to retrieve user SID: %w", err)
	}

	// Retrieve the primary group SID associated with the user
	token, err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_QUERY_SOURCE)
	if err != nil {
		return "", "", fmt.Errorf("failed to open process token: %w", err)
	}
	defer token.Close()

	tokenUser, err := token.GetTokenUser()
	if err != nil {
		return "", "", fmt.Errorf("failed to get token user: %w", err)
	}

	userInfo, err := windows.GetUser(tokenUser.User.Sid)
	if err != nil {
		return "", "", fmt.Errorf("failed to get user information: %w", err)
	}

	primaryGroupId := userInfo.PrimaryGroupId
	primaryGroupSid, _, err := windows.LookupSID(domain, fmt.Sprintf("S-1-5-%d", primaryGroupId))
	if err != nil {
		return "", "", fmt.Errorf("failed to retrieve primary group SID: %w", err)
	}

	return userSid.String(), primaryGroupSid.String(), nil
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
