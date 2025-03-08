package commands

// 실행할 명령을 캡슐화
type Command interface {
	Execute()
}
