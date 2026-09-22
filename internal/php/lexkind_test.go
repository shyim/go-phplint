package php

import "testing"

func TestGeneratedParserSettings(t *testing.T) {
	if yyInitialStackSize < 128 {
		t.Fatalf("parser stack starts at %d entries", yyInitialStackSize)
	}
	if !yyErrorVerbose {
		t.Fatal("syntax errors are not verbose")
	}
}

func TestKeywordBucketsMatchTextLength(t *testing.T) {
	for length, words := range keywordsByLen {
		for _, word := range words {
			if len(word.text) != length {
				t.Fatalf("keyword %q is in the length %d bucket", word.text, length)
			}
		}
	}
}

func TestLiteralFitsInt(t *testing.T) {
	if !literalFitsInt([]byte("1_000"), 10) {
		t.Fatal("1_000 should fit")
	}
	if literalFitsInt([]byte("9223372036854775808"), 10) {
		t.Fatal("2^63 should not fit in a signed int")
	}
	if !literalFitsInt([]byte("9223372036854775807"), 10) {
		t.Fatal("max int64 should fit")
	}
	if literalFitsInt([]byte("8"), 8) {
		t.Fatal("digit 8 is not octal")
	}
}
