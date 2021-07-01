package profile

import (
	"chasqi/data/network"
	"io/ioutil"
	"strconv"
	"strings"
)

type ProfileApi struct {
	Host              string
	GetProfilePath    string
	UpdateProfilePath string
	Connector         *network.NetworkConnector
}

func (p *ProfileApi) GetProfile() {
	println("Getting Profile")
	var url strings.Builder
	url.WriteString(p.Host)
	url.WriteString(p.GetProfilePath)

	resp, err := p.Connector.GetAuthenticated(url.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println("Received Profile")
	println(strconv.Itoa(resp.StatusCode))
	println(string(bytes))
}

func (p *ProfileApi) UpdateProfile(fcmToken string) {
	println("Getting Profile")
	var url strings.Builder
	url.WriteString(p.Host)
	url.WriteString(p.GetProfilePath)
	type FcmTokenUpdate struct {
		FcmToken string `json:"fcm_token"`
	}
	update := FcmTokenUpdate{FcmToken: fcmToken}
	resp, err := p.Connector.PutAuthenticated(url.String(), update)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println("Received Profile")
	println(strconv.Itoa(resp.StatusCode))
	println(string(bytes))
}
