package kargo

import (
	"github.com/carash/kargo-pp/src/core/class/bid"
	"github.com/carash/kargo-pp/src/core/class/job"
	"github.com/carash/kargo-pp/src/core/kargosql"
)

type Kargo struct {
	SQLConn *kargosql.KargoSQL
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
