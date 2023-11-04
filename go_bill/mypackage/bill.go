package mypackage

type Faker struct {
	Name string
}

func (f Faker) getName() string {
	return f.Name
}

