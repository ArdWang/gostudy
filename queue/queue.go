package queue

// An FIFO queue
//数据结构
type Queue [] interface{} //可以输入任意类型

// Pushes the element into the queue
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int)) //可以限定类型
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q) [1:]
	return head.(int) //强制转换 int
}

// Pushes the element into the queue is not
func (q *Queue) IsEmpty() bool  {
	return len(*q) == 0
}
