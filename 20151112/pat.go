package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// httpパッケージのServeMux構造体のmuxEntryがpatHandlerにあたる
type PatternServeMux struct {
	handlers map[string][]*patHandler // {"hoge":{pattern hendler}}みたいな
	// ユーザーに定義されたルーティングのパターンをまとめて保持しておく
}

// パッケージの初期化用のpublicメソッド
func New() *PatternServeMux {
	return &PatternServeMux{make(map[string][]*patHandler)}
}

// httpパッケージでルーティング時handleメソッドでリクエストに一致するhandlerを探し出してServeHTTPが呼ばれることになってる
// 基本的にはurl.pathとルーティングパターンが一致した時どのハンドラを実行するかを定義するだけ
func (p *PatternServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, ph := range p.handlers[r.Method] {
		if params, ok := ph.try(r.URL.Path); ok {
			if len(params) > 0 {
				// クエリパラメタ式に変換
				r.URL.RawQuery = url.Values(params).Encode() + "&" + r.URL.RawQuery
			}
			ph.ServeHTTP(w, r)
			return
		}
	}

	allowed := make([]string, 0, len(p.handlers))
	for meth, handlers := range p.handlers {
		if meth == r.Method {
			continue
		}

		for _, ph := range handlers {
			if _, ok := ph.try(r.URL.Path); ok {
				allowed = append(allowed, meth)
			}
		}
	}

	if len(allowed) == 0 {
		http.NotFound(w, r)
		return
	}

	w.Header().Add("Allow", strings.Join(allowed, ", "))
	http.Error(w, "Method Not Allowed", 405)
}

// HTTPメソッドごとにAddへ引き渡すのみ
func (p *PatternServeMux) Head(pat string, h http.Handler) {
	p.Add("HEAD", pat, h)
}

func (p *PatternServeMux) Get(pat string, h http.Handler) {
	p.Add("HEAD", pat, h)
	p.Add("GET", pat, h)
}

func (p *PatternServeMux) Post(pat string, h http.Handler) {
	p.Add("POST", pat, h)
}

func (p *PatternServeMux) Put(pat string, h http.Handler) {
	p.Add("PUT", pat, h)
}

func (p *PatternServeMux) Del(pat string, h http.Handler) {
	p.Add("DELETE", pat, h)
}

func (p *PatternServeMux) Options(pat string, h http.Handler) {
	p.Add("OPTIONS", pat, h)
}

// 実質ルーティングとハンドラ定義時のpublicインスタンスメソッド
// ex) p.Add("GET", "/hoge/:bar", HogeHandler)
func (p *PatternServeMux) Add(meth, pat string, h http.Handler) {
	p.handlers[meth] = append(p.handlers[meth], &patHandler{pat, h})

	n := len(pat)
	// パスがない時 ex) http://hoge.com/とか
	if n > 0 && pat[n-1] == '/' {
		p.Add(meth, pat[:n-1], http.RedirectHandler(pat, http.StatusMovedPermanently))
	}
}

func Tail(pat, path string) string {
	var i, j int
	for i < len(path) {
		switch {
		case j >= len(pat):
			if pat[len(pat)-1] == '/' {
				return path[i:]
			}
			return ""
		case pat[j] == ':':
			var nextc byte
			_, nextc, j = match(pat, isAlnum, j+1)
			_, _, i = match(path, matchPart(nextc), i)
		case path[i] == pat[j]:
			i++
			j++
		default:
			return ""
		}
	}
	return ""
}

// ルーティングのパターンを受ける構造体
type patHandler struct {
	pat          string // ルーティングパターン
	http.Handler        // ハンドラ
}

func (ph *patHandler) try(path string) (url.Values, bool) {
	p := make(url.Values)
	var i, j int
	for i < len(path) {
		switch {
		case j >= len(ph.pat):
			if ph.pat != "/" && len(ph.pat) > 0 && ph.pat[len(ph.pat)-1] == '/' {
				return p, true
			}
			return nil, false
		case ph.pat[j] == ':':
			var name, val string
			var nextc byte
			name, nextc, j = match(ph.pat, isAlnum, j+1)
			val, _, i = match(path, matchPart(nextc), i)
			p.Add(":"+name, val)
		case path[i] == ph.pat[j]:
			i++
			j++
		default:
			return nil, false
		}
	}
	if j != len(ph.pat) {
		return nil, false
	}
	return p, true
}

func matchPart(b byte) func(byte) bool {
	return func(c byte) bool {
		return c != b && c != '/'
	}
}

func match(s string, f func(byte) bool, i int) (matched string, next byte, j int) {
	j = i
	for j < len(s) && f(s[j]) {
		j++
	}
	if j < len(s) {
		next = s[j]
	}
	return s[i:j], next, j
}

func isAlpha(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' < ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAlnum(ch byte) bool {
	return isAlpha(ch) || isDigit(ch)
}

// テスト用のハンドラ
func HelloServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello,"+r.URL.Query().Get(":name")+"!\n")
}

// このスクリプトを実行してlocalhost:12345/hello/:name で動く
func main() {
	m := New()
	m.Get("/hello/:name", http.HandlerFunc(HelloServer))

	// http.Handle("/", m)
	err := http.ListenAndServe(":12345", m)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
