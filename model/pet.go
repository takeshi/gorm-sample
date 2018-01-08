package model

type Named struct {
	Name     string
	LastName string
}

func (this Named) name() string {
	return this.Name + " " + this.LastName
}

type Cat struct {
	Named
	Id  int
	Toy Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	Named
	Id  int
	Toy Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	Id        int
	OwnerId   int
	OwnerType string
	Name      string
}
