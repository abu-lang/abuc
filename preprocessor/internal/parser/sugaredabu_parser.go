// Code generated from SugaredAbuParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SugaredAbuParser
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

type SugaredAbuParser struct {
	*antlr.BaseParser
}

var sugaredabuparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sugaredabuparserParserInit() {
	staticData := &sugaredabuparserParserStaticData
	staticData.literalNames = []string{
		"", "'and'", "'or'", "'not'", "'absint'", "'absdec'", "'.'", "'+'",
		"'-'", "'/'", "'*'", "'%'", "'=='", "'>'", "'<'", "'>='", "'<='", "'!='",
		"'::'", "':'", "';'", "':='", "'='", "','", "'('", "')'", "'['", "']'",
		"'{'", "'}'", "'has'", "'where'", "'physical'", "'logical'", "'input'",
		"'output'", "'rule'", "", "", "'on'", "'let'", "'in'", "'default'",
		"'for'", "'all'", "'do'", "'owise'", "'define'", "'as'", "'boolean'",
		"'integer'", "'decimal'", "'string'", "'true'", "'false'",
	}
	staticData.symbolicNames = []string{
		"", "AND", "OR", "NOT", "ABSINT", "ABSDEC", "DOT", "PLUS", "MINUS",
		"DIV", "MUL", "MOD", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS",
		"DOUBLECOLON", "COLON", "SEMICOLON", "COLONEQUAL", "EQUALSIGN", "COMMA",
		"RL_BRACKET", "RR_BRACKET", "SL_BRACKET", "SR_BRACKET", "CL_BRACKET",
		"CR_BRACKET", "HAS", "WHERE", "PHYSICAL", "LOGICAL", "INPUT", "OUTPUT",
		"RULE", "THIS", "EXT", "ON", "LET", "IN", "DEFAULT", "FOR", "ALL", "DO",
		"OWISE", "DEFINE", "AS", "BOOLEAN", "INTEGER", "DECIMAL", "STRING",
		"TRUE", "FALSE", "ID", "QUOTED_STRING", "DEC_LITERAL", "INT_LITERAL",
		"WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"program", "typeDecl", "resField", "device", "resList", "resDecl", "compResDecl",
		"type", "ecarule", "letDecl", "task", "actions", "assignment", "expression",
		"mulDivModOperator", "plusMinusOperator", "comparisonOperator", "term",
		"value", "resource", "simpleResource", "nestedResource", "decimalLiteral",
		"integerLiteral", "stringLiteral", "booleanLiteral",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 61, 346, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		5, 0, 54, 8, 0, 10, 0, 12, 0, 57, 9, 0, 1, 0, 4, 0, 60, 8, 0, 11, 0, 12,
		0, 61, 1, 0, 5, 0, 65, 8, 0, 10, 0, 12, 0, 68, 9, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 4, 1, 75, 8, 1, 11, 1, 12, 1, 76, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 95, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 104, 8, 3,
		1, 3, 1, 3, 1, 3, 4, 3, 109, 8, 3, 11, 3, 12, 3, 110, 3, 3, 113, 8, 3,
		1, 4, 4, 4, 116, 8, 4, 11, 4, 12, 4, 117, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 139, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 152, 8, 6, 10, 6, 12, 6, 155, 9, 6, 1, 6,
		1, 6, 3, 6, 159, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 4, 8, 167, 8,
		8, 11, 8, 12, 8, 168, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 175, 8, 8, 1, 8, 4,
		8, 178, 8, 8, 11, 8, 12, 8, 179, 1, 8, 1, 8, 1, 8, 1, 8, 4, 8, 186, 8,
		8, 11, 8, 12, 8, 187, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 196, 8,
		8, 1, 8, 5, 8, 199, 8, 8, 10, 8, 12, 8, 202, 9, 8, 3, 8, 204, 8, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 213, 8, 9, 10, 9, 12, 9, 216,
		9, 9, 1, 10, 1, 10, 3, 10, 220, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 3, 10, 228, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3,
		10, 236, 8, 10, 1, 11, 1, 11, 1, 11, 5, 11, 241, 8, 11, 10, 11, 12, 11,
		244, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 262, 8, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		5, 13, 285, 8, 13, 10, 13, 12, 13, 288, 9, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 17, 1, 17, 3, 17, 298, 8, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 3, 18, 304, 8, 18, 1, 19, 1, 19, 3, 19, 308, 8, 19, 1, 20, 3, 20,
		311, 8, 20, 1, 20, 1, 20, 1, 20, 3, 20, 316, 8, 20, 1, 21, 3, 21, 319,
		8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3,
		21, 330, 8, 21, 1, 22, 3, 22, 333, 8, 22, 1, 22, 1, 22, 1, 23, 3, 23, 338,
		8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 0, 1, 26, 26, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 0, 5, 1, 0, 49, 52, 1, 0, 9, 11, 1, 0, 7, 8, 1,
		0, 12, 17, 1, 0, 53, 54, 367, 0, 55, 1, 0, 0, 0, 2, 69, 1, 0, 0, 0, 4,
		94, 1, 0, 0, 0, 6, 96, 1, 0, 0, 0, 8, 115, 1, 0, 0, 0, 10, 138, 1, 0, 0,
		0, 12, 140, 1, 0, 0, 0, 14, 160, 1, 0, 0, 0, 16, 203, 1, 0, 0, 0, 18, 205,
		1, 0, 0, 0, 20, 235, 1, 0, 0, 0, 22, 237, 1, 0, 0, 0, 24, 245, 1, 0, 0,
		0, 26, 261, 1, 0, 0, 0, 28, 289, 1, 0, 0, 0, 30, 291, 1, 0, 0, 0, 32, 293,
		1, 0, 0, 0, 34, 297, 1, 0, 0, 0, 36, 303, 1, 0, 0, 0, 38, 307, 1, 0, 0,
		0, 40, 315, 1, 0, 0, 0, 42, 329, 1, 0, 0, 0, 44, 332, 1, 0, 0, 0, 46, 337,
		1, 0, 0, 0, 48, 341, 1, 0, 0, 0, 50, 343, 1, 0, 0, 0, 52, 54, 3, 2, 1,
		0, 53, 52, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56,
		1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 60, 3, 6, 3, 0,
		59, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1,
		0, 0, 0, 62, 66, 1, 0, 0, 0, 63, 65, 3, 16, 8, 0, 64, 63, 1, 0, 0, 0, 65,
		68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 1, 1, 0, 0,
		0, 68, 66, 1, 0, 0, 0, 69, 70, 5, 47, 0, 0, 70, 71, 5, 55, 0, 0, 71, 72,
		5, 48, 0, 0, 72, 74, 5, 28, 0, 0, 73, 75, 3, 4, 2, 0, 74, 73, 1, 0, 0,
		0, 75, 76, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78,
		1, 0, 0, 0, 78, 79, 5, 29, 0, 0, 79, 3, 1, 0, 0, 0, 80, 81, 5, 55, 0, 0,
		81, 82, 5, 19, 0, 0, 82, 83, 5, 32, 0, 0, 83, 84, 5, 35, 0, 0, 84, 95,
		3, 14, 7, 0, 85, 86, 5, 55, 0, 0, 86, 87, 5, 19, 0, 0, 87, 88, 5, 32, 0,
		0, 88, 89, 5, 34, 0, 0, 89, 95, 3, 14, 7, 0, 90, 91, 5, 55, 0, 0, 91, 92,
		5, 19, 0, 0, 92, 93, 5, 33, 0, 0, 93, 95, 3, 14, 7, 0, 94, 80, 1, 0, 0,
		0, 94, 85, 1, 0, 0, 0, 94, 90, 1, 0, 0, 0, 95, 5, 1, 0, 0, 0, 96, 97, 5,
		55, 0, 0, 97, 98, 5, 19, 0, 0, 98, 99, 3, 48, 24, 0, 99, 100, 5, 28, 0,
		0, 100, 103, 3, 8, 4, 0, 101, 102, 5, 31, 0, 0, 102, 104, 3, 26, 13, 0,
		103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105,
		112, 5, 29, 0, 0, 106, 108, 5, 30, 0, 0, 107, 109, 5, 55, 0, 0, 108, 107,
		1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0,
		0, 0, 111, 113, 1, 0, 0, 0, 112, 106, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0,
		113, 7, 1, 0, 0, 0, 114, 116, 3, 10, 5, 0, 115, 114, 1, 0, 0, 0, 116, 117,
		1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 9, 1, 0, 0,
		0, 119, 120, 5, 32, 0, 0, 120, 121, 5, 35, 0, 0, 121, 122, 3, 14, 7, 0,
		122, 123, 5, 55, 0, 0, 123, 124, 5, 22, 0, 0, 124, 125, 3, 26, 13, 0, 125,
		139, 1, 0, 0, 0, 126, 127, 5, 32, 0, 0, 127, 128, 5, 34, 0, 0, 128, 129,
		3, 14, 7, 0, 129, 130, 5, 55, 0, 0, 130, 139, 1, 0, 0, 0, 131, 132, 5,
		33, 0, 0, 132, 133, 3, 14, 7, 0, 133, 134, 5, 55, 0, 0, 134, 135, 5, 22,
		0, 0, 135, 136, 3, 26, 13, 0, 136, 139, 1, 0, 0, 0, 137, 139, 3, 12, 6,
		0, 138, 119, 1, 0, 0, 0, 138, 126, 1, 0, 0, 0, 138, 131, 1, 0, 0, 0, 138,
		137, 1, 0, 0, 0, 139, 11, 1, 0, 0, 0, 140, 141, 5, 55, 0, 0, 141, 158,
		5, 55, 0, 0, 142, 143, 5, 22, 0, 0, 143, 144, 5, 24, 0, 0, 144, 145, 5,
		55, 0, 0, 145, 146, 5, 22, 0, 0, 146, 153, 3, 26, 13, 0, 147, 148, 5, 23,
		0, 0, 148, 149, 5, 55, 0, 0, 149, 150, 5, 22, 0, 0, 150, 152, 3, 26, 13,
		0, 151, 147, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153,
		154, 1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157,
		5, 25, 0, 0, 157, 159, 1, 0, 0, 0, 158, 142, 1, 0, 0, 0, 158, 159, 1, 0,
		0, 0, 159, 13, 1, 0, 0, 0, 160, 161, 7, 0, 0, 0, 161, 15, 1, 0, 0, 0, 162,
		163, 5, 36, 0, 0, 163, 164, 5, 55, 0, 0, 164, 166, 5, 39, 0, 0, 165, 167,
		3, 38, 19, 0, 166, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 166, 1,
		0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 174, 1, 0, 0, 0, 170, 171, 5, 40, 0,
		0, 171, 172, 3, 18, 9, 0, 172, 173, 5, 41, 0, 0, 173, 175, 1, 0, 0, 0,
		174, 170, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 177, 1, 0, 0, 0, 176,
		178, 3, 20, 10, 0, 177, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 177,
		1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 204, 1, 0, 0, 0, 181, 182, 5, 36,
		0, 0, 182, 183, 5, 55, 0, 0, 183, 185, 5, 39, 0, 0, 184, 186, 3, 38, 19,
		0, 185, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187,
		188, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 190, 5, 42, 0, 0, 190, 195,
		3, 22, 11, 0, 191, 192, 5, 40, 0, 0, 192, 193, 3, 18, 9, 0, 193, 194, 5,
		41, 0, 0, 194, 196, 1, 0, 0, 0, 195, 191, 1, 0, 0, 0, 195, 196, 1, 0, 0,
		0, 196, 200, 1, 0, 0, 0, 197, 199, 3, 20, 10, 0, 198, 197, 1, 0, 0, 0,
		199, 202, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201,
		204, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 203, 162, 1, 0, 0, 0, 203, 181,
		1, 0, 0, 0, 204, 17, 1, 0, 0, 0, 205, 206, 5, 55, 0, 0, 206, 207, 5, 21,
		0, 0, 207, 214, 3, 26, 13, 0, 208, 209, 5, 20, 0, 0, 209, 210, 5, 55, 0,
		0, 210, 211, 5, 21, 0, 0, 211, 213, 3, 26, 13, 0, 212, 208, 1, 0, 0, 0,
		213, 216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215,
		19, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 217, 219, 5, 43, 0, 0, 218, 220,
		5, 44, 0, 0, 219, 218, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 221, 1, 0,
		0, 0, 221, 222, 3, 26, 13, 0, 222, 223, 5, 45, 0, 0, 223, 224, 3, 22, 11,
		0, 224, 236, 1, 0, 0, 0, 225, 227, 5, 43, 0, 0, 226, 228, 5, 44, 0, 0,
		227, 226, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229,
		230, 3, 26, 13, 0, 230, 231, 5, 45, 0, 0, 231, 232, 3, 22, 11, 0, 232,
		233, 5, 46, 0, 0, 233, 234, 3, 22, 11, 0, 234, 236, 1, 0, 0, 0, 235, 217,
		1, 0, 0, 0, 235, 225, 1, 0, 0, 0, 236, 21, 1, 0, 0, 0, 237, 242, 3, 24,
		12, 0, 238, 239, 5, 23, 0, 0, 239, 241, 3, 24, 12, 0, 240, 238, 1, 0, 0,
		0, 241, 244, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243,
		23, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 245, 246, 3, 38, 19, 0, 246, 247,
		5, 22, 0, 0, 247, 248, 3, 26, 13, 0, 248, 25, 1, 0, 0, 0, 249, 250, 6,
		13, -1, 0, 250, 251, 5, 4, 0, 0, 251, 262, 3, 26, 13, 11, 252, 253, 5,
		5, 0, 0, 253, 262, 3, 26, 13, 10, 254, 255, 5, 3, 0, 0, 255, 262, 3, 26,
		13, 9, 256, 257, 5, 24, 0, 0, 257, 258, 3, 26, 13, 0, 258, 259, 5, 25,
		0, 0, 259, 262, 1, 0, 0, 0, 260, 262, 3, 34, 17, 0, 261, 249, 1, 0, 0,
		0, 261, 252, 1, 0, 0, 0, 261, 254, 1, 0, 0, 0, 261, 256, 1, 0, 0, 0, 261,
		260, 1, 0, 0, 0, 262, 286, 1, 0, 0, 0, 263, 264, 10, 8, 0, 0, 264, 265,
		3, 28, 14, 0, 265, 266, 3, 26, 13, 9, 266, 285, 1, 0, 0, 0, 267, 268, 10,
		7, 0, 0, 268, 269, 3, 30, 15, 0, 269, 270, 3, 26, 13, 8, 270, 285, 1, 0,
		0, 0, 271, 272, 10, 6, 0, 0, 272, 273, 3, 32, 16, 0, 273, 274, 3, 26, 13,
		7, 274, 285, 1, 0, 0, 0, 275, 276, 10, 5, 0, 0, 276, 277, 5, 18, 0, 0,
		277, 285, 3, 26, 13, 6, 278, 279, 10, 4, 0, 0, 279, 280, 5, 1, 0, 0, 280,
		285, 3, 26, 13, 5, 281, 282, 10, 3, 0, 0, 282, 283, 5, 2, 0, 0, 283, 285,
		3, 26, 13, 4, 284, 263, 1, 0, 0, 0, 284, 267, 1, 0, 0, 0, 284, 271, 1,
		0, 0, 0, 284, 275, 1, 0, 0, 0, 284, 278, 1, 0, 0, 0, 284, 281, 1, 0, 0,
		0, 285, 288, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287,
		27, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 289, 290, 7, 1, 0, 0, 290, 29, 1,
		0, 0, 0, 291, 292, 7, 2, 0, 0, 292, 31, 1, 0, 0, 0, 293, 294, 7, 3, 0,
		0, 294, 33, 1, 0, 0, 0, 295, 298, 3, 36, 18, 0, 296, 298, 3, 38, 19, 0,
		297, 295, 1, 0, 0, 0, 297, 296, 1, 0, 0, 0, 298, 35, 1, 0, 0, 0, 299, 304,
		3, 48, 24, 0, 300, 304, 3, 46, 23, 0, 301, 304, 3, 44, 22, 0, 302, 304,
		3, 50, 25, 0, 303, 299, 1, 0, 0, 0, 303, 300, 1, 0, 0, 0, 303, 301, 1,
		0, 0, 0, 303, 302, 1, 0, 0, 0, 304, 37, 1, 0, 0, 0, 305, 308, 3, 40, 20,
		0, 306, 308, 3, 42, 21, 0, 307, 305, 1, 0, 0, 0, 307, 306, 1, 0, 0, 0,
		308, 39, 1, 0, 0, 0, 309, 311, 5, 37, 0, 0, 310, 309, 1, 0, 0, 0, 310,
		311, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 316, 5, 55, 0, 0, 313, 314,
		5, 38, 0, 0, 314, 316, 5, 55, 0, 0, 315, 310, 1, 0, 0, 0, 315, 313, 1,
		0, 0, 0, 316, 41, 1, 0, 0, 0, 317, 319, 5, 37, 0, 0, 318, 317, 1, 0, 0,
		0, 318, 319, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 321, 5, 55, 0, 0, 321,
		322, 5, 26, 0, 0, 322, 323, 5, 55, 0, 0, 323, 330, 5, 27, 0, 0, 324, 325,
		5, 38, 0, 0, 325, 326, 5, 55, 0, 0, 326, 327, 5, 26, 0, 0, 327, 328, 5,
		55, 0, 0, 328, 330, 5, 27, 0, 0, 329, 318, 1, 0, 0, 0, 329, 324, 1, 0,
		0, 0, 330, 43, 1, 0, 0, 0, 331, 333, 5, 8, 0, 0, 332, 331, 1, 0, 0, 0,
		332, 333, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 335, 5, 57, 0, 0, 335,
		45, 1, 0, 0, 0, 336, 338, 5, 8, 0, 0, 337, 336, 1, 0, 0, 0, 337, 338, 1,
		0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 340, 5, 58, 0, 0, 340, 47, 1, 0, 0,
		0, 341, 342, 5, 56, 0, 0, 342, 49, 1, 0, 0, 0, 343, 344, 7, 4, 0, 0, 344,
		51, 1, 0, 0, 0, 36, 55, 61, 66, 76, 94, 103, 110, 112, 117, 138, 153, 158,
		168, 174, 179, 187, 195, 200, 203, 214, 219, 227, 235, 242, 261, 284, 286,
		297, 303, 307, 310, 315, 318, 329, 332, 337,
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

// SugaredAbuParserInit initializes any static state used to implement SugaredAbuParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSugaredAbuParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SugaredAbuParserInit() {
	staticData := &sugaredabuparserParserStaticData
	staticData.once.Do(sugaredabuparserParserInit)
}

// NewSugaredAbuParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSugaredAbuParser(input antlr.TokenStream) *SugaredAbuParser {
	SugaredAbuParserInit()
	this := new(SugaredAbuParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sugaredabuparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SugaredAbuParser.g4"

	return this
}

// SugaredAbuParser tokens.
const (
	SugaredAbuParserEOF           = antlr.TokenEOF
	SugaredAbuParserAND           = 1
	SugaredAbuParserOR            = 2
	SugaredAbuParserNOT           = 3
	SugaredAbuParserABSINT        = 4
	SugaredAbuParserABSDEC        = 5
	SugaredAbuParserDOT           = 6
	SugaredAbuParserPLUS          = 7
	SugaredAbuParserMINUS         = 8
	SugaredAbuParserDIV           = 9
	SugaredAbuParserMUL           = 10
	SugaredAbuParserMOD           = 11
	SugaredAbuParserEQUALS        = 12
	SugaredAbuParserGT            = 13
	SugaredAbuParserLT            = 14
	SugaredAbuParserGTE           = 15
	SugaredAbuParserLTE           = 16
	SugaredAbuParserNOTEQUALS     = 17
	SugaredAbuParserDOUBLECOLON   = 18
	SugaredAbuParserCOLON         = 19
	SugaredAbuParserSEMICOLON     = 20
	SugaredAbuParserCOLONEQUAL    = 21
	SugaredAbuParserEQUALSIGN     = 22
	SugaredAbuParserCOMMA         = 23
	SugaredAbuParserRL_BRACKET    = 24
	SugaredAbuParserRR_BRACKET    = 25
	SugaredAbuParserSL_BRACKET    = 26
	SugaredAbuParserSR_BRACKET    = 27
	SugaredAbuParserCL_BRACKET    = 28
	SugaredAbuParserCR_BRACKET    = 29
	SugaredAbuParserHAS           = 30
	SugaredAbuParserWHERE         = 31
	SugaredAbuParserPHYSICAL      = 32
	SugaredAbuParserLOGICAL       = 33
	SugaredAbuParserINPUT         = 34
	SugaredAbuParserOUTPUT        = 35
	SugaredAbuParserRULE          = 36
	SugaredAbuParserTHIS          = 37
	SugaredAbuParserEXT           = 38
	SugaredAbuParserON            = 39
	SugaredAbuParserLET           = 40
	SugaredAbuParserIN            = 41
	SugaredAbuParserDEFAULT       = 42
	SugaredAbuParserFOR           = 43
	SugaredAbuParserALL           = 44
	SugaredAbuParserDO            = 45
	SugaredAbuParserOWISE         = 46
	SugaredAbuParserDEFINE        = 47
	SugaredAbuParserAS            = 48
	SugaredAbuParserBOOLEAN       = 49
	SugaredAbuParserINTEGER       = 50
	SugaredAbuParserDECIMAL       = 51
	SugaredAbuParserSTRING        = 52
	SugaredAbuParserTRUE          = 53
	SugaredAbuParserFALSE         = 54
	SugaredAbuParserID            = 55
	SugaredAbuParserQUOTED_STRING = 56
	SugaredAbuParserDEC_LITERAL   = 57
	SugaredAbuParserINT_LITERAL   = 58
	SugaredAbuParserWS            = 59
	SugaredAbuParserCOMMENT       = 60
	SugaredAbuParserLINE_COMMENT  = 61
)

// SugaredAbuParser rules.
const (
	SugaredAbuParserRULE_program            = 0
	SugaredAbuParserRULE_typeDecl           = 1
	SugaredAbuParserRULE_resField           = 2
	SugaredAbuParserRULE_device             = 3
	SugaredAbuParserRULE_resList            = 4
	SugaredAbuParserRULE_resDecl            = 5
	SugaredAbuParserRULE_compResDecl        = 6
	SugaredAbuParserRULE_type               = 7
	SugaredAbuParserRULE_ecarule            = 8
	SugaredAbuParserRULE_letDecl            = 9
	SugaredAbuParserRULE_task               = 10
	SugaredAbuParserRULE_actions            = 11
	SugaredAbuParserRULE_assignment         = 12
	SugaredAbuParserRULE_expression         = 13
	SugaredAbuParserRULE_mulDivModOperator  = 14
	SugaredAbuParserRULE_plusMinusOperator  = 15
	SugaredAbuParserRULE_comparisonOperator = 16
	SugaredAbuParserRULE_term               = 17
	SugaredAbuParserRULE_value              = 18
	SugaredAbuParserRULE_resource           = 19
	SugaredAbuParserRULE_simpleResource     = 20
	SugaredAbuParserRULE_nestedResource     = 21
	SugaredAbuParserRULE_decimalLiteral     = 22
	SugaredAbuParserRULE_integerLiteral     = 23
	SugaredAbuParserRULE_stringLiteral      = 24
	SugaredAbuParserRULE_booleanLiteral     = 25
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
	p.RuleIndex = SugaredAbuParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

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

func (s *ProgramContext) AllDevice() []IDeviceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeviceContext); ok {
			len++
		}
	}

	tst := make([]IDeviceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeviceContext); ok {
			tst[i] = t.(IDeviceContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Device(i int) IDeviceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeviceContext); ok {
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

	return t.(IDeviceContext)
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
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *SugaredAbuParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SugaredAbuParserRULE_program)
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
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SugaredAbuParserDEFINE {
		{
			p.SetState(52)
			p.TypeDecl()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SugaredAbuParserID {
		{
			p.SetState(58)
			p.Device()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SugaredAbuParserRULE {
		{
			p.SetState(63)
			p.Ecarule()
		}

		p.SetState(68)
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
	p.RuleIndex = SugaredAbuParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserDEFINE, 0)
}

func (s *TypeDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserID, 0)
}

func (s *TypeDeclContext) AS() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserAS, 0)
}

func (s *TypeDeclContext) CL_BRACKET() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserCL_BRACKET, 0)
}

