package evaluator

import (
	"github.com/kasworld/nonkey/config/builtinfunctions"
	"github.com/kasworld/nonkey/interpreter/object"
)

func init() {
	builtinfunctions.BuiltinFunctions = map[string]*object.Builtin{
		"version":        {Fn: builtinVersion},
		"args":           {Fn: builtinArgs},
		"chmod":          {Fn: builtinChmod},
		"delete":         {Fn: builtinDelete},
		"eval":           {Fn: builtinEval},
		"exit":           {Fn: builtinExit},
		"int":            {Fn: builtinInt},
		"keys":           {Fn: builtinKeys},
		"len":            {Fn: builtinLen},
		"match":          {Fn: builtinMatch},
		"mkdir":          {Fn: builtinMkdir},
		"pragma":         {Fn: builtinPragma},
		"open":           {Fn: builtinOpen},
		"push":           {Fn: builtinPush},
		"puts":           {Fn: builtinPuts},
		"printf":         {Fn: builtinPrintf},
		"set":            {Fn: builtinSet},
		"sprintf":        {Fn: builtinSprintf},
		"stat":           {Fn: builtinStat},
		"string":         {Fn: builtinString},
		"type":           {Fn: builtinType},
		"unlink":         {Fn: builtinUnlink},
		"os.getenv":      {Fn: builtinOsGetEnv},
		"os.setenv":      {Fn: builtinOsSetEnv},
		"os.environment": {Fn: builtinOsEnvironment},
		"directory.glob": {Fn: builtinDirectoryGlob},
		"math.abs":       {Fn: builtinMathAbs},
		"math.random":    {Fn: builtinMathRandom},
		"math.sqrt":      {Fn: builtinMathSqrt},
	}
}
