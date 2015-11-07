package main

import (
    "fmt"
    "net/http"
)

func main() {
    ROVI_API_KEY := "42qpphjg5zc4ts3y9mz2sur8"
    SERVICE_ID := "75836" //Comcast Modesto
    POSTAL_CODE := "95355"

    tv_channels := []int{2,3,4,5,67,8,9,10,11,12,13,14,15,550,551,552,553,554,555,556,557,558,559}

    fmt.Println("Calling to get sourceId mapping for channels.")

    resp, err := http.Get("http://api.rovicorp.com/TVlistings/v9/listings/services/postalcode/" + POSTAL_CODE + "/info?locale=en-US&countrycode=US&msoid=" + SERVICE_ID + "&format=json&apikey=" + ROVI_API_KEY)
    if err != nil {
        panic("Failed on getting service info.")
    }

    fmt.Println(resp)
}
