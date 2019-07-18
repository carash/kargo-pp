package job

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJobLessOrigin(t *testing.T) {
	tc := [][2]Job{
		[2]Job{
			Job{
				Origin: &Location{Name: "Jakarta"},
			},
			Job{
				Origin: &Location{Name: "Surabaya"},
			},
		},
		[2]Job{
			Job{
				Origin: &Location{Name: "Malang"},
			},
			Job{
				Origin: &Location{Name: "Palembang"},
			},
		},
		[2]Job{
			Job{
				Origin: &Location{Name: "Medan"},
			},
			Job{
				Origin: &Location{Name: "Lampung"},
			},
		},
	}

	ans := []bool{
		true,
		true,
		false,
	}

	n := len(tc)
	for i := 0; i < n; i++ {
		j1 := tc[i][0]
		j2 := tc[i][1]
		check := j1.LessOrigin(j2)

		assert.Equal(t, ans[i], check)
	}
}

func TestJobLessDest(t *testing.T) {
	tc := [][2]Job{
		[2]Job{
			Job{
				Dest: &Location{Name: "Jakarta"},
			},
			Job{
				Dest: &Location{Name: "Surabaya"},
			},
		},
		[2]Job{
			Job{
				Dest: &Location{Name: "Medan"},
			},
			Job{
				Dest: &Location{Name: "Lampung"},
			},
		},
		[2]Job{
			Job{
				Dest: &Location{Name: "Malang"},
			},
			Job{
				Dest: &Location{Name: "Palembang"},
			},
		},
	}

	ans := []bool{
		true,
		false,
		true,
	}

	n := len(tc)
	for i := 0; i < n; i++ {
		j1 := tc[i][0]
		j2 := tc[i][1]
		check := j1.LessDest(j2)

		assert.Equal(t, ans[i], check)
	}
}

func TestJobLessBudget(t *testing.T) {
	tc := [][2]Job{
		[2]Job{
			Job{
				Budget: 100,
			},
			Job{
				Budget: 242,
			},
		},
		[2]Job{
			Job{
				Budget: 231,
			},
			Job{
				Budget: 124,
			},
		},
		[2]Job{
			Job{
				Budget: 293,
			},
			Job{
				Budget: 222,
			},
		},
	}

	ans := []bool{
		true,
		true,
		false,
	}

	n := len(tc)
	for i := 0; i < n; i++ {
		j1 := tc[i][0]
		j2 := tc[i][1]
		check := j1.LessBudget(j2)

		assert.Equal(t, ans[i], check)
	}
}

func TestJobLessShipment(t *testing.T) {
	now := time.Now()

	tc := [][2]Job{
		[2]Job{
			Job{
				Shipment: now.Add(100),
			},
			Job{
				Shipment: now.Add(124),
			},
		},
		[2]Job{
			Job{
				Shipment: now.Add(354),
			},
			Job{
				Shipment: now.Add(221),
			},
		},
		[2]Job{
			Job{
				Shipment: now.Add(192),
			},
			Job{
				Shipment: now.Add(492),
			},
		},
	}

	ans := []bool{
		true,
		false,
		true,
	}

	n := len(tc)
	for i := 0; i < n; i++ {
		j1 := tc[i][0]
		j2 := tc[i][1]
		check := j1.LessShipment(j2)

		assert.Equal(t, ans[i], check)
	}
}

func TestJobLessLocation(t *testing.T) {
	tc := [][2]Location{
		[2]Location{
			Location{Name: "Jakarta"},
			Location{Name: "Palembang"},
		},
		[2]Location{
			Location{Name: "Surabaya"},
			Location{Name: "Medan"},
		},
		[2]Location{
			Location{Name: "Malang"},
			Location{Name: "Kalimantan"},
		},
	}

	ans := []bool{
		true,
		false,
		false,
	}

	n := len(tc)
	for i := 0; i < n; i++ {
		l1 := tc[i][0]
		l2 := tc[i][1]
		check := l1.Less(l2)

		assert.Equal(t, ans[i], check)
	}
}
