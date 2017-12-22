package elm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallBack(r *http.Request, w http.ResponseWriter) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	Notify := new(Notify)
	var responesDate map[string]interface{}
	err = Notify.Check(body, &responesDate)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	responesJson, err := json.Marshal(responesDate)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(responesJson)
}
