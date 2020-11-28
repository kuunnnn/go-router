package tool

type Queue interface {
	Peak() interface{}
	Push(ele interface{}) bool
	Shift() interface{}
}

type queue struct {
	list []interface{}
	Size int
	Cap  int
}

func NewQueue(cap int) *queue {
	if cap < 10 {
		cap = 10
	}
	return &queue{
		list: make([]interface{}, 0, cap),
		Cap:  cap,
		Size: 0,
	}
}

func (q *queue) Peak() interface{} {
	if len(q.list) == 0 {
		return nil
	}
	return q.list[0]
}
func (q *queue) Push(e interface{}) bool {
	if q.Size == q.Cap {
		return false
	}
	q.Size += 1
	q.list = append(q.list, e)
	return true
}

func (q *queue) Shift() interface{} {
	if q.Size == 0 {
		return nil
	}
	result := q.list[0]
	q.Size -= 1
	q.list = q.list[1:]
	return result
}
