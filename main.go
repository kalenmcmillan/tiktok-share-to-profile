package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "strings"

    httpf "github.com/bogdanfinn/fhttp"
    tls_client "github.com/bogdanfinn/tls-client"
    "github.com/bogdanfinn/tls-client/profiles"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    jar := tls_client.NewCookieJar()
    options := []tls_client.HttpClientOption{
        tls_client.WithClientProfile(profiles.Chrome_120),
        tls_client.WithRandomTLSExtensionOrder(),
        tls_client.WithCookieJar(jar),
    }

    client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
    if err != nil {
        log.Fatalf("Create Client Fail: %v", err)
    }

    url := "SHARE_LINK_HERE"
    req, err := httpf.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Fatalf("Create Req Fail: %v", err)
    }

    req.Header = httpf.Header{
        "Host":             {"www.tiktok.com"},
        "Sec-Fetch-Dest":   {"document"},
        "User-Agent":       {"Mozilla/5.0 (iPhone; CPU iPhone OS 18_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.1 Mobile/15E148 Safari/604.1"},
        "Accept":           {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
        "Sec-Fetch-Site":   {"none"},
        "Sec-Fetch-Mode":   {"navigate"},
        "Accept-Language":  {"en-US,en;q=0.9"},
        "Priority":         {"u=0, i"},
        "Accept-Encoding":  {"gzip, deflate, br"},
        "Connection":       {"keep-alive"},
    }

    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("Request Fail: %v", err)
    }
    defer resp.Body.Close()

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Read Body Fail: %v", err)
    }
    body := string(bodyBytes)

    doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
    if err != nil {
        log.Fatalf("Parse HTML Fail: %v", err)
    }

    doc.Find("img").Each(func(i int, s *goquery.Selection) {
        if class, exists := s.Attr("class"); exists && strings.Contains(class, "ImgTikTokCoinLogo") {
            if src, ok := s.Attr("src"); ok {
                fmt.Println("Profile Image URL:", src)
            }
        }
    })

    doc.Find("span").Each(func(i int, s *goquery.Selection) {
        if class, exists := s.Attr("class"); exists && strings.Contains(class, "SpanPopupTitle") {
			parsed := strings.TrimSpace(s.Text())
			parts := strings.Split(parsed, " ")
            fmt.Println("Profile Username:", parts[2])
        }
    })
}
