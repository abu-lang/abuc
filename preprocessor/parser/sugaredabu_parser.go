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
		"", "'and'", "'or'", "'not'", "'abs'", "'.'", "'+'", "'-'", "'/'", "'*'",
		"'%'", "'=='", "'>'", "'<'", "'>='", "'<='", "'!='", "'::'", "':'",
		"';'", "':='", "'='", "'('", "')'", "'{'", "'}'", "'has'", "'where'",
		"'physical'", "'logical'", "'input'", "'output'", "'rule'", "", "",
		"'on'", "'let'", "'in'", "'default'", "'for'", "'all'", "'do'", "'owise'",
		"'boolean'", "'integer'", "'decimal'", "'string'", "'true'", "'false'",
	}
	staticData.symbolicNames = []string{
		"", "AND", "OR", "NOT", "ABS", "DOT", "PLUS", "MINUS", "DIV", "MUL",
		"MOD", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "DOUBLECOLON",
		"COLON", "SEMICOLON", "COLONEQUAL", "EQUALSIGN", "RL_BRACKET", "RR_BRACKET",
		"CL_BRACKET", "CR_BRACKET", "HAS", "WHERE", "PHYSICAL", "LOGICAL", "INPUT",
		"OUTPUT", "RULE", "THIS", "EXT", "ON", "LET", "IN", "DEFAULT", "FOR",
		"ALL", "DO", "OWISE", "BOOLEAN", "INTEGER", "DECIMAL", "STRING", "TRUE",
		"FALSE", "ID", "QUOTED_STRING", "DEC_LITERAL", "INT_LITERAL", "WS",
		"COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"program", "device", "resList", "resDecl", "type", "ecarule", "letDecl",
		"task", "actions", "assignment", "expression", "mulDivModOperator",
		"plusMinusOperator", "comparisonOperator", "term", "value", "resource",
		"decimalLiteral", "integerLiteral", "stringLiteral", "booleanLiteral",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 55, 270, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 4, 0, 44, 8, 0, 11, 0, 12, 0, 45, 1, 0, 5, 0, 49, 8, 0, 10, 0, 12, 0,
		52, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 61, 8, 1, 1,
		1, 1, 1, 1, 1, 4, 1, 66, 8, 1, 11, 1, 12, 1, 67, 3, 1, 70, 8, 1, 1, 2,
		4, 2, 73, 8, 2, 11, 2, 12, 2, 74, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 95, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 103, 8, 5, 11, 5,
		12, 5, 104, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 111, 8, 5, 1, 5, 4, 5, 114, 8,
		5, 11, 5, 12, 5, 115, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 122, 8, 5, 11, 5, 12,
		5, 123, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 132, 8, 5, 1, 5, 5, 5,
		135, 8, 5, 10, 5, 12, 5, 138, 9, 5, 3, 5, 140, 8, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 149, 8, 6, 10, 6, 12, 6, 152, 9, 6, 1, 7,
		1, 7, 3, 7, 156, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 164, 8,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 172, 8, 7, 1, 8, 1, 8, 1,
		8, 5, 8, 177, 8, 8, 10, 8, 12, 8, 180, 9, 8, 1, 9, 3, 9, 183, 8, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 192, 8, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 204, 8, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		5, 10, 227, 8, 10, 10, 10, 12, 10, 230, 9, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 3, 14, 240, 8, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 3, 15, 246, 8, 15, 1, 16, 3, 16, 249, 8, 16, 1, 16, 1, 16, 1, 16,
		3, 16, 254, 8, 16, 1, 17, 3, 17, 257, 8, 17, 1, 17, 1, 17, 1, 18, 3, 18,
		262, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 0, 1, 20,
		21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 0, 5, 1, 0, 43, 46, 1, 0, 8, 10, 1, 0, 6, 7, 1, 0, 11, 16,
		1, 0, 47, 48, 287, 0, 43, 1, 0, 0, 0, 2, 53, 1, 0, 0, 0, 4, 72, 1, 0, 0,
		0, 6, 94, 1, 0, 0, 0, 8, 96, 1, 0, 0, 0, 10, 139, 1, 0, 0, 0, 12, 141,
		1, 0, 0, 0, 14, 171, 1, 0, 0, 0, 16, 173, 1, 0, 0, 0, 18, 191, 1, 0, 0,
		0, 20, 203, 1, 0, 0, 0, 22, 231, 1, 0, 0, 0, 24, 233, 1, 0, 0, 0, 26, 235,
		1, 0, 0, 0, 28, 239, 1, 0, 0, 0, 30, 245, 1, 0, 0, 0, 32, 253, 1, 0, 0,
		0, 34, 256, 1, 0, 0, 0, 36, 261, 1, 0, 0, 0, 38, 265, 1, 0, 0, 0, 40, 267,
		1, 0, 0, 0, 42, 44, 3, 2, 1, 0, 43, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0,
		45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 50, 1, 0, 0, 0, 47, 49, 3,
		10, 5, 0, 48, 47, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50,
		51, 1, 0, 0, 0, 51, 1, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 54, 5, 49, 0,
		0, 54, 55, 5, 18, 0, 0, 55, 56, 3, 38, 19, 0, 56, 57, 5, 24, 0, 0, 57,
		60, 3, 4, 2, 0, 58, 59, 5, 27, 0, 0, 59, 61, 3, 20, 10, 0, 60, 58, 1, 0,
		0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 69, 5, 25, 0, 0, 63,
		65, 5, 26, 0, 0, 64, 66, 5, 49, 0, 0, 65, 64, 1, 0, 0, 0, 66, 67, 1, 0,
		0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 70, 1, 0, 0, 0, 69, 63,
		1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 3, 1, 0, 0, 0, 71, 73, 3, 6, 3, 0,
		72, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1,
		0, 0, 0, 75, 5, 1, 0, 0, 0, 76, 77, 5, 28, 0, 0, 77, 78, 5, 31, 0, 0, 78,
		79, 3, 8, 4, 0, 79, 80, 5, 49, 0, 0, 80, 81, 5, 21, 0, 0, 81, 82, 3, 20,
		10, 0, 82, 95, 1, 0, 0, 0, 83, 84, 5, 28, 0, 0, 84, 85, 5, 30, 0, 0, 85,
		86, 3, 8, 4, 0, 86, 87, 5, 49, 0, 0, 87, 95, 1, 0, 0, 0, 88, 89, 5, 29,
		0, 0, 89, 90, 3, 8, 4, 0, 90, 91, 5, 49, 0, 0, 91, 92, 5, 21, 0, 0, 92,
		93, 3, 20, 10, 0, 93, 95, 1, 0, 0, 0, 94, 76, 1, 0, 0, 0, 94, 83, 1, 0,
		0, 0, 94, 88, 1, 0, 0, 0, 95, 7, 1, 0, 0, 0, 96, 97, 7, 0, 0, 0, 97, 9,
		1, 0, 0, 0, 98, 99, 5, 32, 0, 0, 99, 100, 5, 49, 0, 0, 100, 102, 5, 35,
		0, 0, 101, 103, 5, 49, 0, 0, 102, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0,
		104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 110, 1, 0, 0, 0, 106,
		107, 5, 36, 0, 0, 107, 108, 3, 12, 6, 0, 108, 109, 5, 37, 0, 0, 109, 111,
		1, 0, 0, 0, 110, 106, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 113, 1, 0,
		0, 0, 112, 114, 3, 14, 7, 0, 113, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0,
		115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 140, 1, 0, 0, 0, 117,
		118, 5, 32, 0, 0, 118, 119, 5, 49, 0, 0, 119, 121, 5, 35, 0, 0, 120, 122,
		5, 49, 0, 0, 121, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 121, 1, 0,
		0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 5, 38, 0, 0,
		126, 131, 3, 16, 8, 0, 127, 128, 5, 36, 0, 0, 128, 129, 3, 12, 6, 0, 129,
		130, 5, 37, 0, 0, 130, 132, 1, 0, 0, 0, 131, 127, 1, 0, 0, 0, 131, 132,
		1, 0, 0, 0, 132, 136, 1, 0, 0, 0, 133, 135, 3, 14, 7, 0, 134, 133, 1, 0,
		0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0,
		137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 98, 1, 0, 0, 0, 139, 117,
		1, 0, 0, 0, 140, 11, 1, 0, 0, 0, 141, 142, 5, 49, 0, 0, 142, 143, 5, 20,
		0, 0, 143, 150, 3, 20, 10, 0, 144, 145, 5, 19, 0, 0, 145, 146, 5, 49, 0,
		0, 146, 147, 5, 20, 0, 0, 147, 149, 3, 20, 10, 0, 148, 144, 1, 0, 0, 0,
		149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151,
		13, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 155, 5, 39, 0, 0, 154, 156,
		5, 40, 0, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0,
		0, 0, 157, 158, 3, 20, 10, 0, 158, 159, 5, 41, 0, 0, 159, 160, 3, 16, 8,
		0, 160, 172, 1, 0, 0, 0, 161, 163, 5, 39, 0, 0, 162, 164, 5, 40, 0, 0,
		163, 162, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165,
		166, 3, 20, 10, 0, 166, 167, 5, 41, 0, 0, 167, 168, 3, 16, 8, 0, 168, 169,
		5, 42, 0, 0, 169, 170, 3, 16, 8, 0, 170, 172, 1, 0, 0, 0, 171, 153, 1,
		0, 0, 0, 171, 161, 1, 0, 0, 0, 172, 15, 1, 0, 0, 0, 173, 178, 3, 18, 9,
		0, 174, 175, 5, 19, 0, 0, 175, 177, 3, 18, 9, 0, 176, 174, 1, 0, 0, 0,
		177, 180, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179,
		17, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 181, 183, 5, 33, 0, 0, 182, 181,
		1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185, 5, 49,
		0, 0, 185, 186, 5, 21, 0, 0, 186, 192, 3, 20, 10, 0, 187, 188, 5, 34, 0,
		0, 188, 189, 5, 49, 0, 0, 189, 190, 5, 21, 0, 0, 190, 192, 3, 20, 10, 0,
		191, 182, 1, 0, 0, 0, 191, 187, 1, 0, 0, 0, 192, 19, 1, 0, 0, 0, 193, 194,
		6, 10, -1, 0, 194, 195, 5, 4, 0, 0, 195, 204, 3, 20, 10, 10, 196, 197,
		5, 3, 0, 0, 197, 204, 3, 20, 10, 9, 198, 199, 5, 22, 0, 0, 199, 200, 3,
		20, 10, 0, 200, 201, 5, 23, 0, 0, 201, 204, 1, 0, 0, 0, 202, 204, 3, 28,
		14, 0, 203, 193, 1, 0, 0, 0, 203, 196, 1, 0, 0, 0, 203, 198, 1, 0, 0, 0,
		203, 202, 1, 0, 0, 0, 204, 228, 1, 0, 0, 0, 205, 206, 10, 8, 0, 0, 206,
		207, 3, 22, 11, 0, 207, 208, 3, 20, 10, 9, 208, 227, 1, 0, 0, 0, 209, 210,
		10, 7, 0, 0, 210, 211, 3, 24, 12, 0, 211, 212, 3, 20, 10, 8, 212, 227,
		1, 0, 0, 0, 213, 214, 10, 6, 0, 0, 214, 215, 3, 26, 13, 0, 215, 216, 3,
		20, 10, 7, 216, 227, 1, 0, 0, 0, 217, 218, 10, 5, 0, 0, 218, 219, 5, 17,
		0, 0, 219, 227, 3, 20, 10, 6, 220, 221, 10, 4, 0, 0, 221, 222, 5, 1, 0,
		0, 222, 227, 3, 20, 10, 5, 223, 224, 10, 3, 0, 0, 224, 225, 5, 2, 0, 0,
		225, 227, 3, 20, 10, 4, 226, 205, 1, 0, 0, 0, 226, 209, 1, 0, 0, 0, 226,
		213, 1, 0, 0, 0, 226, 217, 1, 0, 0, 0, 226, 220, 1, 0, 0, 0, 226, 223,
		1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0,
		0, 0, 229, 21, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 232, 7, 1, 0, 0,
		232, 23, 1, 0, 0, 0, 233, 234, 7, 2, 0, 0, 234, 25, 1, 0, 0, 0, 235, 236,
		7, 3, 0, 0, 236, 27, 1, 0, 0, 0, 237, 240, 3, 30, 15, 0, 238, 240, 3, 32,
		16, 0, 239, 237, 1, 0, 0, 0, 239, 238, 1, 0, 0, 0, 240, 29, 1, 0, 0, 0,
		241, 246, 3, 38, 19, 0, 242, 246, 3, 36, 18, 0, 243, 246, 3, 34, 17, 0,
		244, 246, 3, 40, 20, 0, 245, 241, 1, 0, 0, 0, 245, 242, 1, 0, 0, 0, 245,
		243, 1, 0, 0, 0, 245, 244, 1, 0, 0, 0, 246, 31, 1, 0, 0, 0, 247, 249, 5,
		33, 0, 0, 248, 247, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 1, 0, 0,
		0, 250, 254, 5, 49, 0, 0, 251, 252, 5, 34, 0, 0, 252, 254, 5, 49, 0, 0,
		253, 248, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 254, 33, 1, 0, 0, 0, 255, 257,
		5, 7, 0, 0, 256, 255, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 1, 0,
		0, 0, 258, 259, 5, 51, 0, 0, 259, 35, 1, 0, 0, 0, 260, 262, 5, 7, 0, 0,
		261, 260, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263,
		264, 5, 52, 0, 0, 264, 37, 1, 0, 0, 0, 265, 266, 5, 50, 0, 0, 266, 39,
		1, 0, 0, 0, 267, 268, 7, 4, 0, 0, 268, 41, 1, 0, 0, 0, 30, 45, 50, 60,
		67, 69, 74, 94, 104, 110, 115, 123, 131, 136, 139, 150, 155, 163, 171,
		178, 182, 191, 203, 226, 228, 239, 245, 248, 253, 256, 261,
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
	SugaredAbuParserABS           = 4
	SugaredAbuParserDOT           = 5
	SugaredAbuParserPLUS          = 6
	SugaredAbuParserMINUS         = 7
	SugaredAbuParserDIV           = 8
	SugaredAbuParserMUL           = 9
	SugaredAbuParserMOD           = 10
	SugaredAbuParserEQUALS        = 11
	SugaredAbuParserGT            = 12
	SugaredAbuParserLT            = 13
	SugaredAbuParserGTE           = 14
	SugaredAbuParserLTE           = 15
	SugaredAbuParserNOTEQUALS     = 16
	SugaredAbuParserDOUBLECOLON   = 17
	SugaredAbuParserCOLON         = 18
	SugaredAbuParserSEMICOLON     = 19
	SugaredAbuParserCOLONEQUAL    = 20
	SugaredAbuParserEQUALSIGN     = 21
	SugaredAbuParserRL_BRACKET    = 22
	SugaredAbuParserRR_BRACKET    = 23
	SugaredAbuParserCL_BRACKET    = 24
	SugaredAbuParserCR_BRACKET    = 25
	SugaredAbuParserHAS           = 26
	SugaredAbuParserWHERE         = 27
	SugaredAbuParserPHYSICAL      = 28
	SugaredAbuParserLOGICAL       = 29
	SugaredAbuParserINPUT         = 30
	SugaredAbuParserOUTPUT        = 31
	SugaredAbuParserRULE          = 32
	SugaredAbuParserTHIS          = 33
	SugaredAbuParserEXT           = 34
	SugaredAbuParserON            = 35
	SugaredAbuParserLET           = 36
	SugaredAbuParserIN            = 37
	SugaredAbuParserDEFAULT       = 38
	SugaredAbuParserFOR           = 39
	SugaredAbuParserALL           = 40
	SugaredAbuParserDO            = 41
	SugaredAbuParserOWISE         = 42
	SugaredAbuParserBOOLEAN       = 43
	SugaredAbuParserINTEGER       = 44
	SugaredAbuParserDECIMAL       = 45
	SugaredAbuParserSTRING        = 46
	SugaredAbuParserTRUE          = 47
	SugaredAbuParserFALSE         = 48
	SugaredAbuParserID            = 49
	SugaredAbuParserQUOTED_STRING = 50
	SugaredAbuParserDEC_LITERAL   = 51
	SugaredAbuParserINT_LITERAL   = 52
	SugaredAbuParserWS            = 53
	SugaredAbuParserCOMMENT       = 54
	SugaredAbuParserLINE_COMMENT  = 55
)

