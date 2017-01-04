package lib
import(
  "sync"
)
func AggregateAllPost()[]Post  {
  var wg sync.WaitGroup
  var posts []Post
  wg.Add(1)
  go func() {
    for _,post := range WittyFeed() {
      posts=append(posts,post)
    }
    wg.Done()
  }()
  wg.Add(1)
  go func() {
    for _,post := range BeepWeep() {
      posts=append(posts,post)
    }
    wg.Done()
  }()
  wg.Wait()
  return posts
}
