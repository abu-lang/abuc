// Code generated from AbuLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type AbuLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AbuLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func abulexerLexerInit() {
	staticData := &AbuLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'and'", "'or'", "'not'", "'absint'", "'absdec'", "'.'", "'+'",
		"'-'", "'/'", "'*'", "'%'", "'=='", "'>'", "'<'", "'>='", "'<='", "'!='",
		"'::'", "':'", "':='", "'='", "','", "'('", "')'", "'['", "']'", "'{'",
		"'}'", "'has'", "'where'", "'physical'", "'logical'", "'input'", "'output'",
		"'rule'", "", "", "'on'", "'for'", "'all'", "'do'", "'define'", "'as'",
		"'boolean'", "'integer'", "'decimal'", "'string'", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "AND", "OR", "NOT", "ABSINT", "ABSDEC", "DOT", "PLUS", "MINUS",
		"DIV", "MUL", "MOD", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS",
		"DOUBLECOLON", "COLON", "COLONEQUAL", "EQUALSIGN", "COMMA", "RL_BRACKET",
		"RR_BRACKET", "SL_BRACKET", "SR_BRACKET", "CL_BRACKET", "CR_BRACKET",
		"HAS", "WHERE", "PHYSICAL", "LOGICAL", "INPUT", "OUTPUT", "RULE", "THIS",
		"EXT", "ON", "FOR", "ALL", "DO", "DEFINE", "AS", "BOOLEAN", "INTEGER",
		"DECIMAL", "STRING", "TRUE", "FALSE", "ID", "QUOTED_STRING", "DEC_LITERAL",
		"INT_LITERAL", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"AND", "OR", "NOT", "ABSINT", "ABSDEC", "DOT", "PLUS", "MINUS", "DIV",
		"MUL", "MOD", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "DOUBLECOLON",
		"COLON", "COLONEQUAL", "EQUALSIGN", "COMMA", "RL_BRACKET", "RR_BRACKET",
		"SL_BRACKET", "SR_BRACKET", "CL_BRACKET", "CR_BRACKET", "HAS", "WHERE",
		"PHYSICAL", "LOGICAL", "INPUT", "OUTPUT", "RULE", "THIS", "EXT", "ON",
		"FOR", "ALL", "DO", "DEFINE", "AS", "BOOLEAN", "INTEGER", "DECIMAL",
		"STRING", "TRUE", "FALSE", "ID", "QUOTED_STRING", "DEC_LITERAL", "INT_LITERAL",
		"WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 56, 378, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1,
		43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1,
		46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 5, 49, 317, 8, 49, 10,
		49, 12, 49, 320, 9, 49, 1, 50, 1, 50, 5, 50, 324, 8, 50, 10, 50, 12, 50,
		327, 9, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 5, 51, 334, 8, 51, 10, 51,
		12, 51, 337, 9, 51, 1, 52, 1, 52, 1, 52, 5, 52, 342, 8, 52, 10, 52, 12,
		52, 345, 9, 52, 3, 52, 347, 8, 52, 1, 53, 4, 53, 350, 8, 53, 11, 53, 12,
		53, 351, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 5, 54, 360, 8, 54, 10,
		54, 12, 54, 363, 9, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55,
		5, 55, 372, 8, 55, 10, 55, 12, 55, 375, 9, 55, 1, 55, 1, 55, 1, 361, 0,
		56, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57,
		29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75,
		38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93,
		47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55,
		111, 56, 1, 0, 7, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65, 90, 97, 122,
		1, 0, 34, 34, 1, 0, 48, 57, 1, 0, 49, 57, 3, 0, 9, 10, 13, 13, 32, 32,
		2, 0, 10, 10, 13, 13, 385, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0,
		0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0,
		0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1,
		0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0,
		105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0,
		0, 0, 1, 113, 1, 0, 0, 0, 3, 117, 1, 0, 0, 0, 5, 120, 1, 0, 0, 0, 7, 124,
		1, 0, 0, 0, 9, 131, 1, 0, 0, 0, 11, 138, 1, 0, 0, 0, 13, 140, 1, 0, 0,
		0, 15, 142, 1, 0, 0, 0, 17, 144, 1, 0, 0, 0, 19, 146, 1, 0, 0, 0, 21, 148,
		1, 0, 0, 0, 23, 150, 1, 0, 0, 0, 25, 153, 1, 0, 0, 0, 27, 155, 1, 0, 0,
		0, 29, 157, 1, 0, 0, 0, 31, 160, 1, 0, 0, 0, 33, 163, 1, 0, 0, 0, 35, 166,
		1, 0, 0, 0, 37, 169, 1, 0, 0, 0, 39, 171, 1, 0, 0, 0, 41, 174, 1, 0, 0,
		0, 43, 176, 1, 0, 0, 0, 45, 178, 1, 0, 0, 0, 47, 180, 1, 0, 0, 0, 49, 182,
		1, 0, 0, 0, 51, 184, 1, 0, 0, 0, 53, 186, 1, 0, 0, 0, 55, 188, 1, 0, 0,
		0, 57, 190, 1, 0, 0, 0, 59, 194, 1, 0, 0, 0, 61, 200, 1, 0, 0, 0, 63, 209,
		1, 0, 0, 0, 65, 217, 1, 0, 0, 0, 67, 223, 1, 0, 0, 0, 69, 230, 1, 0, 0,
		0, 71, 235, 1, 0, 0, 0, 73, 242, 1, 0, 0, 0, 75, 248, 1, 0, 0, 0, 77, 251,
		1, 0, 0, 0, 79, 255, 1, 0, 0, 0, 81, 259, 1, 0, 0, 0, 83, 262, 1, 0, 0,
		0, 85, 269, 1, 0, 0, 0, 87, 272, 1, 0, 0, 0, 89, 280, 1, 0, 0, 0, 91, 288,
		1, 0, 0, 0, 93, 296, 1, 0, 0, 0, 95, 303, 1, 0, 0, 0, 97, 308, 1, 0, 0,
		0, 99, 314, 1, 0, 0, 0, 101, 321, 1, 0, 0, 0, 103, 330, 1, 0, 0, 0, 105,
		346, 1, 0, 0, 0, 107, 349, 1, 0, 0, 0, 109, 355, 1, 0, 0, 0, 111, 369,
		1, 0, 0, 0, 113, 114, 5, 97, 0, 0, 114, 115, 5, 110, 0, 0, 115, 116, 5,
		100, 0, 0, 116, 2, 1, 0, 0, 0, 117, 118, 5, 111, 0, 0, 118, 119, 5, 114,
		0, 0, 119, 4, 1, 0, 0, 0, 120, 121, 5, 110, 0, 0, 121, 122, 5, 111, 0,
		0, 122, 123, 5, 116, 0, 0, 123, 6, 1, 0, 0, 0, 124, 125, 5, 97, 0, 0, 125,
		126, 5, 98, 0, 0, 126, 127, 5, 115, 0, 0, 127, 128, 5, 105, 0, 0, 128,
		129, 5, 110, 0, 0, 129, 130, 5, 116, 0, 0, 130, 8, 1, 0, 0, 0, 131, 132,
		5, 97, 0, 0, 132, 133, 5, 98, 0, 0, 133, 134, 5, 115, 0, 0, 134, 135, 5,
		100, 0, 0, 135, 136, 5, 101, 0, 0, 136, 137, 5, 99, 0, 0, 137, 10, 1, 0,
		0, 0, 138, 139, 5, 46, 0, 0, 139, 12, 1, 0, 0, 0, 140, 141, 5, 43, 0, 0,
		141, 14, 1, 0, 0, 0, 142, 143, 5, 45, 0, 0, 143, 16, 1, 0, 0, 0, 144, 145,
		5, 47, 0, 0, 145, 18, 1, 0, 0, 0, 146, 147, 5, 42, 0, 0, 147, 20, 1, 0,
		0, 0, 148, 149, 5, 37, 0, 0, 149, 22, 1, 0, 0, 0, 150, 151, 5, 61, 0, 0,
		151, 152, 5, 61, 0, 0, 152, 24, 1, 0, 0, 0, 153, 154, 5, 62, 0, 0, 154,
		26, 1, 0, 0, 0, 155, 156, 5, 60, 0, 0, 156, 28, 1, 0, 0, 0, 157, 158, 5,
		62, 0, 0, 158, 159, 5, 61, 0, 0, 159, 30, 1, 0, 0, 0, 160, 161, 5, 60,
		0, 0, 161, 162, 5, 61, 0, 0, 162, 32, 1, 0, 0, 0, 163, 164, 5, 33, 0, 0,
		164, 165, 5, 61, 0, 0, 165, 34, 1, 0, 0, 0, 166, 167, 5, 58, 0, 0, 167,
		168, 5, 58, 0, 0, 168, 36, 1, 0, 0, 0, 169, 170, 5, 58, 0, 0, 170, 38,
		1, 0, 0, 0, 171, 172, 5, 58, 0, 0, 172, 173, 5, 61, 0, 0, 173, 40, 1, 0,
		0, 0, 174, 175, 5, 61, 0, 0, 175, 42, 1, 0, 0, 0, 176, 177, 5, 44, 0, 0,
		177, 44, 1, 0, 0, 0, 178, 179, 5, 40, 0, 0, 179, 46, 1, 0, 0, 0, 180, 181,
		5, 41, 0, 0, 181, 48, 1, 0, 0, 0, 182, 183, 5, 91, 0, 0, 183, 50, 1, 0,
		0, 0, 184, 185, 5, 93, 0, 0, 185, 52, 1, 0, 0, 0, 186, 187, 5, 123, 0,
		0, 187, 54, 1, 0, 0, 0, 188, 189, 5, 125, 0, 0, 189, 56, 1, 0, 0, 0, 190,
		191, 5, 104, 0, 0, 191, 192, 5, 97, 0, 0, 192, 193, 5, 115, 0, 0, 193,
		58, 1, 0, 0, 0, 194, 195, 5, 119, 0, 0, 195, 196, 5, 104, 0, 0, 196, 197,
		5, 101, 0, 0, 197, 198, 5, 114, 0, 0, 198, 199, 5, 101, 0, 0, 199, 60,
		1, 0, 0, 0, 200, 201, 5, 112, 0, 0, 201, 202, 5, 104, 0, 0, 202, 203, 5,
		121, 0, 0, 203, 204, 5, 115, 0, 0, 204, 205, 5, 105, 0, 0, 205, 206, 5,
		99, 0, 0, 206, 207, 5, 97, 0, 0, 207, 208, 5, 108, 0, 0, 208, 62, 1, 0,
		0, 0, 209, 210, 5, 108, 0, 0, 210, 211, 5, 111, 0, 0, 211, 212, 5, 103,
		0, 0, 212, 213, 5, 105, 0, 0, 213, 214, 5, 99, 0, 0, 214, 215, 5, 97, 0,
		0, 215, 216, 5, 108, 0, 0, 216, 64, 1, 0, 0, 0, 217, 218, 5, 105, 0, 0,
		218, 219, 5, 110, 0, 0, 219, 220, 5, 112, 0, 0, 220, 221, 5, 117, 0, 0,
		221, 222, 5, 116, 0, 0, 222, 66, 1, 0, 0, 0, 223, 224, 5, 111, 0, 0, 224,
		225, 5, 117, 0, 0, 225, 226, 5, 116, 0, 0, 226, 227, 5, 112, 0, 0, 227,
		228, 5, 117, 0, 0, 228, 229, 5, 116, 0, 0, 229, 68, 1, 0, 0, 0, 230, 231,
		5, 114, 0, 0, 231, 232, 5, 117, 0, 0, 232, 233, 5, 108, 0, 0, 233, 234,
		5, 101, 0, 0, 234, 70, 1, 0, 0, 0, 235, 236, 5, 116, 0, 0, 236, 237, 5,
		104, 0, 0, 237, 238, 5, 105, 0, 0, 238, 239, 5, 115, 0, 0, 239, 240, 1,
		0, 0, 0, 240, 241, 3, 11, 5, 0, 241, 72, 1, 0, 0, 0, 242, 243, 5, 101,
		0, 0, 243, 244, 5, 120, 0, 0, 244, 245, 5, 116, 0, 0, 245, 246, 1, 0, 0,
		0, 246, 247, 3, 11, 5, 0, 247, 74, 1, 0, 0, 0, 248, 249, 5, 111, 0, 0,
		249, 250, 5, 110, 0, 0, 250, 76, 1, 0, 0, 0, 251, 252, 5, 102, 0, 0, 252,
		253, 5, 111, 0, 0, 253, 254, 5, 114, 0, 0, 254, 78, 1, 0, 0, 0, 255, 256,
		5, 97, 0, 0, 256, 257, 5, 108, 0, 0, 257, 258, 5, 108, 0, 0, 258, 80, 1,
		0, 0, 0, 259, 260, 5, 100, 0, 0, 260, 261, 5, 111, 0, 0, 261, 82, 1, 0,
		0, 0, 262, 263, 5, 100, 0, 0, 263, 264, 5, 101, 0, 0, 264, 265, 5, 102,
		0, 0, 265, 266, 5, 105, 0, 0, 266, 267, 5, 110, 0, 0, 267, 268, 5, 101,
		0, 0, 268, 84, 1, 0, 0, 0, 269, 270, 5, 97, 0, 0, 270, 271, 5, 115, 0,
		0, 271, 86, 1, 0, 0, 0, 272, 273, 5, 98, 0, 0, 273, 274, 5, 111, 0, 0,
		274, 275, 5, 111, 0, 0, 275, 276, 5, 108, 0, 0, 276, 277, 5, 101, 0, 0,
		277, 278, 5, 97, 0, 0, 278, 279, 5, 110, 0, 0, 279, 88, 1, 0, 0, 0, 280,
		281, 5, 105, 0, 0, 281, 282, 5, 110, 0, 0, 282, 283, 5, 116, 0, 0, 283,
		284, 5, 101, 0, 0, 284, 285, 5, 103, 0, 0, 285, 286, 5, 101, 0, 0, 286,
		287, 5, 114, 0, 0, 287, 90, 1, 0, 0, 0, 288, 289, 5, 100, 0, 0, 289, 290,
		5, 101, 0, 0, 290, 291, 5, 99, 0, 0, 291, 292, 5, 105, 0, 0, 292, 293,
		5, 109, 0, 0, 293, 294, 5, 97, 0, 0, 294, 295, 5, 108, 0, 0, 295, 92, 1,
		0, 0, 0, 296, 297, 5, 115, 0, 0, 297, 298, 5, 116, 0, 0, 298, 299, 5, 114,
		0, 0, 299, 300, 5, 105, 0, 0, 300, 301, 5, 110, 0, 0, 301, 302, 5, 103,
		0, 0, 302, 94, 1, 0, 0, 0, 303, 304, 5, 116, 0, 0, 304, 305, 5, 114, 0,
		0, 305, 306, 5, 117, 0, 0, 306, 307, 5, 101, 0, 0, 307, 96, 1, 0, 0, 0,
		308, 309, 5, 102, 0, 0, 309, 310, 5, 97, 0, 0, 310, 311, 5, 108, 0, 0,
		311, 312, 5, 115, 0, 0, 312, 313, 5, 101, 0, 0, 313, 98, 1, 0, 0, 0, 314,
		318, 7, 0, 0, 0, 315, 317, 7, 1, 0, 0, 316, 315, 1, 0, 0, 0, 317, 320,
		1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 100, 1, 0,
		0, 0, 320, 318, 1, 0, 0, 0, 321, 325, 5, 34, 0, 0, 322, 324, 8, 2, 0, 0,
		323, 322, 1, 0, 0, 0, 324, 327, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325,
		326, 1, 0, 0, 0, 326, 328, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 328, 329,
		5, 34, 0, 0, 329, 102, 1, 0, 0, 0, 330, 331, 3, 105, 52, 0, 331, 335, 3,
		11, 5, 0, 332, 334, 7, 3, 0, 0, 333, 332, 1, 0, 0, 0, 334, 337, 1, 0, 0,
		0, 335, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 104, 1, 0, 0, 0, 337,
		335, 1, 0, 0, 0, 338, 347, 5, 48, 0, 0, 339, 343, 7, 4, 0, 0, 340, 342,
		7, 3, 0, 0, 341, 340, 1, 0, 0, 0, 342, 345, 1, 0, 0, 0, 343, 341, 1, 0,
		0, 0, 343, 344, 1, 0, 0, 0, 344, 347, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0,
		346, 338, 1, 0, 0, 0, 346, 339, 1, 0, 0, 0, 347, 106, 1, 0, 0, 0, 348,
		350, 7, 5, 0, 0, 349, 348, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 349,
		1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 354, 6, 53,
		0, 0, 354, 108, 1, 0, 0, 0, 355, 356, 5, 92, 0, 0, 356, 357, 5, 64, 0,
		0, 357, 361, 1, 0, 0, 0, 358, 360, 9, 0, 0, 0, 359, 358, 1, 0, 0, 0, 360,
		363, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 362, 364,
		1, 0, 0, 0, 363, 361, 1, 0, 0, 0, 364, 365, 5, 64, 0, 0, 365, 366, 5, 92,
		0, 0, 366, 367, 1, 0, 0, 0, 367, 368, 6, 54, 1, 0, 368, 110, 1, 0, 0, 0,
		369, 373, 5, 35, 0, 0, 370, 372, 8, 6, 0, 0, 371, 370, 1, 0, 0, 0, 372,
		375, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 376,
		1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 376, 377, 6, 55, 1, 0, 377, 112, 1, 0,
		0, 0, 9, 0, 318, 325, 335, 343, 346, 351, 361, 373, 2, 0, 1, 0, 6, 0, 0,
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

