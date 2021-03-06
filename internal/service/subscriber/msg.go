package subscriber

import (
	"github.com/bitrainforest/filmeta-hic/core/threading"
	"github.com/bitrainforest/filmeta-hic/model"
	"github.com/filecoin-project/lotus/chain/types"
)

type PackMsg struct {
	Msg    model.OneMessage
	AppIds []string
}

func NewPackMsg(msg model.OneMessage, appIds []string) PackMsg {
	return PackMsg{Msg: msg, AppIds: appIds}
}

func RangMsg(trace types.ExecutionTrace) <-chan types.Message {
	list, count := countMsg(trace)
	ch := make(chan types.Message, count)
	threading.GoSafe(func() {
		for _, msg := range list {
			ch <- msg
		}
		close(ch)
	})
	return ch
}

func countMsg(trace types.ExecutionTrace) ([]types.Message, int) {
	var (
		list  []types.Message
		total int
	)
	total += 1
	if trace.Msg != nil {
		list = append(list, *trace.Msg)
	}
	for _, sub := range trace.Subcalls {
		subList, subTotal := countMsg(sub)
		list = append(list, subList...)
		total += subTotal
	}
	return list, total
}
