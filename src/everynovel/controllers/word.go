package controllers

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type WordController struct {
    BaseController
}

type Word struct {
    Id_ 		bson.ObjectId `bson:"_id" json:"id"`
    Tid			string `json:"tid"`
    Pid         string `json:"pid"`
    Title 		string `json:"title"`
    Content     string `json:"content"`
    Date        string `json:"date"`
    User 		string `json:"user"`
}

const (
    URL = "47.90.49.118:28080"
)

func (c *WordController) GetWord() {
    session, err := mgo.Dial(URL)  //连接数据库
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    // id, _ := c.GetInt32("id")
    pid := c.GetString("id")

    collection := session.DB("everynovel").C("en_word")

    infs := []Word{}
    err = collection.Find(&bson.M{"pid":pid}).Sort("-date").All(&infs)

    Id, _ := c.GetInt32("id")
    c.EchoJSON("111", string(Id), infs)
}

func (c *WordController) PostWord() {
    session, err := mgo.Dial(URL)  //连接数据库
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    collection := session.DB("everynovel").C("en_word")

    pid := c.GetString("pid")
    title := c.GetString("title")
    content := c.GetString("content")

    fmt.Println("===")
    fmt.Println(pid)
    fmt.Println(title)
    // fmt.Println(content)

	word := &Word{
		Id_: bson.NewObjectId(),
        Pid: pid,
		Title: title,
		Content: content,
		Date: time.Now().Format("2006-01-02 15:04:05"),
	}

	err = collection.Insert(word)
	if err != nil {
		panic(err)
	}

    c.EchoJSON("111", "success", "")
}