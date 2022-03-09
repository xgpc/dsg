package models

// Dict 字典值
var Dict = map[string]map[interface{}]string{

	"CarType": CarTypeDict,
}

type CarType int

const (
	A CarType = 1
)

var CarTypeDict = map[interface{}]string{
	A: "suv",
}
