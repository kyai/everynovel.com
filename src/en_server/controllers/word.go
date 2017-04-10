package controllers

import (
    // "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type WordController struct {
    BaseController
}

type Word struct {
    Id bson.ObjectId "_id"
    Content     string `json:"content"`
    Date        string `json:"date"`
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

    collection := session.DB("everynovel").C("en_word")

    infs := []Word{}
    err = collection.Find(&bson.M{}).All(&infs)

    Id, _ := c.GetInt32("id")
    c.EchoJSONP("111", string(Id), infs)
}