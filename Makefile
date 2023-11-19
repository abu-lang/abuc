# Copyright 2023 Massimo Comuzzo, Michele Pasqua and Marino Miculan
# SPDX-License-Identifier: Apache-2.0

# GNU Make >= 4.3

P_DIR = parser

ANTLR = antlr4
ANTLRFLAGS = -Dlanguage=Go -o '${P_DIR}' -package '$(notdir ${P_DIR})'

LEX_GRAMMAR = AbuLexer.g4
PAR_GRAMMAR = AbuParser.g4

tokens = $(addprefix ${P_DIR}/,$(addsuffix .interp,$(basename $(1)))  $(addsuffix .tokens,$(basename $(1))))

LTOK = $(call tokens,${LEX_GRAMMAR})
PTOK = $(call tokens,${PAR_GRAMMAR})
LEX = ${P_DIR}/abu_lexer.go
PAR = ${P_DIR}/abu_parser.go ${P_DIR}/abuparser_listener.go ${P_DIR}/abuparser_base_listener.go

.PHONY: clean all
.DELETE_ON_ERROR:

all: ${PAR} ${PTOK} ${LEX} ${LTOK}

${PAR} ${PTOK} &: ${PAR_GRAMMAR} ${LTOK}
	$(ANTLR) ${ANTLRFLAGS} -listener ${PAR_GRAMMAR}

${LEX} ${LTOK} &: ${LEX_GRAMMAR}
	$(ANTLR) ${ANTLRFLAGS} $^

clean:
	rm -f ${PAR} ${PTOK} ${LEX} ${LTOK}
