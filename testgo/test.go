package main

import (
	"fmt"
  "encoding/json"
	"strconv"
)


func main() {
    // url := "http://localhost:8080/v4/threatMatches:find"

    surl := []string {
              "google.com",
              "http://testsafebrowsing.appspot.com/apiv4/ANY_PLATFORM/MALWARE/URL/",
              "http://testsafebrowsing.appspot.com/apiv4/ANY_PLATFORM/SOCIAL_ENGINEERING/URL/",
              "http://testsafebrowsing.appspot.com/apiv4/ANY_PLATFORM/UNWANTED_SOFTWARE/URL/",
    }


    threatEntries := make(map[string][]map[string]string)

    for k, v := range surl {
        tmap := map[string] string{
            "url" : v,
        }
        threatEntries[strconv.Itoa(k)] = append(threatEntries[strconv.Itoa(k)], tmap)
    }

    // prepost := map[string] map {
    //     "threatInfo" : map[string] map {
    //         "threatEntries" : threatEntries
    //     }
    // }

fmt.Println(threatEntries)

    postbody, err := json.Marshal(threatEntries)
    if err != nil {
        fmt.Print(err)
        return
    }

    fmt.Println(postbody)
    //
    // req, err := http.Post(url, "application/json", postbody)

    // req, err := http.NewRequest("POST", url, nil)
    // if err != nil {
		//     fmt.Println("NewRequest: ", err)
	  //     return
    // }
    //
    // client := &http.Client{}
    //
    // resp, err := client.Do(req)
    // if err != nil {
    //     fmt.Println("Do: ", err)
	  //     return
    // }
    //
    // fmt.Print(resp)
    //
    // // defer resp.Body.Close()
    //
    // if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		//     fmt.Println(err)
  	// }
    //
    // fmt.Print(record)
}
