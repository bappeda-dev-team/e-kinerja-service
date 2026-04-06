package roles

func GetRolesServices() ([]Roles, error) {
	return GetAllRoles()
}

func GetRoleServicesID(id string) (Roles, error) {
	return GetId(id)
}

func CreateRoleServices(name string, description string) (Roles, error) {
	return Create(name, description)
}

func UpdateRoleServices(id string, name string, description string) (Roles, error) {
	return Update(id, name, description)
}

func DeleteRoleServices(id string) error {
	return Delete(id)
}
