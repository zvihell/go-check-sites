package controllers

import (
	"check-domain-api/internal/models"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func (h *Handler) Worker() {
	sites := []string{
		"google.com",
		"youtube.com",
		"facebook.com",
		"baidu.com",
		"wikipedia.org",
		"qq.com",
		"taobao.com",
		"yahoo.com",
		"tmall.com",
		"amazon.com",
		"google.co.in",
		"twitter.com",
		"sohu.com",
		"jd.com",
		"live.com",
		"instagram.com",
		"sina.com.cn",
		"weibo.com",
		"google.co.jp",
		"reddit.com",
		"vk.com",
		"360.cn",
		"login.tmall.com",
		"blogspot.com",
		"yandex.ru",
		"google.com.hk",
		"netflix.com",
		"linkedin.com",
		"pornhub.com",
		"google.com.br",
		"twitch.tv",
		"pages.tmall.com",
		"csdn.net",
		"yahoo.co.jp",
		"mail.ru",
		"aliexpress.com",
		"alipay.com",
		"office.com",
		"google.fr",
		"google.ru",
		"google.co.uk",
		"microsoftonline.com",
		"google.de",
		"ebay.com",
		"microsoft.com",
		"livejasmin.com",
		"t.co",
		"bing.com",
		"xvideos.com",
		"google.ca"}

	h.CheckSites(sites)
}

func (h *Handler) CheckSites(sites []string) {
	for {
		var wg sync.WaitGroup
		for _, site := range sites {
			wg.Add(1)
			go func(site string) {
				defer wg.Done()
				resp, err := http.Get("https://" + site)
				if err != nil {
					return
				}
				defer resp.Body.Close()
				t := time.Now()
				elapsed := time.Since(t)
				available := true
				if resp.StatusCode != 200 {
					available = false
				}

				dom := models.Domain{
					Domain:      site,
					Latency:     elapsed,
					Available:   available,
					Last_update: t,
				}

				h.Save(dom)

			}(site)
		}
		wg.Wait()
		time.Sleep(time.Minute)
	}
}

func (h *Handler) Save(dom models.Domain) {
	err := h.domainService.Save(dom)
	if err != nil {
		fmt.Println("An error occurred", err)
		return
	}
}
