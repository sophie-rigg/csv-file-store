package worker

import (
	"github.com/google/uuid"
)

type Job struct {
	data     []byte
	id       string
	attempts int
}

func NewJob() *Job {
	return &Job{
		id: uuid.New().String(),
	}
}

func (j *Job) AddData(data []byte) {
	j.data = data
}

func (j *Job) GetID() string {
	return j.id
}

func (j *Job) GetData() []byte {
	return j.data
}

func (j *Job) IncrementAttempts() {
	j.attempts++
}

func (j *Job) GetAttempts() int {
	return j.attempts
}
