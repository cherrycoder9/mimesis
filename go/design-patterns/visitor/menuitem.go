package main

// 방문자를 받아들이는 Accept 메서드 정의
type MenuItem interface {
	Accept(v Visitor)
}
