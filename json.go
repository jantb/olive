package main

var json = TmLanguage{ScopeName: "source.json", Name: "JSON (Javascript Next)", Comment: "", FileTypes: []string{"json", "sublime-settings", "sublime-menu", "sublime-keymap", "sublime-mousemap", "sublime-theme", "sublime-build", "sublime-project", "sublime-completions"}, FoldingStartMarker: "(?x)       # turn on extended mode\n^        # a line beginning with\n\\s*      # some optional space\n[{\\[]    # the start of an object or array\n(?!      # but not followed by\n.*     # whatever\n[}\\]]  # and the close of an object or array\n,?     # an optional comma\n\\s*    # some optional space\n$      # at the end of the line\n)\n|        # ...or...\n[{\\[]    # the start of an object or array\n\\s*      # some optional space\n$        # at the end of the line", FoldingStopMarker: "(?x)     # turn on extended mode\n^      # a line beginning with\n\\s*    # some optional space\n[}\\]]  # and the close of an object or array", Patterns: []TmLanguagePatterns{TmLanguagePatterns{Comment: "", Match: "", Captures: TmLanguageCaptures{Num0: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num1: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num2: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num3: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num4: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num5: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}}, Name: "", Begin: "", BeginCaptures: TmLanguageCaptures{Num0: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num1: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num2: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num3: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num4: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num5: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}}, Patterns: []TmLanguagePattern(nil), End: "", EndCaptures: TmLanguageCaptures{Num0: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num1: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num2: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num3: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num4: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}, Num5: TmLanguageName{Name: "", Patterns: []TmLanguagePatternsNameMatchInclude(nil)}}, Include: "#value"}}, Repository: TmLanguageRepository{Brackets: struct {
	Patterns []TmLanguagePatterns "json:\"patterns\""
}{Patterns: []TmLanguagePatterns(nil)}, Delimiters: struct {
	Patterns []TmLanguagePatterns "json:\"patterns\""
}{Patterns: []TmLanguagePatterns(nil)}, Keywords: struct {
	Patterns []TmLanguagePatterns "json:\"patterns\""
}{Patterns: []TmLanguagePatterns(nil)}, Operators: struct {
	Patterns []TmLanguagePatterns "json:\"patterns\""
}{Patterns: []TmLanguagePatterns(nil)}, Runes: struct {
	Patterns []TmLanguagePatterns "json:\"patterns\""
}{Patterns: []TmLanguagePatterns(nil)}, StorageTypes: struct {
	Patterns []TmLanguagePatterns "json:\"patterns\""
}{Patterns: []TmLanguagePatterns(nil)}, StringEscapedChar: struct {
	Patterns []TmLanguagePatterns "json:\"patterns\""
}{Patterns: []TmLanguagePatterns(nil)}, StringPlaceholder: struct {
	Patterns []TmLanguagePatterns "json:\"patterns\""
}{Patterns: []TmLanguagePatterns(nil)}, Variables: struct {
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
}{Comment: "", Patterns: []struct {
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
}(nil)}}, Version: ""}
