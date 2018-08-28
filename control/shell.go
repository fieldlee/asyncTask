/**
 *	Author Fanxu(746439274@qq.com)
 */
package control

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"asyncTask/helpers"
	"asyncTask/model"
)

type ShellExecuteParam struct {
	ShellContent []byte `json:"shellContent"`
}

func (this * BaseControl) ShellExecute(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(model.CustomException{500, "params error"})
	}
	var u SdkExecuteParam
	err1 := json.Unmarshal(body,&u)
	if err1 != nil{
		panic(model.CustomException{402, "params error"})
	}
	var command = "rm -rf /tmp/execute.sh"
	res,err :=helpers.ExecShell(command)
	if err != nil{
		panic(model.CustomException{400, err.Error()})
	}
	f,err := os.Create("/tmp/execute.sh")
	defer f.Close()
	if err != nil {
		panic(model.CustomException{410, err.Error()})
	}
	_,err = f.Write(u.ShellContent)
	if err != nil {
		panic(model.CustomException{410, err.Error()})
	}
	command = "chmod a+x /tmp/execute.sh && sh /tmp/execute.sh"
	res,err =helpers.ExecShell(command)
	if err != nil{
		panic(model.CustomException{410, err.Error()})
	}
	helpers.ResponseJson(200,"ok",map[string] interface{}{"logs":res},w)
}