package worker

import (
	"errors"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/miun173/autograd/config"
)

var defaultJobOpt = work.JobOptions{MaxConcurrency: 3, MaxFails: 3}

const cronEvery10Minute = "*/10 * * * *"

// Worker ..
type Worker struct {
	pool       *work.WorkerPool
	redisPool  *redis.Pool
	enqueuer   *work.Enqueuer
	grader     GraderUsecase
	submission Submission
}

// NewWorker ..
func NewWorker(opts ...Option) *Worker {
	wrk := &Worker{}
	for _, opt := range opts {
		opt(wrk)
	}

	return wrk
}

// Start starts worker
func (w *Worker) Start() {
	w.registerJobs()
	w.pool.Start()
}

// Stop stops worker
func (w *Worker) Stop() {
	w.pool.Stop()
}

func (w *Worker) registerJobs() {
	conc := config.WorkerConcurrency()
	nameSpace := config.WorkerNamespace()

	w.pool = work.NewWorkerPool(jobHandler{}, conc, nameSpace, w.redisPool)
	w.pool.Middleware(w.registerJobConfig)

	w.pool.JobWithOptions(jobGradeAssignment, defaultJobOpt, (*jobHandler).handleGradeAssignment)
	w.pool.JobWithOptions(jobCheckAllDueAssignments, defaultJobOpt, (*jobHandler).handleCheckAllDueAssignments)
	w.pool.JobWithOptions(jobGradeSubmission, defaultJobOpt, (*jobHandler).handleGradeSubmission)

	w.pool.PeriodicallyEnqueue(cronEvery10Minute, jobCheckAllDueAssignments)
}

func (w *Worker) registerJobConfig(handler *jobHandler, job *work.Job, next work.NextMiddlewareFunc) error {
	if handler == nil {
		return errors.New("unexpected nil handler")
	}

	handler.pool = w.pool
	handler.redisPool = w.redisPool
	handler.enqueuer = w.enqueuer
	handler.grader = w.grader
	handler.submission = w.submission

	return next()
}