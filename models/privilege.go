package models

// Privilege defines the user privilege to perform actions
type Privilege int8

// Permissions defines the set permitted actions for a user
type Permissions struct {
	CreateEvent bool `json:"createEvent"`
	//TODO: add the rest
}

// GroupPermissions defines the permissions in a group
type GroupPermissions struct {
	GroupName   string `json:"groupName"`
	RoleName    string `json:"roleName"`
	Permissions `json:"permissions"`
}

// Expand populates the permissions from the given privilege bit
func (p Privilege) Expand() Permissions {
	var perm Permissions
	perm.CreateEvent = true
	return perm
}
