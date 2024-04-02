package textprocessing

import (
	"strings"

	"github.com/siddhant-vij/Word-Frequency-Server/pkg/common"
)

type WordCountTask struct {
	Text string
	Word string
}

func (wct *WordCountTask) Execute() int {
	normalizedText := strings.ToLower(wct.Text)
	normalizedWord := strings.ToLower(wct.Word)
	return strings.Count(normalizedText, normalizedWord)
}

func NewWordCountTask(text, word string) common.Task {
	return &WordCountTask{
		Text: text,
		Word: word,
	}
}
