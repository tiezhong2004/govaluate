package govaluate

import (
	"github.com/tiezhong2004/govaluate/in_functions"
)

var FlowFunctions map[string]ExpressionFunction = map[string]ExpressionFunction{
	"isDefined": in_functions.IsDefined,
	"ifnull":    in_functions.IfNull,
	"if":        in_functions.If,
}

var MathFunctions map[string]ExpressionFunction = map[string]ExpressionFunction{
	"cos":   in_functions.Cos,
	"acos":  in_functions.Acos,
	"cosh":  in_functions.Cosh,
	"acosh": in_functions.Acosh,

	"exp":  in_functions.Exp,
	"exp2": in_functions.Exp2,

	"log":   in_functions.Log,
	"log10": in_functions.Log10,

	"round": in_functions.Round,
	"floor": in_functions.Floor,
	"ceil":  in_functions.Ceil,
	"trunc": in_functions.Trunc,

	"sin":   in_functions.Sin,
	"asin":  in_functions.Asin,
	"sinh":  in_functions.Sinh,
	"asinh": in_functions.Asinh,

	"sqrt": in_functions.Sqrt,

	"tan":   in_functions.Tan,
	"atan":  in_functions.Atan,
	"atan2": in_functions.Atan2,
	"tanh":  in_functions.Tanh,
	"atanh": in_functions.Atanh,

	"max": in_functions.Max,
	"min": in_functions.Min,
}

var Functions map[string]ExpressionFunction

func init() {
	Functions = make(map[string]ExpressionFunction)
	for k, v := range MathFunctions {
		Functions[k] = v
	}
	for k, v := range FlowFunctions {
		Functions[k] = v
	}
}
