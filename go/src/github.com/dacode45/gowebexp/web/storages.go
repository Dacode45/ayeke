package web

import (
  "github.com/dakool45/gowebexp/pages"
  "errors"
  "strconv"
  "strings"
)

type Storage interface{
  AddPage(page pages.Page)
  GetPageBySlug(slug string) (pages.Page, error)
  String() string
  CountPage() int
}

//Mem storage
type MemStorage struct{
  Pages []pages.Page
}

// instance
func NewMemStorage() (s storage, err error){
  ms := &MemStorage{}
  ms.Init()
  return ms, nil
}

func (ms MemStorage) String() (str string){
  switch len(ms.Pages){
  case 0:
    str = "No pages"
  case 1:
    str = "1 page : " + ms.Pages[0].Name
  default:
    var page_names []string
    for _, p := range ms.Pages{
      page_names = append(page_names, p.Name)
    }
    str = strconv.Itoa(len(ms.Pages)) + " pages : " + strings.Join(page_names, ", ")
  }
  return str
}

func (ms *MemStorage) CountPage() int{
  return len(ms.Pages)
}

func (ms *MemStorage) AddPage(page pages.Page) err{
  for index, dup := range ms.Pages{
    if dup.Slug == page.Slug{
      ms.Pages[index] = page
      return
    }
  }
  ms.Pages = append(ms.Pages, page)
}

func (ms *MemStorage) GetPageBySlug(slug string) (page pages.Page, err error){
  for _, page := range ms.Pages{
    if page.Slug == slug{
      return page, nil
    }
  }
  return pages.Page{}, errors.New("storage.GetPageBySlug This page does not exist")
}

//Init Create 2 dummy pages
func (ms *MemStorage) Init(){
  ms.AddPage(pages.Page{
    Name: "foo",
    Slug: "foo",
    Content: "#foo \nHello World"
  })
  ms.AddPage(pages.Page{
    Name: "bar",
    Slug: "bar",
    Content: "#bar \nHello World"
  })
}
