package bookstore_test

import (
	"testing"

	"bookstore"
	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "Some Title",
		Author: "Someone",
		Copies: 2,
	}

	want := 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuyErrorIfNoBooksLeft(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Copies: 0,
	}

	_, err := bookstore.Buy(b)

	if err == nil {
		t.Error("want error from buying after 0 copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "Some Title", Author: "Someone", Copies: 2},
		{Title: "Other Title", Author: "Someone else", Copies: 2},
	}

	want := []bookstore.Book{
		{Title: "Some Title", Author: "Someone", Copies: 2},
		{Title: "Other Title", Author: "Someone else", Copies: 2},
	}
	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()

	book := []bookstore.Book{
		{Id: 1, Author: "Someone"},
		{Id: 2, Author: "Someone else"},
	}

	get := bookstore.GetBook(book, 2)
	want := bookstore.Book{
		Id: 2, Author: "Someone else",
	}

	if !cmp.Equal(get, want) {
		t.Error(cmp.Diff(get, want))
	}
}
