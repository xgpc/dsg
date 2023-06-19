package dsg

type Response struct {
	Code  int
	Msg   string
	Data  interface{}
	List  []interface{}
	Total int
}
