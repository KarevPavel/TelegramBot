package rutracker

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const BaseUrl string = "http://rutrackerripnext.onion"
const LoginUrl = BaseUrl + "/forum/login.php"
const SearchUrl = BaseUrl + "/forum/tracker.php?nm=${searchString}"

const Login string = "zwy706"
const Password string = "Vjqghjgecr09559c"

type LoginRequest struct {
	redirect string
	loginUser string
	loginPassword string
	login string
}

type Rutracker struct {
	customProxy string
	client *http.Client
	cookie *cookiejar.Jar
}

//http://rutrackerripnext.onion/forum/login.php

/*POST
http://rutrackerripnext.onion/forum/login.php
redirect=index.php&login_username=zwy706&login_password=%09Vjqghjgecr09559c&login=%C2%F5%EE%E4*/


func (loginRequest *LoginRequest) urlEncode() string {
	sb := strings.Builder{};
	sb.WriteString("redirect=")
	if len(loginRequest.redirect) > 0 {
		sb.WriteString(loginRequest.redirect)
	} else {
		sb.WriteString("index.php")
	}
	sb.WriteString("&")
	sb.WriteString("loginUser=")
	sb.WriteString(loginRequest.loginUser)
	sb.WriteString("&")
	sb.WriteString("loginPassword=")
	sb.WriteString("%09")
	sb.WriteString(loginRequest.loginPassword)
	sb.WriteString("&")
	sb.WriteString("login=")
	sb.WriteString("%C2%F5%EE%E4")
	return sb.String()
}

func NewRutracker(customProxy string) *Rutracker {
	proxyURL, err := url.Parse(customProxy)
	if err != nil {
		log.Println(err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	options := cookiejar.Options{
		PublicSuffixList: nil,
	}

	jar, _ := cookiejar.New(&options)

	client := &http.Client{
		Transport: transport,
		Jar:       jar,
	}

	tracker := &Rutracker{
		customProxy: customProxy,
		client:      client,
		cookie:      jar,
	}

	tracker.login(LoginRequest{
		loginUser:     Login,
		loginPassword: Password,
	})
	return tracker
}


func (tracker *Rutracker) poll() {
	for {
		ch := make(chan interface{})
		go tracker.dummyRequest(ch)
	}
}

func (tracker *Rutracker) dummyRequest(respChan chan interface{}) {
	response, _ := tracker.client.Get(BaseUrl)
	respChan <- response
}

func (tracker *Rutracker) login(login LoginRequest) {
	log.Println("Try to login into rutracker")
	req, _ := http.NewRequest("POST", LoginUrl, strings.NewReader(login.urlEncode()))
	response, httpError := tracker.client.Do(req)
	if httpError != nil {
		log.Panic("Error!", httpError)
	}
	url, _ := url.Parse(BaseUrl)
	tracker.cookie.SetCookies(url, response.Cookies())
	log.Printf("Save coockie: %s", response.Cookies())
	log.Println("Alright! Start session!")
}

func (tracker *Rutracker) Search(searchString string)  {
	req, _ := http.NewRequest("POST", SearchUrl, strings.NewReader("max=1&nm=test"))
	response, httpError := tracker.client.Do(req)
	if httpError != nil {
		log.Panic("Error!", httpError)
	}
	log.Println("response, %s" , response)
}

func (tracker *Rutracker) getTorrent()  {

}

func (tracker *Rutracker) getMagnet()  {

}