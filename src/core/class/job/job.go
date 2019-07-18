package job

import (
	"time"

	"github.com/carash/kargo-pp/src/core/class/user"
)

type Job struct {
	ID          int
	User        *user.User
	Origin      *Location
	Dest        *Location
	Budget      int
	Shipment    time.Time
	Description string
}

func (j Job) LessOrigin(o Job) bool {
	return j.Origin.Less(*o.Origin)
}

func (j Job) LessDest(o Job) bool {
	return j.Dest.Less(*o.Dest)
}

func (j Job) LessBudget(o Job) bool {
	return j.Budget < o.Budget
}

func (j Job) LessShipment(o Job) bool {
	return j.Shipment.Before(o.Shipment)
}

type Location struct {
	ID   int
	Name string
	Code string
}

func (l Location) Less(o Location) bool {
	return l.Name < o.Name
}
