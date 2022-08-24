package main

import (
	"errors"
	"fmt"
)

const (
	MAX_FLASHCARD = 1<<32 - 1
)

var (
	ErrInvalidNotebook   = errors.New("notebook must have a title")
	ErrEmptyFlashcard    = errors.New("empty flashcard. it must have both word and content")
	ErrNonExistFlashcard = errors.New("the flashcard does not exist")
	ErrFullFlashcards    = errors.New("no more flashcards could be inserted in the notebook")
)

type Notebook struct {
	ID          int
	Title       string
	Description string
	Flashcards  map[int]*Flashcard
	nextID      int
}

type Flashcard struct {
	Word     string
	Content  string
	Mastered bool
}

func NewNotebook(id int, title, description string) (*Notebook, error) {
	if title == "" {
		return nil, ErrInvalidNotebook
	}
	nb := &Notebook{
		ID:          id,
		Title:       title,
		Description: description,
		Flashcards:  map[int]*Flashcard{},
		nextID:      0,
	}
	return nb, nil
}

func (nb *Notebook) Add(word, content string) (int, error) {
	if word == "" || content == "" {
		return 0, ErrEmptyFlashcard
	}
	if nb.nextID >= MAX_FLASHCARD {
		return MAX_FLASHCARD, ErrFullFlashcards
	}
	card := &Flashcard{
		Word:     word,
		Content:  content,
		Mastered: false,
	}
	nb.Flashcards[nb.nextID] = card
	nb.nextID = nb.nextID + 1
	return nb.nextID - 1, nil
}

func (nb *Notebook) Delete(id int) error {
	_, exist := nb.Flashcards[id]
	if !exist {
		return ErrNonExistFlashcard
	}
	delete(nb.Flashcards, id)
	return nil
}

func (nb *Notebook) Edit(id int, word, content string, mastered bool) error {
	_, exist := nb.Flashcards[id]
	if !exist {
		return ErrNonExistFlashcard
	}
	c := nb.Flashcards[id]
	c.Word = word
	c.Content = content
	c.Mastered = mastered
	nb.Flashcards[id] = c
	return nil
}

func main() {
	fmt.Println("Hello world")
}
