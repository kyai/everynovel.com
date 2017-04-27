package controllers

import (
    // "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    // "time"
)

type UserController struct {
    BaseController
}

type User struct {
    Id_ 		bson.ObjectId `bson:"_id" json:"id"`
    Username    string `json:"username"`
    Password    string `json:"password"`
}

// const (
//     URL = "47.90.49.118:28080"
// )

func (c *UserController) GetUser() {
    session, err := mgo.Dial(URL)  //连接数据库
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    // id, _ := c.GetInt32("id")
    username := c.GetString("username")
    password := c.GetString("password")

    collection := session.DB("everynovel").C("en_user")

    infs := User{}
    err = collection.Find(&bson.M{"username":username,"password":password}).One(&infs)

    if err == nil{
        var rdata = make(map[string]interface{})
        rdata["id"] = infs.Id_
        rdata["username"] = infs.Username

        // session
        c.SetSession("uid", infs.Id_)
        c.SetSession("uname", infs.Username)

        c.EchoJSON("1001", "222", rdata)
    }else{
        c.EchoJSON("1002", "验证失败", "")
    }
}

// func (c *WordController) PostUser() {
//     session, err := mgo.Dial(URL)  //连接数据库
//     if err != nil {
//         panic(err)
//     }
//     defer session.Close()
//     session.SetMode(mgo.Monotonic, true)

//     collection := session.DB("everynovel").C("en_word")

//     pid := c.GetString("pid")
//     title := c.GetString("title")
//     content := c.GetString("content")

//     fmt.Println("===")
//     fmt.Println(pid)
//     fmt.Println(title)
//     // fmt.Println(content)

// 	word := &Word{
// 		Id_: bson.NewObjectId(),
//         Pid: pid,
// 		Title: title,
// 		Content: content,
// 		Date: time.Now().Format("2006-01-02 15:04:05"),
// 	}

// 	err = collection.Insert(word)
// 	if err != nil {
// 		panic(err)
// 	}

//     c.EchoJSON("111", "success", "")
// }