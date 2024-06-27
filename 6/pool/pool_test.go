package pool_test

import (
	"alexandernizov/lessons/6/pool"
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestWithoutError(t *testing.T) {
	//Тест на утечку горутин
	defer goleak.VerifyNone(t)

	jobsCount := 20

	jobs := []pool.Task{}
	for i := 0; i < jobsCount; i++ {
		jobs = append(jobs, jobWithoutError)
	}

	err := pool.Run(jobs, 1, 0)
	assert.Equal(t, nil, err, "expected 0 error")
	err = pool.Run(jobs, 8, 0)
	assert.Equal(t, nil, err, "expected 0 error")
	err = pool.Run(jobs, 1, 1)
	assert.Equal(t, nil, err, "expected 0 error")
	err = pool.Run(jobs, 8, 1)
	assert.Equal(t, nil, err, "expected 0 error")
}

func TestWithAllError(t *testing.T) {
	//Тест на утечку горутин
	defer goleak.VerifyNone(t)

	jobsCount := 20

	jobs := []pool.Task{}
	for i := 0; i < jobsCount; i++ {
		jobs = append(jobs, jobWithError)
	}

	err := pool.Run(jobs, 1, 0)
	assert.Equal(t, nil, err, "expected 0 error")
	err = pool.Run(jobs, 8, 0)
	assert.Equal(t, nil, err, "expected 0 error")
	err = pool.Run(jobs, 1, 1)
	assert.Equal(t, pool.ErrErrorsLimitExceeded, err, "expected error")
	err = pool.Run(jobs, 8, 1)
	assert.Equal(t, pool.ErrErrorsLimitExceeded, err, "expected error")
	err = pool.Run(jobs, 1, 2)
	assert.Equal(t, pool.ErrErrorsLimitExceeded, err, "expected error")
	err = pool.Run(jobs, 8, 2)
	assert.Equal(t, pool.ErrErrorsLimitExceeded, err, "expected error")
}

func TestWithPartuallyError(t *testing.T) {
	//Тест на утечку горутин
	defer goleak.VerifyNone(t)

	jobsCount := 20

	jobs := []pool.Task{}
	for i := 0; i < (jobsCount / 2); i++ {
		jobs = append(jobs, jobWithError)
	}
	for i := 0; i < (jobsCount / 2); i++ {
		jobs = append(jobs, jobWithoutError)
	}

	err := pool.Run(jobs, 1, 0)
	assert.Equal(t, nil, err, "expected 0 error")
	err = pool.Run(jobs, 8, 0)
	assert.Equal(t, nil, err, "expected 0 error")
	err = pool.Run(jobs, 1, 1)
	assert.Equal(t, pool.ErrErrorsLimitExceeded, err, "expected error")
	err = pool.Run(jobs, 8, 1)
	assert.Equal(t, pool.ErrErrorsLimitExceeded, err, "expected error")
	err = pool.Run(jobs, 1, 2)
	assert.Equal(t, pool.ErrErrorsLimitExceeded, err, "expected error")
	err = pool.Run(jobs, 8, 2)
	assert.Equal(t, pool.ErrErrorsLimitExceeded, err, "expected error")
}

func jobWithError() error {
	i := 1_000_000_000
	rnd := int64(rand.Intn(i))

	//some difficult work
	for i := 0; i < int(rnd); i++ {
		_ = i
	}

	return errors.New("error")
}

func jobWithoutError() error {
	i := 1_000_000_000
	rnd := int64(rand.Intn(i))

	//some difficult work
	for i := 0; i < int(rnd); i++ {
		_ = i
	}

	return nil
}
