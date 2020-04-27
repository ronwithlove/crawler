package model

//这里另起model并没有放在zhenai下，是因为这个爬虫也可以正对其他网站
//相亲网
type Profile struct {
	UserId string
	Name string
	Gender string
	Age	int
	Height	int
	Income string
	Marriage string
	Education string
	Occupation string
	City string
	Jiguan string
}
