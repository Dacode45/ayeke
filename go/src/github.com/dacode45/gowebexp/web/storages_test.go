package web

import (
  "github.com/dacode45/gowebexp/pages"
  "fmt"
  "testing"
)

func TestNewMemStorage(t *testing.T){
  s, err := NewMemStorage()
  if err != nill{
    t.Error("An error occured while creating a MemStorage", err)
  }
  fmt.Println(s)
}

func TestNewMemStorage_AddPage(t *testing.T){
  s, err := NewMemStorage()
  if err != nil{
    t.Error("An error occured while creating a MemStorage", err)
  }
  s.AddPage(pages.Page{
    Name: "ham",
    Slug: "ham",
    Content: "#ham\nHello World",
  })
  if s.CountPage() != 3{
    t.Error("We should have 3 page in Storage but we found", s.CountPage())
  }
  fmt.Println(s)
  s.AddPage(pages.Page{
    Name: "dam",
    Slug: "dam",
    Content: "#dam\nHello World",
  })
  if s.CountPage() != 4{
    t.Error("We should have 4 page in Storage but we found", s.CountPage())
  }
  fmt.Println(s)
}

func TestNewMemStorage_GetPageBySlug(t *testing.T){
  s, err := NewMemStorage()
  page, err := s.GetPageBySlug("foo")
  if err != nil{
    t.Error("An error occured while retrieving the page \"foo\" by slug", err)
  }
  if page.Name != "foo"{
    t.Error("An error occured because page.Name is different than \"foo\"", page.Name)
  }
}
