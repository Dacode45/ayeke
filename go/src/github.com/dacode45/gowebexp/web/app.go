package web

import (
  "github.com/dacode45/gowebexp/pages"
  "crypto/rand"
  "encoding/base64"
  "fmt"
  "github.com/gorilla/context"
  "github.com/gorrila/mux"
  "github.com/gorrila/sessions"
  "html/template"
  "log"
  "net/http"
  "os"
  "path/filepath"
)

//Check existence of file or a directory
func Exists(name string) bool{
  if _, err := os.Stat(name); err != nil{
    if os.IsNotExist(err){
      log.Fatal(name, " does not exist.")
    }
    return false
  }
  return true
}

func templatePath(templateDir, name string) string{
  return filepath.Join(templateDir, name)
}

func parseTemplates(templateDir string, filenames ...string) (templates map[string]*template.Template){
  if len(filenames) == 0{
    log.Fatal("You must pass at least one file to parseTemplates")
  }
  templates = make(map[string]*template.Template)

  for _, filename := range filenames {
    templates[filename] = template.Must(template.ParseFiles(
      templatePath(templateDir, filename),
      templatePath(templateDir, "base.html")
      ))
  }
  return templates
}

//Implement a singleton WebApp
type WebApp struct{
  Router *mux.Router
  Storage Storage
  StaticDir string
  TemplateDir string
  Templates map[string]*template.Template
  CookieStore *sessions.CookieStore
}

func newWebApp() *WebApp{
  storage, err := NewMemStorage()
  if err != nil {
    fmt.Println("An error occured while creating the storage : ", err)
  }

  //Read configuration from environ variable
  StaticDir := os.Getenv("GOWEBEXP_STATIC")
  if staticDir == "" || !Exists(StaticDir){
    log.Println("StaticDir: ", staticDir)
    log.Fatal("`export GOWEB_STATIC=<a=path>` must be set in the environ")
  }
  TemplateDir := os.Getenv("GOWEBEXP_TEMPLATES")
  if TemplateDir == "" || !Exists(TemplateDir){
    log.Println("TemplateDir: ", TemplateDir)
    log.Fatal("`export GOWEB_EXP_TEMPLATES=<path>` must be set in the environ")
  }

  app := &WebApp{
    Router: mux.NewRouter(),
    Storage: storage,
    StaticDir: StaticDir,
    TemplateDir: TemplateDir,
    Templates: parseTemplates(TemplateDir, "index.html", "page_list.html", "page_detail.html")
    CookieStore:= sessions.NewCookieStore([]byte{"something very secret"})
  }
  return app
}

func(app *WebApp) GenerateCSRFToken() string{
  token := make([]byte, 80)
  rand.Read(token)
  base64.StdEncoding.EncodeToString(token)
  return token
}

func(app *WebApp) SetRoute(){
  // request multiplexer
	app.Router.HandleFunc("/inspect/{slug}/", RequestInspector).Name("inspector")
	app.Router.HandleFunc("/pages/{slug}/", PageDetail).Name("page_detail")
	app.Router.HandleFunc("/pages/", PageList).Name("page_list")
	app.Router.HandleFunc("/", Index).Name("index")
}

func (app *WebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the session and write a DummyKey
	session, _ := App.CookieStore.Get(r, "gowebexp")
	csrf_token, ok := session.Values["csrf_token"]
	if !ok {
		csrf_token = App.GenerateCsrfToken()
		session.Values["csrf_token"] = csrf_token
		App.CookieStore.Save(r, w, session)
	}

	context.Set(r, "csrf_token", csrf_token)
	if r.Method == "POST" && r.FormValue("csrf_token") != csrf_token {
		http.Error(w, "Fobidden", http.StatusForbidden)
	} else {
		app.Router.ServeHTTP(w, r)
	}
}

var App = newWebApp()

func RequestInspector(w http.ResponsWriter, r *http.Request){
  fmt.Fprintf(w, "URL: %s, method: %s, vars: %s", r.URL, r.Method, mux.Vals(r))
}



func Index(w http.ResponseWriter, r *http.Request) {
	pages_url, err := App.Router.Get("page_list").URL()
	if err != nil {
		panic(err)
	}
	tmpl := App.Templates["index.html"]
	ctx := make(map[string]interface{})
	ctx["pages_url"] = pages_url
	ctx["title"] = "Web experiment with GO"
	err = tmpl.ExecuteTemplate(w, "base", ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PageList(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})
	ctx["csrf_token"] = context.Get(r, "csrf_token")
	if r.Method == "POST" {
		page := pages.Page{
			Name:    r.FormValue("name"),
			Slug:    r.FormValue("slug"),
			Content: r.FormValue("content")}

		validationErrors, err := page.Validate()
		if err != nil {
			ctx["validationErrors"] = validationErrors
		} else {
			App.Storage.AddPage(page)
			// Redirect to /pages/ after the creation of the page
			// retrieve the url from the Router
			pages_url, err := App.Router.Get("page_list").URL()
			if err != nil {
				panic(err)
			}
			http.Redirect(w, r, pages_url.String(), 302)
		}
	}

	tmpl := App.Templates["page_list.html"]
	ctx["storage"] = App.Storage
	err := tmpl.ExecuteTemplate(w, "base", ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PageDetail(w http.ResponseWriter, r *http.Request) {
	tmpl := App.Templates["page_detail.html"]
	ctx := make(map[string]interface{})
	page, err := App.Storage.GetPageBySlug(mux.Vars(r)["slug"])
	ctx["page"] = page
	if err != nil {
		http.NotFound(w, r)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
