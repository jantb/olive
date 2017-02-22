package main

var dark = Theme{Name: "Dark", Settings: []struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "", Scopes: []string{"entity.name.function", "support.function", "support.constant.handlebars"}, Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#DCDCAA", Background: ""}, Name: "Function declarations"}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "", Scopes: []string{"meta.return-type", "support.class", "support.type", "entity.name.type", "entity.name.class", "entity.name.package", "storage.type.cs", "storage.type.generic.cs", "storage.type.modifier.cs", "storage.type.variable.cs", "storage.type.annotation.java", "storage.type.generic.java", "storage.type.java", "storage.type.object.array.java", "storage.type.primitive.array.java", "storage.type.primitive.java", "storage.type.token.java", "storage.type.groovy", "storage.type.annotation.groovy", "storage.type.parameters.groovy", "storage.type.generic.groovy", "storage.type.object.array.groovy", "storage.type.primitive.array.groovy", "storage.type.primitive.groovy"}, Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#4EC9B0", Background: ""}, Name: "Types declaration and references"}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "", Scopes: []string{"meta.return.type", "meta.type.cast.expr", "meta.type.new.expr", "support.constant.math", "support.constant.dom", "support.constant.json", "entity.other.inherited-class"}, Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#4EC9B0", Background: ""}, Name: "Types declaration and references, TS grammar specific"}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword.control", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#C586C0", Background: ""}, Name: "Control flow keywords"}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "", Scopes: []string{"variable", "meta.definition.variable.name", "support.variable"}, Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#9CDCFE", Background: ""}, Name: "Variable and parameter name"}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "", Scopes: []string{"meta.object-literal.key", "meta.object-literal.key entity.name.function"}, Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#9CDCFE", Background: ""}, Name: "Object keys, TS grammar specific"}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "", Scopes: []string{"support.constant.property-value", "support.constant.font-name", "support.constant.media-type", "support.constant.media", "constant.other.color.rgb-value", "constant.other.rgb-value", "support.constant.color"}, Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#CE9178", Background: ""}, Name: "CSS property value"}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "emphasis", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "italic", Foreground: "", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "strong", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "bold", Foreground: "", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "header", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#000080", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "comment", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#608b4e", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "constant.language", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "constant.numeric", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#b5cea8", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "constant.regexp", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#646695", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "constant.rgb-value", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d4d4d4", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.name.tag", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.name.function", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d4d4d4", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.name.class", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d4d4d4", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.name.selector", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d7ba7d", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.other.attribute-name", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#9cdcfe", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "entity.other.attribute-name.css", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d7ba7d", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "invalid", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#f44747", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.underline", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "underline", Foreground: "", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.bold", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "bold", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.heading", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "bold", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.italic", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "italic", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.inserted", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#b5cea8", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.deleted", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#ce9178", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.changed", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.punctuation.quote", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#608b4e", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.punctuation.list", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#6796e6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "", Scopes: []string{"markup.quote", "markup.list"}, Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d4d4d4", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "markup.inline.raw", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#ce9178", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "meta.selector", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d7ba7d", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "", Scopes: []string{"meta.tag", "punctuation.tag.js", "punctuation.tag.tsx"}, Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#808080", Background: ""}, Name: "brackets of XML tags"}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "meta.preprocessor", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "meta.preprocessor.string", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#ce9178", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "meta.preprocessor.numeric", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#b5cea8", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "meta.structure.dictionary.key.python", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#9cdcfe", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "meta.header.diff", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "storage", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "storage.type", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "storage.modifier", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "string", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#ce9178", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "string.tag", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#ce9178", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "string.value", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#ce9178", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "string.regexp", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d16969", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "support.type.property-name", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#9cdcfe", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword.control", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword.control.less", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d7ba7d", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword.operator", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d4d4d4", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword.operator.new", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "keyword.other.unit", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#b5cea8", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "metatag.php", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#569cd6", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "support.function.git-rebase", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#9cdcfe", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "constant.sha.git-rebase", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#b5cea8", Background: ""}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "editor.ruler", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#5A5A5A", Background: "#1E1E1E"}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "editor", Scopes: []string(nil), Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#FFFFFF", Background: "#1E1E1E"}, Name: ""}, struct {
	Scope    string   "json:\"scope\""
	Scopes   []string "json:\"scopes\""
	Settings struct {
		FontStyle  string "json:\"fontStyle,omitempty\""
		Foreground string "json:\"foreground,omitempty\""
		Background string "json:\"background,omitempty\""
	} "json:\"settings\""
	Name string "json:\"name,omitempty\""
}{Scope: "", Scopes: []string{"storage.modifier.import.java", "storage.modifier.package.java"}, Settings: struct {
	FontStyle  string "json:\"fontStyle,omitempty\""
	Foreground string "json:\"foreground,omitempty\""
	Background string "json:\"background,omitempty\""
}{FontStyle: "", Foreground: "#d4d4d4", Background: ""}, Name: "coloring of the Java import and package identifiers"}}}
