package worker

type WorkerBuilder struct {
	task Task
}

func NewBuilder() *WorkerBuilder {
	return new(WorkerBuilder)
}

func (b *WorkerBuilder) Task(t Task) *WorkerBuilder {
	b.task = t
	return b
}

func (b *WorkerBuilder) Build() *Worker {
	return &Worker{
		task: b.task,
	}
}
