package main

import (
    "fmt"
    //"github.com/pravj/geopattern"
    //"picgen/geopattern"
    "picgen/geopattern/svg"
    //"os"
    //"regexp"
)

func main() {
    /*email := os.Args[1]
    var validated_email string
    if isEmailValid(email) {
        validated_email = email
    }
    args := map[string]string{"phrase": validated_email }
    gp := geopattern.Generate(args)*/
    rect_args := make(map[string]interface{})
    rect_args["fill"] = "#ae44dc"
    Svg := new(svg.SVG)
    Svg.SetHeight(400)
    Svg.SetWidth(400)
    Svg.Rect("0","0","100%","100%", rect_args)
    fmt.Println(Svg.Str())
    //fmt.Println(gp)
}
/*
func isEmailValid(e string) bool {
    emailRegex := regexp.MustCompile("^[a-zA-Z-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    if len(e) < 3 || len(e) > 254 {
        return false
    }

    return emailRegex.MatchString(e)
}*/
