// package record provides chat record
package record

var (
	// 聊天记录字典
	// k is vxid
	// v is tokens array
	historyMap map[string]*CircularBuffer

	// 最近聊天记录映射
	lastMap map[string]string

	// 历史记录保留条数
	maxNums int = 100
)

// NewRecord
// size 历史记录保留条数
func NewRecord(size int) {
	maxNums = size
	historyMap = make(map[string]*CircularBuffer)
	lastMap = make(map[string]string)
}

// HistoricalById
// 返回指定用户的聊天记录(所有)
func HistoricalById(i string) (h []string) {
	var (
		buf *CircularBuffer
		ok  bool
	)

	if buf, ok = historyMap[i]; !ok {
		return
	}
	return buf.Iterate()
}

// SetTxtById
// 追加与指定用户聊天记录
// id means vxid
func SetTxtById(i, s string) {
	if _, ok := historyMap[i]; !ok {
		historyMap[i] = NewCircularBuffer(maxNums)
	}
	historyMap[i].Append(s)
	SetLastById(i, s)
}

// SetLastById
//
// 注册联系人的最后一条聊天记录
func SetLastById(i, s string) {
	lastMap[i] = s
}
