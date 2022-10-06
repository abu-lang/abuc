// Code generated from AbuParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AbuParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AbuParser struct {
	*antlr.BaseParser
}

var abuparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func abuparserParserInit() {
	staticData := &abuparserParserStaticData
	staticData.literalNames = []string{
		"", "'and'", "'or'", "'not'", "'abs'", "'.'", "'+'", "'-'", "'/'", "'*'",
		"'%'", "'=='", "'>'", "'<'", "'>='", "'<='", "'!='", "'::'", "':'",
		"';'", "':='", "'='", "','", "'('", "')'", "'['", "']'", "'{'", "'}'",
		"'has'", "'where'", "'physical'", "'logical'", "'input'", "'output'",
		"'rule'", "", "", "'on'", "'for'", "'all'", "'do'", "'define'", "'as'",
		"'boolean'", "'integer'", "'decimal'", "'string'", "'true'", "'false'",
	}
	staticData.symbolicNames = []string{
		"", "AND", "OR", "NOT", "ABS", "DOT", "PLUS", "MINUS", "DIV", "MUL",
		"MOD", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "DOUBLECOLON",
		"COLON", "SEMICOLON", "COLONEQUAL", "EQUALSIGN", "COMMA", "RL_BRACKET",
		"RR_BRACKET", "SL_BRACKET", "SR_BRACKET", "CL_BRACKET", "CR_BRACKET",
		"HAS", "WHERE", "PHYSICAL", "LOGICAL", "INPUT", "OUTPUT", "RULE", "THIS",
		"EXT", "ON", "FOR", "ALL", "DO", "DEFINE", "AS", "BOOLEAN", "INTEGER",
		"DECIMAL", "STRING", "TRUE", "FALSE", "ID", "QUOTED_STRING", "DEC_LITERAL",
		"INT_LITERAL", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"program", "typeDecl", "resField", "device", "resList", "resDecl", "compResDecl",
		"type", "ecarule", "task", "actions", "assignment", "expression", "mulDivModOperator",
		"plusMinusOperator", "comparisonOperator", "term", "value", "resource",
		"simpleResource", "nestedResource", "decimalLiteral", "integerLiteral",
		"stringLiteral", "booleanLiteral",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 56, 284, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 5, 0, 52, 8,
		0, 10, 0, 12, 0, 55, 9, 0, 1, 0, 1, 0, 5, 0, 59, 8, 0, 10, 0, 12, 0, 62,
		9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 69, 8, 1, 11, 1, 12, 1, 70, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 89, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 98, 8, 3, 1, 3, 1, 3, 1, 3, 4, 3, 103, 8, 3, 11, 3, 12, 3,
		104, 3, 3, 107, 8, 3, 1, 4, 4, 4, 110, 8, 4, 11, 4, 12, 4, 111, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 133, 8, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 146, 8, 6, 10, 6, 12,
		6, 149, 9, 6, 1, 6, 1, 6, 3, 6, 153, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 4, 8, 161, 8, 8, 11, 8, 12, 8, 162, 1, 8, 4, 8, 166, 8, 8, 11, 8,
		12, 8, 167, 1, 9, 1, 9, 3, 9, 172, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 5, 10, 181, 8, 10, 10, 10, 12, 10, 184, 9, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 3, 12, 200, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 223, 8, 12, 10, 12, 12, 12, 226,
		9, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 3, 16, 236,
		8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 242, 8, 17, 1, 18, 1, 18, 3,
		18, 246, 8, 18, 1, 19, 3, 19, 249, 8, 19, 1, 19, 1, 19, 1, 19, 3, 19, 254,
		8, 19, 1, 20, 3, 20, 257, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 3, 20, 268, 8, 20, 1, 21, 3, 21, 271, 8, 21, 1,
		21, 1, 21, 1, 22, 3, 22, 276, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 0, 1, 24, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 5, 1, 0, 44, 47,
		1, 0, 8, 10, 1, 0, 6, 7, 1, 0, 11, 16, 1, 0, 48, 49, 296, 0, 53, 1, 0,
		0, 0, 2, 63, 1, 0, 0, 0, 4, 88, 1, 0, 0, 0, 6, 90, 1, 0, 0, 0, 8, 109,
		1, 0, 0, 0, 10, 132, 1, 0, 0, 0, 12, 134, 1, 0, 0, 0, 14, 154, 1, 0, 0,
		0, 16, 156, 1, 0, 0, 0, 18, 169, 1, 0, 0, 0, 20, 177, 1, 0, 0, 0, 22, 185,
		1, 0, 0, 0, 24, 199, 1, 0, 0, 0, 26, 227, 1, 0, 0, 0, 28, 229, 1, 0, 0,
		0, 30, 231, 1, 0, 0, 0, 32, 235, 1, 0, 0, 0, 34, 241, 1, 0, 0, 0, 36, 245,
		1, 0, 0, 0, 38, 253, 1, 0, 0, 0, 40, 267, 1, 0, 0, 0, 42, 270, 1, 0, 0,
		0, 44, 275, 1, 0, 0, 0, 46, 279, 1, 0, 0, 0, 48, 281, 1, 0, 0, 0, 50, 52,
		3, 2, 1, 0, 51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0,
		53, 54, 1, 0, 0, 0, 54, 56, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 60, 3,
		6, 3, 0, 57, 59, 3, 16, 8, 0, 58, 57, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60,
		58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 1, 1, 0, 0, 0, 62, 60, 1, 0, 0,
		0, 63, 64, 5, 42, 0, 0, 64, 65, 5, 50, 0, 0, 65, 66, 5, 43, 0, 0, 66, 68,
		5, 27, 0, 0, 67, 69, 3, 4, 2, 0, 68, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0,
		70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 5,
		28, 0, 0, 73, 3, 1, 0, 0, 0, 74, 75, 5, 50, 0, 0, 75, 76, 5, 18, 0, 0,
		76, 77, 5, 31, 0, 0, 77, 78, 5, 34, 0, 0, 78, 89, 3, 14, 7, 0, 79, 80,
		5, 50, 0, 0, 80, 81, 5, 18, 0, 0, 81, 82, 5, 31, 0, 0, 82, 83, 5, 33, 0,
		0, 83, 89, 3, 14, 7, 0, 84, 85, 5, 50, 0, 0, 85, 86, 5, 18, 0, 0, 86, 87,
		5, 32, 0, 0, 87, 89, 3, 14, 7, 0, 88, 74, 1, 0, 0, 0, 88, 79, 1, 0, 0,
		0, 88, 84, 1, 0, 0, 0, 89, 5, 1, 0, 0, 0, 90, 91, 5, 50, 0, 0, 91, 92,
		5, 18, 0, 0, 92, 93, 3, 46, 23, 0, 93, 94, 5, 27, 0, 0, 94, 97, 3, 8, 4,
		0, 95, 96, 5, 30, 0, 0, 96, 98, 3, 24, 12, 0, 97, 95, 1, 0, 0, 0, 97, 98,
		1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 106, 5, 28, 0, 0, 100, 102, 5, 29,
		0, 0, 101, 103, 5, 50, 0, 0, 102, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0,
		104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 107, 1, 0, 0, 0, 106,
		100, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 7, 1, 0, 0, 0, 108, 110, 3,
		10, 5, 0, 109, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 109, 1, 0, 0,
		0, 111, 112, 1, 0, 0, 0, 112, 9, 1, 0, 0, 0, 113, 114, 5, 31, 0, 0, 114,
		115, 5, 34, 0, 0, 115, 116, 3, 14, 7, 0, 116, 117, 5, 50, 0, 0, 117, 118,
		5, 21, 0, 0, 118, 119, 3, 24, 12, 0, 119, 133, 1, 0, 0, 0, 120, 121, 5,
		31, 0, 0, 121, 122, 5, 33, 0, 0, 122, 123, 3, 14, 7, 0, 123, 124, 5, 50,
		0, 0, 124, 133, 1, 0, 0, 0, 125, 126, 5, 32, 0, 0, 126, 127, 3, 14, 7,
		0, 127, 128, 5, 50, 0, 0, 128, 129, 5, 21, 0, 0, 129, 130, 3, 24, 12, 0,
		130, 133, 1, 0, 0, 0, 131, 133, 3, 12, 6, 0, 132, 113, 1, 0, 0, 0, 132,
		120, 1, 0, 0, 0, 132, 125, 1, 0, 0, 0, 132, 131, 1, 0, 0, 0, 133, 11, 1,
		0, 0, 0, 134, 135, 5, 50, 0, 0, 135, 152, 5, 50, 0, 0, 136, 137, 5, 21,
		0, 0, 137, 138, 5, 23, 0, 0, 138, 139, 5, 50, 0, 0, 139, 140, 5, 21, 0,
		0, 140, 147, 3, 24, 12, 0, 141, 142, 5, 22, 0, 0, 142, 143, 5, 50, 0, 0,
		143, 144, 5, 21, 0, 0, 144, 146, 3, 24, 12, 0, 145, 141, 1, 0, 0, 0, 146,
		149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 150,
		1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 151, 5, 24, 0, 0, 151, 153, 1, 0,
		0, 0, 152, 136, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 13, 1, 0, 0, 0,
		154, 155, 7, 0, 0, 0, 155, 15, 1, 0, 0, 0, 156, 157, 5, 35, 0, 0, 157,
		158, 5, 50, 0, 0, 158, 160, 5, 38, 0, 0, 159, 161, 3, 36, 18, 0, 160, 159,
		1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0,
		0, 0, 163, 165, 1, 0, 0, 0, 164, 166, 3, 18, 9, 0, 165, 164, 1, 0, 0, 0,
		166, 167, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168,
		17, 1, 0, 0, 0, 169, 171, 5, 39, 0, 0, 170, 172, 5, 40, 0, 0, 171, 170,
		1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 3, 24,
		12, 0, 174, 175, 5, 41, 0, 0, 175, 176, 3, 20, 10, 0, 176, 19, 1, 0, 0,
		0, 177, 182, 3, 22, 11, 0, 178, 179, 5, 19, 0, 0, 179, 181, 3, 22, 11,
		0, 180, 178, 1, 0, 0, 0, 181, 184, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182,
		183, 1, 0, 0, 0, 183, 21, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 185, 186, 3,
		36, 18, 0, 186, 187, 5, 21, 0, 0, 187, 188, 3, 24, 12, 0, 188, 23, 1, 0,
		0, 0, 189, 190, 6, 12, -1, 0, 190, 191, 5, 4, 0, 0, 191, 200, 3, 24, 12,
		10, 192, 193, 5, 3, 0, 0, 193, 200, 3, 24, 12, 9, 194, 195, 5, 23, 0, 0,
		195, 196, 3, 24, 12, 0, 196, 197, 5, 24, 0, 0, 197, 200, 1, 0, 0, 0, 198,
		200, 3, 32, 16, 0, 199, 189, 1, 0, 0, 0, 199, 192, 1, 0, 0, 0, 199, 194,
		1, 0, 0, 0, 199, 198, 1, 0, 0, 0, 200, 224, 1, 0, 0, 0, 201, 202, 10, 8,
		0, 0, 202, 203, 3, 26, 13, 0, 203, 204, 3, 24, 12, 9, 204, 223, 1, 0, 0,
		0, 205, 206, 10, 7, 0, 0, 206, 207, 3, 28, 14, 0, 207, 208, 3, 24, 12,
		8, 208, 223, 1, 0, 0, 0, 209, 210, 10, 6, 0, 0, 210, 211, 3, 30, 15, 0,
		211, 212, 3, 24, 12, 7, 212, 223, 1, 0, 0, 0, 213, 214, 10, 5, 0, 0, 214,
		215, 5, 17, 0, 0, 215, 223, 3, 24, 12, 6, 216, 217, 10, 4, 0, 0, 217, 218,
		5, 1, 0, 0, 218, 223, 3, 24, 12, 5, 219, 220, 10, 3, 0, 0, 220, 221, 5,
		2, 0, 0, 221, 223, 3, 24, 12, 4, 222, 201, 1, 0, 0, 0, 222, 205, 1, 0,
		0, 0, 222, 209, 1, 0, 0, 0, 222, 213, 1, 0, 0, 0, 222, 216, 1, 0, 0, 0,
		222, 219, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224,
		225, 1, 0, 0, 0, 225, 25, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 228, 7,
		1, 0, 0, 228, 27, 1, 0, 0, 0, 229, 230, 7, 2, 0, 0, 230, 29, 1, 0, 0, 0,
		231, 232, 7, 3, 0, 0, 232, 31, 1, 0, 0, 0, 233, 236, 3, 34, 17, 0, 234,
		236, 3, 36, 18, 0, 235, 233, 1, 0, 0, 0, 235, 234, 1, 0, 0, 0, 236, 33,
		1, 0, 0, 0, 237, 242, 3, 46, 23, 0, 238, 242, 3, 44, 22, 0, 239, 242, 3,
		42, 21, 0, 240, 242, 3, 48, 24, 0, 241, 237, 1, 0, 0, 0, 241, 238, 1, 0,
		0, 0, 241, 239, 1, 0, 0, 0, 241, 240, 1, 0, 0, 0, 242, 35, 1, 0, 0, 0,
		243, 246, 3, 38, 19, 0, 244, 246, 3, 40, 20, 0, 245, 243, 1, 0, 0, 0, 245,
		244, 1, 0, 0, 0, 246, 37, 1, 0, 0, 0, 247, 249, 5, 36, 0, 0, 248, 247,
		1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 254, 5, 50,
		0, 0, 251, 252, 5, 37, 0, 0, 252, 254, 5, 50, 0, 0, 253, 248, 1, 0, 0,
		0, 253, 251, 1, 0, 0, 0, 254, 39, 1, 0, 0, 0, 255, 257, 5, 36, 0, 0, 256,
		255, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 259,
		5, 50, 0, 0, 259, 260, 5, 25, 0, 0, 260, 261, 5, 50, 0, 0, 261, 268, 5,
		26, 0, 0, 262, 263, 5, 37, 0, 0, 263, 264, 5, 50, 0, 0, 264, 265, 5, 25,
		0, 0, 265, 266, 5, 50, 0, 0, 266, 268, 5, 26, 0, 0, 267, 256, 1, 0, 0,
		0, 267, 262, 1, 0, 0, 0, 268, 41, 1, 0, 0, 0, 269, 271, 5, 7, 0, 0, 270,
		269, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 273,
		5, 52, 0, 0, 273, 43, 1, 0, 0, 0, 274, 276, 5, 7, 0, 0, 275, 274, 1, 0,
		0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 5, 53, 0, 0,
		278, 45, 1, 0, 0, 0, 279, 280, 5, 51, 0, 0, 280, 47, 1, 0, 0, 0, 281, 282,
		7, 4, 0, 0, 282, 49, 1, 0, 0, 0, 27, 53, 60, 70, 88, 97, 104, 106, 111,
		132, 147, 152, 162, 167, 171, 182, 199, 222, 224, 235, 241, 245, 248, 253,
		256, 267, 270, 275,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// AbuParserInit initializes any static state used to implement AbuParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAbuParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AbuParserInit() {
	staticData := &abuparserParserStaticData
	staticData.once.Do(abuparserParserInit)
}

// NewAbuParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAbuParser(input antlr.TokenStream) *AbuParser {
	AbuParserInit()
	this := new(AbuParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &abuparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AbuParser.g4"

	return this
}

// AbuParser tokens.
const (
	AbuParserEOF           = antlr.TokenEOF
	AbuParserAND           = 1
	AbuParserOR            = 2
	AbuParserNOT           = 3
	AbuParserABS           = 4
	AbuParserDOT           = 5
	AbuParserPLUS          = 6
	AbuParserMINUS         = 7
	AbuParserDIV           = 8
	AbuParserMUL           = 9
	AbuParserMOD           = 10
	AbuParserEQUALS        = 11
	AbuParserGT            = 12
	AbuParserLT            = 13
	AbuParserGTE           = 14
	AbuParserLTE           = 15
	AbuParserNOTEQUALS     = 16
	AbuParserDOUBLECOLON   = 17
	AbuParserCOLON         = 18
	AbuParserSEMICOLON     = 19
	AbuParserCOLONEQUAL    = 20
	AbuParserEQUALSIGN     = 21
	AbuParserCOMMA         = 22
	AbuParserRL_BRACKET    = 23
	AbuParserRR_BRACKET    = 24
	AbuParserSL_BRACKET    = 25
	AbuParserSR_BRACKET    = 26
	AbuParserCL_BRACKET    = 27
	AbuParserCR_BRACKET    = 28
	AbuParserHAS           = 29
	AbuParserWHERE         = 30
	AbuParserPHYSICAL      = 31
	AbuParserLOGICAL       = 32
	AbuParserINPUT         = 33
	AbuParserOUTPUT        = 34
	AbuParserRULE          = 35
	AbuParserTHIS          = 36
	AbuParserEXT           = 37
	AbuParserON            = 38
	AbuParserFOR           = 39
	AbuParserALL           = 40
	AbuParserDO            = 41
	AbuParserDEFINE        = 42
	AbuParserAS            = 43
	AbuParserBOOLEAN       = 44
	AbuParserINTEGER       = 45
	AbuParserDECIMAL       = 46
	AbuParserSTRING        = 47
	AbuParserTRUE          = 48
	AbuParserFALSE         = 49
	AbuParserID            = 50
	AbuParserQUOTED_STRING = 51
	AbuParserDEC_LITERAL   = 52
	AbuParserINT_LITERAL   = 53
	AbuParserWS            = 54
	AbuParserCOMMENT       = 55
	AbuParserLINE_COMMENT  = 56
)

// AbuParser rules.
const (
	AbuParserRULE_program            = 0
	AbuParserRULE_typeDecl           = 1
	AbuParserRULE_resField           = 2
	AbuParserRULE_device             = 3
	AbuParserRULE_resList            = 4
	AbuParserRULE_resDecl            = 5
	AbuParserRULE_compResDecl        = 6
	AbuParserRULE_type               = 7
	AbuParserRULE_ecarule            = 8
	AbuParserRULE_task               = 9
	AbuParserRULE_actions            = 10
	AbuParserRULE_assignment         = 11
	AbuParserRULE_expression         = 12
	AbuParserRULE_mulDivModOperator  = 13
	AbuParserRULE_plusMinusOperator  = 14
	AbuParserRULE_comparisonOperator = 15
	AbuParserRULE_term               = 16
	AbuParserRULE_value              = 17
	AbuParserRULE_resource           = 18
	AbuParserRULE_simpleResource     = 19
	AbuParserRULE_nestedResource     = 20
	AbuParserRULE_decimalLiteral     = 21
	AbuParserRULE_integerLiteral     = 22
	AbuParserRULE_stringLiteral      = 23
	AbuParserRULE_booleanLiteral     = 24
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Device() IDeviceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeviceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeviceContext)
}

