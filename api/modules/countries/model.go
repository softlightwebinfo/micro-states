package countries

type Country struct {
	Id int `json:"id"`
	CountryCreate
}

type CountryCreate struct {
	Iso  *string `json:"iso"`
	Name *string `json:"name"`
}

type Countries struct {
	Countries []*Country `json:"countries"`
}
