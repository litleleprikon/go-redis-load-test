package runner

import (
	"context"
	"errors"

	"github.com/litleleprikon/go-redis-load-test/internal/reporter"
)

// Executor is an interface that should be implemented by each redis storage layer
type Executor interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Add(ctx context.Context, key string, values ...string) error
	Sdiff(ctx context.Context, key1, key2 string) ([]string, error)
	Pub(ctx context.Context, key string, value string) error
	Sub(ctx context.Context, key string, callback func(string, error) error)
	Smembers(ctx context.Context, key string) ([]string, error)
	Zadd(key string, score int, value string) error
	Zremrangebyscore(key string) (map[int]string, error)
}

// NewRunner is a constructor function for Runner
func NewRunner(executor Executor, rep reporter.Reporter) *Runner {
	return &Runner{
		executor: executor,
		reporter: rep,
	}
}

// Runner is a struct that emulate different behavior of application and tries to load redis with different operations
type Runner struct {
	executor Executor
	reporter reporter.Reporter
}

// Run starts an experiment. Loads database, call reporter
func (r *Runner) Run(ctx context.Context) error {
	return errors.New("not implemented")
}
