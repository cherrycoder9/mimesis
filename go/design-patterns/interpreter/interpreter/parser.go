package interpreter

import (
	"fmt"
	"strconv"
)

// 토큰을 기반으로 AST(Abstract Syntax Tree)를 구성하는 파서 구현

// 토큰 슬라이스와 현재 파싱 위치를 저장하는 구조체
type Parser struct {
	tokens []Token
	pos    int
}

// 토큰 슬라이스를 받아 Parser 구조체를 생성
func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

// 현재 위치의 토큰을 반환하는 메서드
func (p *Parser) currentToken() Token {
	return p.tokens[p.pos]
}

// 현재 토큰이 기대하는 타입과 일치하면 다음 토큰으로 진행
// 만약 타입이 다르면 에러를 발생시킴
func (p *Parser) eat(tokenType TokenType) {
	if p.currentToken().Type == tokenType {
		p.pos++
	} else {
		panic(fmt.Sprintf("토큰 타입 오류: 예상 %s, 현재 %s", tokenType, p.currentToken().Type))
	}
}

// 아래의 문법 규칙을 처리하는 메서드
// factor = NUMBER | '(' expr ')'
func (p *Parser) factor() Expression {
	token := p.currentToken()
	if token.Type == NUMBER {
		p.eat(NUMBER)
		value, err := strconv.Atoi(token.Value)
		if err != nil {
			panic("숫자 변환 오류: " + token.Value)
		}
		return &NumberExpression{Value: value}
	} else if token.Type == LPAREN {
		p.eat(LPAREN)
		expr := p.expr() // 괄호 안의 표현식을 파싱함
		p.eat(RPAREN)
		return expr
	} else {
		panic("유효하지 않은 토큰: " + string(token.Type))
	}
}

// 아래의 문법 규칙을 처리하는 함수
// term = factor { (*|/) factor }
func (p *Parser) term() Expression {
	node := p.factor()
	for p.currentToken().Type == MUL || p.currentToken().Type == DIV {
		token := p.currentToken()
		if token.Type == MUL {
			p.eat(MUL)
			node = &MultiplyExpression{Left: node, Right: p.factor()}
		} else if token.Type == DIV {
			p.eat(DIV)
			node = &DivideExpression{Left: node, Right: p.factor()}
		}
	}
	return node
}

// 아래 문법 규칙을 처리하는 메서드
// expr = term { (+|-) term }
func (p *Parser) expr() Expression {
	node := p.term()
	for p.currentToken().Type == PLUS || p.currentToken().Type == MINUS {
		token := p.currentToken()
		if token.Type == PLUS {
			p.eat(PLUS)
			node = &AddExpression{Left: node, Right: p.term()}
		} else if token.Type == MINUS {
			p.eat(MINUS)
			node = &SubtractExpression{Left: node, Right: p.term()}
		}
	}
	return node
}

// 전체 입력에 대해 파싱을 수행하고 최종 Expression(AST의 루트)를 반환하는 함수
func (p *Parser) Parse() Expression {
	return p.expr()
}
