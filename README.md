# TASKINATOR 🤖

Welcome to the TASKINATOR! This package provides a robust utility in Golang that helps to set up and manage tasks that executes based on the duration specified by the user.

## Key Features 📌

1. Schedule multiple tasks with given name, schedule time, maximum retry count, retry delay, and action function.

1. Robust automatic retry logic in place in case of task failures.

1. The ability to execute multiple tasks concurrently.

______________________________________________________________________

## How to Use 🚀

__1. Get the package__

```bash
go get -u github.com/kylix31/taskinator 
```

__2. Import the package__

In your Go file, import the `tasksgenerator` package.

```go
import "github.com/kylix31/taskinator"
```

__3. Usage__

Create a Scheduler object using the `NewScheduler` function.

```go
scheduler := tasksgenerator.NewScheduler()
```

Add tasks to the Scheduler using the `AddTask` method. The `Task` struct must include the following fields:

- `Name` (string): Name of the Task
- `Schedule` (time.Duration): Interval at which the task will be triggered - if nil will run forever
- `MaxRetries` (int): Maximum number of retries that will be done before task is considered failed
- `RetryDelay` (time.Duration): Delay before each retry
- `Action` (func() error): Function to be performed by the Task

```go
task := tasksgenerator.Task{
	Name: "Sample Task",
	Schedule: 5 * time.Second,
	MaxRetries: 3,
	RetryDelay: 2 * time.Second,
	Action: func() error {
		// your task code
	},
}

scheduler.AddTask(task)
```

Finally, start the Scheduler using the `Start` method where `Duration` is the total time throughout which tasks should run.

```go
scheduler.Start(60 * time.Second)  // Run tasks for 60 seconds
```

And that's it! Now, TASKINATOR is at your service, doing all those tasks you've scheduled on the desired interval.

______________________________________________________________________

## Contribution ✨

Feel free to report issues or create a pull-request. All contributions are welcome.

______________________________________________________________________

ACKNOWLEDGEMENT 💖

Thanks goes to those wonderful people who inspired or contributed in any way towards the development of TASKINATOR.

______________________________________________________________________

Let TASKINATOR bring success to your task scheduling needs! 🎉

Made with 💙 and Go.
