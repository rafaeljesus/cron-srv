package scheduler

import (
	"github.com/EmpregoLigado/cron-srv/models"
	"github.com/EmpregoLigado/cron-srv/repo"
)

type SchedulerMock struct {
	Created bool
	Updated bool
	Deleted bool
}

func NewMock() *SchedulerMock {
	return &SchedulerMock{
		Created: false,
		Updated: false,
		Deleted: false,
	}
}

func (s *SchedulerMock) Create(cron *models.Cron) (err error) {
	s.Created = true
	return
}

func (s *SchedulerMock) Update(cron *models.Cron) (err error) {
	s.Updated = true
	return
}

func (s *SchedulerMock) Delete(id uint) (err error) {
	s.Deleted = true
	return
}

func (sm *SchedulerMock) ScheduleAll(repo repo.Repo) (err error) {
	return
}