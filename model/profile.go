package model

import "encoding/json"

//这里另起model并没有放在zhenai下，是因为这个爬虫也可以正对其他网站
//相亲网
type Profile struct {
	UserId string
	Name string
	Age	string
	Marriage string
}


func FromJsonObj(o interface{}) (Profile, error){
	var profile Profile
	s,err:=json.Marshal(o)//转成string
	if err != nil{
		return profile,err
	}

	err=json.Unmarshal(s,&profile)//在转成json放到这个profile中
	return profile,err
}