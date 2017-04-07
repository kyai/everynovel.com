package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) BgetString(data string) string {
	return strings.TrimSpace(c.GetString(data))
}

//stoken验证
func (c *BaseController) SToken() {
	//SToken := strings.TrimSpace(this.GetString("stoken")) //前台ID

	//url := this.Ctx.Conten.Uri()
	//
	//p := this.Ctx.Request.URL.Path
	fdata := c.GetString("username")

	// var dat map[string]interface{}
	// if err := json.Unmarshal([]byte(fdata), &dat); err == nil {
	// 	fmt.Println(dat)
	// } else {
	// 	fmt.Println(err)
	// }

	// // var result map[string]interface{}
	// // err := json.Unmarshal([]byte(uri), &result)

	// // // url := this.Ctx.Uri()

	// //fmt.Println(result)

	fmt.Println(fdata)

	// fmt.Println(this.Ctx.Request.FormValue("index_id"))

	//fmt.Println(this.Ctx.Request.URL)
	//fmt.Println(SToken)

}

func (c *BaseController) EchoJSON(code, msg string, data interface{}) {
	var jdata = make(map[string]interface{})
	jdata["Code"] = code
	jdata["Msg"] = msg
	jdata["Data"] = data
	adata_json, _ := json.Marshal(jdata)
	c.Ctx.WriteString(string(adata_json))
	return
}

func init() {

}
