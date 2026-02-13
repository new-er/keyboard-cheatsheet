package ast

import "new-er/keyboard-cheatsheet/data/ast/nodes"

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) parsePrimary() nodes.Node {
	tok := p.tokens[p.pos]
	p.pos++

	var node nodes.Node

	switch tok.Type {
	case TokenKey:
		node = nodes.KeyNode{Name: tok.Value}
	case TokenLParen:
		expr := p.parseExpression()
		p.pos++ // Consume ')'
		node = expr
	default:
		panic("Unexpected token")
	}

	if p.tokens[p.pos].Type == TokenQuestion {
		p.pos++
		node = nodes.OptionalNode{Key: node}
	} else if p.tokens[p.pos].Type == TokenColon {
		p.pos++
		alias := p.tokens[p.pos]
		p.pos++
		node = nodes.AliasNode{Key: node, Alias: alias.Value}
	}

	return node
}

func (p *Parser) parseExpression() nodes.Node {
	left := p.parsePrimary()

	for p.pos < len(p.tokens) {
		tok := p.tokens[p.pos]

		switch tok.Type {
		case TokenPlus:
			p.pos++
			right := p.parsePrimary()
			left = nodes.NewComboNode(left, right)
		case TokenPipe:
			p.pos++
			right := p.parsePrimary()
			left = nodes.NewAnyOfNode(left, right)
		case TokenSequence:
			p.pos++
			right := p.parsePrimary()
			left = nodes.NewSequenceNode(left, right)
		case TokenQuestion:
			p.pos++
			left = nodes.NewOptionalNode(left)
		case TokenColon:
			p.pos++
			alias := p.tokens[p.pos]
			p.pos++
			left = nodes.NewAliasNode(left, alias.Value)
		default:
			return left
		}
	}
	return left
}

func (p *Parser) Parse() nodes.Node {
	return p.parseExpression()
}