func (s *ProgramContext) AllTypeDecl() []ITypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclContext); ok {
			tst[i] = t.(ITypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) TypeDecl(i int) ITypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *ProgramContext) AllEcarule() []IEcaruleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEcaruleContext); ok {
			len++
		}
	}

	tst := make([]IEcaruleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEcaruleContext); ok {
			tst[i] = t.(IEcaruleContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Ecarule(i int) IEcaruleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEcaruleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEcaruleContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *AbuParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AbuParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AbuParserDEFINE {
		{
			p.SetState(50)
			p.TypeDecl()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Device()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AbuParserRULE {
		{
			p.SetState(57)
			p.Ecarule()
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(AbuParserDEFINE, 0)
}

func (s *TypeDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(AbuParserID, 0)
}

func (s *TypeDeclContext) AS() antlr.TerminalNode {
	return s.GetToken(AbuParserAS, 0)
}

func (s *TypeDeclContext) CL_BRACKET() antlr.TerminalNode {
	return s.GetToken(AbuParserCL_BRACKET, 0)
}

func (s *TypeDeclContext) CR_BRACKET() antlr.TerminalNode {
	return s.GetToken(AbuParserCR_BRACKET, 0)
}

func (s *TypeDeclContext) AllResField() []IResFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResFieldContext); ok {
			len++
		}
	}

	tst := make([]IResFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResFieldContext); ok {
			tst[i] = t.(IResFieldContext)
			i++
		}
	}

	return tst
}

func (s *TypeDeclContext) ResField(i int) IResFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResFieldContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (p *AbuParser) TypeDecl() (localctx ITypeDeclContext) {
	this := p
	_ = this

	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AbuParserRULE_typeDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(AbuParserDEFINE)
	}
	{
		p.SetState(64)
		p.Match(AbuParserID)
	}
	{
		p.SetState(65)
		p.Match(AbuParserAS)
	}
	{
		p.SetState(66)
		p.Match(AbuParserCL_BRACKET)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AbuParserID {
		{
			p.SetState(67)
			p.ResField()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(AbuParserCR_BRACKET)
	}

	return localctx
}

// IResFieldContext is an interface to support dynamic dispatch.
type IResFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResFieldContext differentiates from other interfaces.
	IsResFieldContext()
}

type ResFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResFieldContext() *ResFieldContext {
	var p = new(ResFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_resField
	return p
}

func (*ResFieldContext) IsResFieldContext() {}

func NewResFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResFieldContext {
	var p = new(ResFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_resField

	return p
}

func (s *ResFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ResFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(AbuParserID, 0)
}

func (s *ResFieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(AbuParserCOLON, 0)
}

func (s *ResFieldContext) PHYSICAL() antlr.TerminalNode {
	return s.GetToken(AbuParserPHYSICAL, 0)
}

func (s *ResFieldContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(AbuParserOUTPUT, 0)
}

func (s *ResFieldContext) Type() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ResFieldContext) INPUT() antlr.TerminalNode {
	return s.GetToken(AbuParserINPUT, 0)
}

func (s *ResFieldContext) LOGICAL() antlr.TerminalNode {
	return s.GetToken(AbuParserLOGICAL, 0)
}

func (s *ResFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterResField(s)
	}
}

func (s *ResFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitResField(s)
	}
}

