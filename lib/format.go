package gotprint

type Fix struct {
	Pre  string
	Post string
}

type FormatSettings struct {
	separator   string
	pad         string
	fixes       Fix
	structfixes []Fix
}

var defaultformat = FormatSettings{
	separator: " ",
	pad:       " ",
	fixes: Fix{
		Pre:  "",
		Post: "",
	},
	structfixes: []Fix{
		{Pre: "", Post: ""},
		{Pre: "[", Post: "]"},
		{Pre: "(", Post: ")"},
		{Pre: "{", Post: "}"},
	},
}

var currentformat FormatSettings

func SetDefaultFormat() {
	currentformat = defaultformat
}

func Format() *FormatSettings {
	return &currentformat
}

func (fs *FormatSettings) SetSeparator(sep string) *FormatSettings {
	fs.separator = sep
	return fs
}

func (fs *FormatSettings) SetPadding(pad string) *FormatSettings {
	fs.pad = pad
	return fs
}

func (fs *FormatSettings) SetFixes(f Fix) *FormatSettings {
	fs.fixes = f
	return fs
}

func (fs *FormatSettings) SetStructFixes(f []Fix) *FormatSettings {
	fs.structfixes = f
	return fs
}