func (s *TypeDeclContext) CR_BRACKET() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserCR_BRACKET, 0)
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
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (p *SugaredAbuParser) TypeDecl() (localctx ITypeDeclContext) {
	this := p
	_ = this

	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SugaredAbuParserRULE_typeDecl)
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
		p.SetState(69)
		p.Match(SugaredAbuParserDEFINE)
	}
	{
		p.SetState(70)
		p.Match(SugaredAbuParserID)
	}
	{
		p.SetState(71)
		p.Match(SugaredAbuParserAS)
	}
	{
		p.SetState(72)
		p.Match(SugaredAbuParserCL_BRACKET)
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SugaredAbuParserID {
		{
			p.SetState(73)
			p.ResField()
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
		p.Match(SugaredAbuParserCR_BRACKET)
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
	p.RuleIndex = SugaredAbuParserRULE_resField
	return p
}

func (*ResFieldContext) IsResFieldContext() {}

func NewResFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResFieldContext {
	var p = new(ResFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_resField

	return p
}

func (s *ResFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ResFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserID, 0)
}

func (s *ResFieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserCOLON, 0)
}

func (s *ResFieldContext) PHYSICAL() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserPHYSICAL, 0)
}

func (s *ResFieldContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserOUTPUT, 0)
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
	return s.GetToken(SugaredAbuParserINPUT, 0)
}

