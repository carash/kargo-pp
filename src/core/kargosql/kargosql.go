package kargosql

import (
	"time"

	"github.com/carash/kargo-pp/src/core/class/bid"
	"github.com/carash/kargo-pp/src/core/class/job"
	"github.com/carash/kargo-pp/src/core/class/user"
)

type KargoSQL struct {
	Bids []bid.Bid
	Jobs []job.Job
}

func Connect() *KargoSQL {
	users := []user.User{
		user.User{
			ID:   1,
			Name: "Some Name",
		},
		user.User{
			ID:   2,
			Name: "Other Name",
		},
		user.User{
			ID:   3,
			Name: "My Name",
		},
		user.User{
			ID:   4,
			Name: "Your Name",
		},
		user.User{
			ID:   5,
			Name: "His Name",
		},
	}

	locs := []job.Location{
		job.Location{Name: "Jakarta"},
		job.Location{Name: "Surabaya"},
		job.Location{Name: "Bogor"},
	}

	now := time.Now()

	jobs := []job.Job{
		job.Job{
			User:     &users[0],
			Origin:   &locs[0],
			Dest:     &locs[1],
			Budget:   100,
			Shipment: now.Add(1000),
		},
		job.Job{
			User:     &users[0],
			Origin:   &locs[1],
			Dest:     &locs[0],
			Budget:   200,
			Shipment: now.Add(2000),
		},
		job.Job{
			User:     &users[1],
			Origin:   &locs[1],
			Dest:     &locs[2],
			Budget:   125,
			Shipment: now.Add(2000),
		},
		job.Job{
			User:     &users[1],
			Origin:   &locs[2],
			Dest:     &locs[0],
			Budget:   273,
			Shipment: now.Add(500),
		},
	}

	vhcs := []bid.Vehicle{
		bid.Vehicle{Name: "Truck"},
		bid.Vehicle{Name: "Tronton"},
		bid.Vehicle{Name: "CDE"},
		bid.Vehicle{Name: "CDD"},
	}

	bids := []bid.Bid{
		bid.Bid{
			User:    &users[2],
			Job:     &jobs[0],
			Price:   100,
			Vehicle: &vhcs[0],
		},
		bid.Bid{
			User:    &users[3],
			Job:     &jobs[0],
			Price:   120,
			Vehicle: &vhcs[0],
		},
		bid.Bid{
			User:    &users[4],
			Job:     &jobs[0],
			Price:   120,
			Vehicle: &vhcs[2],
		},
		bid.Bid{
			User:    &users[2],
			Job:     &jobs[1],
			Price:   210,
			Vehicle: &vhcs[2],
		},
		bid.Bid{
			User:    &users[3],
			Job:     &jobs[1],
			Price:   200,
			Vehicle: &vhcs[1],
		},
		bid.Bid{
			User:    &users[4],
			Job:     &jobs[2],
			Price:   120,
			Vehicle: &vhcs[1],
		},
		bid.Bid{
			User:    &users[3],
			Job:     &jobs[2],
			Price:   130,
			Vehicle: &vhcs[3],
		},
		bid.Bid{
			User:    &users[4],
			Job:     &jobs[3],
			Price:   300,
			Vehicle: &vhcs[2],
		},
	}

	return &KargoSQL{
		Jobs: jobs,
		Bids: bids,
	}
}

func (k *KargoSQL) GetBid(userId int) ([]bid.Bid, error) {
	return k.Bids, nil
}

func (k *KargoSQL) GetJob(userId int) ([]job.Job, error) {
	return k.Jobs, nil
}