func (p *AbuParser) ResField() (localctx IResFieldContext) {
	this := p
	_ = this

	localctx = NewResFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AbuParserRULE_resField)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Match(AbuParserID)
		}
		{
			p.SetState(75)
			p.Match(AbuParserCOLON)
		}
		{
			p.SetState(76)
			p.Match(AbuParserPHYSICAL)
		}
		{
			p.SetState(77)
			p.Match(AbuParserOUTPUT)
		}
		{
			p.SetState(78)
			p.Type()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Match(AbuParserID)
		}
		{
			p.SetState(80)
			p.Match(AbuParserCOLON)
		}
		{
			p.SetState(81)
			p.Match(AbuParserPHYSICAL)
		}
		{
			p.SetState(82)
			p.Match(AbuParserINPUT)
		}
		{
			p.SetState(83)
			p.Type()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.Match(AbuParserID)
		}
		{
			p.SetState(85)
			p.Match(AbuParserCOLON)
		}
		{
			p.SetState(86)
			p.Match(AbuParserLOGICAL)
		}
		{
			p.SetState(87)
			p.Type()
		}

	}

	return localctx
}

// IDeviceContext is an interface to support dynamic dispatch.
type IDeviceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeviceContext differentiates from other interfaces.
	IsDeviceContext()
}

type DeviceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeviceContext() *DeviceContext {
	var p = new(DeviceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_device
	return p
}

func (*DeviceContext) IsDeviceContext() {}

func NewDeviceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeviceContext {
	var p = new(DeviceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_device

	return p
}

func (s *DeviceContext) GetParser() antlr.Parser { return s.parser }

func (s *DeviceContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AbuParserID)
}

func (s *DeviceContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AbuParserID, i)
}

func (s *DeviceContext) COLON() antlr.TerminalNode {
	return s.GetToken(AbuParserCOLON, 0)
}

func (s *DeviceContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *DeviceContext) CL_BRACKET() antlr.TerminalNode {
	return s.GetToken(AbuParserCL_BRACKET, 0)
}

func (s *DeviceContext) ResList() IResListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResListContext)
}

func (s *DeviceContext) CR_BRACKET() antlr.TerminalNode {
	return s.GetToken(AbuParserCR_BRACKET, 0)
}

func (s *DeviceContext) WHERE() antlr.TerminalNode {
	return s.GetToken(AbuParserWHERE, 0)
}

func (s *DeviceContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeviceContext) HAS() antlr.TerminalNode {
	return s.GetToken(AbuParserHAS, 0)
}

func (s *DeviceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeviceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeviceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterDevice(s)
	}
}

func (s *DeviceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitDevice(s)
	}
}