func (s *ResFieldContext) LOGICAL() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserLOGICAL, 0)
}

func (s *ResFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterResField(s)
	}
}

func (s *ResFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitResField(s)
	}
}

func (p *SugaredAbuParser) ResField() (localctx IResFieldContext) {
	this := p
	_ = this

	localctx = NewResFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SugaredAbuParserRULE_resField)

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

	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(81)
			p.Match(SugaredAbuParserCOLON)
		}
		{
			p.SetState(82)
			p.Match(SugaredAbuParserPHYSICAL)
		}
		{
			p.SetState(83)
			p.Match(SugaredAbuParserOUTPUT)
		}
		{
			p.SetState(84)
			p.Type()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(86)
			p.Match(SugaredAbuParserCOLON)
		}
		{
			p.SetState(87)
			p.Match(SugaredAbuParserPHYSICAL)
		}
		{
			p.SetState(88)
			p.Match(SugaredAbuParserINPUT)
		}
		{
			p.SetState(89)
			p.Type()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(91)
			p.Match(SugaredAbuParserCOLON)
		}
		{
			p.SetState(92)
			p.Match(SugaredAbuParserLOGICAL)
		}
		{
			p.SetState(93)
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
	p.RuleIndex = SugaredAbuParserRULE_device
	return p
}