// SugaredAbuParser rules.
const (
	SugaredAbuParserRULE_program            = 0
	SugaredAbuParserRULE_device             = 1
	SugaredAbuParserRULE_resList            = 2
	SugaredAbuParserRULE_resDecl            = 3
	SugaredAbuParserRULE_type               = 4
	SugaredAbuParserRULE_ecarule            = 5
	SugaredAbuParserRULE_letDecl            = 6
	SugaredAbuParserRULE_task               = 7
	SugaredAbuParserRULE_actions            = 8
	SugaredAbuParserRULE_assignment         = 9
	SugaredAbuParserRULE_expression         = 10
	SugaredAbuParserRULE_mulDivModOperator  = 11
	SugaredAbuParserRULE_plusMinusOperator  = 12
	SugaredAbuParserRULE_comparisonOperator = 13
	SugaredAbuParserRULE_term               = 14
	SugaredAbuParserRULE_value              = 15
	SugaredAbuParserRULE_resource           = 16
	SugaredAbuParserRULE_decimalLiteral     = 17
	SugaredAbuParserRULE_integerLiteral     = 18
	SugaredAbuParserRULE_stringLiteral      = 19
	SugaredAbuParserRULE_booleanLiteral     = 20
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
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SugaredAbuParserID {
		{
			p.SetState(42)
			p.Device()
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SugaredAbuParserRULE {
		{
			p.SetState(47)
			p.Ecarule()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 2, SugaredAbuParserRULE_device)
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
		p.SetState(53)
		p.Match(SugaredAbuParserID)
	}
	{
		p.SetState(54)
		p.Match(SugaredAbuParserCOLON)
	}
	{
		p.SetState(55)
		p.StringLiteral()
	}
	{
		p.SetState(56)
		p.Match(SugaredAbuParserCL_BRACKET)
	}
	{
		p.SetState(57)
		p.ResList()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SugaredAbuParserWHERE {
		{
			p.SetState(58)
			p.Match(SugaredAbuParserWHERE)
		}
		{
			p.SetState(59)
			p.expression(0)
		}

	}
	{
		p.SetState(62)
		p.Match(SugaredAbuParserCR_BRACKET)
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SugaredAbuParserHAS {
		{
			p.SetState(63)
			p.Match(SugaredAbuParserHAS)
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(64)
					p.Match(SugaredAbuParserID)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 4, SugaredAbuParserRULE_resList)
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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SugaredAbuParserPHYSICAL || _la == SugaredAbuParserLOGICAL {
		{
			p.SetState(71)
			p.ResDecl()
		}

		p.SetState(74)
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
	p.EnterRule(localctx, 6, SugaredAbuParserRULE_resDecl)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Match(SugaredAbuParserPHYSICAL)
		}
		{
			p.SetState(77)
			p.Match(SugaredAbuParserOUTPUT)
		}
		{
			p.SetState(78)
			p.Type()
		}
		{
			p.SetState(79)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(80)
			p.Match(SugaredAbuParserEQUALSIGN)
		}
		{
			p.SetState(81)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Match(SugaredAbuParserPHYSICAL)
		}
		{
			p.SetState(84)
			p.Match(SugaredAbuParserINPUT)
		}
		{
			p.SetState(85)
			p.Type()
		}
		{
			p.SetState(86)
			p.Match(SugaredAbuParserID)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Match(SugaredAbuParserLOGICAL)
		}
		{
			p.SetState(89)
			p.Type()
		}
		{
			p.SetState(90)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(91)
			p.Match(SugaredAbuParserEQUALSIGN)
		}
		{
			p.SetState(92)
			p.expression(0)
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
	p.EnterRule(localctx, 8, SugaredAbuParserRULE_type)
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
		p.SetState(96)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(SugaredAbuParserBOOLEAN-43))|(1<<(SugaredAbuParserINTEGER-43))|(1<<(SugaredAbuParserDECIMAL-43))|(1<<(SugaredAbuParserSTRING-43)))) != 0) {
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

func (s *EcaruleContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SugaredAbuParserID)
}

func (s *EcaruleContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserID, i)
}

func (s *EcaruleContext) ON() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserON, 0)
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
	p.EnterRule(localctx, 10, SugaredAbuParserRULE_ecarule)
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

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Match(SugaredAbuParserRULE)
		}
		{
			p.SetState(99)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(100)
			p.Match(SugaredAbuParserON)
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SugaredAbuParserID {
			{
				p.SetState(101)
				p.Match(SugaredAbuParserID)
			}

			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SugaredAbuParserLET {
			{
				p.SetState(106)
				p.Match(SugaredAbuParserLET)
			}
			{
				p.SetState(107)
				p.LetDecl()
			}
			{
				p.SetState(108)
				p.Match(SugaredAbuParserIN)
			}

		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SugaredAbuParserFOR {
			{
				p.SetState(112)
				p.Task()
			}

			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(SugaredAbuParserRULE)
		}
		{
			p.SetState(118)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(119)
			p.Match(SugaredAbuParserON)
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SugaredAbuParserID {
			{
				p.SetState(120)
				p.Match(SugaredAbuParserID)
			}

			p.SetState(123)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(125)
			p.Match(SugaredAbuParserDEFAULT)
		}
		{
			p.SetState(126)
			p.Actions()
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SugaredAbuParserLET {
			{
				p.SetState(127)
				p.Match(SugaredAbuParserLET)
			}
			{
				p.SetState(128)
				p.LetDecl()
			}
			{
				p.SetState(129)
				p.Match(SugaredAbuParserIN)
			}

		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SugaredAbuParserFOR {
			{
				p.SetState(133)
				p.Task()
			}

			p.SetState(138)
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
	p.EnterRule(localctx, 12, SugaredAbuParserRULE_letDecl)
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
		p.SetState(141)
		p.Match(SugaredAbuParserID)
	}
	{
		p.SetState(142)
		p.Match(SugaredAbuParserCOLONEQUAL)
	}
	{
		p.SetState(143)
		p.expression(0)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SugaredAbuParserSEMICOLON {
		{
			p.SetState(144)
			p.Match(SugaredAbuParserSEMICOLON)
		}
		{
			p.SetState(145)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(146)
			p.Match(SugaredAbuParserCOLONEQUAL)
		}
		{
			p.SetState(147)
			p.expression(0)
		}

		p.SetState(152)
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
	p.EnterRule(localctx, 14, SugaredAbuParserRULE_task)
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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Match(SugaredAbuParserFOR)
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SugaredAbuParserALL {
			{
				p.SetState(154)
				p.Match(SugaredAbuParserALL)
			}

		}
		{
			p.SetState(157)
			p.expression(0)
		}
		{
			p.SetState(158)
			p.Match(SugaredAbuParserDO)
		}
		{
			p.SetState(159)
			p.Actions()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.Match(SugaredAbuParserFOR)
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SugaredAbuParserALL {
			{
				p.SetState(162)
				p.Match(SugaredAbuParserALL)
			}

		}
		{
			p.SetState(165)
			p.expression(0)
		}
		{
			p.SetState(166)
			p.Match(SugaredAbuParserDO)
		}
		{
			p.SetState(167)
			p.Actions()
		}
		{
			p.SetState(168)
			p.Match(SugaredAbuParserOWISE)
		}
		{
			p.SetState(169)
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

func (s *ActionsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(SugaredAbuParserSEMICOLON)
}

func (s *ActionsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserSEMICOLON, i)
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
	p.EnterRule(localctx, 16, SugaredAbuParserRULE_actions)
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
		p.SetState(173)
		p.Assignment()
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SugaredAbuParserSEMICOLON {
		{
			p.SetState(174)
			p.Match(SugaredAbuParserSEMICOLON)
		}
		{
			p.SetState(175)
			p.Assignment()
		}

		p.SetState(180)
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

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserID, 0)
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

func (s *AssignmentContext) THIS() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserTHIS, 0)
}

func (s *AssignmentContext) EXT() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserEXT, 0)
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
	p.EnterRule(localctx, 18, SugaredAbuParserRULE_assignment)
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

	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SugaredAbuParserTHIS, SugaredAbuParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SugaredAbuParserTHIS {
			{
				p.SetState(181)
				p.Match(SugaredAbuParserTHIS)
			}

		}
		{
			p.SetState(184)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(185)
			p.Match(SugaredAbuParserEQUALSIGN)
		}
		{
			p.SetState(186)
			p.expression(0)
		}

	case SugaredAbuParserEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.Match(SugaredAbuParserEXT)
		}
		{
			p.SetState(188)
			p.Match(SugaredAbuParserID)
		}
		{
			p.SetState(189)
			p.Match(SugaredAbuParserEQUALSIGN)
		}
		{
			p.SetState(190)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *ExpressionContext) ABS() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserABS, 0)
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, SugaredAbuParserRULE_expression, _p)

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
	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SugaredAbuParserABS:
		{
			p.SetState(194)
			p.Match(SugaredAbuParserABS)
		}
		{
			p.SetState(195)
			p.expression(10)
		}

	case SugaredAbuParserNOT:
		{
			p.SetState(196)
			p.Match(SugaredAbuParserNOT)
		}
		{
			p.SetState(197)
			p.expression(9)
		}

	case SugaredAbuParserRL_BRACKET:
		{
			p.SetState(198)
			p.Match(SugaredAbuParserRL_BRACKET)
		}
		{
			p.SetState(199)
			p.expression(0)
		}
		{
			p.SetState(200)
			p.Match(SugaredAbuParserRR_BRACKET)
		}

	case SugaredAbuParserMINUS, SugaredAbuParserTHIS, SugaredAbuParserEXT, SugaredAbuParserTRUE, SugaredAbuParserFALSE, SugaredAbuParserID, SugaredAbuParserQUOTED_STRING, SugaredAbuParserDEC_LITERAL, SugaredAbuParserINT_LITERAL:
		{
			p.SetState(202)
			p.Term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SugaredAbuParserRULE_expression)
				p.SetState(205)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(206)
					p.MulDivModOperator()
				}
				{
					p.SetState(207)
					p.expression(9)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SugaredAbuParserRULE_expression)
				p.SetState(209)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(210)
					p.PlusMinusOperator()
				}
				{
					p.SetState(211)
					p.expression(8)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SugaredAbuParserRULE_expression)
				p.SetState(213)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(214)
					p.ComparisonOperator()
				}
				{
					p.SetState(215)
					p.expression(7)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SugaredAbuParserRULE_expression)
				p.SetState(217)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(218)
					p.Match(SugaredAbuParserDOUBLECOLON)
				}
				{
					p.SetState(219)
					p.expression(6)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SugaredAbuParserRULE_expression)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(221)
					p.Match(SugaredAbuParserAND)
				}
				{
					p.SetState(222)
					p.expression(5)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SugaredAbuParserRULE_expression)
				p.SetState(223)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(224)
					p.Match(SugaredAbuParserOR)
				}
				{
					p.SetState(225)
					p.expression(4)
				}

			}

		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 22, SugaredAbuParserRULE_mulDivModOperator)
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
	p.EnterRule(localctx, 24, SugaredAbuParserRULE_plusMinusOperator)
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
		p.SetState(233)
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
	p.EnterRule(localctx, 26, SugaredAbuParserRULE_comparisonOperator)
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
		p.SetState(235)
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
	p.EnterRule(localctx, 28, SugaredAbuParserRULE_term)

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

	p.SetState(239)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SugaredAbuParserMINUS, SugaredAbuParserTRUE, SugaredAbuParserFALSE, SugaredAbuParserQUOTED_STRING, SugaredAbuParserDEC_LITERAL, SugaredAbuParserINT_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(237)
			p.Value()
		}

	case SugaredAbuParserTHIS, SugaredAbuParserEXT, SugaredAbuParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
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
	p.EnterRule(localctx, 30, SugaredAbuParserRULE_value)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.StringLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.IntegerLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(243)
			p.DecimalLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(244)
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