func (p *AbuParser) Device() (localctx IDeviceContext) {
	this := p
	_ = this

	localctx = NewDeviceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AbuParserRULE_device)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(AbuParserID)
	}
	{
		p.SetState(91)
		p.Match(AbuParserCOLON)
	}
	{
		p.SetState(92)
		p.StringLiteral()
	}
	{
		p.SetState(93)
		p.Match(AbuParserCL_BRACKET)
	}
	{
		p.SetState(94)
		p.ResList()
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbuParserWHERE {
		{
			p.SetState(95)
			p.Match(AbuParserWHERE)
		}
		{
			p.SetState(96)
			p.expression(0)
		}

	}
	{
		p.SetState(99)
		p.Match(AbuParserCR_BRACKET)
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbuParserHAS {
		{
			p.SetState(100)
			p.Match(AbuParserHAS)
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AbuParserID {
			{
				p.SetState(101)
				p.Match(AbuParserID)
			}

			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IResListContext is an interface to support dynamic dispatch.
type IResListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResListContext differentiates from other interfaces.
	IsResListContext()
}

type ResListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResListContext() *ResListContext {
	var p = new(ResListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_resList
	return p
}

func (*ResListContext) IsResListContext() {}

func NewResListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResListContext {
	var p = new(ResListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_resList

	return p
}

func (s *ResListContext) GetParser() antlr.Parser { return s.parser }

func (s *ResListContext) AllResDecl() []IResDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResDeclContext); ok {
			len++
		}
	}

	tst := make([]IResDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResDeclContext); ok {
			tst[i] = t.(IResDeclContext)
			i++
		}
	}

	return tst
}

func (s *ResListContext) ResDecl(i int) IResDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResDeclContext)
}

func (s *ResListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterResList(s)
	}
}

func (s *ResListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitResList(s)
	}
}

func (p *AbuParser) ResList() (localctx IResListContext) {
	this := p
	_ = this

	localctx = NewResListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AbuParserRULE_resList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(AbuParserPHYSICAL-31))|(1<<(AbuParserLOGICAL-31))|(1<<(AbuParserID-31)))) != 0) {
		{
			p.SetState(108)
			p.ResDecl()
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IResDeclContext is an interface to support dynamic dispatch.
type IResDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResDeclContext differentiates from other interfaces.
	IsResDeclContext()
}

type ResDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResDeclContext() *ResDeclContext {
	var p = new(ResDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_resDecl
	return p
}

func (*ResDeclContext) IsResDeclContext() {}

func NewResDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResDeclContext {
	var p = new(ResDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_resDecl

	return p
}

func (s *ResDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ResDeclContext) PHYSICAL() antlr.TerminalNode {
	return s.GetToken(AbuParserPHYSICAL, 0)
}

func (s *ResDeclContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(AbuParserOUTPUT, 0)
}

func (s *ResDeclContext) Type() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ResDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(AbuParserID, 0)
}

func (s *ResDeclContext) EQUALSIGN() antlr.TerminalNode {
	return s.GetToken(AbuParserEQUALSIGN, 0)
}

func (s *ResDeclContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ResDeclContext) INPUT() antlr.TerminalNode {
	return s.GetToken(AbuParserINPUT, 0)
}

func (s *ResDeclContext) LOGICAL() antlr.TerminalNode {
	return s.GetToken(AbuParserLOGICAL, 0)
}

func (s *ResDeclContext) CompResDecl() ICompResDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompResDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompResDeclContext)
}

func (s *ResDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterResDecl(s)
	}
}

func (s *ResDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitResDecl(s)
	}
}

func (p *AbuParser) ResDecl() (localctx IResDeclContext) {
	this := p
	_ = this

	localctx = NewResDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AbuParserRULE_resDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.Match(AbuParserPHYSICAL)
		}
		{
			p.SetState(114)
			p.Match(AbuParserOUTPUT)
		}
		{
			p.SetState(115)
			p.Type()
		}
		{
			p.SetState(116)
			p.Match(AbuParserID)
		}
		{
			p.SetState(117)
			p.Match(AbuParserEQUALSIGN)
		}
		{
			p.SetState(118)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Match(AbuParserPHYSICAL)
		}
		{
			p.SetState(121)
			p.Match(AbuParserINPUT)
		}
		{
			p.SetState(122)
			p.Type()
		}
		{
			p.SetState(123)
			p.Match(AbuParserID)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Match(AbuParserLOGICAL)
		}
		{
			p.SetState(126)
			p.Type()
		}
		{
			p.SetState(127)
			p.Match(AbuParserID)
		}
		{
			p.SetState(128)
			p.Match(AbuParserEQUALSIGN)
		}
		{
			p.SetState(129)
			p.expression(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			p.CompResDecl()
		}

	}

	return localctx
}

// ICompResDeclContext is an interface to support dynamic dispatch.
type ICompResDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompResDeclContext differentiates from other interfaces.
	IsCompResDeclContext()
}

type CompResDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompResDeclContext() *CompResDeclContext {
	var p = new(CompResDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_compResDecl
	return p
}

func (*CompResDeclContext) IsCompResDeclContext() {}

func NewCompResDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompResDeclContext {
	var p = new(CompResDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_compResDecl

	return p
}

func (s *CompResDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *CompResDeclContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AbuParserID)
}

func (s *CompResDeclContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AbuParserID, i)
}

func (s *CompResDeclContext) AllEQUALSIGN() []antlr.TerminalNode {
	return s.GetTokens(AbuParserEQUALSIGN)
}

func (s *CompResDeclContext) EQUALSIGN(i int) antlr.TerminalNode {
	return s.GetToken(AbuParserEQUALSIGN, i)
}

func (s *CompResDeclContext) RL_BRACKET() antlr.TerminalNode {
	return s.GetToken(AbuParserRL_BRACKET, 0)
}

func (s *CompResDeclContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CompResDeclContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompResDeclContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(AbuParserRR_BRACKET, 0)
}

func (s *CompResDeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(AbuParserCOMMA)
}

func (s *CompResDeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(AbuParserCOMMA, i)
}

func (s *CompResDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompResDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompResDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterCompResDecl(s)
	}
}

func (s *CompResDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitCompResDecl(s)
	}
}