func (*DeviceContext) IsDeviceContext() {}

func NewDeviceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeviceContext {
	var p = new(DeviceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_device

	return p
}

func (s *DeviceContext) GetParser() antlr.Parser { return s.parser }

func (s *DeviceContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SugaredAbuParserID)
}

func (s *DeviceContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserID, i)
}

func (s *DeviceContext) COLON() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserCOLON, 0)
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
	return s.GetToken(SugaredAbuParserCL_BRACKET, 0)
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
	return s.GetToken(SugaredAbuParserCR_BRACKET, 0)
}

func (s *DeviceContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserWHERE, 0)
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
	return s.GetToken(SugaredAbuParserHAS, 0)
}

func (s *DeviceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeviceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeviceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterDevice(s)
	}
}

func (s *DeviceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitDevice(s)
	}
}

func (p *SugaredAbuParser) Device() (localctx IDeviceContext) {
	this := p
	_ = this

	localctx = NewDeviceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SugaredAbuParserRULE_device)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(SugaredAbuParserID)
	}
	{
		p.SetState(97)
		p.Match(SugaredAbuParserCOLON)
	}
	{
		p.SetState(98)
		p.StringLiteral()
	}
	{
		p.SetState(99)
		p.Match(SugaredAbuParserCL_BRACKET)
	}
	{
		p.SetState(100)
		p.ResList()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SugaredAbuParserWHERE {
		{
			p.SetState(101)
			p.Match(SugaredAbuParserWHERE)
		}
		{
			p.SetState(102)
			p.expression(0)
		}

	}
	{
		p.SetState(105)
		p.Match(SugaredAbuParserCR_BRACKET)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SugaredAbuParserHAS {
		{
			p.SetState(106)
			p.Match(SugaredAbuParserHAS)
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(107)
					p.Match(SugaredAbuParserID)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	p.RuleIndex = SugaredAbuParserRULE_resList
	return p
}

func (*ResListContext) IsResListContext() {}

func NewResListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResListContext {
	var p = new(ResListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_resList

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
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterResList(s)
	}
}

func (s *ResListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitResList(s)
	}
}

func (p *SugaredAbuParser) ResList() (localctx IResListContext) {
	this := p
	_ = this

	localctx = NewResListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SugaredAbuParserRULE_resList)
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
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(SugaredAbuParserPHYSICAL-32))|(1<<(SugaredAbuParserLOGICAL-32))|(1<<(SugaredAbuParserID-32)))) != 0) {
		{
			p.SetState(114)
			p.ResDecl()
		}

		p.SetState(117)
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
	p.RuleIndex = SugaredAbuParserRULE_resDecl
	return p
}

