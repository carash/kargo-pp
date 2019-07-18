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

func (b Job) LessOrigin(o Job) bool {
	return true
}

func (b Job) LessDest(o Job) bool {
	return true
}

func (b Job) LessBudget(o Job) bool {
	return true
}

func (b Job) LessShipment(o Job) bool {
	return true
}

type Location struct {
	ID   int
	Name string
	Code string
}

func (l Location) Less(o Location) bool {
	return l.Name < o.Name
}