func (p *AbuParser) CompResDecl() (localctx ICompResDeclContext) {
	this := p
	_ = this

	localctx = NewCompResDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AbuParserRULE_compResDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(AbuParserID)
	}
	{
		p.SetState(135)
		p.Match(AbuParserID)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbuParserEQUALSIGN {
		{
			p.SetState(136)
			p.Match(AbuParserEQUALSIGN)
		}
		{
			p.SetState(137)
			p.Match(AbuParserRL_BRACKET)
		}
		{
			p.SetState(138)
			p.Match(AbuParserID)
		}
		{
			p.SetState(139)
			p.Match(AbuParserEQUALSIGN)
		}
		{
			p.SetState(140)
			p.expression(0)
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AbuParserCOMMA {
			{
				p.SetState(141)
				p.Match(AbuParserCOMMA)
			}
			{
				p.SetState(142)
				p.Match(AbuParserID)
			}
			{
				p.SetState(143)
				p.Match(AbuParserEQUALSIGN)
			}
			{
				p.SetState(144)
				p.expression(0)
			}

			p.SetState(149)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(150)
			p.Match(AbuParserRR_BRACKET)
		}

	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(AbuParserBOOLEAN, 0)
}

func (s *TypeContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(AbuParserINTEGER, 0)
}

func (s *TypeContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(AbuParserDECIMAL, 0)
}

func (s *TypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(AbuParserSTRING, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *AbuParser) Type() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AbuParserRULE_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(AbuParserBOOLEAN-44))|(1<<(AbuParserINTEGER-44))|(1<<(AbuParserDECIMAL-44))|(1<<(AbuParserSTRING-44)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEcaruleContext is an interface to support dynamic dispatch.
type IEcaruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEcaruleContext differentiates from other interfaces.
	IsEcaruleContext()
}

type EcaruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEcaruleContext() *EcaruleContext {
	var p = new(EcaruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_ecarule
	return p
}

func (*EcaruleContext) IsEcaruleContext() {}

func NewEcaruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EcaruleContext {
	var p = new(EcaruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_ecarule

	return p
}

func (s *EcaruleContext) GetParser() antlr.Parser { return s.parser }

func (s *EcaruleContext) RULE() antlr.TerminalNode {
	return s.GetToken(AbuParserRULE, 0)
}

func (s *EcaruleContext) ID() antlr.TerminalNode {
	return s.GetToken(AbuParserID, 0)
}

func (s *EcaruleContext) ON() antlr.TerminalNode {
	return s.GetToken(AbuParserON, 0)
}

func (s *EcaruleContext) AllResource() []IResourceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IResourceContext); ok {
			len++
		}
	}

	tst := make([]IResourceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IResourceContext); ok {
			tst[i] = t.(IResourceContext)
			i++
		}
	}

	return tst
}

func (s *EcaruleContext) Resource(i int) IResourceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResourceContext)
}

func (s *EcaruleContext) AllTask() []ITaskContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITaskContext); ok {
			len++
		}
	}

	tst := make([]ITaskContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITaskContext); ok {
			tst[i] = t.(ITaskContext)
			i++
		}
	}

	return tst
}

func (s *EcaruleContext) Task(i int) ITaskContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITaskContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITaskContext)
}

func (s *EcaruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EcaruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EcaruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterEcarule(s)
	}
}

func (s *EcaruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitEcarule(s)
	}
}

func (p *AbuParser) Ecarule() (localctx IEcaruleContext) {
	this := p
	_ = this

	localctx = NewEcaruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AbuParserRULE_ecarule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(AbuParserRULE)
	}
	{
		p.SetState(157)
		p.Match(AbuParserID)
	}
	{
		p.SetState(158)
		p.Match(AbuParserON)
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(AbuParserTHIS-36))|(1<<(AbuParserEXT-36))|(1<<(AbuParserID-36)))) != 0) {
		{
			p.SetState(159)
			p.Resource()
		}

		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AbuParserFOR {
		{
			p.SetState(164)
			p.Task()
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITaskContext is an interface to support dynamic dispatch.
type ITaskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTaskContext differentiates from other interfaces.
	IsTaskContext()
}

type TaskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTaskContext() *TaskContext {
	var p = new(TaskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_task
	return p
}

func (*TaskContext) IsTaskContext() {}

func NewTaskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TaskContext {
	var p = new(TaskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_task

	return p
}

func (s *TaskContext) GetParser() antlr.Parser { return s.parser }

func (s *TaskContext) FOR() antlr.TerminalNode {
	return s.GetToken(AbuParserFOR, 0)
}

func (s *TaskContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TaskContext) DO() antlr.TerminalNode {
	return s.GetToken(AbuParserDO, 0)
}

func (s *TaskContext) Actions() IActionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionsContext)
}

func (s *TaskContext) ALL() antlr.TerminalNode {
	return s.GetToken(AbuParserALL, 0)
}

func (s *TaskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TaskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TaskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterTask(s)
	}
}

func (s *TaskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitTask(s)
	}
}

func (p *AbuParser) Task() (localctx ITaskContext) {
	this := p
	_ = this

	localctx = NewTaskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AbuParserRULE_task)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(AbuParserFOR)
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbuParserALL {
		{
			p.SetState(170)
			p.Match(AbuParserALL)
		}

	}
	{
		p.SetState(173)
		p.expression(0)
	}
	{
		p.SetState(174)
		p.Match(AbuParserDO)
	}
	{
		p.SetState(175)
		p.Actions()
	}

	return localctx
}

// IActionsContext is an interface to support dynamic dispatch.
type IActionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionsContext differentiates from other interfaces.
	IsActionsContext()
}

type ActionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionsContext() *ActionsContext {
	var p = new(ActionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_actions
	return p
}

func (*ActionsContext) IsActionsContext() {}

func NewActionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionsContext {
	var p = new(ActionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_actions

	return p
}

func (s *ActionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionsContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *ActionsContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ActionsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(AbuParserSEMICOLON)
}

func (s *ActionsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(AbuParserSEMICOLON, i)
}

func (s *ActionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterActions(s)
	}
}

func (s *ActionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitActions(s)
	}
}

func (p *AbuParser) Actions() (localctx IActionsContext) {
	this := p
	_ = this

	localctx = NewActionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AbuParserRULE_actions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Assignment()
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AbuParserSEMICOLON {
		{
			p.SetState(178)
			p.Match(AbuParserSEMICOLON)
		}
		{
			p.SetState(179)
			p.Assignment()
		}

		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Resource() IResourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResourceContext)
}

func (s *AssignmentContext) EQUALSIGN() antlr.TerminalNode {
	return s.GetToken(AbuParserEQUALSIGN, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *AbuParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AbuParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Resource()
	}
	{
		p.SetState(186)
		p.Match(AbuParserEQUALSIGN)
	}
	{
		p.SetState(187)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ABS() antlr.TerminalNode {
	return s.GetToken(AbuParserABS, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(AbuParserNOT, 0)
}

func (s *ExpressionContext) RL_BRACKET() antlr.TerminalNode {
	return s.GetToken(AbuParserRL_BRACKET, 0)
}

func (s *ExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(AbuParserRR_BRACKET, 0)
}

func (s *ExpressionContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpressionContext) MulDivModOperator() IMulDivModOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulDivModOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulDivModOperatorContext)
}

func (s *ExpressionContext) PlusMinusOperator() IPlusMinusOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPlusMinusOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPlusMinusOperatorContext)
}