func (*ResDeclContext) IsResDeclContext() {}

func NewResDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResDeclContext {
	var p = new(ResDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_resDecl

	return p
}

func (s *ResDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ResDeclContext) PHYSICAL() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserPHYSICAL, 0)
}

func (s *ResDeclContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserOUTPUT, 0)
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
	return s.GetToken(SugaredAbuParserID, 0)
}

func (s *ResDeclContext) EQUALSIGN() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserEQUALSIGN, 0)
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
	return s.GetToken(SugaredAbuParserINPUT, 0)
}

func (s *ResDeclContext) LOGICAL() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserLOGICAL, 0)
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
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterResDecl(s)
	}
}

func (s *ResDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitResDecl(s)
	}
}

func (p *SugaredAbuParser) ResDecl() (localctx IResDeclContext) {
	this := p
	_ = this

	localctx = NewResDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SugaredAbuParserRULE_resDecl)

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

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(SugaredAbuParserPHYSICAL)
		}
		{
			p.SetState(120)
			p.Match(SugaredAbuParserOUTPUT)
		}
		{
			p.SetState(121)
			p.Type()
		}
		{
			p.SetState(122)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(123)
			p.Match(SugaredAbuParserEQUALSIGN)
		}
		{
			p.SetState(124)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.Match(SugaredAbuParserPHYSICAL)
		}
		{
			p.SetState(127)
			p.Match(SugaredAbuParserINPUT)
		}
		{
			p.SetState(128)
			p.Type()
		}
		{
			p.SetState(129)
			p.Match(SugaredAbuParserID)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
			p.Match(SugaredAbuParserLOGICAL)
		}
		{
			p.SetState(132)
			p.Type()
		}
		{
			p.SetState(133)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(134)
			p.Match(SugaredAbuParserEQUALSIGN)
		}
		{
			p.SetState(135)
			p.expression(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(137)
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
	p.RuleIndex = SugaredAbuParserRULE_compResDecl
	return p
}

func (*CompResDeclContext) IsCompResDeclContext() {}

func NewCompResDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompResDeclContext {
	var p = new(CompResDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_compResDecl

	return p
}

func (s *CompResDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *CompResDeclContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SugaredAbuParserID)
}

func (s *CompResDeclContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserID, i)
}

func (s *CompResDeclContext) AllEQUALSIGN() []antlr.TerminalNode {
	return s.GetTokens(SugaredAbuParserEQUALSIGN)
}

func (s *CompResDeclContext) EQUALSIGN(i int) antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserEQUALSIGN, i)
}

func (s *CompResDeclContext) RL_BRACKET() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserRL_BRACKET, 0)
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
	return s.GetToken(SugaredAbuParserRR_BRACKET, 0)
}

func (s *CompResDeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SugaredAbuParserCOMMA)
}

func (s *CompResDeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserCOMMA, i)
}

func (s *CompResDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompResDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompResDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterCompResDecl(s)
	}
}

func (s *CompResDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitCompResDecl(s)
	}
}

func (p *SugaredAbuParser) CompResDecl() (localctx ICompResDeclContext) {
	this := p
	_ = this

	localctx = NewCompResDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SugaredAbuParserRULE_compResDecl)
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
		p.SetState(140)
		p.Match(SugaredAbuParserID)
	}
	{
		p.SetState(141)
		p.Match(SugaredAbuParserID)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SugaredAbuParserEQUALSIGN {
		{
			p.SetState(142)
			p.Match(SugaredAbuParserEQUALSIGN)
		}
		{
			p.SetState(143)
			p.Match(SugaredAbuParserRL_BRACKET)
		}
		{
			p.SetState(144)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(145)
			p.Match(SugaredAbuParserEQUALSIGN)
		}
		{
			p.SetState(146)
			p.expression(0)
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SugaredAbuParserCOMMA {
			{
				p.SetState(147)
				p.Match(SugaredAbuParserCOMMA)
			}
			{
				p.SetState(148)
				p.Match(SugaredAbuParserID)
			}
			{
				p.SetState(149)
				p.Match(SugaredAbuParserEQUALSIGN)
			}
			{
				p.SetState(150)
				p.expression(0)
			}

			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(156)
			p.Match(SugaredAbuParserRR_BRACKET)
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
	p.RuleIndex = SugaredAbuParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserBOOLEAN, 0)
}

func (s *TypeContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserINTEGER, 0)
}

func (s *TypeContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserDECIMAL, 0)
}

