package queue

import "errors"

type ArrayQueue struct {
	values   []int
	front    int
	back     int
	inserted int
}

func (queue *ArrayQueue) Init(size int) error {
	if size <= 0 {
		return errors.New("size não pode ser <= 0")
	}
	queue.values = make([]int, size)
	return nil
}

func (queue *ArrayQueue) doubleArray() {
	newArray := make([]int, 2*len(queue.values))

	for i := 0; i < queue.inserted; i++ {
		newArray[i] = queue.values[(i+queue.front)%len(queue.values)]
	}
	queue.values = newArray

	queue.front = 0
	queue.back = queue.inserted
}

func (queue *ArrayQueue) Enqueue(value int) {
	if queue.inserted == len(queue.values) {
		queue.doubleArray()
	}

	queue.values[queue.back] = value
	queue.back = (queue.back + 1) % len(queue.values)

	queue.inserted++
}

func (queue *ArrayQueue) Dequeue() (int, error) {
	if queue.inserted == 0 {
		return -1, errors.New("Não se pode dar Dequeue em uma queue vazia")
	}

	value := queue.values[queue.front]
	queue.front = (queue.front + 1) % len(queue.values)

	queue.inserted--
	return value, nil
}

func (queue *ArrayQueue) Front() (int, error) {
	if queue.inserted == 0 {
		return -1, errors.New("Nao se pode acessar Front em uma queue vazia")
	}

	return queue.values[queue.front], nil
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.inserted == 0
}

func (queue *ArrayQueue) Size() int {
	return queue.inserted
}
