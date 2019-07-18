package kargosql

import (
	"github.com/carash/kargo-pp/src/core/class/bid"
	"github.com/carash/kargo-pp/src/core/class/job"
)

type KargoSQL struct {
}

func Connect() *KargoSQL {
	return &KargoSQL{}
}

func (k *KargoSQL) GetBid(userId int) ([]bid.Bid, error) {
	return nil, nil
}

func (k *KargoSQL) GetJob(userId int) ([]job.Job, error) {
	return nil, nil
}