func (s *ExpressionContext) ComparisonOperator() IComparisonOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *ExpressionContext) DOUBLECOLON() antlr.TerminalNode {
	return s.GetToken(AbuParserDOUBLECOLON, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(AbuParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(AbuParserOR, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AbuParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *AbuParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, AbuParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbuParserABS:
		{
			p.SetState(190)
			p.Match(AbuParserABS)
		}
		{
			p.SetState(191)
			p.expression(10)
		}

	case AbuParserNOT:
		{
			p.SetState(192)
			p.Match(AbuParserNOT)
		}
		{
			p.SetState(193)
			p.expression(9)
		}

	case AbuParserRL_BRACKET:
		{
			p.SetState(194)
			p.Match(AbuParserRL_BRACKET)
		}
		{
			p.SetState(195)
			p.expression(0)
		}
		{
			p.SetState(196)
			p.Match(AbuParserRR_BRACKET)
		}

	case AbuParserMINUS, AbuParserTHIS, AbuParserEXT, AbuParserTRUE, AbuParserFALSE, AbuParserID, AbuParserQUOTED_STRING, AbuParserDEC_LITERAL, AbuParserINT_LITERAL:
		{
			p.SetState(198)
			p.Term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AbuParserRULE_expression)
				p.SetState(201)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(202)
					p.MulDivModOperator()
				}
				{
					p.SetState(203)
					p.expression(9)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AbuParserRULE_expression)
				p.SetState(205)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(206)
					p.PlusMinusOperator()
				}
				{
					p.SetState(207)
					p.expression(8)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AbuParserRULE_expression)
				p.SetState(209)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(210)
					p.ComparisonOperator()
				}
				{
					p.SetState(211)
					p.expression(7)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AbuParserRULE_expression)
				p.SetState(213)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(214)
					p.Match(AbuParserDOUBLECOLON)
				}
				{
					p.SetState(215)
					p.expression(6)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AbuParserRULE_expression)
				p.SetState(216)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(217)
					p.Match(AbuParserAND)
				}
				{
					p.SetState(218)
					p.expression(5)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AbuParserRULE_expression)
				p.SetState(219)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(220)
					p.Match(AbuParserOR)
				}
				{
					p.SetState(221)
					p.expression(4)
				}

			}

		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IMulDivModOperatorContext is an interface to support dynamic dispatch.
type IMulDivModOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMulDivModOperatorContext differentiates from other interfaces.
	IsMulDivModOperatorContext()
}

type MulDivModOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulDivModOperatorContext() *MulDivModOperatorContext {
	var p = new(MulDivModOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_mulDivModOperator
	return p
}

func (*MulDivModOperatorContext) IsMulDivModOperatorContext() {}

func NewMulDivModOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulDivModOperatorContext {
	var p = new(MulDivModOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_mulDivModOperator

	return p
}

func (s *MulDivModOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MulDivModOperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(AbuParserMUL, 0)
}

func (s *MulDivModOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(AbuParserDIV, 0)
}

func (s *MulDivModOperatorContext) MOD() antlr.TerminalNode {
	return s.GetToken(AbuParserMOD, 0)
}

func (s *MulDivModOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulDivModOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterMulDivModOperator(s)
	}
}

func (s *MulDivModOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitMulDivModOperator(s)
	}
}

func (p *AbuParser) MulDivModOperator() (localctx IMulDivModOperatorContext) {
	this := p
	_ = this

	localctx = NewMulDivModOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AbuParserRULE_mulDivModOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AbuParserDIV)|(1<<AbuParserMUL)|(1<<AbuParserMOD))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPlusMinusOperatorContext is an interface to support dynamic dispatch.
type IPlusMinusOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlusMinusOperatorContext differentiates from other interfaces.
	IsPlusMinusOperatorContext()
}

type PlusMinusOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlusMinusOperatorContext() *PlusMinusOperatorContext {
	var p = new(PlusMinusOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_plusMinusOperator
	return p
}

func (*PlusMinusOperatorContext) IsPlusMinusOperatorContext() {}

func NewPlusMinusOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlusMinusOperatorContext {
	var p = new(PlusMinusOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_plusMinusOperator

	return p
}

func (s *PlusMinusOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *PlusMinusOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(AbuParserPLUS, 0)
}

func (s *PlusMinusOperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(AbuParserMINUS, 0)
}

func (s *PlusMinusOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusMinusOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlusMinusOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterPlusMinusOperator(s)
	}
}

func (s *PlusMinusOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitPlusMinusOperator(s)
	}
}

func (p *AbuParser) PlusMinusOperator() (localctx IPlusMinusOperatorContext) {
	this := p
	_ = this

	localctx = NewPlusMinusOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AbuParserRULE_plusMinusOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AbuParserPLUS || _la == AbuParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(AbuParserEQUALS, 0)
}

func (s *ComparisonOperatorContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(AbuParserNOTEQUALS, 0)
}

func (s *ComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(AbuParserLT, 0)
}

func (s *ComparisonOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(AbuParserLTE, 0)
}

func (s *ComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(AbuParserGT, 0)
}

func (s *ComparisonOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(AbuParserGTE, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (p *AbuParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	this := p
	_ = this

	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AbuParserRULE_comparisonOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AbuParserEQUALS)|(1<<AbuParserGT)|(1<<AbuParserLT)|(1<<AbuParserGTE)|(1<<AbuParserLTE)|(1<<AbuParserNOTEQUALS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *TermContext) Resource() IResourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResourceContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *AbuParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AbuParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(235)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbuParserMINUS, AbuParserTRUE, AbuParserFALSE, AbuParserQUOTED_STRING, AbuParserDEC_LITERAL, AbuParserINT_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(233)
			p.Value()
		}

	case AbuParserTHIS, AbuParserEXT, AbuParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)
			p.Resource()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) StringLiteral() IStringLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ValueContext) IntegerLiteral() IIntegerLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *ValueContext) DecimalLiteral() IDecimalLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecimalLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *ValueContext) BooleanLiteral() IBooleanLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *AbuParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AbuParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(237)
			p.StringLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.IntegerLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(239)
			p.DecimalLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(240)
			p.BooleanLiteral()
		}

	}

	return localctx
}

// IResourceContext is an interface to support dynamic dispatch.
type IResourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResourceContext differentiates from other interfaces.
	IsResourceContext()
}

type ResourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourceContext() *ResourceContext {
	var p = new(ResourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_resource
	return p
}

func (*ResourceContext) IsResourceContext() {}

func NewResourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceContext {
	var p = new(ResourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_resource

	return p
}

func (s *ResourceContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceContext) SimpleResource() ISimpleResourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleResourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleResourceContext)
}

