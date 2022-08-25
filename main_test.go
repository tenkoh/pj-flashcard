package main_test

import (
	"testing"

	fc "github.com/tenkoh/pj-flashcard"
)

func Test_CardBinder(t *testing.T) {
	binder, err := fc.NewCardBinder("test", "test")
	if err != nil {
		t.Fatal(err)
	}
	n, err := binder.Add("card1", "card1 content")
	if err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatal(err)
	}
	if err := binder.Edit(0, "card1", "card1 content", true); err != nil {
		t.Fatal(err)
	}
	if !binder.Cards[0].Mastered {
		t.Error("failed to edit an existing flashcard. want Mastered=true, but got false")
	}
	if err := binder.Delete(0); err != nil {
		t.Fatal(err)
	}
}
