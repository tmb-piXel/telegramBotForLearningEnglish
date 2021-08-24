package services

import (
	"math/rand"

	m "github.com/tmb-piXel/LearnEnglishBot/pkg/models"
	r "github.com/tmb-piXel/LearnEnglishBot/pkg/repositories"
	"github.com/tmb-piXel/LearnEnglishBot/pkg/storage"
)

func NewUser(chatID int64) {
	u := m.NewUser(chatID)
	r.SaveUser(u)
}

func SaveWords(chatID int64, original, transleted string) {
	r.SaveWords(chatID, original, transleted)
}

func NewWord(chatID int64) (word string) {
	u := r.FindUser(chatID)
	words := storage.GetTransletedWords(u.GetLanguage(), u.GetTopic())
	if u.GetIsToRu() {
		words = storage.GetOriginalWords(u.GetLanguage(), u.GetTopic())
	}
	i := randomIter(*words)
	word = (*words)[i]
	u.SetIterWord(i)
	r.UpdateUser(u)
	return
}

func randomIter(a []string) int {
	size := len(a)
	r := rand.Intn(size)
	return r
}

func Word(chatID int64) (word string) {
	u := r.FindUser(chatID)
	words := storage.GetOriginalWords(u.GetLanguage(), u.GetTopic())
	if u.GetIsToRu() {
		words = storage.GetTransletedWords(u.GetLanguage(), u.GetTopic())
	}
	word = (*words)[u.GetIterWord()]
	return
}

func SetIsToRu(chatID int64, isToRu bool) {
	u := r.FindUser(chatID)
	u.SetIsToRu(isToRu)
	r.UpdateUser(u)
}

func SetLanguage(chatID int64, language string) {
	u := r.FindUser(chatID)
	u.SetLanguage(language)
	r.UpdateUser(u)
}

func Language(chaID int64) string {
	u := r.FindUser(chaID)
	return u.GetLanguage()
}

func SetTopic(chatID int64, topic string) {
	u := r.FindUser(chatID)
	u.SetTopic(topic)
	r.UpdateUser(u)
}

func ListWords(chatID int64) (list string) {
	u := r.FindUser(chatID)
	o := storage.GetOriginalWords(u.GetLanguage(), u.GetTopic())
	t := storage.GetTransletedWords(u.GetLanguage(), u.GetTopic())
	for i, w := range *o {
		list += w + " - " + (*t)[i] + "\n"
	}
	return
}
