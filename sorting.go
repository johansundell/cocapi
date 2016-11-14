package cocapi

import "math"

type DonationRatio []Member

func (c DonationRatio) Len() int {
	return len(c)
}

func (c DonationRatio) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c DonationRatio) Less(i, j int) bool {
	var d = make([]float64, 4)
	d[0], d[1], d[2], d[3] = float64(c[i].Donations), float64(c[i].DonationsReceived), float64(c[j].Donations), float64(c[j].DonationsReceived)
	for k, n := range d {
		if n == 0 {
			d[k] = math.SmallestNonzeroFloat64
		}
	}
	if (d[0] / d[1]) == (d[2] / d[3]) {
		return d[0] > d[2]
	}
	return (d[0] / d[1]) > (d[2] / d[3])
}

type Roles []Member

func (c Roles) Len() int {
	return len(c)
}

func (c Roles) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Roles) Less(i, j int) bool {
	if c[i].Role == c[j].Role {
		return c[i].ExpLevel > c[j].ExpLevel
	}
	roles := map[string]int{
		"leader":   4,
		"coLeader": 3,
		"admin":    2,
		"member":   1,
	}
	return roles[c[i].Role] > roles[c[j].Role]
}
