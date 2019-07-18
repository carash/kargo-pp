package main

import (
	"github.com/gin-gonic/gin"

	"github.com/carash/kargo-pp/src/core/handler"
	"github.com/carash/kargo-pp/src/core/kargo"
	"github.com/carash/kargo-pp/src/core/kargosql"
)

func main() {
	r := gin.Default()

	sqlconn := kargosql.Connect()
	kargocont := &kargo.Kargo{sqlconn}
	kargohndl := &handler.KargoHandler{kargocont}

	r.GET("/job", kargohndl.HandleGetJob)
	r.GET("/bid", kargohndl.HandleGetBid)

	r.Run("1234")
}
