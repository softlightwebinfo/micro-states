package towns

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.Group("/towns")
}
