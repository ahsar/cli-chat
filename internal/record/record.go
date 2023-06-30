// package record provides chat record
package record

// 聊天记录字典
// k is user id
// v is tokens array
var historyMap map[string]*CircularBuffer

// 历史记录保留条数
var maxNums int = 100

// NewRecord
// size 历史记录保留条数
func NewRecord(size int) {
	maxNums = size
	historyMap = make(map[string]*CircularBuffer)
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
func SetTxtById(i, s string) {
	if _, ok := historyMap[i]; !ok {
		historyMap[i] = NewCircularBuffer(maxNums)
	}
	historyMap[i].Append(s)
}
