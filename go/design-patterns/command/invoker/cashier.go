package invoker

import "command/commands"

// 명령을 받아 실행하는 직원 역할 (Invoker)
type Cashier struct {
	CommandQueue []commands.Command
}

// 명령을 받아 큐에 저장
func (c *Cashier) TakeOrder(cmd commands.Command) {
	c.CommandQueue = append(c.CommandQueue, cmd)
}

// 모든 명령을 순서대로 실행
func (c *Cashier) ExecuteOrders() {
	for _, cmd := range c.CommandQueue {
		cmd.Execute()
	}
	// 실행후 명령 큐 초기화
	c.CommandQueue = []commands.Command{}
}
