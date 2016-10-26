package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
)

type homeHandlers struct{}

var pageTmpl = `<!DOCTYPE html>
<html>
<head>
<title>{{.title}}</title>
<style>
html{
	box-sizing: 'border-box';
}
.container{
	width: 100%;
	min-height: 100vh;
}
.container > div{width: 100%;float:left;}
.col-md-4{ width: 33.33%;}
.col-md-offset{margin-left: 33.33%;}
</style>
</head>
<body>
	<div class="container">{{template "content" .}}</div>
</body>
</html>
`

var homePage = `{{define "content"}}<div class="row">
<div class="col-md-4 col-md-offset-4">
<h2>{{.message}}</h2>
<div class="panel">
<ul>
<li><a href="/about">About</a></li>
<li><a href="/hello/" class="btn">Hello</a></li>
<li><input type="text" name="username">  <a href="/hello/" class="btn" id="gobtn">Go</a></li>
</ul>
<p id="msg"></p>
</div>
</div>
</div>

<script>
var ip = document.querySelector("input[name='username']")
var btn = document.getElementById("gobtn")
var msg = document.getElementById("msg")
var baseurl = String(btn.href)
ip.addEventListener('input',function(e){
	btn.href = baseurl+ip.value
})

document.querySelectorAll("a.btn").forEach(a=>{
	a.addEventListener("click",function(e){
		e.preventDefault()
		console.log(e)
		GetResult(e.target.href)
		.then(res=>{return res.text()})
		.then(txt=>{msg.innerHTML = e.target.href+"  :  "+txt})
		return false
	})
})

function GetResult(path){
	return fetch(path)

}
</script>
{{end}}`

var aboutPage = `{{define "content"}}<div class="row">
<div class="col-md-4 col-md-offset-4">
<h2>{{.message}}</h2>
</div>
</div>
{{end}}`

func (h homeHandlers) homePage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, int64(0), map[string]string{"title": "Mux Test Home", "message": "HOME"})
	newr := r.WithContext(ctx)
	h.render(homePage, w, newr)
}

func (h homeHandlers) aboutPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, int64(0), map[string]string{"title": "Mux Test Home", "message": "ABOUT"})
	newr := r.WithContext(ctx)
	h.render(aboutPage, w, newr)
}

func (h homeHandlers) render(tmpl string, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := ctx.Value(int64(0))
	t, err := template.New("page").Parse(pageTmpl + tmpl)
	if err != nil {
		log.Println(err)
		h.errPage500(w, r)
		return
	}
	w.Header().Add("Content-type", "text/html")
	t.ExecuteTemplate(w, "page", data)
}

func (h homeHandlers) errPage404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Sorry cant find that."))
}
func (h homeHandlers) errPage500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Sorry had a bit of a problem."))
}
