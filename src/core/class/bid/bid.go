package bid

import (
	"github.com/carash/kargo-pp/src/core/class/job"
	"github.com/carash/kargo-pp/src/core/class/user"
)

type Bid struct {
	ID      int
	User    *user.User
	Job     *job.Job
	Price   int
	Vehicle *Vehicle
}

func (b Bid) Less(o Bid) bool {
	return true
}

type Vehicle struct {
	ID   int
	Name string
	Code string
}

func (v Vehicle) Less(o Vehicle) bool {
	return true
}