func (s *ResourceContext) NestedResource() INestedResourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedResourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedResourceContext)
}

func (s *ResourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterResource(s)
	}
}

func (s *ResourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitResource(s)
	}
}

func (p *AbuParser) Resource() (localctx IResourceContext) {
	this := p
	_ = this

	localctx = NewResourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, AbuParserRULE_resource)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(243)
			p.SimpleResource()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)
			p.NestedResource()
		}

	}

	return localctx
}

// ISimpleResourceContext is an interface to support dynamic dispatch.
type ISimpleResourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleResourceContext differentiates from other interfaces.
	IsSimpleResourceContext()
}

type SimpleResourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleResourceContext() *SimpleResourceContext {
	var p = new(SimpleResourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_simpleResource
	return p
}

func (*SimpleResourceContext) IsSimpleResourceContext() {}

func NewSimpleResourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleResourceContext {
	var p = new(SimpleResourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_simpleResource

	return p
}

func (s *SimpleResourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleResourceContext) ID() antlr.TerminalNode {
	return s.GetToken(AbuParserID, 0)
}

func (s *SimpleResourceContext) THIS() antlr.TerminalNode {
	return s.GetToken(AbuParserTHIS, 0)
}

func (s *SimpleResourceContext) EXT() antlr.TerminalNode {
	return s.GetToken(AbuParserEXT, 0)
}

func (s *SimpleResourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleResourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleResourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterSimpleResource(s)
	}
}

func (s *SimpleResourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitSimpleResource(s)
	}
}

func (p *AbuParser) SimpleResource() (localctx ISimpleResourceContext) {
	this := p
	_ = this

	localctx = NewSimpleResourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, AbuParserRULE_simpleResource)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(253)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbuParserTHIS, AbuParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbuParserTHIS {
			{
				p.SetState(247)
				p.Match(AbuParserTHIS)
			}

		}
		{
			p.SetState(250)
			p.Match(AbuParserID)
		}

	case AbuParserEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)
			p.Match(AbuParserEXT)
		}
		{
			p.SetState(252)
			p.Match(AbuParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INestedResourceContext is an interface to support dynamic dispatch.
type INestedResourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNestedResourceContext differentiates from other interfaces.
	IsNestedResourceContext()
}

type NestedResourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedResourceContext() *NestedResourceContext {
	var p = new(NestedResourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_nestedResource
	return p
}

func (*NestedResourceContext) IsNestedResourceContext() {}

func NewNestedResourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedResourceContext {
	var p = new(NestedResourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_nestedResource

	return p
}

func (s *NestedResourceContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedResourceContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AbuParserID)
}

func (s *NestedResourceContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AbuParserID, i)
}

func (s *NestedResourceContext) SL_BRACKET() antlr.TerminalNode {
	return s.GetToken(AbuParserSL_BRACKET, 0)
}

func (s *NestedResourceContext) SR_BRACKET() antlr.TerminalNode {
	return s.GetToken(AbuParserSR_BRACKET, 0)
}

func (s *NestedResourceContext) THIS() antlr.TerminalNode {
	return s.GetToken(AbuParserTHIS, 0)
}

func (s *NestedResourceContext) EXT() antlr.TerminalNode {
	return s.GetToken(AbuParserEXT, 0)
}

func (s *NestedResourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedResourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedResourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterNestedResource(s)
	}
}

func (s *NestedResourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitNestedResource(s)
	}
}

func (p *AbuParser) NestedResource() (localctx INestedResourceContext) {
	this := p
	_ = this

	localctx = NewNestedResourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, AbuParserRULE_nestedResource)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(267)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbuParserTHIS, AbuParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbuParserTHIS {
			{
				p.SetState(255)
				p.Match(AbuParserTHIS)
			}

		}
		{
			p.SetState(258)
			p.Match(AbuParserID)
		}
		{
			p.SetState(259)
			p.Match(AbuParserSL_BRACKET)
		}
		{
			p.SetState(260)
			p.Match(AbuParserID)
		}
		{
			p.SetState(261)
			p.Match(AbuParserSR_BRACKET)
		}

	case AbuParserEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(262)
			p.Match(AbuParserEXT)
		}
		{
			p.SetState(263)
			p.Match(AbuParserID)
		}
		{
			p.SetState(264)
			p.Match(AbuParserSL_BRACKET)
		}
		{
			p.SetState(265)
			p.Match(AbuParserID)
		}
		{
			p.SetState(266)
			p.Match(AbuParserSR_BRACKET)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDecimalLiteralContext is an interface to support dynamic dispatch.
type IDecimalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalLiteralContext differentiates from other interfaces.
	IsDecimalLiteralContext()
}

type DecimalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalLiteralContext() *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_decimalLiteral
	return p
}

func (*DecimalLiteralContext) IsDecimalLiteralContext() {}

func NewDecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_decimalLiteral

	return p
}

func (s *DecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalLiteralContext) DEC_LITERAL() antlr.TerminalNode {
	return s.GetToken(AbuParserDEC_LITERAL, 0)
}

func (s *DecimalLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(AbuParserMINUS, 0)
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (p *AbuParser) DecimalLiteral() (localctx IDecimalLiteralContext) {
	this := p
	_ = this

	localctx = NewDecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AbuParserRULE_decimalLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbuParserMINUS {
		{
			p.SetState(269)
			p.Match(AbuParserMINUS)
		}

	}
	{
		p.SetState(272)
		p.Match(AbuParserDEC_LITERAL)
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(AbuParserINT_LITERAL, 0)
}

func (s *IntegerLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(AbuParserMINUS, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (p *AbuParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	this := p
	_ = this

	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AbuParserRULE_integerLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbuParserMINUS {
		{
			p.SetState(274)
			p.Match(AbuParserMINUS)
		}

	}
	{
		p.SetState(277)
		p.Match(AbuParserINT_LITERAL)
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(AbuParserQUOTED_STRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (p *AbuParser) StringLiteral() (localctx IStringLiteralContext) {
	this := p
	_ = this

	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AbuParserRULE_stringLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(AbuParserQUOTED_STRING)
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbuParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbuParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(AbuParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(AbuParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbuParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (p *AbuParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	this := p
	_ = this

	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AbuParserRULE_booleanLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AbuParserTRUE || _la == AbuParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *AbuParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AbuParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
