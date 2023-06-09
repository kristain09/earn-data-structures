package structures

type Stack struct {
	items []int
}

type Queue struct {
	Queueitems []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (q *Queue) Pop() int {
	l := len(q.Queueitems) - 1
	toREmove := q.Queueitems[l]
	q.Queueitems = q.Queueitems[:l]
	return toREmove
}

func (q *Queue) PushQueue(i int) {
	q.Queueitems = append(q.Queueitems, i)
}

func (q *Queue) PopQueue() int {
	toREmove := q.Queueitems[0]

	q.Queueitems = q.Queueitems[1:]
	return toREmove
}
