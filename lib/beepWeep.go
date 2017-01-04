package lib
import (
  "github.com/PuerkitoBio/goquery"
  "time"
)
type Post struct {
  Name string `bson:"name"`
  Href string `bson:"href"`
  Created_at time.Time `bson:"created_at"`
}
func BeepWeep()[]Post {
  var posts []Post;
  doc, _ := goquery.NewDocument("http://beepweep.com/")
  latest:=doc.Find("ul.widget-full1")
  latest.Find("li").Each(func (i int, s *goquery.Selection)  {
    widgetBox:=s.Find(".widget-full-list-text").Find("a")
    header:=widgetBox.Text()
    href,_:=widgetBox.Attr("href")
    posts=append(posts,Post{header, href, time.Now()})
  })
  return posts
}
