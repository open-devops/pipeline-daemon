package types

type Permission struct {
	AccountId   string `accountId,omitempty`
	RoleId      string `roleId,omitempty`
	AccountName string
	AccountMail string
	AccessToken string
	RoleName    string
}
