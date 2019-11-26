package utils

type QueuedFunctions struct {
	queue []func() (string, error)
}

func (q *QueuedFunctions) Apply() (string, error) {
	for _, action := range q.queue {
		output, err := action()
		if err != nil {
			return output, err
		}
	}
	return "", nil
}

func (q *QueuedFunctions) Append(action func() (string, error)) {
	q.queue = append(q.queue, action)
}
