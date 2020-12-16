/*
   This job processor is designed to take a list of jobs, allowcate to container/workers while 
   maintaining a maximum number of workers (containers) and exit when all jobs have completed.

 1. create the job request slice
 2. pass the slice to job processor
    *. loop through containers to get a free list (idle)
    *. if there are idle continers,
            loop through queued jobs and to the maximum number of conncurrent; 
            mark job as running
    *. sleep n
    *. insert completed jobs into completed
 4. evaluate failed to determine if they should be re-submitted

 ref: https://www.opsdash.com/blog/job-queues-in-go.html

*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Job struct {
    id  int
    name string
    status string // queued, working, completed
    container string
    attempts int
    startTime time.Time
    endTime  time.Time
    failed bool
    response map[string]interface{}
}

type JobQueue struct {
    requested []*Job
    completed []*Job
    failed    []*Job
}

const (
    Queued = "queued"
    Working = "working"
    Completed = "completed"
)

var (
    rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
    jobid = 100
    maxConcurrent = 4
)

func scanContainers(containers map[string]*Job) (idle []string, working []string) {
    idle = make([]string, 0)
    working = make([]string, 0)

    for container, job := range containers {
        if job == nil || job.IsCompleted() {
            idle = append(idle, container)
        } else if job.IsWorking() {
            working = append(working, container)
        }
    }

    return idle, working
}

func (q JobQueue) getQueuedJobs() []*Job {
    jobs := make([]*Job, 0)

    for _, job := range q.requested {
        if job.IsQueued() {
            jobs = append(jobs, job)
        }
    }

    return jobs
}

func (q JobQueue) process(started, done chan bool, containers map[string]*Job) {
    defer func() {
        done<- true
    }()

    started <- true

    for {
        idle, working := scanContainers(containers)
        queue := q.getQueuedJobs()

        if len(containers) == len(idle) && len(queue) == 0 {
            break
        }

        loop := 0

        workCount := len(working)
        for {
            idleCount := len(idle)
            queueCount := len(queue)
            loop++

            fmt.Printf("%d: idle: %d, working: %d, queued: %d\n", loop, idleCount, workCount, queueCount)

            if idleCount < 1 || workCount >= maxConcurrent || queueCount < 1 {
                break
            }

            container := idle[0]
            job := queue[0]

            containers[container] = job
            job.container = container

            idle = idle[1:]
            queue = queue[1:]

            // simulate the job...
            fn := func(done chan bool, response map[string]interface{}) {
                n := rnd.Intn(5000) + 5000
                time.Sleep(time.Millisecond * time.Duration(n))

                response["status"] = "ok"
                response["errors"] = make([]string, 0)

                done <- true
            }

            // run the job
            jobStarted := make(chan bool)
            go job.run(jobStarted, fn)

            <- jobStarted

            workCount++
        }

        time.Sleep(time.Second * 1)
    }

    done<- true
}

func (j *Job) run(started chan bool, fn func(chan bool, map[string]interface{})) {
    defer j.SetCompleted()

    j.SetWorking()
    j.startTime = time.Now()
    j.attempts++

    started <- true

    done := make(chan bool)
    go fn(done, j.response)

    <- done
    j.endTime = time.Now()

    fmt.Println("job done, response: ", j.response)
}

// loop forever creating and submitting random jobs at random times...
func createJob(id int) *Job {
    job := Job{
        id:id,
        name:fmt.Sprintf("My Job# %d", id),
        status:Queued,
        attempts:0,
        failed: true,
        response:make(map[string]interface{}),
    }

    return &job
}

func (j Job) IsQueued() bool {
    return j.status == Queued
}

func (j *Job) SetWorking() string {
    j.status = Working
    fmt.Printf("working -> %v\n", j)
    return j.status
}

func (j *Job) IsWorking() bool {
    return j.status == Working
}

func (j *Job) SetCompleted() string {
    j.status = Completed
    fmt.Printf("complete -> %v\n", j)
    return j.status
}

func (j *Job) IsCompleted() bool {
    return j.status == Completed
}

func generateJobs() []*Job {
    jobCount := 8
    jobs := make([]*Job, jobCount)
    for i := 0; i < jobCount; i++ {
        jobid++
        jobs[i] = createJob(jobid)
    }

    return jobs
}

func NewJobQueue(jobs []*Job) *JobQueue {
    q := JobQueue{}

    q.requested = jobs
    q.completed = make([]*Job, 0)
    q.failed = make([]*Job, 0)

    return &q
}

// create a new map of containers/jobs where all jobs are nil
func getContainers() map[string]*Job {
    c := make(map[string]*Job)

    for i := 1; i <= 10; i++ {
        name := fmt.Sprintf("service-runner-%d", i)
        c[name] = nil
    }

    return c
}

func main() {
    jobs := generateJobs()
    // fmt.Println(jobs)

    q := NewJobQueue(jobs)

    started := make(chan bool)
    done := make(chan bool)

    containers := getContainers()
    // fmt.Println(containers)

    go q.process(started, done, containers)

    <-started
    fmt.Println("job processor has started...")

    <-done
    fmt.Println("job processor has completed...")
}

