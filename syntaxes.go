package main

var syntaxes = []map[string]interface{}{map[string]interface{}{"name": "Go", "comment": "Go language", "fileTypes": []interface{}{"go"}, "foldingStartMarker": "({|\\()\\s*$", "foldingStopMarker": "(}|\\))\\s*$", "patterns": []interface{}{map[string]interface{}{"include": "#comments"}, map[string]interface{}{"end": "\"", "endCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.string.end.go"}}, "name": "string.quoted.double.go", "patterns": []interface{}{map[string]interface{}{"include": "#string_escaped_char"}, map[string]interface{}{"include": "#string_placeholder"}}, "comment": "Interpreted string literals", "begin": "\"", "beginCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.string.begin.go"}}}, map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"include": "#string_placeholder"}}, "comment": "Raw string literals", "begin": "`", "beginCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.string.begin.go"}}, "end": "`", "endCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.string.end.go"}}, "name": "string.quoted.raw.go"}, map[string]interface{}{"match": "<\\-([\\t ]+)chan\\b", "captures": map[string]interface{}{"1": map[string]interface{}{"name": "invalid.illegal.receive-channel.go"}}, "comment": "Syntax error receiving channels"}, map[string]interface{}{"comment": "Syntax error sending channels", "match": "\\bchan([\\t ]+)<-", "captures": map[string]interface{}{"1": map[string]interface{}{"name": "invalid.illegal.send-channel.go"}}}, map[string]interface{}{"match": "\\[\\](\\s+)", "captures": map[string]interface{}{"1": map[string]interface{}{"name": "invalid.illegal.slice.go"}}, "comment": "Syntax error using slices"}, map[string]interface{}{"comment": "Syntax error numeric literals", "match": "\\b0[0-7]*[89]\\d*\\b", "name": "invalid.illegal.numeric.go"}, map[string]interface{}{"comment": "Built-in functions", "match": "\\b(append|cap|close|complex|copy|delete|imag|len|make|new|panic|print|println|real|recover)\\b(\\()", "name": "support.function.builtin.go"}, map[string]interface{}{"comment": "Function declarations", "match": "^(\\bfunc\\b)(?:\\s+(\\([^\\)]+\\)\\s+)?([a-zA-Z_]\\w*)(\\())?", "captures": map[string]interface{}{"1": map[string]interface{}{"name": "keyword.function.go"}, "2": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"include": "#brackets"}, map[string]interface{}{"include": "#operators"}}}, "3": map[string]interface{}{"name": "entity.name.function.go"}}}, map[string]interface{}{"comment": "Functions", "match": "(\\bfunc\\b)|([a-zA-Z_]\\w*)(\\()", "captures": map[string]interface{}{"1": map[string]interface{}{"name": "keyword.function.go"}, "2": map[string]interface{}{"name": "support.function.go"}}}, map[string]interface{}{"comment": "Floating-point literals", "match": "(\\.\\d+([Ee][-+]\\d+)?i?)\\b|\\b\\d+\\.\\d*(([Ee][-+]\\d+)?i?\\b)?", "name": "constant.numeric.floating-point.go"}, map[string]interface{}{"match": "\\b((0x[0-9a-fA-F]+)|(0[0-7]+i?)|(\\d+([Ee]\\d+)?i?)|(\\d+[Ee][-+]\\d+i?))\\b", "name": "constant.numeric.integer.go", "comment": "Integers"}, map[string]interface{}{"name": "constant.language.go", "comment": "Language constants", "match": "\\b(true|false|nil|iota)\\b"}, map[string]interface{}{"comment": "Package declarations", "match": "(=package)\\s+([a-zA-Z_]\\w*)", "captures": map[string]interface{}{"1": map[string]interface{}{"name": "entity.name.package.go"}}}, map[string]interface{}{"comment": "Single line import declarations", "match": "(=import)(\\s+((\\s+\")[^\\s]*)?\\s*)((\")([^\"]*)(\"))", "captures": map[string]interface{}{"6": map[string]interface{}{"name": "punctuation.definition.string.end.go"}, "2": map[string]interface{}{"name": "entity.alias.import.go"}, "3": map[string]interface{}{"name": "string.quoted.double.go"}, "4": map[string]interface{}{"name": "punctuation.definition.string.begin.go"}, "5": map[string]interface{}{"name": "entity.name.import.go"}}}, map[string]interface{}{"beginCaptures": map[string]interface{}{"1": map[string]interface{}{"name": "punctuation.other.bracket.round.go"}}, "patterns": []interface{}{map[string]interface{}{"match": "((\\s+\")[^\\s]*)?\\s+((\")([^\"]*)(\"))", "captures": map[string]interface{}{"1": map[string]interface{}{"name": "entity.alias.import.go"}, "2": map[string]interface{}{"name": "string.quoted.double.go"}, "3": map[string]interface{}{"name": "punctuation.definition.string.begin.go"}, "4": map[string]interface{}{"name": "entity.name.import.go"}, "5": map[string]interface{}{"name": "punctuation.definition.string.end.go"}}}, map[string]interface{}{"include": "#comments"}}, "end": "\\)", "endCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.other.bracket.round.go"}}, "comment": "Multiline import declarations", "begin": "(=import)\\s+(\\()"}, map[string]interface{}{"comment": "Type declarations", "match": "(=type)\\s+([a-zA-Z_]\\w*)", "captures": map[string]interface{}{"1": map[string]interface{}{"name": "entity.name.type.go"}}}, map[string]interface{}{"captures": map[string]interface{}{"0": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"match": "[a-zA-Z_]\\w*", "name": "variable.other.assignment.go"}, map[string]interface{}{"include": "#delimiters"}}}}, "comment": "Shorthand variable declaration and assignments", "match": "[a-zA-Z_]\\w*(?:,\\s*[a-zA-Z_]\\w*)*(\\s*:=)"}, map[string]interface{}{"comment": "Assignments to existing variables", "match": "(!var )\\s*([a-zA-Z_]\\w*(?:,\\s*[a-zA-Z_]\\w*)*)(\\s*=[^=])", "captures": map[string]interface{}{"1": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"match": "[a-zA-Z_]\\w*", "name": "variable.other.assignment.go"}, map[string]interface{}{"include": "#delimiters"}}}}}, map[string]interface{}{"comment": "Single line variable declarations/assignments", "match": "(=var)\\s+(.*)$", "captures": map[string]interface{}{"1": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"include": "#variables"}}}}}, map[string]interface{}{"end": "\\)", "endCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.other.bracket.round.go"}}, "patterns": []interface{}{map[string]interface{}{"include": "#variables"}, map[string]interface{}{"include": "$self"}}, "comment": "Multiline variable declarations/assignments", "begin": "(\\bvar\\b)\\s+(\\()", "beginCaptures": map[string]interface{}{"1": map[string]interface{}{"name": "keyword.var.go"}, "2": map[string]interface{}{"name": "punctuation.other.bracket.round.go"}}}, map[string]interface{}{"comment": "Terminators", "match": ";", "name": "punctuation.terminator.go"}, map[string]interface{}{"include": "#brackets"}, map[string]interface{}{"include": "#delimiters"}, map[string]interface{}{"include": "#keywords"}, map[string]interface{}{"include": "#operators"}, map[string]interface{}{"include": "#runes"}, map[string]interface{}{"include": "#storage_types"}}, "repository": map[string]interface{}{"delimiters": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"name": "punctuation.other.comma.go", "match": ","}, map[string]interface{}{"match": "\\.(\\.\\.)", "name": "punctuation.other.period.go"}, map[string]interface{}{"match": ":(=)", "name": "punctuation.other.colon.go"}}}, "keywords": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"comment": "Flow control keywords", "match": "\\b(break|case|continue|default|defer|else|fallthrough|for|go|goto|if|range|return|select|switch)\\b", "name": "keyword.control.go"}, map[string]interface{}{"match": "\\bchan\\b", "name": "keyword.channel.go"}, map[string]interface{}{"name": "keyword.const.go", "match": "\\bconst\\b"}, map[string]interface{}{"match": "\\bfunc\\b", "name": "keyword.function.go"}, map[string]interface{}{"match": "\\binterface\\b", "name": "keyword.interface.go"}, map[string]interface{}{"match": "\\bimport\\b", "name": "keyword.import.go"}, map[string]interface{}{"name": "keyword.map.go", "match": "\\bmap\\b"}, map[string]interface{}{"match": "\\bpackage\\b", "name": "keyword.package.go"}, map[string]interface{}{"match": "\\bstruct\\b", "name": "keyword.struct.go"}, map[string]interface{}{"match": "\\btype\\b", "name": "keyword.type.go"}, map[string]interface{}{"match": "\\bvar\\b", "name": "keyword.var.go"}}}, "operators": map[string]interface{}{"comment": "Note that the order here is very important!", "patterns": []interface{}{map[string]interface{}{"match": "(\\*|&)(\\w)", "name": "keyword.operator.address.go"}, map[string]interface{}{"match": "<\\-", "name": "keyword.operator.channel.go"}, map[string]interface{}{"match": "\\-\\-", "name": "keyword.operator.decrement.go"}, map[string]interface{}{"match": "\\+\\+", "name": "keyword.operator.increment.go"}, map[string]interface{}{"match": "(==|!=|<=|>=|<[^<]|>[^>])", "name": "keyword.operator.comparison.go"}, map[string]interface{}{"match": "(&&|\\|\\||!)", "name": "keyword.operator.logical.go"}, map[string]interface{}{"match": "(=|\\+=|\\-=|\\|=|\\^=|\\*=|/=|:=|%=|<<=|>>=|&\\^=|&=)", "name": "keyword.operator.assignment.go"}, map[string]interface{}{"match": "(\\+|\\-|\\*|/|%)", "name": "keyword.operator.arithmetic.go"}, map[string]interface{}{"match": "(&(\\^)|\\||\\^|&\\^|<<|>>)", "name": "keyword.operator.arithmetic.bitwise.go"}, map[string]interface{}{"match": "\\.\\.\\.", "name": "keyword.operator.ellipsis.go"}}}, "runes": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"match": "\\'(\\\\([0-7]{3}|[abfnrtv\\\\'\"]|x[0-9a-fA-F]{2}|u[0-9a-fA-F]{4}|U[0-9a-fA-F]{8})|\\p{Any})\\'", "name": "constant.other.rune.go"}, map[string]interface{}{"match": "\\'.*\\'", "name": "invalid.illegal.unknown-rune.go"}}}, "string_escaped_char": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"match": "\\\\([0-7]{3}|[abfnrtv\\\\'\"]|x[0-9a-fA-F]{2}|u[0-9a-fA-F]{4}|U[0-9a-fA-F]{8})", "name": "constant.character.escape.go"}, map[string]interface{}{"match": "\\\\[^0-7xuUabfnrtv\\'\"]", "name": "invalid.illegal.unknown-escape.go"}}}, "variables": map[string]interface{}{"comment": "First add tests and make sure existing tests still pass when changing anything here!", "patterns": []interface{}{map[string]interface{}{"match": "([a-zA-Z_]\\w*(?:,\\s*[a-zA-Z_]\\w*)*)\\s*(=.*)", "captures": map[string]interface{}{"1": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"name": "variable.other.assignment.go", "match": "[a-zA-Z_]\\w*"}, map[string]interface{}{"include": "#delimiters"}}}, "2": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"include": "$self"}}}}}, map[string]interface{}{"match": "([a-zA-Z_]\\w*(?:,\\s*[a-zA-Z_]\\w*)*)(\\s+[\\*]?[a-zA-Z_]\\w*\\s*)(=.*)", "captures": map[string]interface{}{"1": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"match": "[a-zA-Z_]\\w*", "name": "variable.other.assignment.go"}, map[string]interface{}{"include": "#delimiters"}}}, "2": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"include": "$self"}}}, "3": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"include": "$self"}}}}}, map[string]interface{}{"match": "([a-zA-Z_]\\w*(?:,\\s*[a-zA-Z_]\\w*)*)(\\s+[\\[\\]\\*]{0,3}[a-zA-Z_]\\w*\\s*[^=].*)", "captures": map[string]interface{}{"1": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"match": "[a-zA-Z_]\\w*", "name": "variable.other.declaration.go"}, map[string]interface{}{"include": "#delimiters"}}}, "2": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"include": "$self"}}}}}}}, "comments": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"begin": "/\\*", "end": "\\*/", "captures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.comment.go"}}, "name": "comment.block.go"}, map[string]interface{}{"begin": "//", "beginCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.comment.go"}}, "end": "$", "name": "comment.line.double-slash.go"}}}, "storage_types": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"match": "\\bbool\\b", "name": "storage.type.boolean.go"}, map[string]interface{}{"match": "\\bbyte\\b", "name": "storage.type.byte.go"}, map[string]interface{}{"match": "\\berror\\b", "name": "storage.type.error.go"}, map[string]interface{}{"match": "\\b(complex(64|128)|float(32|64)|u?int(8|16|32|64)?)\\b", "name": "storage.type.numeric.go"}, map[string]interface{}{"name": "storage.type.rune.go", "match": "\\brune\\b"}, map[string]interface{}{"match": "\\bstring\\b", "name": "storage.type.string.go"}, map[string]interface{}{"match": "\\buintptr\\b", "name": "storage.type.uintptr.go"}}}, "string_placeholder": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"match": "%(\\[\\d+\\])?([\\+#\\-0\\x20]{,2}((\\d+|\\*)?(\\.?(\\d+|\\*|(\\[\\d+\\])\\*?)?(\\[\\d+\\])?)?))?[vT%tbcdoqxXUbeEfFgGsp]", "name": "constant.other.placeholder.go"}}}, "brackets": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"match": "\\{|\\}", "name": "punctuation.other.bracket.curly.go"}, map[string]interface{}{"match": "\\(|\\)", "name": "punctuation.other.bracket.round.go"}, map[string]interface{}{"match": "\\[|\\]", "name": "punctuation.other.bracket.square.go"}}}}, "scopeName": "source.go"}, map[string]interface{}{"scopeName": "source.json", "name": "JSON", "fileTypes": []interface{}{"avsc", "babelrc", "bowerrc", "composer.lock", "geojson", "gltf", "ipynb", "jscsrc", "jshintrc", "jslintrc", "json", "jsonl", "jsonld", "languagebabel", "ldj", "ldjson", "schema", "stylintrc", "template", "tern-config", "tern-project", "tfstate", "tfstate.backup", "topojson", "webapp", "webmanifest"}, "patterns": []interface{}{map[string]interface{}{"include": "#value"}}, "repository": map[string]interface{}{"constant": map[string]interface{}{"match": "\\b(true|false|null)\\b", "name": "constant.language.json"}, "number": map[string]interface{}{"match": "-?([1-9]|0(\\d))\\d+(\\.\\d+)?([eE][+-]?\\d+)?", "name": "constant.numeric.json"}, "object": map[string]interface{}{"beginCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.dictionary.begin.json"}}, "end": "}", "endCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.dictionary.end.json"}}, "name": "meta.structure.dictionary.json", "patterns": []interface{}{map[string]interface{}{"begin": "(\")", "end": "(=\")", "name": "meta.structure.dictionary.key.json", "patterns": []interface{}{map[string]interface{}{"include": "#string"}}}, map[string]interface{}{"begin": ":", "beginCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.separator.dictionary.key-value.json"}}, "end": "(,)([\\s\\n]*})|(,)|(})", "endCaptures": map[string]interface{}{"1": map[string]interface{}{"name": "invalid.illegal.trailing-dictionary-separator.json"}, "2": map[string]interface{}{"name": "punctuation.separator.dictionary.pair.json"}}, "name": "meta.structure.dictionary.value.json", "patterns": []interface{}{map[string]interface{}{"include": "#value"}, map[string]interface{}{"match": "[^\\s,]", "name": "invalid.illegal.expected-dictionary-separator.json"}}}, map[string]interface{}{"match": "[^\\s}]", "name": "invalid.illegal.expected-dictionary-separator.json"}}, "begin": "{"}, "string": map[string]interface{}{"name": "string.quoted.double.json", "patterns": []interface{}{map[string]interface{}{"match": "()\n\\\\                # a literal backslash\n(                   # followed by\n  [\"\\\\/bfnrt]     # one of these characters\n  |                 # or\n  u[0-9a-fA-F]{4}   # a u and four hex digits\n)", "name": "constant.character.escape.json"}, map[string]interface{}{"match": "\\\\.", "name": "invalid.illegal.unrecognized-string-escape.json"}}, "begin": "\"", "beginCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.string.begin.json"}}, "end": "\"", "endCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.string.end.json"}}}, "value": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"include": "#constant"}, map[string]interface{}{"include": "#number"}, map[string]interface{}{"include": "#string"}, map[string]interface{}{"include": "#array"}, map[string]interface{}{"include": "#object"}}}, "array": map[string]interface{}{"end": "(,)?[\\s\\n]*(\\])", "endCaptures": map[string]interface{}{"1": map[string]interface{}{"name": "invalid.illegal.trailing-array-separator.json"}, "2": map[string]interface{}{"name": "punctuation.definition.array.end.json"}}, "name": "meta.structure.array.json", "patterns": []interface{}{map[string]interface{}{"include": "#value"}, map[string]interface{}{"match": ",", "name": "punctuation.separator.array.json"}, map[string]interface{}{"name": "invalid.illegal.expected-array-separator.json", "match": "[^\\s\\]]"}}, "begin": "\\[", "beginCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.array.begin.json"}}}}}}