func (s *TypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserSTRING, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *SugaredAbuParser) Type() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SugaredAbuParserRULE_type)
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
		p.SetState(160)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-49)&-(0x1f+1)) == 0 && ((1<<uint((_la-49)))&((1<<(SugaredAbuParserBOOLEAN-49))|(1<<(SugaredAbuParserINTEGER-49))|(1<<(SugaredAbuParserDECIMAL-49))|(1<<(SugaredAbuParserSTRING-49)))) != 0) {
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
	p.RuleIndex = SugaredAbuParserRULE_ecarule
	return p
}

func (*EcaruleContext) IsEcaruleContext() {}

func NewEcaruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EcaruleContext {
	var p = new(EcaruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_ecarule

	return p
}

func (s *EcaruleContext) GetParser() antlr.Parser { return s.parser }

func (s *EcaruleContext) RULE() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserRULE, 0)
}

func (s *EcaruleContext) ID() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserID, 0)
}

func (s *EcaruleContext) ON() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserON, 0)
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

func (s *EcaruleContext) LET() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserLET, 0)
}

func (s *EcaruleContext) LetDecl() ILetDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetDeclContext)
}

func (s *EcaruleContext) IN() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserIN, 0)
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

func (s *EcaruleContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserDEFAULT, 0)
}

func (s *EcaruleContext) Actions() IActionsContext {
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

func (s *EcaruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EcaruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EcaruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterEcarule(s)
	}
}

func (s *EcaruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitEcarule(s)
	}
}

func (p *SugaredAbuParser) Ecarule() (localctx IEcaruleContext) {
	this := p
	_ = this

	localctx = NewEcaruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SugaredAbuParserRULE_ecarule)
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

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(SugaredAbuParserRULE)
		}
		{
			p.SetState(163)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(164)
			p.Match(SugaredAbuParserON)
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(SugaredAbuParserTHIS-37))|(1<<(SugaredAbuParserEXT-37))|(1<<(SugaredAbuParserID-37)))) != 0) {
			{
				p.SetState(165)
				p.Resource()
			}

			p.SetState(168)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SugaredAbuParserLET {
			{
				p.SetState(170)
				p.Match(SugaredAbuParserLET)
			}
			{
				p.SetState(171)
				p.LetDecl()
			}
			{
				p.SetState(172)
				p.Match(SugaredAbuParserIN)
			}

		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SugaredAbuParserFOR {
			{
				p.SetState(176)
				p.Task()
			}

			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Match(SugaredAbuParserRULE)
		}
		{
			p.SetState(182)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(183)
			p.Match(SugaredAbuParserON)
		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(SugaredAbuParserTHIS-37))|(1<<(SugaredAbuParserEXT-37))|(1<<(SugaredAbuParserID-37)))) != 0) {
			{
				p.SetState(184)
				p.Resource()
			}

			p.SetState(187)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(189)
			p.Match(SugaredAbuParserDEFAULT)
		}
		{
			p.SetState(190)
			p.Actions()
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SugaredAbuParserLET {
			{
				p.SetState(191)
				p.Match(SugaredAbuParserLET)
			}
			{
				p.SetState(192)
				p.LetDecl()
			}
			{
				p.SetState(193)
				p.Match(SugaredAbuParserIN)
			}

		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SugaredAbuParserFOR {
			{
				p.SetState(197)
				p.Task()
			}

			p.SetState(202)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// ILetDeclContext is an interface to support dynamic dispatch.
type ILetDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLetDeclContext differentiates from other interfaces.
	IsLetDeclContext()
}

type LetDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetDeclContext() *LetDeclContext {
	var p = new(LetDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SugaredAbuParserRULE_letDecl
	return p
}

func (*LetDeclContext) IsLetDeclContext() {}

func NewLetDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetDeclContext {
	var p = new(LetDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_letDecl

	return p
}

func (s *LetDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *LetDeclContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SugaredAbuParserID)
}

func (s *LetDeclContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserID, i)
}

func (s *LetDeclContext) AllCOLONEQUAL() []antlr.TerminalNode {
	return s.GetTokens(SugaredAbuParserCOLONEQUAL)
}

func (s *LetDeclContext) COLONEQUAL(i int) antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserCOLONEQUAL, i)
}

func (s *LetDeclContext) AllExpression() []IExpressionContext {
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

func (s *LetDeclContext) Expression(i int) IExpressionContext {
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

func (s *LetDeclContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(SugaredAbuParserSEMICOLON)
}

func (s *LetDeclContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserSEMICOLON, i)
}

func (s *LetDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterLetDecl(s)
	}
}

func (s *LetDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitLetDecl(s)
	}
}

func (p *SugaredAbuParser) LetDecl() (localctx ILetDeclContext) {
	this := p
	_ = this

	localctx = NewLetDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SugaredAbuParserRULE_letDecl)
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
		p.SetState(205)
		p.Match(SugaredAbuParserID)
	}
	{
		p.SetState(206)
		p.Match(SugaredAbuParserCOLONEQUAL)
	}
	{
		p.SetState(207)
		p.expression(0)
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SugaredAbuParserSEMICOLON {
		{
			p.SetState(208)
			p.Match(SugaredAbuParserSEMICOLON)
		}
		{
			p.SetState(209)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(210)
			p.Match(SugaredAbuParserCOLONEQUAL)
		}
		{
			p.SetState(211)
			p.expression(0)
		}

		p.SetState(216)
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
	p.RuleIndex = SugaredAbuParserRULE_task
	return p
}

func (*TaskContext) IsTaskContext() {}

func NewTaskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TaskContext {
	var p = new(TaskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_task

	return p
}

func (s *TaskContext) GetParser() antlr.Parser { return s.parser }

func (s *TaskContext) FOR() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserFOR, 0)
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
	return s.GetToken(SugaredAbuParserDO, 0)
}

func (s *TaskContext) AllActions() []IActionsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IActionsContext); ok {
			len++
		}
	}

	tst := make([]IActionsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IActionsContext); ok {
			tst[i] = t.(IActionsContext)
			i++
		}
	}

	return tst
}

func (s *TaskContext) Actions(i int) IActionsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionsContext); ok {
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

	return t.(IActionsContext)
}

func (s *TaskContext) ALL() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserALL, 0)
}

func (s *TaskContext) OWISE() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserOWISE, 0)
}

func (s *TaskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TaskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TaskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterTask(s)
	}
}

func (s *TaskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitTask(s)
	}
}

func (p *SugaredAbuParser) Task() (localctx ITaskContext) {
	this := p
	_ = this

	localctx = NewTaskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SugaredAbuParserRULE_task)
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

	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Match(SugaredAbuParserFOR)
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SugaredAbuParserALL {
			{
				p.SetState(218)
				p.Match(SugaredAbuParserALL)
			}

		}
		{
			p.SetState(221)
			p.expression(0)
		}
		{
			p.SetState(222)
			p.Match(SugaredAbuParserDO)
		}
		{
			p.SetState(223)
			p.Actions()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(225)
			p.Match(SugaredAbuParserFOR)
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SugaredAbuParserALL {
			{
				p.SetState(226)
				p.Match(SugaredAbuParserALL)
			}

		}
		{
			p.SetState(229)
			p.expression(0)
		}
		{
			p.SetState(230)
			p.Match(SugaredAbuParserDO)
		}
		{
			p.SetState(231)
			p.Actions()
		}
		{
			p.SetState(232)
			p.Match(SugaredAbuParserOWISE)
		}
		{
			p.SetState(233)
			p.Actions()
		}

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
	p.RuleIndex = SugaredAbuParserRULE_actions
	return p
}

