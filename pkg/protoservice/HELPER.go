package protoservice

import (
	"sync/atomic"

	"github.com/ohanan/card_proto/pkg/protoservice/proto"
)

func (x *pluginXServerAdapter) InitHelper(c *HostXClient) {
	x.Helper = &Helper{HostXClient: c}
}

type Helper struct {
	*HostXClient
	thisID uint64
}

func (h *Helper) NextID() uint64 {
	return atomic.AddUint64(&h.thisID, 1)
}
func (h *Helper) NewOptionBuilder(action *proto.Action) *OptionBuilder {
	return &OptionBuilder{
		helper:      h,
		onSelectMap: map[uint64]func(){},
	}
}

type OptionBuilder struct {
	helper      *Helper
	onSelectMap map[uint64]func()
}

func (ab *OptionBuilder) ID(onSelect func(action proto.UserAction)) *OptionBuilder {
	option.Id = ab.helper.NextID()
	ab.action.Options = append(ab.action.Options, option)
	ab.onSelectMap[option.Id] = onSelect
	return ab
}

func (ab *OptionBuilder) Wait() {
	ab.onSelectMap[ab.helper.AskAction(ab.action).Result.Id]()
}
