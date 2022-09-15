// Package render
// @Author:        asus
// @Description:   $
// @File:          response
// @Data:          2022/2/2818:27
//
package render

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

type ResList struct {
    List []interface{}
    Total  int

}