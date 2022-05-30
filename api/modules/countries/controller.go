package countries

import (
	"fmt"
	states "micro-states/modules/states"
	towns "micro-states/modules/towns"
	"micro-states/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAll(c *gin.Context) {
	row, err := settings.DB.Query("SELECT id, iso, name FROM countries ORDER BY name asc")

	countries := Countries{}

	if err != nil {
		fmt.Println("error:countries:controller:getAll:1", err)
		c.JSON(http.StatusNoContent, gin.H{"countries": nil})
		return
	}

	for row.Next() {
		var country Country

		err := row.Scan(
			&country.Id,
			&country.Iso,
			&country.Name,
		)

		if err != nil {
			fmt.Println("error:countries:controller:getAll:2", err)
			continue
		}

		countries.Countries = append(countries.Countries, &country)
	}

	c.JSON(http.StatusOK, countries)
}

func GetCountry(c *gin.Context) {
	country := c.Params.ByName("country")
	var countryData Country
	err := settings.DB.QueryRow("SELECT id, iso, name FROM countries WHERE id=$1 ORDER BY name asc", &country).
		Scan(
			&countryData.Id,
			&countryData.Iso,
			&countryData.Name,
		)

	if err != nil {
		fmt.Println("error:countries:controller:GetCountry:1", err)
		c.JSON(http.StatusNotFound, gin.H{"country": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"country": countryData})
}

func getStates(c *gin.Context) {
	country := c.Params.ByName("country")
	row, err := settings.DB.Query("SELECT id, name, fk_country_id, slug, id_state FROM states WHERE fk_country_id=$1 ORDER BY name asc", &country)

	countries := states.States{}

	if err != nil {
		fmt.Println("error:countries:controller:getStates:1", err)
		c.JSON(http.StatusNotFound, gin.H{"states": nil})
		return
	}

	for row.Next() {
		var country states.State

		err := row.Scan(
			&country.Id,
			&country.Name,
			&country.FkCountryId,
			&country.Slug,
			&country.IdState,
		)

		if err != nil {
			fmt.Println("error:countries:controller:getAll:2", err)
			continue
		}

		countries.States = append(countries.States, &country)
	}

	c.JSON(http.StatusOK, countries)
}

func getState(c *gin.Context) {
	countryId := c.Params.ByName("country")
	stateId := c.Params.ByName("state")

	state := states.State{}

	err := settings.DB.QueryRow(
		"SELECT id, name, fk_country_id, slug, id_state FROM states WHERE fk_country_id=$1 and id_state=$2 ORDER BY name asc",
		&countryId, &stateId,
	).Scan(
		&state.Id,
		&state.Name,
		&state.FkCountryId,
		&state.Slug,
		&state.IdState,
	)

	if err != nil {
		fmt.Println("error:countries:controller:getState:1", err)
		c.JSON(http.StatusNotFound, gin.H{"state": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"state": state})
}

func getTowns(c *gin.Context) {
	state := c.Params.ByName("state")
	row, err := settings.DB.Query("SELECT id_town,fk_state_id,cod_town,dc,name, slug FROM towns WHERE fk_state_id=$1 ORDER BY name asc", &state)

	townsRes := towns.Towns{}

	if err != nil {
		fmt.Println("error:countries:controller:getTowns:1", err)
		c.JSON(http.StatusNotFound, gin.H{"states": nil})
		return
	}

	for row.Next() {
		var town towns.Town

		err := row.Scan(
			&town.IdTown,
			&town.FkStateId,
			&town.CodTown,
			&town.Dc,
			&town.Name,
			&town.Slug,
		)

		if err != nil {
			fmt.Println("error:countries:controller:getTowns:2", err)
			continue
		}

		townsRes.Towns = append(townsRes.Towns, &town)
	}

	c.JSON(http.StatusOK, townsRes)
}

func getTown(c *gin.Context) {
	stateId := c.Params.ByName("state")
	townId := c.Params.ByName("town")

	townsRes := towns.Town{}

	err := settings.DB.QueryRow(
		"SELECT id_town,fk_state_id,cod_town,dc,name, slug FROM towns WHERE fk_state_id=$1 and id_town=$2 ORDER BY name asc",
		&stateId, &townId,
	).Scan(
		&townsRes.IdTown,
		&townsRes.FkStateId,
		&townsRes.CodTown,
		&townsRes.Dc,
		&townsRes.Name,
		&townsRes.Slug,
	)

	if err != nil {
		fmt.Println("error:countries:controller:getTown:1", err)
		c.JSON(http.StatusNotFound, gin.H{"states": nil})
		return
	}
	c.JSON(http.StatusOK, townsRes)
}