func (s *ResourceContext) ID() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserID, 0)
}

func (s *ResourceContext) THIS() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserTHIS, 0)
}

func (s *ResourceContext) EXT() antlr.TerminalNode {
	return s.GetToken(SugaredAbuParserEXT, 0)
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
	p.EnterRule(localctx, 32, SugaredAbuParserRULE_resource)
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
	case SugaredAbuParserTHIS, SugaredAbuParserID:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SugaredAbuParserTHIS {
			{
				p.SetState(247)
				p.Match(SugaredAbuParserTHIS)
			}

		}
		{
			p.SetState(250)
			p.Match(SugaredAbuParserID)
		}

	case SugaredAbuParserEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)
			p.Match(SugaredAbuParserEXT)
		}
		{
			p.SetState(252)
			p.Match(SugaredAbuParserID)
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
	p.EnterRule(localctx, 34, SugaredAbuParserRULE_decimalLiteral)
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
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SugaredAbuParserMINUS {
		{
			p.SetState(255)
			p.Match(SugaredAbuParserMINUS)
		}

	}
	{
		p.SetState(258)
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
	p.EnterRule(localctx, 36, SugaredAbuParserRULE_integerLiteral)
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
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SugaredAbuParserMINUS {
		{
			p.SetState(260)
			p.Match(SugaredAbuParserMINUS)
		}

	}
	{
		p.SetState(263)
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
	p.EnterRule(localctx, 38, SugaredAbuParserRULE_stringLiteral)

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
		p.SetState(265)
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
	p.EnterRule(localctx, 40, SugaredAbuParserRULE_booleanLiteral)
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
		p.SetState(267)
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
	case 10:
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
