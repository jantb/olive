package main

var golang = TmLanguage{
	ScopeName: "source.go", Name: "Go", Comment: "Go language", FileTypes: []string{"go"}, FoldingStartMarker: "({|\\()\\s*$", FoldingStopMarker: "(}|\\))\\s*$", Patterns: []struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Block comments", Begin: "/\\*", End: "\\*/", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: "punctuation.definition.comment.go"}}, Name: "comment.block.go", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Line comments", Begin: "//", End: "$", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "comment.line.double-slash.go", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: "punctuation.definition.comment.go"}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Interpreted string literals", Begin: "\"", End: "\"", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "string.quoted.double.go", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: "punctuation.definition.string.begin.go"}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: "punctuation.definition.string.end.go"}}, Patterns: []struct {
		Include string "json:\"include\""
	}{struct {
		Include string "json:\"include\""
	}{Include: "#string_escaped_char"}, struct {
		Include string "json:\"include\""
	}{Include: "#string_placeholder"}}, Match: "", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Raw string literals", Begin: "`", End: "`", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "string.quoted.raw.go", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: "punctuation.definition.string.begin.go"}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: "punctuation.definition.string.end.go"}}, Patterns: []struct {
		Include string "json:\"include\""
	}{struct {
		Include string "json:\"include\""
	}{Include: "#string_placeholder"}}, Match: "", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Syntax error receiving channels", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "<\\-([\\t ]+)chan\\b", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Syntax error sending channels", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "\\bchan([\\t ]+)<-", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Syntax error using slices", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "\\[\\](\\s+)", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Syntax error numeric literals", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "invalid.illegal.numeric.go", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "\\b0[0-7]*[89]\\d*\\b", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Built-in functions", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "support.function.builtin.go", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "\\b(append|cap|close|complex|copy|delete|imag|len|make|new|panic|print|println|real|recover)\\b(?=\\()", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Function declarations", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "^(\\bfunc\\b)(?:\\s+(\\([^\\)]+\\)\\s+)?([a-zA-Z_]\\w*)(?=\\())?", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Functions", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "(\\bfunc\\b)|([a-zA-Z_]\\w*)(?=\\()", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Floating-point literals", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "constant.numeric.floating-point.go", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "(\\.\\d+([Ee][-+]\\d+)?i?)\\b|\\b\\d+\\.\\d*(([Ee][-+]\\d+)?i?\\b)?", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Integers", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "constant.numeric.integer.go", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "\\b((0x[0-9a-fA-F]+)|(0[0-7]+i?)|(\\d+([Ee]\\d+)?i?)|(\\d+[Ee][-+]\\d+i?))\\b", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Language constants", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "constant.language.go", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "\\b(true|false|nil|iota)\\b", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Package declarations", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "(?<=package)\\s+([a-zA-Z_]\\w*)", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Type declarations", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "(?<=type)\\s+([a-zA-Z_]\\w*)", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Shorthand variable declaration and assignments", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "[a-zA-Z_]\\w*(?:,\\s*[a-zA-Z_]\\w*)*(?=\\s*:=)", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Assignments to existing variables", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "(?<!var )\\s*([a-zA-Z_]\\w*(?:,\\s*[a-zA-Z_]\\w*)*)(?=\\s*=[^=])", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Single line variable declarations/assignments", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "(?<=var)\\s+(.*)$", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Multiline variable declarations/assignments", Begin: "(\\bvar\\b)\\s+(\\()", End: "\\)", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: "punctuation.other.bracket.round.go"}}, Patterns: []struct {
		Include string "json:\"include\""
	}{struct {
		Include string "json:\"include\""
	}{Include: "#variables"}, struct {
		Include string "json:\"include\""
	}{Include: "$self"}}, Match: "", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "Terminators", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "punctuation.terminator.go", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: ";", Include: ""}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "", Include: "#brackets"}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "", Include: "#delimiters"}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "", Include: "#keywords"}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "", Include: "#operators"}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "", Include: "#runes"}, struct {
		Comment  string "json:\"comment,omitempty\""
		Begin    string "json:\"begin,omitempty\""
		End      string "json:\"end,omitempty\""
		Captures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"captures,omitempty\""
		Name          string "json:\"name,omitempty\""
		BeginCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"beginCaptures,omitempty\""
		EndCaptures struct {
			Num0 struct {
				Name string "json:\"name\""
			} "json:\"0\""
		} "json:\"endCaptures,omitempty\""
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns,omitempty\""
		Match   string "json:\"match,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Comment: "", Begin: "", End: "", Captures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Name: "", BeginCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, EndCaptures: struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	}{Num0: struct {
		Name string "json:\"name\""
	}{Name: ""}}, Patterns: []struct {
		Include string "json:\"include\""
	}(nil), Match: "", Include: "#storage_types"}}, Repository: struct {
		Brackets struct {
			Patterns []struct {
				Match string "json:\"match\""
				Name  string "json:\"name\""
			} "json:\"patterns\""
		} "json:\"brackets\""
		Delimiters struct {
			Patterns []struct {
				Match string "json:\"match\""
				Name  string "json:\"name\""
			} "json:\"patterns\""
		} "json:\"delimiters\""
		Keywords struct {
			Patterns []struct {
				Comment string "json:\"comment,omitempty\""
				Match   string "json:\"match\""
				Name    string "json:\"name\""
			} "json:\"patterns\""
		} "json:\"keywords\""
		Operators struct {
			Comment  string "json:\"comment\""
			Patterns []struct {
				Match string "json:\"match\""
				Name  string "json:\"name\""
			} "json:\"patterns\""
		} "json:\"operators\""
		Runes struct {
			Patterns []struct {
				Match string "json:\"match\""
				Name  string "json:\"name\""
			} "json:\"patterns\""
		} "json:\"runes\""
		StorageTypes struct {
			Patterns []struct {
				Match string "json:\"match\""
				Name  string "json:\"name\""
			} "json:\"patterns\""
		} "json:\"storage_types\""
		StringEscapedChar struct {
			Patterns []struct {
				Match string "json:\"match\""
				Name  string "json:\"name\""
			} "json:\"patterns\""
		} "json:\"string_escaped_char\""
		StringPlaceholder struct {
			Patterns []struct {
				Match string "json:\"match\""
				Name  string "json:\"name\""
			} "json:\"patterns\""
		} "json:\"string_placeholder\""
		Variables struct {
			Comment  string "json:\"comment\""
			Patterns []struct {
				Match    string "json:\"match\""
				Captures struct {
					Num1 struct {
						Patterns []struct {
							Match   string "json:\"match,omitempty\""
							Name    string "json:\"name,omitempty\""
							Include string "json:\"include,omitempty\""
						} "json:\"patterns\""
					} "json:\"1\""
					Num2 struct {
						Patterns []struct {
							Include string "json:\"include\""
						} "json:\"patterns\""
					} "json:\"2\""
				} "json:\"captures\""
			} "json:\"patterns\""
		} "json:\"variables\""
	}{Brackets: struct {
		Patterns []struct {
			Match string "json:\"match\""
			Name  string "json:\"name\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\{|\\}", Name: "punctuation.other.bracket.curly.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\(|\\)", Name: "punctuation.other.bracket.round.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\[|\\]", Name: "punctuation.other.bracket.square.go"}}}, Delimiters: struct {
		Patterns []struct {
			Match string "json:\"match\""
			Name  string "json:\"name\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: ",", Name: "punctuation.other.comma.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\.(?!\\.\\.)", Name: "punctuation.other.period.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: ":(?!=)", Name: "punctuation.other.colon.go"}}}, Keywords: struct {
		Patterns []struct {
			Comment string "json:\"comment,omitempty\""
			Match   string "json:\"match\""
			Name    string "json:\"name\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Comment string "json:\"comment,omitempty\""
		Match   string "json:\"match\""
		Name    string "json:\"name\""
	}{struct {
		Comment string "json:\"comment,omitempty\""
		Match   string "json:\"match\""
		Name    string "json:\"name\""
	}{Comment: "Flow control keywords", Match: "\\b(break|case|continue|default|defer|else|fallthrough|for|go|goto|if|range|return|select|switch)\\b", Name: "keyword.control.go"}, struct {
		Comment string "json:\"comment,omitempty\""
		Match   string "json:\"match\""
		Name    string "json:\"name\""
	}{Comment: "", Match: "\\bchan\\b", Name: "keyword.channel.go"}, struct {
		Comment string "json:\"comment,omitempty\""
		Match   string "json:\"match\""
		Name    string "json:\"name\""
	}{Comment: "", Match: "\\bconst\\b", Name: "keyword.const.go"}, struct {
		Comment string "json:\"comment,omitempty\""
		Match   string "json:\"match\""
		Name    string "json:\"name\""
	}{Comment: "", Match: "\\bfunc\\b", Name: "keyword.function.go"}, struct {
		Comment string "json:\"comment,omitempty\""
		Match   string "json:\"match\""
		Name    string "json:\"name\""
	}{Comment: "", Match: "\\binterface\\b", Name: "keyword.interface.go"}, struct {
		Comment string "json:\"comment,omitempty\""
		Match   string "json:\"match\""
		Name    string "json:\"name\""
	}{Comment: "", Match: "\\bimport\\b", Name: "keyword.import.go"}, struct {
		Comment string "json:\"comment,omitempty\""
		Match   string "json:\"match\""
		Name    string "json:\"name\""
	}{Comment: "", Match: "\\bmap\\b", Name: "keyword.map.go"}, struct {
		Comment string "json:\"comment,omitempty\""
		Match   string "json:\"match\""
		Name    string "json:\"name\""
	}{Comment: "", Match: "\\bpackage\\b", Name: "keyword.package.go"}, struct {
		Comment string "json:\"comment,omitempty\""
		Match   string "json:\"match\""
		Name    string "json:\"name\""
	}{Comment: "", Match: "\\bstruct\\b", Name: "keyword.struct.go"}, struct {
		Comment string "json:\"comment,omitempty\""
		Match   string "json:\"match\""
		Name    string "json:\"name\""
	}{Comment: "", Match: "\\btype\\b", Name: "keyword.type.go"}, struct {
		Comment string "json:\"comment,omitempty\""
		Match   string "json:\"match\""
		Name    string "json:\"name\""
	}{Comment: "", Match: "\\bvar\\b", Name: "keyword.var.go"}}}, Operators: struct {
		Comment  string "json:\"comment\""
		Patterns []struct {
			Match string "json:\"match\""
			Name  string "json:\"name\""
		} "json:\"patterns\""
	}{Comment: "Note that the order here is very important!", Patterns: []struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "(\\*|&)(?=\\w)", Name: "keyword.operator.address.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "<\\-", Name: "keyword.operator.channel.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\-\\-", Name: "keyword.operator.decrement.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\+\\+", Name: "keyword.operator.increment.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "(==|!=|<=|>=|<[^<]|>[^>])", Name: "keyword.operator.comparison.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "(&&|\\|\\||!)", Name: "keyword.operator.logical.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "(=|\\+=|\\-=|\\|=|\\^=|\\*=|/=|:=|%=|<<=|>>=|&\\^=|&=)", Name: "keyword.operator.assignment.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "(\\+|\\-|\\*|/|%)", Name: "keyword.operator.arithmetic.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "(&(?!\\^)|\\||\\^|&\\^|<<|>>)", Name: "keyword.operator.arithmetic.bitwise.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\.\\.\\.", Name: "keyword.operator.ellipsis.go"}}}, Runes: struct {
		Patterns []struct {
			Match string "json:\"match\""
			Name  string "json:\"name\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\'(\\\\([0-7]{3}|[abfnrtv\\\\'\"]|x[0-9a-fA-F]{2}|u[0-9a-fA-F]{4}|U[0-9a-fA-F]{8})|\\p{Any})\\'", Name: "constant.other.rune.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\'.*\\'", Name: "invalid.illegal.unknown-rune.go"}}}, StorageTypes: struct {
		Patterns []struct {
			Match string "json:\"match\""
			Name  string "json:\"name\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\bbool\\b", Name: "storage.type.boolean.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\bbyte\\b", Name: "storage.type.byte.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\berror\\b", Name: "storage.type.error.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\b(complex(64|128)|float(32|64)|u?int(8|16|32|64)?)\\b", Name: "storage.type.numeric.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\brune\\b", Name: "storage.type.rune.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\bstring\\b", Name: "storage.type.string.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\buintptr\\b", Name: "storage.type.uintptr.go"}}}, StringEscapedChar: struct {
		Patterns []struct {
			Match string "json:\"match\""
			Name  string "json:\"name\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\\\([0-7]{3}|[abfnrtv\\\\'\"]|x[0-9a-fA-F]{2}|u[0-9a-fA-F]{4}|U[0-9a-fA-F]{8})", Name: "constant.character.escape.go"}, struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "\\\\[^0-7xuUabfnrtv\\'\"]", Name: "invalid.illegal.unknown-escape.go"}}}, StringPlaceholder: struct {
		Patterns []struct {
			Match string "json:\"match\""
			Name  string "json:\"name\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{struct {
		Match string "json:\"match\""
		Name  string "json:\"name\""
	}{Match: "%(\\[\\d+\\])?([\\+#\\-0\\x20]{,2}((\\d+|\\*)?(\\.?(\\d+|\\*|(\\[\\d+\\])\\*?)?(\\[\\d+\\])?)?))?[vT%tbcdoqxXUbeEfFgGsp]", Name: "constant.other.placeholder.go"}}}, Variables: struct {
		Comment  string "json:\"comment\""
		Patterns []struct {
			Match    string "json:\"match\""
			Captures struct {
				Num1 struct {
					Patterns []struct {
						Match   string "json:\"match,omitempty\""
						Name    string "json:\"name,omitempty\""
						Include string "json:\"include,omitempty\""
					} "json:\"patterns\""
				} "json:\"1\""
				Num2 struct {
					Patterns []struct {
						Include string "json:\"include\""
					} "json:\"patterns\""
				} "json:\"2\""
			} "json:\"captures\""
		} "json:\"patterns\""
	}{Comment: "First add tests and make sure existing tests still pass when changing anything here!", Patterns: []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Patterns []struct {
					Match   string "json:\"match,omitempty\""
					Name    string "json:\"name,omitempty\""
					Include string "json:\"include,omitempty\""
				} "json:\"patterns\""
			} "json:\"1\""
			Num2 struct {
				Patterns []struct {
					Include string "json:\"include\""
				} "json:\"patterns\""
			} "json:\"2\""
		} "json:\"captures\""
	}{struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Patterns []struct {
					Match   string "json:\"match,omitempty\""
					Name    string "json:\"name,omitempty\""
					Include string "json:\"include,omitempty\""
				} "json:\"patterns\""
			} "json:\"1\""
			Num2 struct {
				Patterns []struct {
					Include string "json:\"include\""
				} "json:\"patterns\""
			} "json:\"2\""
		} "json:\"captures\""
	}{Match: "([a-zA-Z_]\\w*(?:,\\s*[a-zA-Z_]\\w*)*)\\s*(=.*)", Captures: struct {
		Num1 struct {
			Patterns []struct {
				Match   string "json:\"match,omitempty\""
				Name    string "json:\"name,omitempty\""
				Include string "json:\"include,omitempty\""
			} "json:\"patterns\""
		} "json:\"1\""
		Num2 struct {
			Patterns []struct {
				Include string "json:\"include\""
			} "json:\"patterns\""
		} "json:\"2\""
	}{Num1: struct {
		Patterns []struct {
			Match   string "json:\"match,omitempty\""
			Name    string "json:\"name,omitempty\""
			Include string "json:\"include,omitempty\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Match   string "json:\"match,omitempty\""
		Name    string "json:\"name,omitempty\""
		Include string "json:\"include,omitempty\""
	}{struct {
		Match   string "json:\"match,omitempty\""
		Name    string "json:\"name,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Match: "[a-zA-Z_]\\w*", Name: "variable.other.assignment.go", Include: ""}, struct {
		Match   string "json:\"match,omitempty\""
		Name    string "json:\"name,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Match: "", Name: "", Include: "#delimiters"}}}, Num2: struct {
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Include string "json:\"include\""
	}{struct {
		Include string "json:\"include\""
	}{Include: "$self"}}}}}, struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Patterns []struct {
					Match   string "json:\"match,omitempty\""
					Name    string "json:\"name,omitempty\""
					Include string "json:\"include,omitempty\""
				} "json:\"patterns\""
			} "json:\"1\""
			Num2 struct {
				Patterns []struct {
					Include string "json:\"include\""
				} "json:\"patterns\""
			} "json:\"2\""
		} "json:\"captures\""
	}{Match: "([a-zA-Z_]\\w*(?:,\\s*[a-zA-Z_]\\w*)*)(\\s+[\\*]?[a-zA-Z_]\\w*\\s*)(=.*)", Captures: struct {
		Num1 struct {
			Patterns []struct {
				Match   string "json:\"match,omitempty\""
				Name    string "json:\"name,omitempty\""
				Include string "json:\"include,omitempty\""
			} "json:\"patterns\""
		} "json:\"1\""
		Num2 struct {
			Patterns []struct {
				Include string "json:\"include\""
			} "json:\"patterns\""
		} "json:\"2\""
	}{Num1: struct {
		Patterns []struct {
			Match   string "json:\"match,omitempty\""
			Name    string "json:\"name,omitempty\""
			Include string "json:\"include,omitempty\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Match   string "json:\"match,omitempty\""
		Name    string "json:\"name,omitempty\""
		Include string "json:\"include,omitempty\""
	}{struct {
		Match   string "json:\"match,omitempty\""
		Name    string "json:\"name,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Match: "[a-zA-Z_]\\w*", Name: "variable.other.assignment.go", Include: ""}, struct {
		Match   string "json:\"match,omitempty\""
		Name    string "json:\"name,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Match: "", Name: "", Include: "#delimiters"}}}, Num2: struct {
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Include string "json:\"include\""
	}{struct {
		Include string "json:\"include\""
	}{Include: "$self"}}}}}, struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Patterns []struct {
					Match   string "json:\"match,omitempty\""
					Name    string "json:\"name,omitempty\""
					Include string "json:\"include,omitempty\""
				} "json:\"patterns\""
			} "json:\"1\""
			Num2 struct {
				Patterns []struct {
					Include string "json:\"include\""
				} "json:\"patterns\""
			} "json:\"2\""
		} "json:\"captures\""
	}{Match: "([a-zA-Z_]\\w*(?:,\\s*[a-zA-Z_]\\w*)*)(\\s+[\\[\\]\\*]{0,3}[a-zA-Z_]\\w*\\s*[^=].*)", Captures: struct {
		Num1 struct {
			Patterns []struct {
				Match   string "json:\"match,omitempty\""
				Name    string "json:\"name,omitempty\""
				Include string "json:\"include,omitempty\""
			} "json:\"patterns\""
		} "json:\"1\""
		Num2 struct {
			Patterns []struct {
				Include string "json:\"include\""
			} "json:\"patterns\""
		} "json:\"2\""
	}{Num1: struct {
		Patterns []struct {
			Match   string "json:\"match,omitempty\""
			Name    string "json:\"name,omitempty\""
			Include string "json:\"include,omitempty\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Match   string "json:\"match,omitempty\""
		Name    string "json:\"name,omitempty\""
		Include string "json:\"include,omitempty\""
	}{struct {
		Match   string "json:\"match,omitempty\""
		Name    string "json:\"name,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Match: "[a-zA-Z_]\\w*", Name: "variable.other.declaration.go", Include: ""}, struct {
		Match   string "json:\"match,omitempty\""
		Name    string "json:\"name,omitempty\""
		Include string "json:\"include,omitempty\""
	}{Match: "", Name: "", Include: "#delimiters"}}}, Num2: struct {
		Patterns []struct {
			Include string "json:\"include\""
		} "json:\"patterns\""
	}{Patterns: []struct {
		Include string "json:\"include\""
	}{struct {
		Include string "json:\"include\""
	}{Include: "$self"}}}}}}}}, Version: "https://github.com/atom/language-go/commit/93594dfb138a664f0914d54e408527e136899fb2"}
