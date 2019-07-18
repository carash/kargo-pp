package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/carash/kargo-pp/src/core/class/bid"
	"github.com/carash/kargo-pp/src/core/class/job"
	"github.com/carash/kargo-pp/src/core/kargo"
	"github.com/gin-gonic/gin"
)

type KargoHandler struct {
	KargoController *kargo.Kargo
}

type JobResponse struct {
	ErrorCode int       `json:"err_code,omitempty"`
	ErrorMsg  string    `json:"err_msg,omitempty"`
	Jobs      []job.Job `json:"jobs,omitempty"`
}

func (k *KargoHandler) HandleGetJob(c *gin.Context) {
	var statusCode int = http.StatusOK
	var errcode int
	var errmsg string
	var jobs []job.Job
	defer func() {
		data := JobResponse{
			ErrorCode: errcode,
			ErrorMsg:  errmsg,
			Jobs:      jobs,
		}

		s, _ := json.Marshal(data)
		c.Data(statusCode, "application/json", s)
	}()

	uidString, ok := c.GetQuery("user_id")
	if !ok {
		statusCode = http.StatusBadRequest
		errcode = 1
		errmsg = "need user id"
		return
	}

	uid, err := strconv.Atoi(uidString)
	if err != nil {
		statusCode = http.StatusBadRequest
		errcode = 2
		errmsg = "user id must be int"
		return
	}

	param := kargo.JobParam{UserID: uid}

	sortBy := c.Query("sort_by")
	switch sortBy {
	case "vehicle", "price":
		param.SortBy = sortBy
	default:
		param.SortBy = "price"
	}

	j, err := k.KargoController.GetJob(param)
	if err != nil {
		statusCode = http.StatusInternalServerError
		errcode = 3
		errmsg = "error retrieving bids"
		return
	}

	jobs = j
	return
}

type BidResponse struct {
	ErrorCode int       `json:"err_code,omitempty"`
	ErrorMsg  string    `json:"err_msg,omitempty"`
	Bids      []bid.Bid `json:"bids,omitempty"`
}

func (k *KargoHandler) HandleGetBid(c *gin.Context) {
	var statusCode int = http.StatusOK
	var errcode int
	var errmsg string
	var bids []bid.Bid
	defer func() {
		data := BidResponse{
			ErrorCode: errcode,
			ErrorMsg:  errmsg,
			Bids:      bids,
		}

		s, _ := json.Marshal(data)
		c.Data(statusCode, "application/json", s)
	}()

	uidString, ok := c.GetQuery("user_id")
	if !ok {
		statusCode = http.StatusBadRequest
		errcode = 1
		errmsg = "need user id"
		return
	}

	uid, err := strconv.Atoi(uidString)
	if err != nil {
		statusCode = http.StatusBadRequest
		errcode = 2
		errmsg = "user id must be int"
		return
	}

	param := kargo.BidParam{UserID: uid}

	sortBy := c.Query("sort_by")
	switch sortBy {
	case "origin", "destination", "budget", "shipment":
		param.SortBy = sortBy
	default:
		param.SortBy = "shipment"
	}

	b, err := k.KargoController.GetBid(param)
	if err != nil {
		statusCode = http.StatusInternalServerError
		errcode = 3
		errmsg = "error retrieving bids"
		return
	}

	bids = b
	return
}
