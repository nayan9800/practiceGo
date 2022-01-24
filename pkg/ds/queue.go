package ds

import "fmt"

/*Queue*/
type Queue struct{ data []int }

/*creates new Queue*/
func NewQueue() Queue { return Queue{data: []int{}} }

/*gives length of Queue*/
func (q *Queue) Len() int { return len(q.data) }

/*Enqueue value from queue*/
func (q *Queue) Enqueue(val int) { q.data = append(q.data, val) }

/*Dequeue value form queue*/
func (q *Queue) Dequeue() (ans int) {
	ans = q.data[0]
	q.data = q.data[1:]
	return
}

/*Peeks the current value in queue*/
func (q *Queue) Peek() int { return q.data[0] }

/*Checks if queue is empty*/
func (q *Queue) IsEmpty() bool { return q.Len() == 0 }

/*Prints queue on standard output*/
func (q *Queue) Show() { fmt.Println(q.data) }