// AbuLexerInit initializes any static state used to implement AbuLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAbuLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AbuLexerInit() {
	staticData := &AbuLexerLexerStaticData
	staticData.once.Do(abulexerLexerInit)
}

// NewAbuLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAbuLexer(input antlr.CharStream) *AbuLexer {
	AbuLexerInit()
	l := new(AbuLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AbuLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "AbuLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AbuLexer tokens.
const (
	AbuLexerAND           = 1
	AbuLexerOR            = 2
	AbuLexerNOT           = 3
	AbuLexerABSINT        = 4
	AbuLexerABSDEC        = 5
	AbuLexerDOT           = 6
	AbuLexerPLUS          = 7
	AbuLexerMINUS         = 8
	AbuLexerDIV           = 9
	AbuLexerMUL           = 10
	AbuLexerMOD           = 11
	AbuLexerEQUALS        = 12
	AbuLexerGT            = 13
	AbuLexerLT            = 14
	AbuLexerGTE           = 15
	AbuLexerLTE           = 16
	AbuLexerNOTEQUALS     = 17
	AbuLexerDOUBLECOLON   = 18
	AbuLexerCOLON         = 19
	AbuLexerCOLONEQUAL    = 20
	AbuLexerEQUALSIGN     = 21
	AbuLexerCOMMA         = 22
	AbuLexerRL_BRACKET    = 23
	AbuLexerRR_BRACKET    = 24
	AbuLexerSL_BRACKET    = 25
	AbuLexerSR_BRACKET    = 26
	AbuLexerCL_BRACKET    = 27
	AbuLexerCR_BRACKET    = 28
	AbuLexerHAS           = 29
	AbuLexerWHERE         = 30
	AbuLexerPHYSICAL      = 31
	AbuLexerLOGICAL       = 32
	AbuLexerINPUT         = 33
	AbuLexerOUTPUT        = 34
	AbuLexerRULE          = 35
	AbuLexerTHIS          = 36
	AbuLexerEXT           = 37
	AbuLexerON            = 38
	AbuLexerFOR           = 39
	AbuLexerALL           = 40
	AbuLexerDO            = 41
	AbuLexerDEFINE        = 42
	AbuLexerAS            = 43
	AbuLexerBOOLEAN       = 44
	AbuLexerINTEGER       = 45
	AbuLexerDECIMAL       = 46
	AbuLexerSTRING        = 47
	AbuLexerTRUE          = 48
	AbuLexerFALSE         = 49
	AbuLexerID            = 50
	AbuLexerQUOTED_STRING = 51
	AbuLexerDEC_LITERAL   = 52
	AbuLexerINT_LITERAL   = 53
	AbuLexerWS            = 54
	AbuLexerCOMMENT       = 55
	AbuLexerLINE_COMMENT  = 56
)
