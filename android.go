package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"strings"
	"time"

	http "github.com/bogdanfinn/fhttp"
)

type AndroidClient struct {
	Fid                  string
	AndroidRegisterToken string
}

func (emp *AndroidClient) Install() WarpRegisterInput {
	var response WarpRegisterInput
	client := HttpClient()

	currentTime := time.Now()
	tosTimestamp := currentTime.Format("2006-01-02T15:04:05.00-07:00")

	emp.Fid = NewFid()
	InstallInput := FirebaseInstallRegisterInput{
		Fid:         emp.Fid,
		AppID:       AppId,
		AuthVersion: AuthVersion,
		SdkVersion:  SdkVersion,
	}
	jsonBytes, _ := json.Marshal(InstallInput)
	data := strings.NewReader(string(jsonBytes))

	req, err := http.NewRequest("POST", fmt.Sprintf("https://firebaseinstallations.googleapis.com/v1/projects/project-%s/installations", FirebaseProjectId), data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type":               {"application/json"},
		"Accept":                     {"application/json"},
		"Cache-Control":              {"no-cache"},
		"X-Android-Package":          {"com.cloudflare.onedotonedotonedotone"},
		"x-firebase-client":          {"kotlin/1.4.32 fire-installations/16.3.3 fire-analytics/17.6.0 fire-cls-ndk/17.0.0 fire-iid/20.3.0 fire-fcm/20.1.7_1p fire-android/ fire-cls/17.2.2 fire-core/19.3.1"},
		"x-firebase-client-log-type": {"3"},
		"X-Android-Cert":             {"3A595E52DD381BCEE86A82A089C9BDC78FD459BF"}, // Probably is dynamic, might need changing later on
		"x-goog-api-key":             {"AIzaSyD8EGrWU54WutcvV_JdaK5w5IlTFsxU7Nc"},  // Probably is dynamic, might need changing later on
		"User-Agent":                 {"Dalvik/2.1.0 (Linux; U; Android 9; SM-S908N Build/PQ3A.190705.06091305)"},
		"Host":                       {"firebaseinstallations.googleapis.com"},
		http.HeaderOrderKey: {
			"Content-Type",
			"Accept",
			"Cache-Control",
			"X-Android-Package",
			"x-firebase-client",
			"x-firebase-client-log-type",
			"X-Android-Cert",
			"x-goog-api-key",
			"User-Agent",
			"Host",
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

	var installResponse FirebaseInstallRegisterResponse

	err = json.Unmarshal(readBytes, &installResponse)
	if err != nil {
		log.Println(err)
		return response
	}

	emp.Fid = installResponse.Fid

	data = strings.NewReader(`X-subtype=1003331445306&sender=1003331445306&X-app_ver=3092&X-osv=28&X-cliv=fiid-20.3.0&X-gmsv=210613022&X-appid=` + url.QueryEscape(emp.Fid) + `&X-scope=*&X-Goog-Firebase-Installations-Auth=` + url.QueryEscape(installResponse.AuthToken.Token) + `&X-gmp_app_id=` + url.QueryEscape(AppId) + `&X-Firebase-Client=kotlin%2F1.4.32+fire-installations%2F16.3.3+fire-analytics%2F17.6.0+fire-cls-ndk%2F17.0.0+fire-iid%2F20.3.0+fire-fcm%2F20.1.7_1p+fire-android%2F+fire-cls%2F17.2.2+fire-core%2F19.3.1&X-firebase-app-name-hash=R1dAH9Ui7M-ynoznwBdw01tLxhI&X-Firebase-Client-Log-Type=1&X-app_ver_name=6.23&app=com.cloudflare.onedotonedotonedotone&device=3591402226202727333&app_ver=3092&info=g4in4TzMUk4TIJ42qPW9kfDjy_ckERk&gcm_ver=210613022&plat=0&cert=3a595e52dd381bcee86a82a089c9bdc78fd459bf&target_ver=32`) // Some parameter values might be dynamic, don't think it matters for now though
	req, err = http.NewRequest("POST", "https://android.apis.google.com/c2dm/register3", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Authorization": {"AidLogin 3591402226202727333:3026631415424630535"},
		"app":           {"com.cloudflare.onedotonedotonedotone"},
		"gcm_ver":       {"210613022"},
		"User-Agent":    {"Android-GCM/1.5 (gracelte PQ3A.190705.06091305)"},
		"content-type":  {"application/x-www-form-urlencoded"},
		"Host":          {"android.apis.google.com"},
		http.HeaderOrderKey: {
			"Authorization",
			"app",
			"gcm_ver",
			"User-Agent",
			"content-type",
			"Host",
		},
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return response
	}
	defer resp.Body.Close()

	readBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return response
	}

	emp.AndroidRegisterToken = strings.Split(string(readBytes), "=")[1]

	keyData := NewX25519KeyPair()

	response.Key = keyData.RegisterKey
	response.InstallID = emp.Fid
	response.FcmToken = emp.AndroidRegisterToken
	response.Tos = tosTimestamp
	response.Model = "Samsung SM-S908N"
	response.SerialNumber = emp.Fid
	response.Locale = "en_US"

	return response
}
