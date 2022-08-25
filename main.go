package main

import (
	"errors"
	"fmt"
)

const (
	MAX_FLASHCARD = 1<<32 - 1
)

var (
	ErrInvalidCardBinder = errors.New("cardBinder must have a title")
	ErrFullCardBinder    = errors.New("no more cards could be inserted in the cardBinder")
	ErrEmptyCard         = errors.New("empty card. it must have both word and content")
	ErrNonExistCard      = errors.New("the card does not exist")
)

type App struct {
	Binders map[int]*CardBinder
}

type CardBinder struct {
	Title       string
	Description string
	Cards       map[int]*Card
	nextID      int
}

type Card struct {
	Word     string
	Content  string
	Mastered bool
}

func NewApp() *App {
	return &App{
		Binders: map[int]*CardBinder{},
	}
}

func NewCardBinder(title, description string) (*CardBinder, error) {
	if title == "" {
		return nil, ErrInvalidCardBinder
	}
	nb := &CardBinder{
		Title:       title,
		Description: description,
		Cards:       map[int]*Card{},
		nextID:      0,
	}
	return nb, nil
}

func (binder *CardBinder) Add(word, content string) (int, error) {
	if word == "" || content == "" {
		return 0, ErrEmptyCard
	}
	if binder.nextID >= MAX_FLASHCARD {
		return MAX_FLASHCARD, ErrFullCardBinder
	}
	card := &Card{
		Word:     word,
		Content:  content,
		Mastered: false,
	}
	binder.Cards[binder.nextID] = card
	binder.nextID = binder.nextID + 1
	return binder.nextID - 1, nil
}

func (binder *CardBinder) Delete(id int) error {
	_, exist := binder.Cards[id]
	if !exist {
		return ErrNonExistCard
	}
	delete(binder.Cards, id)
	return nil
}

func (binder *CardBinder) Edit(id int, word, content string, mastered bool) error {
	_, exist := binder.Cards[id]
	if !exist {
		return ErrNonExistCard
	}
	c := binder.Cards[id]
	c.Word = word
	c.Content = content
	c.Mastered = mastered
	binder.Cards[id] = c
	return nil
}

func main() {
	fmt.Println("Hello world")
}