func (*ActionsContext) IsActionsContext() {}

func NewActionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionsContext {
	var p = new(ActionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_actions

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

func (s *ActionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SugaredAbuParserCOMMA)
}

func (s *ActionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserCOMMA, i)
}

func (s *ActionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterActions(s)
	}
}

func (s *ActionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitActions(s)
	}
}

func (p *SugaredAbuParser) Actions() (localctx IActionsContext) {
	this := p
	_ = this

	localctx = NewActionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SugaredAbuParserRULE_actions)
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
		p.SetState(237)
		p.Assignment()
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SugaredAbuParserCOMMA {
		{
			p.SetState(238)
			p.Match(SugaredAbuParserCOMMA)
		}
		{
			p.SetState(239)
			p.Assignment()
		}

		p.SetState(244)
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
	p.RuleIndex = SugaredAbuParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_assignment

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
	return s.GetToken(SugaredAbuParserEQUALSIGN, 0)
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
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *SugaredAbuParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SugaredAbuParserRULE_assignment)

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
		p.SetState(245)
		p.Resource()
	}
	{
		p.SetState(246)
		p.Match(SugaredAbuParserEQUALSIGN)
	}
	{
		p.SetState(247)
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
	p.RuleIndex = SugaredAbuParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ABSINT() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserABSINT, 0)
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

func (s *ExpressionContext) ABSDEC() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserABSDEC, 0)
}

func (s *ExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserNOT, 0)
}

func (s *ExpressionContext) RL_BRACKET() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserRL_BRACKET, 0)
}

func (s *ExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserRR_BRACKET, 0)
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
	return s.GetToken(SugaredAbuParserDOUBLECOLON, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserOR, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SugaredAbuParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *SugaredAbuParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, SugaredAbuParserRULE_expression, _p)

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
	p.SetState(261)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SugaredAbuParserABSINT:
		{
			p.SetState(250)
			p.Match(SugaredAbuParserABSINT)
		}
		{
			p.SetState(251)
			p.expression(11)
		}

	case SugaredAbuParserABSDEC:
		{
			p.SetState(252)
			p.Match(SugaredAbuParserABSDEC)
		}
		{
			p.SetState(253)
			p.expression(10)
		}

	case SugaredAbuParserNOT:
		{
			p.SetState(254)
			p.Match(SugaredAbuParserNOT)
		}
		{
			p.SetState(255)
			p.expression(9)
		}

	case SugaredAbuParserRL_BRACKET:
		{
			p.SetState(256)
			p.Match(SugaredAbuParserRL_BRACKET)
		}
		{
			p.SetState(257)
			p.expression(0)
		}
		{
			p.SetState(258)
			p.Match(SugaredAbuParserRR_BRACKET)
		}

	case SugaredAbuParserMINUS, SugaredAbuParserTHIS, SugaredAbuParserEXT, SugaredAbuParserTRUE, SugaredAbuParserFALSE, SugaredAbuParserID, SugaredAbuParserQUOTED_STRING, SugaredAbuParserDEC_LITERAL, SugaredAbuParserINT_LITERAL:
		{
			p.SetState(260)
			p.Term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(284)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SugaredAbuParserRULE_expression)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(264)
					p.MulDivModOperator()
				}
				{
					p.SetState(265)
					p.expression(9)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SugaredAbuParserRULE_expression)
				p.SetState(267)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(268)
					p.PlusMinusOperator()
				}
				{
					p.SetState(269)
					p.expression(8)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SugaredAbuParserRULE_expression)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(272)
					p.ComparisonOperator()
				}
				{
					p.SetState(273)
					p.expression(7)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SugaredAbuParserRULE_expression)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(276)
					p.Match(SugaredAbuParserDOUBLECOLON)
				}
				{
					p.SetState(277)
					p.expression(6)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SugaredAbuParserRULE_expression)
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(279)
					p.Match(SugaredAbuParserAND)
				}
				{
					p.SetState(280)
					p.expression(5)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SugaredAbuParserRULE_expression)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(282)
					p.Match(SugaredAbuParserOR)
				}
				{
					p.SetState(283)
					p.expression(4)
				}

			}

		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
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
	p.RuleIndex = SugaredAbuParserRULE_mulDivModOperator
	return p
}

func (*MulDivModOperatorContext) IsMulDivModOperatorContext() {}

func NewMulDivModOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulDivModOperatorContext {
	var p = new(MulDivModOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_mulDivModOperator

	return p
}

func (s *MulDivModOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MulDivModOperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserMUL, 0)
}

func (s *MulDivModOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserDIV, 0)
}

func (s *MulDivModOperatorContext) MOD() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserMOD, 0)
}

func (s *MulDivModOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulDivModOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterMulDivModOperator(s)
	}
}

func (s *MulDivModOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitMulDivModOperator(s)
	}
}

func (p *SugaredAbuParser) MulDivModOperator() (localctx IMulDivModOperatorContext) {
	this := p
	_ = this

	localctx = NewMulDivModOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SugaredAbuParserRULE_mulDivModOperator)
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
		p.SetState(289)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SugaredAbuParserDIV)|(1<<SugaredAbuParserMUL)|(1<<SugaredAbuParserMOD))) != 0) {
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
	p.RuleIndex = SugaredAbuParserRULE_plusMinusOperator
	return p
}

func (*PlusMinusOperatorContext) IsPlusMinusOperatorContext() {}

func NewPlusMinusOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlusMinusOperatorContext {
	var p = new(PlusMinusOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_plusMinusOperator

	return p
}

func (s *PlusMinusOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *PlusMinusOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserPLUS, 0)
}

func (s *PlusMinusOperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserMINUS, 0)
}

func (s *PlusMinusOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusMinusOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlusMinusOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterPlusMinusOperator(s)
	}
}

func (s *PlusMinusOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitPlusMinusOperator(s)
	}
}

func (p *SugaredAbuParser) PlusMinusOperator() (localctx IPlusMinusOperatorContext) {
	this := p
	_ = this

	localctx = NewPlusMinusOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SugaredAbuParserRULE_plusMinusOperator)
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
		p.SetState(291)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SugaredAbuParserPLUS || _la == SugaredAbuParserMINUS) {
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
	p.RuleIndex = SugaredAbuParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserEQUALS, 0)
}

func (s *ComparisonOperatorContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserNOTEQUALS, 0)
}

func (s *ComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserLT, 0)
}

func (s *ComparisonOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserLTE, 0)
}

func (s *ComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserGT, 0)
}

