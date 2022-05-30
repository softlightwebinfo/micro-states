package states

type State struct {
	Id int `json:"id"`
	StateCreate
}

type StateCreate struct {
	Name        *string `json:"name"`
	FkCountryId *string `json:"fk_country_id"`
	Slug        *string `json:"slug"`
	IdState     *string `json:"id_state"`
}

type States struct {
	States []*State `json:"states"`
}
