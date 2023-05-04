package touch_down

import "context"

type TouchDown interface {
	Send(ctx context.Context) error           // 同步发送
	AsyncSend(ctx context.Context) error      // 异步发送
	BatchSend(ctx context.Context) error      // 批量发送
	BatchAsyncSend(ctx context.Context) error // 批量异步发送
	Revoke(ctx context.Context) error         // 撤销
	Stat(ctx context.Context) error           // 统计
}
