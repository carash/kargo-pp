package bid

type Bid struct {
}

func (b Bid) Less(o Bid) bool {
	return true
}

type Vehicle struct {
}

func (v Vehicle) Less(o Vehicle) bool {
	return true
}
