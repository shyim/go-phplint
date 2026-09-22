package php

import "strings"

// humanizeSyntax rewrites yacc token identifiers into source-oriented words.
func humanizeSyntax(message string) string {
	if !strings.Contains(message, "T_") {
		return message
	}
	var b strings.Builder
	b.Grow(len(message))
	for i := 0; i < len(message); {
		if message[i] == 'T' && i+1 < len(message) && message[i+1] == '_' && (i == 0 || !isTokenNameChar(message[i-1])) {
			j := i + 2
			for j < len(message) && isTokenNameChar(message[j]) {
				j++
			}
			name := message[i:j]
			if nice, ok := syntaxNames[name]; ok {
				b.WriteString(nice)
			} else {
				b.WriteString("token")
			}
			i = j
			continue
		}
		b.WriteByte(message[i])
		i++
	}
	return b.String()
}

func isTokenNameChar(c byte) bool {
	return c == '_' || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

var syntaxNames = map[string]string{
	"T_INCLUDE":                             "include",
	"T_INCLUDE_ONCE":                        "include_once",
	"T_EXIT":                                "exit",
	"T_IF":                                  "if",
	"T_LNUMBER":                             "integer",
	"T_DNUMBER":                             "number",
	"T_STRING":                              "identifier",
	"T_STRING_VARNAME":                      "variable name",
	"T_VARIABLE":                            "variable",
	"T_NUM_STRING":                          "number",
	"T_INLINE_HTML":                         "inline HTML",
	"T_CHARACTER":                           "character",
	"T_BAD_CHARACTER":                       "character",
	"T_ENCAPSED_AND_WHITESPACE":             "string",
	"T_CONSTANT_ENCAPSED_STRING":            "string",
	"T_ECHO":                                "echo",
	"T_DO":                                  "do",
	"T_WHILE":                               "while",
	"T_ENDWHILE":                            "endwhile",
	"T_FOR":                                 "for",
	"T_ENDFOR":                              "endfor",
	"T_FOREACH":                             "foreach",
	"T_ENDFOREACH":                          "endforeach",
	"T_DECLARE":                             "declare",
	"T_ENDDECLARE":                          "enddeclare",
	"T_AS":                                  "as",
	"T_SWITCH":                              "switch",
	"T_ENDSWITCH":                           "endswitch",
	"T_CASE":                                "case",
	"T_DEFAULT":                             "default",
	"T_BREAK":                               "break",
	"T_CONTINUE":                            "continue",
	"T_GOTO":                                "goto",
	"T_FUNCTION":                            "function",
	"T_FN":                                  "fn",
	"T_CONST":                               "const",
	"T_RETURN":                              "return",
	"T_TRY":                                 "try",
	"T_CATCH":                               "catch",
	"T_FINALLY":                             "finally",
	"T_THROW":                               "throw",
	"T_USE":                                 "use",
	"T_INSTEADOF":                           "insteadof",
	"T_GLOBAL":                              "global",
	"T_VAR":                                 "var",
	"T_UNSET":                               "unset",
	"T_ISSET":                               "isset",
	"T_EMPTY":                               "empty",
	"T_HALT_COMPILER":                       "__halt_compiler",
	"T_CLASS":                               "class",
	"T_TRAIT":                               "trait",
	"T_INTERFACE":                           "interface",
	"T_EXTENDS":                             "extends",
	"T_IMPLEMENTS":                          "implements",
	"T_OBJECT_OPERATOR":                     "->",
	"T_DOUBLE_ARROW":                        "=>",
	"T_LIST":                                "list",
	"T_ARRAY":                               "array",
	"T_CALLABLE":                            "callable",
	"T_CLASS_C":                             "__CLASS__",
	"T_TRAIT_C":                             "__TRAIT__",
	"T_METHOD_C":                            "__METHOD__",
	"T_FUNC_C":                              "__FUNCTION__",
	"T_LINE":                                "__LINE__",
	"T_FILE":                                "__FILE__",
	"T_COMMENT":                             "comment",
	"T_DOC_COMMENT":                         "comment",
	"T_OPEN_TAG":                            "<?php",
	"T_OPEN_TAG_WITH_ECHO":                  "<?=",
	"T_CLOSE_TAG":                           "?>",
	"T_WHITESPACE":                          "whitespace",
	"T_START_HEREDOC":                       "heredoc",
	"T_END_HEREDOC":                         "heredoc end",
	"T_DOLLAR_OPEN_CURLY_BRACES":            "${",
	"T_CURLY_OPEN":                          "{$",
	"T_PAAMAYIM_NEKUDOTAYIM":                "::",
	"T_NAMESPACE":                           "namespace",
	"T_NS_C":                                "__NAMESPACE__",
	"T_DIR":                                 "__DIR__",
	"T_NS_SEPARATOR":                        "\\",
	"T_ELLIPSIS":                            "...",
	"T_EVAL":                                "eval",
	"T_REQUIRE":                             "require",
	"T_REQUIRE_ONCE":                        "require_once",
	"T_LOGICAL_OR":                          "or",
	"T_LOGICAL_XOR":                         "xor",
	"T_LOGICAL_AND":                         "and",
	"T_INSTANCEOF":                          "instanceof",
	"T_NEW":                                 "new",
	"T_CLONE":                               "clone",
	"T_ELSEIF":                              "elseif",
	"T_ELSE":                                "else",
	"T_ENDIF":                               "endif",
	"T_PRINT":                               "print",
	"T_YIELD":                               "yield",
	"T_STATIC":                              "static",
	"T_ABSTRACT":                            "abstract",
	"T_FINAL":                               "final",
	"T_PRIVATE":                             "private",
	"T_PROTECTED":                           "protected",
	"T_PUBLIC":                              "public",
	"T_INC":                                 "++",
	"T_DEC":                                 "--",
	"T_YIELD_FROM":                          "yield from",
	"T_INT_CAST":                            "(int)",
	"T_DOUBLE_CAST":                         "(float)",
	"T_STRING_CAST":                         "(string)",
	"T_ARRAY_CAST":                          "(array)",
	"T_OBJECT_CAST":                         "(object)",
	"T_BOOL_CAST":                           "(bool)",
	"T_UNSET_CAST":                          "(unset)",
	"T_VOID_CAST":                           "(void)",
	"T_COALESCE":                            "??",
	"T_SPACESHIP":                           "<=>",
	"T_NOELSE":                              "else",
	"T_PLUS_EQUAL":                          "+=",
	"T_MINUS_EQUAL":                         "-=",
	"T_MUL_EQUAL":                           "*=",
	"T_POW_EQUAL":                           "**=",
	"T_DIV_EQUAL":                           "/=",
	"T_CONCAT_EQUAL":                        ".=",
	"T_MOD_EQUAL":                           "%=",
	"T_AND_EQUAL":                           "&=",
	"T_OR_EQUAL":                            "|=",
	"T_XOR_EQUAL":                           "^=",
	"T_SL_EQUAL":                            "<<=",
	"T_SR_EQUAL":                            ">>=",
	"T_COALESCE_EQUAL":                      "??=",
	"T_BOOLEAN_OR":                          "||",
	"T_BOOLEAN_AND":                         "&&",
	"T_POW":                                 "**",
	"T_SL":                                  "<<",
	"T_SR":                                  ">>",
	"T_IS_IDENTICAL":                        "===",
	"T_IS_NOT_IDENTICAL":                    "!==",
	"T_IS_EQUAL":                            "==",
	"T_IS_NOT_EQUAL":                        "!=",
	"T_IS_SMALLER_OR_EQUAL":                 "<=",
	"T_IS_GREATER_OR_EQUAL":                 ">=",
	"T_NULLSAFE_OBJECT_OPERATOR":            "?->",
	"T_MATCH":                               "match",
	"T_ATTRIBUTE":                           "#[",
	"T_NAME_RELATIVE":                       "name",
	"T_NAME_QUALIFIED":                      "name",
	"T_NAME_FULLY_QUALIFIED":                "name",
	"T_READONLY":                            "readonly",
	"T_ENUM":                                "enum",
	"T_AMPERSAND_FOLLOWED_BY_VAR_OR_VARARG": "&",
	"T_AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG": "&",
	"T_PIPE":           "|>",
	"T_PUBLIC_SET":     "public(set)",
	"T_PROTECTED_SET":  "protected(set)",
	"T_PRIVATE_SET":    "private(set)",
	"T_PROPERTY_HOOKS": "{",
}
