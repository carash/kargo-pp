package handler

import (
	"github.com/carash/kargo-pp/src/core/kargo"
	"github.com/gin-gonic/gin"
)

type KargoHandler struct {
	KargoController *kargo.Kargo
}

func (k *KargoHandler) HandleGetJob(c *gin.Context) {}

func (k *KargoHandler) HandleGetBid(c *gin.Context) {}
