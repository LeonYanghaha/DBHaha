package util

//var  //config map[string]types.Interface{}
var config map[string]interface{} = make(map[string]interface{})

func init()  {
	config["port"] = 8000
}
