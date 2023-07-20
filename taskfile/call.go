package taskfile

// Call is the parameters to a task call
type Call struct {
	Task   string
	Vars   *Vars
	Envs   *Vars
	Silent bool
	Direct bool // Was the task called directly or via another task?
}
