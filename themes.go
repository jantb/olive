package main

var dark = Theme{Name: "Dark", Settings: []struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "emphasis", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "italic", Foreground: "", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "strong", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "bold", Foreground: "", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "header", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#000080", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "comment", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#608b4e", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "constant.language", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "constant.numeric", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#b5cea8", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "constant.regexp", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#646695", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "constant.rgb-value", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d4d4d4", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.name.tag", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.name.function", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d4d4d4", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.name.class", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d4d4d4", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.name.selector", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d7ba7d", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.other.attribute-name", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#9cdcfe", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.other.attribute-name.css", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d7ba7d", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "invalid", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#f44747", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.underline", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "underline", Foreground: "", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.bold", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "bold", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.heading", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "bold", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.italic", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "italic", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.inserted", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#b5cea8", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.deleted", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#ce9178", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.changed", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.punctuation.quote", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#608b4e", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.punctuation.list", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#6796e6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d4d4d4", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.inline.raw", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#ce9178", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "meta.selector", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d7ba7d", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#808080", Background: ""}, Name: "brackets of XML tags"}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "meta.preprocessor", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "meta.preprocessor.string", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#ce9178", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "meta.preprocessor.numeric", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#b5cea8", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "meta.structure.dictionary.key.python", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#9cdcfe", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "meta.header.diff", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "storage", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "storage.type", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "storage.modifier", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "string", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#ce9178", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "string.tag", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#ce9178", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "string.value", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#ce9178", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "string.regexp", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d16969", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "support.type.property-name", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#9cdcfe", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword.control", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword.control.less", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d7ba7d", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword.operator", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d4d4d4", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword.operator.new", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword.other.unit", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#b5cea8", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "metatag.php", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "support.function.git-rebase", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#9cdcfe", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "constant.sha.git-rebase", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#b5cea8", Background: ""}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "editor.ruler", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#5A5A5A", Background: "#1E1E1E"}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "editor", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#FFFFFF", Background: "#1E1E1E"}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.name.package.go", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#D7BA7D", Background: "#1E1E1E"}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.name.type.go", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#698857", Background: "#1E1E1E"}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword.function.go", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#DCDDA7", Background: "#1E1E1E"}, Name: ""}, struct {
	Scope    string "json:\"scope\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "", Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d4d4d4", Background: ""}, Name: "coloring of the Java import and package identifiers"}}}
