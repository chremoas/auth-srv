package model

func (Alliance) TableName() string {
	return "alliances"
}

func (Corporation) TableName() string {
	return "corporations"
}

func (Character) TableName() string {
	return "characters"
}

func (Role) TableName() string {
	return "roles"
}

func (User) TableName() string {
	return "users"
}
