package logclient

import "sync"

type (
	LogQueue struct {
		sync.Mutex
		queue [QUEUE_SIZE + 1][]byte // 有效 log 数据 size 为 24
		tail  int                    // tail 指向的位置不包括 log 数据
	}
)

const QUEUE_SIZE = 2

func NewLogQueue() *LogQueue {
	return &LogQueue{
		queue: [QUEUE_SIZE + 1][]byte{},
		tail:  0,
	}
}

func (q *LogQueue) Len() int {
	q.Lock()
	defer q.Unlock()
	// 通过 tail 计算得到
	return q.tail
}

func (q *LogQueue) Size() int {
	// 通过 len(queue) 计算得到
	return len(q.queue) - 1
}

// Pack 打包当前 queue 的日志, 打包后会重置 queue
func (q *LogQueue) Pack() [][]byte {
	q.Lock()
	defer q.Unlock()
	data := q.queue[:q.tail]
	q.tail = 0
	return data
}

// Enqueue 入队一条日志, 返回当前 queue 是否已满
// 如果已满, 需要调用 Pack, 否则数据会被覆盖
func (q *LogQueue) Enqueue(data []byte) bool {
	q.Lock()
	defer q.Unlock()
	// 判断 tail 是否大于等于 size, 将其设为 0
	if q.tail >= q.Size() {
		q.tail = 0
	}
	q.queue[q.tail] = data
	q.tail += 1
	return q.tail == q.Size()
}
