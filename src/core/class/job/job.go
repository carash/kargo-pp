package job

type Job struct {
}

func (b Job) Less(o Job) bool {
	return true
}

type Location struct {
}

func (v Location) Less(o Location) bool {
	return true
}
