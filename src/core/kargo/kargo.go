package kargo

import (
	"github.com/carash/kargo-pp/src/core/class/bid"
	"github.com/carash/kargo-pp/src/core/class/job"
	"github.com/carash/kargo-pp/src/core/kargosql"
	"github.com/carash/kargo-pp/src/util/kargosort"
)

type Kargo struct {
	SQLConn *kargosql.KargoSQL
}

type BidParam struct {
	UserID int
	SortBy string
}

func (k *Kargo) GetBid(param BidParam) ([]bid.Bid, error) {
	bids, _ := k.SQLConn.GetBid(param.UserID)

	switch param.SortBy {
	case "vehicle":
		kargosort.Sort(bids, func(i, j int) bool { return bids[i].LessVehicle(bids[j]) })
	case "price":
		kargosort.Sort(bids, func(i, j int) bool { return bids[i].LessPrice(bids[j]) })
	}

	return bids, nil
}

type JobParam struct {
	UserID int
	SortBy string
}

func (k *Kargo) GetJob(param JobParam) ([]job.Job, error) {
	jobs, _ := k.SQLConn.GetJob(param.UserID)

	switch param.SortBy {
	case "origin":
		kargosort.Sort(jobs, func(i, j int) bool { return jobs[i].LessOrigin(jobs[j]) })
	case "destination":
		kargosort.Sort(jobs, func(i, j int) bool { return jobs[i].LessDest(jobs[j]) })
	case "budget":
		kargosort.Sort(jobs, func(i, j int) bool { return jobs[i].LessBudget(jobs[j]) })
	case "shipment":
		kargosort.Sort(jobs, func(i, j int) bool { return jobs[i].LessShipment(jobs[j]) })
	}

	return jobs, nil
}
