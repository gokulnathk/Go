package main

import ("net/http"; "io"; "os"; "log")

func main() {
    out, err := os.Create("D:\\output.rar")
    defer out.Close()
    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Get("http://192.168.1.25:8888/Fileshare_1.rar")
    defer resp.Body.Close()
    if err != nil {
        log.Fatal(err)
    }

    n, err := io.Copy(out, resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Fatal(n)
}