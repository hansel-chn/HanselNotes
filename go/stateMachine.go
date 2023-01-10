package main

/*
不同状态存在fsm中，例如register，heartbeat，synchronization三个状态
nextStatus(corresponding to currentStatus)->map[nextStatus]FsmFunInterface->Action()->newNextStatus
*/
//type StateMachine struct {
//	TimeSecond time.Duration
//	nextStatus enum.FsmStatus
//	errStatus  enum.FsmStatus
//	fsm        map[enum.FsmStatus]FsmFunInterface
//}
