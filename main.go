package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
    // CORSを許可するためのヘッダーを追加
    w.Header().Set("Access-Control-Allow-Origin", "*")

    omikuji := []string{"大吉", "中吉", "小吉", "吉", "末吉", "凶", "大凶"}
    rand.Seed(time.Now().UnixNano())
    result := omikuji[rand.Intn(len(omikuji))]

    fmt.Fprintf(w, "おみくじの結果は: %s", result)
}

func main() {
    http.HandleFunc("/omikuji", omikujiHandler)
    fmt.Println("サーバーがポート8080で起動しました...")
    http.ListenAndServe(":8080", nil)
}
