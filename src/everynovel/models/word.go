package models

import (
	"github.com/astaxie/beego/orm"
)

func GetWord(Id, PId int) (float64, error) {
	// o := orm.NewOrm()
	// o.Using("default")
	// //wparam := make([]interface{}, 0)
	// //wparam = append(wparam, SiteId, Username)
	// var uMoney float64
	// err := o.Raw("SELECT money FROM `l_user` WHERE site_id = '" + SiteId + "' AND username = '" + Username + "' Limit 1").QueryRow(&uMoney)
	// if err == nil {
	// 	return uMoney, err
	// } else {
	// 	return 0.00, err
	// }
}

func init() {

}
