package countries

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.Group("/countries").
		GET("", getAll).
		GET("/:country/states/:state/towns/:town", getTown).
		GET("/:country/states/:state/towns", getTowns).
		GET("/:country/states/:state", getState).
		GET("/:country/states", getStates).
		GET("/:country", GetCountry)
}
