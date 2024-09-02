package main

import (
	"fmt"
	"math"
	"testing"
)

type Expr interface {
	Eval(env Env) float64
}

type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

type unary struct {
	op rune
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

type binary struct {
	op   rune
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	left, right := b.x.Eval(env), b.y.Eval(env)
	switch b.op {
	case '+':
		return left + right
	case '-':
		return left - right
	case '*':
		return left * right
	case '/':
		return left / right
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

type call struct {
	fn   string
	args []Expr
}

func (c call) Eval(env Env) float64 {
	values := make([]float64, len(c.args))
	for ndx, each := range c.args {
		values[ndx] = each.Eval(env)
	}
	switch c.fn {
	case "pow":
		return math.Pow(values[0], values[1])
	case "sin":
		return math.Sin(values[0])
	case "sqrt":
		return math.Sqrt(values[0])
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

type Env map[Var]float64

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"}
	}
	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %s = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
