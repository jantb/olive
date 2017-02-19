package main

var golang = TmLanguage{ScopeName: "source.go", Name: "Go", Comment: "Go language", FileTypes: []string{"go"}, FoldingStartMarker: "({|\\()\\s*$", FoldingStopMarker: "(}|\\))\\s*$", Patterns: []struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Block comments", Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "comment.block.go", Begin: "/\\*", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "\\*/", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Line comments", Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "comment.line.double-slash.go", Begin: "//", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "$", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Interpreted string literals", Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "string.quoted.double.go", Begin: "\"", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}{struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}{Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
	Num2 struct {
		Name string "json:\"name\""
	} "json:\"2\""
	Num3 struct {
		Name string "json:\"name\""
	} "json:\"3\""
	Num4 struct {
		Name string "json:\"name\""
	} "json:\"4\""
	Num5 struct {
		Name string "json:\"name\""
	} "json:\"5\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}, Num2: struct {
	Name string "json:\"name\""
}{Name: ""}, Num3: struct {
	Name string "json:\"name\""
}{Name: ""}, Num4: struct {
	Name string "json:\"name\""
}{Name: ""}, Num5: struct {
	Name string "json:\"name\""
}{Name: ""}}}, struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}{Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
	Num2 struct {
		Name string "json:\"name\""
	} "json:\"2\""
	Num3 struct {
		Name string "json:\"name\""
	} "json:\"3\""
	Num4 struct {
		Name string "json:\"name\""
	} "json:\"4\""
	Num5 struct {
		Name string "json:\"name\""
	} "json:\"5\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}, Num2: struct {
	Name string "json:\"name\""
}{Name: ""}, Num3: struct {
	Name string "json:\"name\""
}{Name: ""}, Num4: struct {
	Name string "json:\"name\""
}{Name: ""}, Num5: struct {
	Name string "json:\"name\""
}{Name: ""}}}}, End: "\"", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: "punctuation.definition.string.end.go"}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Raw string literals", Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "string.quoted.raw.go", Begin: "`", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}{struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}{Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
	Num2 struct {
		Name string "json:\"name\""
	} "json:\"2\""
	Num3 struct {
		Name string "json:\"name\""
	} "json:\"3\""
	Num4 struct {
		Name string "json:\"name\""
	} "json:\"4\""
	Num5 struct {
		Name string "json:\"name\""
	} "json:\"5\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}, Num2: struct {
	Name string "json:\"name\""
}{Name: ""}, Num3: struct {
	Name string "json:\"name\""
}{Name: ""}, Num4: struct {
	Name string "json:\"name\""
}{Name: ""}, Num5: struct {
	Name string "json:\"name\""
}{Name: ""}}}}, End: "`", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: "punctuation.definition.string.end.go"}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Syntax error receiving channels", Match: "<\\-([\\t ]+)chan\\b", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: "invalid.illegal.receive-channel.go"}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Syntax error sending channels", Match: "\\bchan([\\t ]+)<-", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: "invalid.illegal.send-channel.go"}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Syntax error using slices", Match: "\\[\\](\\s+)", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: "invalid.illegal.slice.go"}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Syntax error numeric literals", Match: "\\b0[0-7]*[89]\\d*\\b", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "invalid.illegal.numeric.go", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Built-in functions", Match: "\\b(append|cap|close|complex|copy|delete|imag|len|make|new|panic|print|println|real|recover)\\b\\(", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "support.function.builtin.go", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Function declarations", Match: "^(\\bfunc\\b)", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: "keyword.function.go"}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Functions", Match: "(\\bfunc\\b)|([a-zA-Z_]\\w*)(\\()", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: "keyword.function.go"}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Floating-point literals", Match: "(\\.\\d+([Ee][-+]\\d+)?i?)\\b|\\b\\d+\\.\\d*(([Ee][-+]\\d+)?i?\\b)?", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "constant.numeric.floating-point.go", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Integers", Match: "\\b((0x[0-9a-fA-F]+)|(0[0-7]+i?)|(\\d+([Ee]\\d+)?i?)|(\\d+[Ee][-+]\\d+i?))\\b", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "constant.numeric.integer.go", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Language constants", Match: "\\b(true|false|nil|iota)\\b", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "constant.language.go", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Package declarations", Match: "(package)\\s+([a-zA-Z_]\\w*)", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: "entity.name.package.go"}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Single line import declarations", Match: "(import)(\\s+((!\\s+\")[^\\s]*)?\\s*)((\")([^\"]*)(\"))", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Multiline import declarations", Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "", Begin: "(?<=import)\\s+(\\()", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: "punctuation.other.bracket.round.go"}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}{struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}{Match: "((?!\\s+\")[^\\s]*)?\\s+((\")([^\"]*)(\"))", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
	Num2 struct {
		Name string "json:\"name\""
	} "json:\"2\""
	Num3 struct {
		Name string "json:\"name\""
	} "json:\"3\""
	Num4 struct {
		Name string "json:\"name\""
	} "json:\"4\""
	Num5 struct {
		Name string "json:\"name\""
	} "json:\"5\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: "entity.alias.import.go"}, Num2: struct {
	Name string "json:\"name\""
}{Name: "string.quoted.double.go"}, Num3: struct {
	Name string "json:\"name\""
}{Name: "punctuation.definition.string.begin.go"}, Num4: struct {
	Name string "json:\"name\""
}{Name: "entity.name.import.go"}, Num5: struct {
	Name string "json:\"name\""
}{Name: "punctuation.definition.string.end.go"}}}}, End: "\\)", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: "punctuation.other.bracket.round.go"}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Type declarations", Match: "(type)\\s+([a-zA-Z_]\\w*)", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: "entity.name.type.go"}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Shorthand variable declaration and assignments", Match: "[a-zA-Z_]\\w*(?:,\\s*[a-zA-Z_]\\w*)*(\\s*)", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Assignments to existing variables", Match: "(!var )\\s*([a-zA-Z_]\\w*(?:,\\s*[a-zA-Z_]\\w*)*)(\\s*=[^=])", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Single line variable declarations/assignments", Match: "(var)\\s+(.*)$", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Multiline variable declarations/assignments", Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "", Begin: "(\\bvar\\b)\\s+(\\()", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: "keyword.var.go"}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}{struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}{Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
	Num2 struct {
		Name string "json:\"name\""
	} "json:\"2\""
	Num3 struct {
		Name string "json:\"name\""
	} "json:\"3\""
	Num4 struct {
		Name string "json:\"name\""
	} "json:\"4\""
	Num5 struct {
		Name string "json:\"name\""
	} "json:\"5\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}, Num2: struct {
	Name string "json:\"name\""
}{Name: ""}, Num3: struct {
	Name string "json:\"name\""
}{Name: ""}, Num4: struct {
	Name string "json:\"name\""
}{Name: ""}, Num5: struct {
	Name string "json:\"name\""
}{Name: ""}}}, struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}{Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
	Num2 struct {
		Name string "json:\"name\""
	} "json:\"2\""
	Num3 struct {
		Name string "json:\"name\""
	} "json:\"3\""
	Num4 struct {
		Name string "json:\"name\""
	} "json:\"4\""
	Num5 struct {
		Name string "json:\"name\""
	} "json:\"5\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}, Num2: struct {
	Name string "json:\"name\""
}{Name: ""}, Num3: struct {
	Name string "json:\"name\""
}{Name: ""}, Num4: struct {
	Name string "json:\"name\""
}{Name: ""}, Num5: struct {
	Name string "json:\"name\""
}{Name: ""}}}}, End: "\\)", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: "punctuation.other.bracket.round.go"}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "Terminators", Match: ";", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "punctuation.terminator.go", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: ""}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "", Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: "#brackets"}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "", Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: "#delimiters"}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "", Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: "#keywords"}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "", Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: "#operators"}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "", Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: "#runes"}, struct {
	Comment  string "json:\"comment,omitempty\""
	Match    string "json:\"match,omitempty\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"captures,omitempty\""
	Name          string "json:\"name,omitempty\""
	Begin         string "json:\"begin,omitempty\""
	BeginCaptures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
	} "json:\"beginCaptures,omitempty\""
	Patterns []struct {
		Match    string "json:\"match\""
		Captures struct {
			Num1 struct {
				Name string "json:\"name\""
			} "json:\"1\""
			Num2 struct {
				Name string "json:\"name\""
			} "json:\"2\""
			Num3 struct {
				Name string "json:\"name\""
			} "json:\"3\""
			Num4 struct {
				Name string "json:\"name\""
			} "json:\"4\""
			Num5 struct {
				Name string "json:\"name\""
			} "json:\"5\""
		} "json:\"captures\""
	} "json:\"patterns,omitempty\""
	End         string "json:\"end,omitempty\""
	EndCaptures struct {
		Num0 struct {
			Name string "json:\"name\""
		} "json:\"0\""
	} "json:\"endCaptures,omitempty\""
	Include string "json:\"include,omitempty\""
}{Comment: "", Match: "", Captures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Name: "", Begin: "", BeginCaptures: struct {
	Num1 struct {
		Name string "json:\"name\""
	} "json:\"1\""
}{Num1: struct {
	Name string "json:\"name\""
}{Name: ""}}, Patterns: []struct {
	Match    string "json:\"match\""
	Captures struct {
		Num1 struct {
			Name string "json:\"name\""
		} "json:\"1\""
		Num2 struct {
			Name string "json:\"name\""
		} "json:\"2\""
		Num3 struct {
			Name string "json:\"name\""
		} "json:\"3\""
		Num4 struct {
			Name string "json:\"name\""
		} "json:\"4\""
		Num5 struct {
			Name string "json:\"name\""
		} "json:\"5\""
	} "json:\"captures\""
}(nil), End: "", EndCaptures: struct {
	Num0 struct {
		Name string "json:\"name\""
	} "json:\"0\""
}{Num0: struct {
	Name string "json:\"name\""
}{Name: ""}}, Include: "#storage_types"}}, Repository: struct {
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
}{Include: "$self"}}}}}}}}, Version: "https://github.com/atom/language-go/commit/d941ce3155b500e65b4d7fbc53ea51b9c92ec1cb"}
