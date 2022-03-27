package dailygo

/*
You run an e-commerce website and want to record the last N order ids in a log.
Implement a data structure to accomplish this, with the following API:
record(order_id): adds the order_id to the log
get_last(i): gets the ith last element from the log. i is guaranteed to be
smaller than or equal to N.
You should be as efficient with time and space as possible.
*/
func _20220327(size int, logs []int, last int) int {
	nLog := _20220327_newLogStore(size)
	for _, log := range logs {
		nLog._20220327_record(log)
	}
	return nLog._20220327_getLast(last)
}

type _20220327_nLog struct {
	logIds     []int // circular log with max size
	size       int
	currentNdx int
}

func _20220327_newLogStore(size int) *_20220327_nLog {
	return &_20220327_nLog{
		logIds:     make([]int, size),
		size:       size,
		currentNdx: 0,
	}
}

func (n *_20220327_nLog) _20220327_record(orderId int) {
	n.currentNdx = (n.currentNdx + 1) % n.size
	n.logIds[n.currentNdx] = orderId
}

func (n *_20220327_nLog) _20220327_getLast(i int) int {
	if i > n.size {
		panic("i must be <= n.size")
	}
	return n.logIds[(n.currentNdx+n.size-i)%n.size]
}
