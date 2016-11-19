package utils

import (
    "encoding/json"
    "net/http"
)

type Result struct{
    Ret int
    Reason string
    Data interface{}
}


func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
    out := &Result{ret, reason, i}
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}