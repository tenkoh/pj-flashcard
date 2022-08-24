package main_test

import (
	"testing"

	fc "github.com/tenkoh/pj-flashcard"
)

func Test_Flashcards(t *testing.T) {
	nb, err := fc.NewNotebook(1, "test", "test")
	if err != nil {
		t.Fatal(err)
	}
	n, err := nb.Add("card1", "card1 content")
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatal(err)
	}
	if err := nb.Edit(0, "card1", "card1 content", true); err != nil {
		t.Fatal(err)
	}
	if !nb.Flashcards[0].Mastered {
		t.Error("failed to edit an existing flashcard. want Mastered=true, but got false")
	}
	if err := nb.Delete(0); err != nil {
		t.Fatal(err)
	}
}
