package conllu

import (
	"bufio"
	"os"
	"testing"
)

var (
	SAMPLE_PARAGRAPH = []string{
		"# text = тест, тест, тест",
		"# source = corpus.txt 1",
		"1	Алгоритм	алгоритм	NOUN	_	Animacy=Inan|Case=Nom|Gender=Masc|Number=Sing	12	nsubj	12:nsubj	SpaceAfter=No",
		"2	,	,	PUNCT	_	_	1	punct	1:punct	_",
		"3	от	от	ADP	_	_	4	case	4:case	_",
		"4	имени	имя	NOUN	_	Animacy=Inan|Case=Gen|Gender=Neut|Number=Sing	1	conj	1:conj	_",
	}
)

func TestParseParagraph(t *testing.T) {
	sentence, err := ParseParagraph(SAMPLE_PARAGRAPH)
	if err != nil {
		t.Fatal(err)
	}
	if sentence.Metadata["text"] != "тест, тест, тест" {
		t.Fail()
	}
	if sentence.Metadata["source"] != "corpus.txt 1" {
		t.Fail()
	}
}

func TestParse(t *testing.T) {
	file, _ := os.Open("test_data/syntagrus.conllu")
	defer file.Close()
	rd := bufio.NewReader(file)
	sentences, err := Parse(rd)
	if err != nil {
		t.Fatal(err)
	}
	if len(sentences) != 5 {
		t.Fail()
	}
}

func BenchmarkParseParagraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseParagraph(SAMPLE_PARAGRAPH)
	}
}
