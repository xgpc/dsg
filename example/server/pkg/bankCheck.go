package pkg

import (
	"fmt"
	"github.com/xgpc/dsg/v2/pkg/json"
	"io"
	"net/http"
	"strings"
)

type RespCode string

var RespCodeMD = map[string]string{
	"0000": "结果匹配,信息准确，三个要素完全匹配",
	"0001": "开户名不能为空,开户名不能为空",
	"0002": "银行卡号格式错误,银行卡号格式错误",
	"0003": "身份证号格式错误,身份证号格式错误",
	"0006": "银行卡号不能为空,银行卡号不能为空",
	"0007": "身份证号不能为空,身份证号不能为空",
	"0008": "信息不匹配,信息不匹配",
}

func (p *RespCode) ErrorValue() string {
	return RespCodeMD[string(*p)]
}

type ResBankAuth struct {
	Name        string   `json:"name"`
	CardNo      string   `json:"cardNo"`
	IdNo        string   `json:"idNo"`
	RespMessage string   `json:"respMessage"`
	RespCode    RespCode `json:"respCode"`
	BankName    string   `json:"bankName"`
	BankKind    string   `json:"bankKind"`
	BankType    string   `json:"bankType"`
	BankCode    string   `json:"bankCode"`
}

func BankAuthenticate(APPCODE, name, cardNo, idNo string) (*ResBankAuth, error) {
	var resBankAuth ResBankAuth
	url := "https://yunybank34.market.alicloudapi.com/bankAuthenticate3"
	method := "POST"

	s := strings.Builder{}
	s.WriteString("name=")
	s.WriteString(name)
	s.WriteString("&cardNo=")
	s.WriteString(cardNo)
	s.WriteString("&idNo=")
	s.WriteString(idNo)

	payload := strings.NewReader(s.String())

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "APPCODE "+APPCODE)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			return
		}
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if res.StatusCode != 200 {
		if res.StatusCode != 202 {
			return nil, fmt.Errorf(res.Status)
		}
	}

	err = json.Decode(body, &resBankAuth)
	if err != nil {
		return nil, err
	}

	return &resBankAuth, nil
}
