package conllu

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

type Metadata map[string]string

const (
	METADATA_PREFIX            byte   = '#'
	METADATA_DELIMITER         string = "="
	TOKEN_ATTRIBUTE_COUNT      int    = 10
	TOKEN_ATTRIBUTE_DELIMITER  string = "\t"
	TOKEN_ATTRIBUTE_EMPTY      string = "_"
	TOKEN_ARRAY_DELIMITER      string = "|"
	TOKEN_FEATURE_DELIMITER    string = "="
	TOKEN_DEPENDENCY_DELIMITER string = ":"
	TOKEN_RANGE_DELIMITER      string = "-"
	TOKEN_EMPTY_DELIMITER      string = "."
)

type Feature struct {
	Name  string
	Value string
}

type Dependency struct {
	Head     string
	Relation string
}

type Token struct {
	ID      string       // Word index, integer starting at 1 for each new sentence; may be a range for multiword tokens; may be a decimal number for empty nodes.
	Form    string       // Word form or punctuation symbol.
	Lemma   string       // Lemma or stem of word form.
	UPosTag string       // Universal part-of-speech tag.
	XPosTag string       // Language-specific part-of-speech tag; underscore if not available.
	Feats   []Feature    // List of morphological features from the universal feature inventory or from a defined language-specific extension; underscore if not available.
	Head    string       // Head of the current word, which is either a value of ID or zero (0).
	DepRel  string       // Universal dependency relation to the HEAD (root iff HEAD = 0) or a defined language-specific subtype of one.
	Deps    []Dependency // Enhanced dependency graph in the form of a list of head-deprel pairs.
	Misc    string       // Any other annotation.
}

func (t *Token) IsMultiword() bool {
	if strings.Index(t.ID, TOKEN_RANGE_DELIMITER) != -1 {
		return true
	} else {
		return false
	}
}

func (t *Token) IsEmptyNode() bool {
	if strings.Index(t.ID, TOKEN_EMPTY_DELIMITER) != -1 {
		return true
	} else {
		return false
	}
}

type Sentence struct {
	Tokens   []Token
	Metadata Metadata
}

func parseAttributeValueAsArray(value string) []string {
	if value == TOKEN_ATTRIBUTE_EMPTY {
		return []string{}
	} else {
		return strings.Split(value, TOKEN_ARRAY_DELIMITER)
	}
}

func parseFeature(value string) Feature {
	items := strings.Split(value, TOKEN_FEATURE_DELIMITER)
	return Feature{
		Name:  strings.TrimSpace(items[0]),
		Value: strings.TrimSpace(items[1]),
	}
}

func parseTokenFeatures(value string) []Feature {
	features := make([]Feature, 0)
	items := parseAttributeValueAsArray(value)
	for i := 0; i < len(items); i++ {
		features = append(features, parseFeature(items[i]))
	}
	return features
}

func parseDependency(value string) Dependency {
	items := strings.Split(value, TOKEN_DEPENDENCY_DELIMITER)
	return Dependency{
		Head:     items[0],
		Relation: items[1],
	}
}

func parseTokenDependencies(value string) []Dependency {
	dependencies := make([]Dependency, 0)
	items := parseAttributeValueAsArray(value)
	for i := 0; i < len(items); i++ {
		dependencies = append(dependencies, parseDependency(items[i]))
	}
	return dependencies
}

func ParseParagraph(lines []string) (Sentence, error) {
	tokens := make([]Token, 0)
	metadata := make(Metadata)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line[0] == METADATA_PREFIX {
			items := strings.Split(line[1:], METADATA_DELIMITER)
			if len(items) != 2 {
				return Sentence{}, errors.New(fmt.Sprintf("Invalid metadata line: '%v'", line))
			}
			key, value := strings.TrimSpace(items[0]), strings.TrimSpace(items[1])
			metadata[key] = value
		} else {
			items := strings.Split(line, TOKEN_ATTRIBUTE_DELIMITER)
			if len(items) != TOKEN_ATTRIBUTE_COUNT {
				return Sentence{}, errors.New(fmt.Sprintf("Invalid token line: '%+v'", line))
			}
			tokens = append(tokens,
				Token{
					ID:      items[0],
					Form:    items[1],
					Lemma:   items[2],
					UPosTag: items[3],
					XPosTag: items[4],
					Feats:   parseTokenFeatures(items[5]),
					Head:    items[6],
					DepRel:  items[7],
					Deps:    parseTokenDependencies(items[8]),
					Misc:    items[9],
				},
			)
		}
	}
	return Sentence{
		Tokens:   tokens,
		Metadata: metadata,
	}, nil
}

func Parse(rd *bufio.Reader) ([]Sentence, error) {
	buf := make([]string, 0)
	sentences := make([]Sentence, 0)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		} else {
			if (line == "\n") || (line == "") {
				if len(buf) > 0 {
					sentence, err := ParseParagraph(buf)
					if err != nil {
						return sentences, err
					}
					sentences = append(sentences, sentence)
					buf = make([]string, 0)
				}
			} else {
				buf = append(buf, line)
			}
		}
	}
	return sentences, nil
}
