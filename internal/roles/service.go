package roles

func GetRolesServices() ([]Roles, error) {
	return GetAllRoles()
}

func GetRoleServicesID(id string) (Roles, error) {
	return GetId(id)
}