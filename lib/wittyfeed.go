package lib
import (
  "os/exec"
  "github.com/PuerkitoBio/goquery"
  "strings"
  "time"
)

func WittyFeed()[]Post  {
  var posts []Post
  cmd:=exec.Command("phantomjs","lib/serverRender.js","https://www.wittyfeed.com/")
  out,_:=cmd.Output()
  htmlReader:=strings.NewReader(string(out))
  doc,_:=goquery.NewDocumentFromReader(htmlReader)
  doc.Find("div.stryTitle-wrapper").Find("div.stryTitle-caption").Each(func (i int, selection *goquery.Selection)  {
    widgetBox:=selection.Find("h3").Find("a")
    header:=widgetBox.Text()
    href,_:=widgetBox.Attr("href")
    posts=append(posts,Post{header, "https://wittyfeed.com"+href, time.Now()})
  })
  return posts
}
