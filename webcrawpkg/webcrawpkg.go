package webcrawpkg

import (
	"log"
	"net/http"
	"os"
	// "github.com/ofabry/go-callvis/examples/main/webcrawpkg/job"
)

var l = log.New(os.Stderr, "mypkg", 0)

var T = new(WebSpider)

type WebSpider struct{}

func init() {
	go func() {
		l.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

func NewWebSpider() *WebSpider {
	return &WebSpider{}
}

func (t *WebSpider) Identifier() {}
func (t *WebSpider) ScrapData() {}
func (t *WebSpider) PraseData() {}
func (t *WebSpider) WriteInDB() {}
func (t *WebSpider) SendEmail() {}

func RunTask() {
	defer deferred()
	go Concurrent(T)
	T.StartTask()
}
func deferred() {}

func Concurrent(T *WebSpider) {
	T.ScrapData()
	T.PraseData()
	T.WriteInDB()
	T.SendEmail()
}
