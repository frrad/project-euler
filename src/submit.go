package main

import (
	"euler"
	"fmt"
	"io/ioutil"
"time"
	"net/http"
"os/exec"
	"net/url"
"strconv"
"strings"
	)

var settings map[string]string
const permissions = 777
const setPath = "../eulerdata/settings.dat"

type myjar struct {
	jar map[string][]*http.Cookie
}

func (p *myjar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	//fmt.Printf("The URL is : %s\n", u.String())
	//fmt.Printf("The cookie being set is : %s\n", cookies)
	p.jar[u.Host] = cookies
}

func (p *myjar) Cookies(u *url.URL) []*http.Cookie {
	//fmt.Printf("The URL is : %s\n", u.String())
	//fmt.Printf("Cookie being returned is : %s\n", p.jar[u.Host])
	return p.jar[u.Host]
}

//given an authenticated client writes status.html to given path
func getStatus(client *http.Client, path string){

	resp, err := client.Get(settings["statPath"])
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	ioutil.WriteFile(path, b, permissions)

	//fmt.Println(string(b))
}

func auth(client *http.Client, uname, pass string){
	form := make(url.Values)
	form.Set("username", "antest")
	form.Set("password", "password")
	form.Set("remember", "1")
	form.Set("login","Login")

	// Authenticate
	_, err := client.PostForm("http://projecteuler.net/login", form)
	if err != nil {
		fmt.Printf("Error Authenticating: %s", err)
	}
}

func getSettings(path string) map[string]string{
	sets := euler.Import(path)

	out:= make(map[string]string)

	for _, line:= range sets{
		two := strings.SplitN(line,":",2)
		out[two[0]] = two[1]
	}	

	return out
}

//takes authenticated client, problem number and solution: submits answer
func submit(client *http.Client, problem int, solution string){
	pname := strconv.Itoa(problem)
	url := "http://projecteuler.net/problem="+pname

	fmt.Println("Fetching Problem...",problem)
	resp, err := client.Get(url)
	fmt.Println("Page Downloaded.")

	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	page := string(b)

	capStart := strings.Index(page, "<img src=\"captcha")
	if capStart == -1{
		panic("no captcha")
	}
	capEnd := strings.Index(page[capStart+10:],"\"")
	capURL := page[capStart + 10:capStart + 10 + capEnd]

	fmt.Println("Downloading Captcha...")
	resp, err = client.Get("http://projecteuler.net/"+capURL)
	fmt.Println("Captcha Downloaded.")	

	if err != nil {
		fmt.Printf("Error : %s", err)
	}	



	b, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Println("Cracking Captcha...")
	captcha := crackCap(b)

	fmt.Println("Captcha Solved:",captcha)


}

func crackCap( b []byte) (crack int){
	timeStr := strconv.FormatInt((time.Now().Unix()),10)
	path := settings["capPath"]+ timeStr+ ".png"

	ioutil.WriteFile(path, b, permissions)

	do := exec.Command("ristretto",path)
	do.Start()	
	

	fmt.Println("Please Input Captcha:")
	fmt.Scan(&crack)	

	return
}

func main() {

	client := &http.Client{}

	jar := &myjar{}
	jar.jar = make(map[string][]*http.Cookie)
	client.Jar = jar



	

	settings = make(map[string]string)
	settings["statPath"] = "http://projecteuler.net/progress"
	settings["capPath"]= "../eulerdata/captcha/" //trailing slash!

	fmt.Println("Reading settings from file...")
	fileSets := getSettings(setPath)
	for key, val := range fileSets{
		//settings from file overwrite defaults
		settings[key] = val
	}


	fmt.Println("Authenticating...")
	auth(client, settings["username"],settings["password"])
	fmt.Println("Authentication Complete.")

	submit(client, 11, "")


	//getStatus(client, "../eulerdata/status.html")

}
