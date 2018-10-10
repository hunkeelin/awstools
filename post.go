package main

import (
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

func post(w http.ResponseWriter, r *http.Request) error {
    var p payload
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("unable to read response body post")
        return err
    }
    err = json.Unmarshal(b, &p)
    if err != nil {
        fmt.Println("unable to unmarshal json post")
        return err
    }
    switch strings.ToLower(p.Action) {
    case "reboot":
        err = reboot(p.Nametag)
        if err != nil {
            return err
        }
    default:
        return errors.New("Invalid Action please specify correct action")
    }
    return nil
}

