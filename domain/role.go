package domain

import "strings"

type RolePermissions struct {
	rolePermissions map[string][]string
}

func (p RolePermissions) IsAuthorizedFor(role string, routeName string) bool {
	permissions := p.rolePermissions[role]

	for _, v := range permissions {
		if v == strings.TrimSpace(routeName) {
			return true
		}
	}
	return false
}

func GetRolePermissions() RolePermissions {
	return RolePermissions{map[string][]string{
		"admin": {"GetAllCustomers", "GetAllAccounts", "CreateCustomer", "GetCustomerById", "UpdateCustomerByID", "DeleteCustomerById", "CreateAccount", "GetAccountsByCustomerId", "GetAccountsByAccountId", "UpdateAccountByAccountId", "DeleteAccountById", "CreateTransaction", "GetAllTransactionsByAccountId"},
		"user":  {"GetCustomerById", "GetAccountsByCustomerId"},
	}}
}
