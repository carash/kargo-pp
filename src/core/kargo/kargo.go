package kargo

import (
	"github.com/carash/kargo-pp/src/core/class/bid"
	"github.com/carash/kargo-pp/src/core/class/job"
	"github.com/carash/kargo-pp/src/core/kargosql"
)

type Kargo struct {
	SQLConn *kargosql.KargoSQL
}

type BidParam struct {
	UserID int
}

func (k *Kargo) GetBid(param BidParam) ([]bid.Bid, error) {
	return nil, nil
}

type JobParam struct {
	UserID int
}

func (k *Kargo) GetJob(param JobParam) ([]job.Job, error) {
	return nil, nil
}
