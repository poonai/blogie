package main
import (
  "blogscraper/lib"
  "log"
  mgo "gopkg.in/mgo.v2"
  "time"
  "gopkg.in/mgo.v2/bson"
  "os"
)

func main()  {
  for _= range time.NewTicker(time.Hour*3).C {
    Job()
  }
}

func Job()  {
  session, _ := mgo.Dial(os.Getenv("MLAB"))
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  c:=session.DB("blogscrapper").C("posts")
  posts:=lib.AggregateAllPost()
  for _,post:=range posts{
    exists,_:=c.Find(bson.M{"name":post.Name}).Count()
    if exists == 0 {
      log.Print("inserting")
      c.Insert(&post)
    } else {
      log.Print("already exist")
    }
  }
}
