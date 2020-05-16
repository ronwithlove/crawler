package worker

//序列化Parser
type SerializedParser struct{
	Name string//方法名 ,如：ParseCityList，ProfileParser
	Args interface{}//这里ParseCityList对应的就是nil,ProfileParser对应的就是传入username
}
//{"ParseCityList",nil},{"ProfileParser",username}


