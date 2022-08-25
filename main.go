package main

import (
	"embed"
	"errors"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	MAX_FLASHCARD = 1<<32 - 1
)

//all directive is required to embed _files
//go:embed all:frontend/out/*
var embededFiles embed.FS

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

func getFileSystem() (http.FileSystem, error) {
	fsys, err := fs.Sub(embededFiles, "frontend/out")
	if err != nil {
		return nil, err
	}
	return http.FS(fsys), nil
}

func main() {
	e := echo.New()
	fsys, err := getFileSystem()
	if err != nil {
		panic(err)
	}
	assetHandler := http.FileServer(fsys)
	e.GET("/", echo.WrapHandler(assetHandler))
	e.GET("/_next/*", echo.WrapHandler(assetHandler))
	e.Logger.Fatal(e.Start(":1323"))
}
