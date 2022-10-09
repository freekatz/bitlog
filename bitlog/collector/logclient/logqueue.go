package logclient

type (
	LogQueue struct {
		queue      [26][]byte
		head, tail int
		statusCode queueStatusCode // 标记当前 logQueue 的状态
	}
	queueStatusCode int
)

const (
	queueStatus_Reset queueStatusCode = iota
)

func (q *LogQueue) Len() int {
	// 通过 len(queue), head, tail 计算得到
	return 0
}

func (q *LogQueue) Pack() [][]byte {
	// 通过 queue, head, tail 得到
	return nil
}

// Enqueue 入队一条日志, 返回当前 queue 是否已满
func (q *LogQueue) Enqueue([]byte) bool {
	return true
}
