package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	http "github.com/bogdanfinn/fhttp"
)

type WarpClient struct {
	Id        string
	AuthToken string
	License   string
}

func (emp *WarpClient) Register() WarpRegisterResponse {
	var response WarpRegisterResponse

	androidClient := AndroidClient{}
	warpRegisterInput := androidClient.Install()
	fmt.Println(warpRegisterInput)

	client := HttpClient()

	jsonBytes, _ := json.Marshal(warpRegisterInput)
	data := strings.NewReader(string(jsonBytes))

	fmt.Println(string(jsonBytes))

	req, err := http.NewRequest("POST", "https://api.cloudflareclient.com/v0a3092/reg", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"CF-Client-Version": {"a-6.23-3092"},
		"Content-Type":      {"application/json; charset=UTF-8"},
		"Host":              {"api.cloudflareclient.com"},
		"User-Agent":        {"okhttp/3.12.1"},
		http.HeaderOrderKey: {
			"CF-Client-Version",
			"Content-Type",
			"Host",
			"User-Agent",
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer resp.Body.Close()

	readBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response
	}

	if resp.StatusCode != 200 {
		log.Println(string(readBytes), resp.StatusCode)
		return response
	}

	err = json.Unmarshal(readBytes, &response)
	if err != nil {
		log.Println(err)
		return response
	}

	emp.Id = response.ID
	emp.AuthToken = response.Token
	emp.License = response.Account.License

	return response
}

func (emp *WarpClient) PatchReferrer(id string) bool {
	client := HttpClient()

	var data = strings.NewReader(`{"referrer":"` + id + `"}`)
	req, err := http.NewRequest("PATCH", "https://api.cloudflareclient.com/v0a3092/reg/"+emp.Id, data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"CF-Client-Version": {"a-6.23-3092"},
		"Content-Type":      {"application/json; charset=UTF-8"},
		"Host":              {"api.cloudflareclient.com"},
		"User-Agent":        {"okhttp/3.12.1"},
		"Authorization":     {"Bearer " + emp.AuthToken},
		http.HeaderOrderKey: {
			"CF-Client-Version",
			"Content-Type",
			"Host",
			"User-Agent",
			"Authorization",
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		readBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return false
		}

		fmt.Println(string(readBytes))
		return false
	}

	return true
}

func ResolveReferralLink(referralLink string) string {
	client := HttpClient()
	client.SetFollowRedirect(true)

	req, err := http.NewRequest("GET", referralLink+"?_imcp=1", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"accept-language":           {"en-US,en;q=0.9"},
		"cache-control":             {"no-cache"},
		"pragma":                    {"no-cache"},
		"priority":                  {"u=0, i"},
		"sec-ch-ua":                 {`"Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"`},
		"sec-ch-ua-mobile":          {"?0"},
		"sec-ch-ua-platform":        {`"macOS"`},
		"sec-fetch-dest":            {"document"},
		"sec-fetch-mode":            {"navigate"},
		"sec-fetch-site":            {"none"},
		"sec-fetch-user":            {"?1"},
		"upgrade-insecure-requests": {"1"},
		"user-agent":                {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"},
		http.HeaderOrderKey: {
			"accept",
			"accept-language",
			"cache-control",
			"pragma",
			"priority",
			"sec-ch-ua",
			"sec-ch-ua-mobile",
			"sec-ch-ua-platform",
			"sec-fetch-dest",
			"sec-fetch-mode",
			"sec-fetch-site",
			"sec-fetch-user",
			"upgrade-insecure-requests",
			"user-agent",
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != 404 {
		log.Fatalln("unable to resolve referral link")
	}
	return resp.Request.URL.Query().Get("referrer")
}
