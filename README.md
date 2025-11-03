# TikTok Profile Scraper
This Golang program fetches and parses TikTok share links and extracts information like profile image URL and username.

**Features**
- Parses TikTok HTML using goquery <a href="https://github.com/PuerkitoBio/goquery">PuerkitoBio/goquery</a>
- Extracts Profile Image (`<img>` with class containing `ImgTikTokCoinLogo`) and Username (`<span>` with class containing `SpanPopupTitle`) 
## Installation & Setup
**1.** Clone the local repository
```
git clone https://github.com/kalenmcmillan/tiktok-share-to-profile.git
```
**2.** Initialize a Go module and install dependencies:
```
go mod init share-to-profile
go get github.com/bogdanfinn/tls-client
go get github.com/bogdanfinn/fhttp
go get github.com/bogdanfinn/tls-client/profiles
go get github.com/PuerkitoBio/goquery
go mod tidy
```
**3.** Replaced `SHARE_LINK_HERE` in `main.go` with a real shared video link.
### License
MIT License © 2025 Kalen McMillan