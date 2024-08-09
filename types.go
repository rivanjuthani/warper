package main

type FirebaseInstallRegisterInput struct {
	Fid         string `json:"fid"`
	AppID       string `json:"appId"`
	AuthVersion string `json:"authVersion"`
	SdkVersion  string `json:"sdkVersion"`
}

type FirebaseInstallRegisterResponse struct {
	Name         string `json:"name"`
	Fid          string `json:"fid"`
	RefreshToken string `json:"refreshToken"`
	AuthToken    struct {
		Token     string `json:"token"`
		ExpiresIn string `json:"expiresIn"`
	} `json:"authToken"`
}

type WarpRegisterInput struct {
	Key          string `json:"key"`
	InstallID    string `json:"install_id"`
	FcmToken     string `json:"fcm_token"`
	Tos          string `json:"tos"`
	Model        string `json:"model"`
	SerialNumber string `json:"serial_number"`
	Locale       string `json:"locale"`
}

type WarpRegisterResponse struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Model   string `json:"model"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Account struct {
		ID                       string `json:"id"`
		AccountType              string `json:"account_type"`
		Created                  string `json:"created"`
		Updated                  string `json:"updated"`
		PremiumData              int    `json:"premium_data"`
		Quota                    int    `json:"quota"`
		Usage                    int    `json:"usage"`
		WarpPlus                 bool   `json:"warp_plus"`
		ReferralCount            int    `json:"referral_count"`
		ReferralRenewalCountdown int    `json:"referral_renewal_countdown"`
		Role                     string `json:"role"`
		License                  string `json:"license"`
		TTL                      string `json:"ttl"`
	} `json:"account"`
	Config struct {
		ClientID string `json:"client_id"`
		Peers    []struct {
			PublicKey string `json:"public_key"`
			Endpoint  struct {
				V4    string `json:"v4"`
				V6    string `json:"v6"`
				Host  string `json:"host"`
				Ports []int  `json:"ports"`
			} `json:"endpoint"`
		} `json:"peers"`
		Interface struct {
			Addresses struct {
				V4 string `json:"v4"`
				V6 string `json:"v6"`
			} `json:"addresses"`
		} `json:"interface"`
		Services struct {
			HTTPProxy string `json:"http_proxy"`
		} `json:"services"`
	} `json:"config"`
	Token           string `json:"token"`
	WarpEnabled     bool   `json:"warp_enabled"`
	WaitlistEnabled bool   `json:"waitlist_enabled"`
	Created         string `json:"created"`
	Updated         string `json:"updated"`
	Tos             string `json:"tos"`
	Place           int    `json:"place"`
	Locale          string `json:"locale"`
	Enabled         bool   `json:"enabled"`
	InstallID       string `json:"install_id"`
	FcmToken        string `json:"fcm_token"`
	SerialNumber    string `json:"serial_number"`
}
