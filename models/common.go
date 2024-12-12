package models

type RegexReplace struct {
	Pattern string
	With    string
}

type DocReplacement struct {
	Match   string
	Select  []string
	Replace string
}
