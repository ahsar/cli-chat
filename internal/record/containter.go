package record

type CircularBuffer struct {
	data     []string
	size     int // 缓冲区的大小
	position int // 当前位置
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		data:     make([]string, size),
		size:     size,
		position: 0,
	}
}

func (buf *CircularBuffer) Append(value string) {
	buf.data[buf.position] = value
	buf.position = (buf.position + 1) % buf.size
}

func (buf *CircularBuffer) GetEarliest() string {
	earliestPos := buf.position
	return buf.data[earliestPos]
}

func (buf *CircularBuffer) Iterate() []string {
	elements := make([]string, buf.size)
	startPos := buf.position
	for i := 0; i < buf.size; i++ {
		elements[i] = buf.data[startPos]
		startPos = (startPos + 1) % buf.size
	}
	return elements
}
