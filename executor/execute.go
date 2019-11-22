package executor

import (
	"fmt"
	"github.com/ui-kreinhard/go-setuper/log"
	"os"
	"time"
)

type ToRun struct {
	Name          string
	FunctionToRun func() (string, error)
}

type Executor struct {
	functionsToRun []ToRun
	name           string
}

type IExecutor interface {
	Plan(nameOfTask string, deferred func() (string, error)) IExecutor
	PlannedSleep(name string, duration time.Duration) IExecutor
	PlanIgnoreError(nameOfTask string, deferred func() (string, error)) IExecutor
	Run()
}

func IgnoreError(funcWithErrorToIgnore ToRun) ToRun {
	return ToRun{
		funcWithErrorToIgnore.Name,
		func() (string, error) {
			output, err := funcWithErrorToIgnore.FunctionToRun()
			log.Println("Ignoring error on", funcWithErrorToIgnore.Name)
			log.LogOutputError(output, err)
			return output, nil
		},
	}
}

func (e *Executor) PlannedSleep(name string, duration time.Duration) IExecutor {
	e.functionsToRun = append(e.functionsToRun,
		ToRun{
			name,
			func() (string, error) {
				time.Sleep(duration)
				return "", nil
			},
		},
	)
	return e
}

func (e *Executor) PlanIgnoreError(nameOfTask string, deferred func() (string, error)) IExecutor {
	e.functionsToRun = append(e.functionsToRun, IgnoreError(ToRun{nameOfTask, deferred}))
	return e
}

func (e *Executor) Plan(nameOfTask string, deferred func() (string, error)) IExecutor {
	e.functionsToRun = append(e.functionsToRun, ToRun{nameOfTask, deferred})
	return e
}

func (e *Executor) Run() {
	log.Println("=============== Starting", e.name, "task ==============")
	start := time.Now()

	for _, toRunElement := range e.functionsToRun {
		log.Println("Running", toRunElement.Name)
		output, err := toRunElement.FunctionToRun()
		log.LogOutputError(output, err)
		log.Println("Ran", toRunElement.Name)
		log.Println()
		if err != nil {
			log.Println("Stopping execution on", toRunElement.Name, "error recognized")
			os.Exit(1)
		}
	}
	duration := time.Since(start)
	log.Println("=============== Finished", e.name, "took", duration, "===============")
	fmt.Println()
}

func NewExecutor(name string) IExecutor {
	return &Executor{name: name}
}
