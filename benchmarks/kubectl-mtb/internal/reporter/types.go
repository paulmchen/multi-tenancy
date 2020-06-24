package reporter

import (
	"fmt"
	"time"

	"github.com/creasty/defaults"
	"sigs.k8s.io/multi-tenancy/benchmarks/kubectl-mtb/pkg/benchmark"
	benchmarksuite "sigs.k8s.io/multi-tenancy/benchmarks/kubectl-mtb/pkg/benchmark_suite"
)

type SuiteSummary struct {
	NumberOfTotalTests   int
	NumberOfPassedTests  int
	NumberOfFailedTests  int
	NumberOfSkippedTests int
	RunTime              time.Duration
	Suite                *benchmarksuite.BenchmarkSuite
}

type TestSummary struct {
	Validation      bool `default:"true"`
	ValidationError error
	Test            bool `default:"true"`
	RunTime         time.Duration
	Benchmark       *benchmark.Benchmark
}

// SetDefaults usage := https://github.com/creasty/defaults#usage
func (t *TestSummary) SetDefaults() error {
	if err := defaults.Set(t); err != nil {
		return fmt.Errorf("it should not return an error: %v", err)
	}
	return nil
}
