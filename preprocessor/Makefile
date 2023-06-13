# Copyright 2023 Massimo Comuzzo, Michele Pasqua and Marino Miculan
# SPDX-License-Identifier: Apache-2.0

# GNU Make >= 4.3

P_DIR = internal/parser

ANTLR = antlr4
ANTLRFLAGS = -Dlanguage=Go -o '${P_DIR}' -package '$(notdir ${P_DIR})'

TOK = ${P_DIR}/SugaredAbuLexer.interp ${P_DIR}/SugaredAbuLexer.tokens
PTOKS = ${P_DIR}/SugaredAbuParser.interp ${P_DIR}/SugaredAbuParser.tokens
LEX = ${P_DIR}/sugaredabu_lexer.go
PAR = ${P_DIR}/sugaredabu_parser.go ${P_DIR}/sugaredabuparser_listener.go ${P_DIR}/sugaredabuparser_base_listener.go

.PHONY: clean all
.DELETE_ON_ERROR:

all: ${PAR} ${PTOKS} ${LEX} ${TOK}

${PAR} ${PTOKS} &: SugaredAbuParser.g4 ${TOK}
	$(ANTLR) ${ANTLRFLAGS} -listener SugaredAbuParser.g4

${LEX} ${TOK} &: SugaredAbuLexer.g4
	$(ANTLR) ${ANTLRFLAGS} $^

clean:
	rm -f ${PAR} ${PTOKS} ${LEX} ${TOK}
