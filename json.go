package main

var json = map[string]interface{}{"repository": map[string]interface{}{"array": map[string]interface{}{"begin": "\\[", "beginCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.array.begin.json"}}, "end": "(,)?[\\s\\n]*(\\])", "endCaptures": map[string]interface{}{"1": map[string]interface{}{"name": "invalid.illegal.trailing-array-separator.json"}, "2": map[string]interface{}{"name": "punctuation.definition.array.end.json"}}, "name": "meta.structure.array.json", "patterns": []interface{}{map[string]interface{}{"include": "#value"}, map[string]interface{}{"match": ",", "name": "punctuation.separator.array.json"}, map[string]interface{}{"match": "[^\\s\\]]", "name": "invalid.illegal.expected-array-separator.json"}}}, "constant": map[string]interface{}{"match": "\\b(true|false|null)\\b", "name": "constant.language.json"}, "number": map[string]interface{}{"match": "-?([1-9]|0(\\d))\\d+(\\.\\d+)?([eE][+-]?\\d+)?", "name": "constant.numeric.json"}, "object": map[string]interface{}{"begin": "{", "beginCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.dictionary.begin.json"}}, "end": "}", "endCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.dictionary.end.json"}}, "name": "meta.structure.dictionary.json", "patterns": []interface{}{map[string]interface{}{"end": "(=\")", "name": "meta.structure.dictionary.key.json", "patterns": []interface{}{map[string]interface{}{"include": "#string"}}, "begin": "(\")"}, map[string]interface{}{"name": "meta.structure.dictionary.value.json", "patterns": []interface{}{map[string]interface{}{"include": "#value"}, map[string]interface{}{"match": "[^\\s,]", "name": "invalid.illegal.expected-dictionary-separator.json"}}, "begin": ":", "beginCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.separator.dictionary.key-value.json"}}, "end": "(,)([\\s\\n]*})|(,)|(})", "endCaptures": map[string]interface{}{"2": map[string]interface{}{"name": "punctuation.separator.dictionary.pair.json"}, "1": map[string]interface{}{"name": "invalid.illegal.trailing-dictionary-separator.json"}}}, map[string]interface{}{"match": "[^\\s}]", "name": "invalid.illegal.expected-dictionary-separator.json"}}}, "string": map[string]interface{}{"begin": "\"", "beginCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.string.begin.json"}}, "end": "\"", "endCaptures": map[string]interface{}{"0": map[string]interface{}{"name": "punctuation.definition.string.end.json"}}, "name": "string.quoted.double.json", "patterns": []interface{}{map[string]interface{}{"match": "(?x)\n\\\\                # a literal backslash\n(                   # followed by\n  [\"\\\\/bfnrt]     # one of these characters\n  |                 # or\n  u[0-9a-fA-F]{4}   # a u and four hex digits\n)", "name": "constant.character.escape.json"}, map[string]interface{}{"name": "invalid.illegal.unrecognized-string-escape.json", "match": "\\\\."}}}, "value": map[string]interface{}{"patterns": []interface{}{map[string]interface{}{"include": "#constant"}, map[string]interface{}{"include": "#number"}, map[string]interface{}{"include": "#string"}, map[string]interface{}{"include": "#array"}, map[string]interface{}{"include": "#object"}}}}, "scopeName": "source.json", "name": "JSON", "fileTypes": []interface{}{"avsc", "babelrc", "bowerrc", "composer.lock", "geojson", "gltf", "ipynb", "jscsrc", "jshintrc", "jslintrc", "json", "jsonl", "jsonld", "languagebabel", "ldj", "ldjson", "schema", "stylintrc", "template", "tern-config", "tern-project", "tfstate", "tfstate.backup", "topojson", "webapp", "webmanifest"}, "patterns": []interface{}{map[string]interface{}{"include": "#value"}}}
