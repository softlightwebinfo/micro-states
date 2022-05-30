package towns

type Town struct {
	IdTown int `json:"id_town"`
	TownCreate
}

type TownCreate struct {
	Name      string  `json:"name"`
	FkStateId *string `json:"fk_state_id"`
	Slug      *string `json:"slug"`
	CodTown   *string `json:"cod_town"`
	Dc        *string `json:"dc"`
}

type Towns struct {
	Towns []*Town `json:"towns"`
}
