package kargo

import (
	"github.com/carash/kargo-pp/src/core/class/bid"
	"github.com/carash/kargo-pp/src/core/class/job"
)

type Kargo struct {
}

type ShipperBidParam struct {
}

func (k *Kargo) ShipperGetBid(param ShipperBidParam) ([]bid.Bid, error) {
	return nil, nil
}

type ShipperJobParam struct {
}

func (k *Kargo) ShipperGetJob(param ShipperJobParam) ([]job.Job, error) {
	return nil, nil
}

type TransporterJobParam struct {
}

func (k *Kargo) TransporterGetJob(param TransporterJobParam) ([]job.Job, error) {
	return nil, nil
}
