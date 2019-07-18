package bid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBidLessPrice(t *testing.T) {
	tc := [][2]Bid{
		[2]Bid{
			Bid{
				Price: 100,
			},
			Bid{
				Price: 123,
			},
		},
		[2]Bid{
			Bid{
				Price: 928,
			},
			Bid{
				Price: 243,
			},
		},
		[2]Bid{
			Bid{
				Price: 123,
			},
			Bid{
				Price: 424,
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
		b1 := tc[i][0]
		b2 := tc[i][1]
		check := b1.LessPrice(b2)

		assert.Equal(t, ans[i], check)
	}
}

func TestBidLessVehicle(t *testing.T) {
	tc := [][2]Bid{
		[2]Bid{
			Bid{
				Vehicle: &Vehicle{
					Name: "Truck",
				},
			},
			Bid{
				Vehicle: &Vehicle{
					Name: "Fuso",
				},
			},
		},
		[2]Bid{
			Bid{
				Vehicle: &Vehicle{
					Name: "CDE",
				},
			},
			Bid{
				Vehicle: &Vehicle{
					Name: "Tronton",
				},
			},
		},
		[2]Bid{
			Bid{
				Vehicle: &Vehicle{
					Name: "Pickup",
				},
			},
			Bid{
				Vehicle: &Vehicle{
					Name: "CDD",
				},
			},
		},
	}

	ans := []bool{
		false,
		true,
		false,
	}

	n := len(tc)
	for i := 0; i < n; i++ {
		b1 := tc[i][0]
		b2 := tc[i][1]
		check := b1.LessVehicle(b2)

		assert.Equal(t, ans[i], check)
	}
}

func TestVehicleLess(t *testing.T) {
	tc := [][2]Vehicle{
		[2]Vehicle{
			Vehicle{
				Name: "Truck",
			},
			Vehicle{
				Name: "Fuso",
			},
		},
		[2]Vehicle{
			Vehicle{
				Name: "Pickup",
			},
			Vehicle{
				Name: "CDD",
			},
		},
		[2]Vehicle{
			Vehicle{
				Name: "CDE",
			},
			Vehicle{
				Name: "Tronton",
			},
		},
	}

	ans := []bool{
		false,
		false,
		true,
	}

	n := len(tc)
	for i := 0; i < n; i++ {
		v1 := tc[i][0]
		v2 := tc[i][1]
		check := v1.Less(v2)

		assert.Equal(t, ans[i], check)
	}
}