func (s *ComparisonOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserGTE, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (p *SugaredAbuParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	this := p
	_ = this

	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SugaredAbuParserRULE_comparisonOperator)
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
		p.SetState(293)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SugaredAbuParserEQUALS)|(1<<SugaredAbuParserGT)|(1<<SugaredAbuParserLT)|(1<<SugaredAbuParserGTE)|(1<<SugaredAbuParserLTE)|(1<<SugaredAbuParserNOTEQUALS))) != 0) {
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
	p.RuleIndex = SugaredAbuParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_term

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
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *SugaredAbuParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SugaredAbuParserRULE_term)

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

	p.SetState(297)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SugaredAbuParserMINUS, SugaredAbuParserTRUE, SugaredAbuParserFALSE, SugaredAbuParserQUOTED_STRING, SugaredAbuParserDEC_LITERAL, SugaredAbuParserINT_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(295)
			p.Value()
		}

	case SugaredAbuParserTHIS, SugaredAbuParserEXT, SugaredAbuParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(296)
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
	p.RuleIndex = SugaredAbuParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_value

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
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *SugaredAbuParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SugaredAbuParserRULE_value)

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

	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(299)
			p.StringLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(300)
			p.IntegerLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(301)
			p.DecimalLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(302)
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
	p.RuleIndex = SugaredAbuParserRULE_resource
	return p
}

func (*ResourceContext) IsResourceContext() {}

func NewResourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceContext {
	var p = new(ResourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_resource

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
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterResource(s)
	}
}

func (s *ResourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitResource(s)
	}
}

func (p *SugaredAbuParser) Resource() (localctx IResourceContext) {
	this := p
	_ = this

	localctx = NewResourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SugaredAbuParserRULE_resource)

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

	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(305)
			p.SimpleResource()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(306)
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
	p.RuleIndex = SugaredAbuParserRULE_simpleResource
	return p
}

func (*SimpleResourceContext) IsSimpleResourceContext() {}

func NewSimpleResourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleResourceContext {
	var p = new(SimpleResourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_simpleResource

	return p
}

func (s *SimpleResourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleResourceContext) ID() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserID, 0)
}

func (s *SimpleResourceContext) THIS() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserTHIS, 0)
}

func (s *SimpleResourceContext) EXT() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserEXT, 0)
}

func (s *SimpleResourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleResourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleResourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterSimpleResource(s)
	}
}

func (s *SimpleResourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitSimpleResource(s)
	}
}

func (p *SugaredAbuParser) SimpleResource() (localctx ISimpleResourceContext) {
	this := p
	_ = this

	localctx = NewSimpleResourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SugaredAbuParserRULE_simpleResource)
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

	p.SetState(315)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SugaredAbuParserTHIS, SugaredAbuParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SugaredAbuParserTHIS {
			{
				p.SetState(309)
				p.Match(SugaredAbuParserTHIS)
			}

		}
		{
			p.SetState(312)
			p.Match(SugaredAbuParserID)
		}

	case SugaredAbuParserEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)
			p.Match(SugaredAbuParserEXT)
		}
		{
			p.SetState(314)
			p.Match(SugaredAbuParserID)
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
	p.RuleIndex = SugaredAbuParserRULE_nestedResource
	return p
}

func (*NestedResourceContext) IsNestedResourceContext() {}

func NewNestedResourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedResourceContext {
	var p = new(NestedResourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_nestedResource

	return p
}

func (s *NestedResourceContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedResourceContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SugaredAbuParserID)
}

func (s *NestedResourceContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserID, i)
}

func (s *NestedResourceContext) SL_BRACKET() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserSL_BRACKET, 0)
}

func (s *NestedResourceContext) SR_BRACKET() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserSR_BRACKET, 0)
}

func (s *NestedResourceContext) THIS() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserTHIS, 0)
}

func (s *NestedResourceContext) EXT() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserEXT, 0)
}

func (s *NestedResourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedResourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedResourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterNestedResource(s)
	}
}

func (s *NestedResourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitNestedResource(s)
	}
}

func (p *SugaredAbuParser) NestedResource() (localctx INestedResourceContext) {
	this := p
	_ = this

	localctx = NewNestedResourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SugaredAbuParserRULE_nestedResource)
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

	p.SetState(329)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SugaredAbuParserTHIS, SugaredAbuParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SugaredAbuParserTHIS {
			{
				p.SetState(317)
				p.Match(SugaredAbuParserTHIS)
			}

		}
		{
			p.SetState(320)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(321)
			p.Match(SugaredAbuParserSL_BRACKET)
		}
		{
			p.SetState(322)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(323)
			p.Match(SugaredAbuParserSR_BRACKET)
		}

	case SugaredAbuParserEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.Match(SugaredAbuParserEXT)
		}
		{
			p.SetState(325)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(326)
			p.Match(SugaredAbuParserSL_BRACKET)
		}
		{
			p.SetState(327)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(328)
			p.Match(SugaredAbuParserSR_BRACKET)
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
	p.RuleIndex = SugaredAbuParserRULE_decimalLiteral
	return p
}

func (*DecimalLiteralContext) IsDecimalLiteralContext() {}

func NewDecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_decimalLiteral

	return p
}

func (s *DecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalLiteralContext) DEC_LITERAL() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserDEC_LITERAL, 0)
}

func (s *DecimalLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserMINUS, 0)
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (p *SugaredAbuParser) DecimalLiteral() (localctx IDecimalLiteralContext) {
	this := p
	_ = this

	localctx = NewDecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SugaredAbuParserRULE_decimalLiteral)
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
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SugaredAbuParserMINUS {
		{
			p.SetState(331)
			p.Match(SugaredAbuParserMINUS)
		}

	}
	{
		p.SetState(334)
		p.Match(SugaredAbuParserDEC_LITERAL)
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
	p.RuleIndex = SugaredAbuParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserINT_LITERAL, 0)
}

func (s *IntegerLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserMINUS, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (p *SugaredAbuParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	this := p
	_ = this

	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SugaredAbuParserRULE_integerLiteral)
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
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SugaredAbuParserMINUS {
		{
			p.SetState(336)
			p.Match(SugaredAbuParserMINUS)
		}

	}
	{
		p.SetState(339)
		p.Match(SugaredAbuParserINT_LITERAL)
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
	p.RuleIndex = SugaredAbuParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserQUOTED_STRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (p *SugaredAbuParser) StringLiteral() (localctx IStringLiteralContext) {
	this := p
	_ = this

	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SugaredAbuParserRULE_stringLiteral)

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
		p.SetState(341)
		p.Match(SugaredAbuParserQUOTED_STRING)
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
	p.RuleIndex = SugaredAbuParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SugaredAbuParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SugaredAbuParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (p *SugaredAbuParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	this := p
	_ = this

	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SugaredAbuParserRULE_booleanLiteral)
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
		p.SetState(343)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SugaredAbuParserTRUE || _la == SugaredAbuParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *SugaredAbuParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SugaredAbuParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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
