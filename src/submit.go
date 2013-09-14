package main

import (
	//"euler"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	)

type myjar struct {
	jar map[string][]*http.Cookie
}

func (p *myjar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	fmt.Printf("The URL is : %s\n", u.String())
	fmt.Printf("The cookie being set is : %s\n", cookies)
	p.jar[u.Host] = cookies
}

func (p *myjar) Cookies(u *url.URL) []*http.Cookie {
	fmt.Printf("The URL is : %s\n", u.String())
	fmt.Printf("Cookie being returned is : %s\n", p.jar[u.Host])
	return p.jar[u.Host]
}

func getStatus(client *http.Client, path string){


	resp, err := client.Get("http://projecteuler.net/progress")
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	ioutil.WriteFile(path, b, 777)

	//fmt.Println(string(b))
}

func main() {

	client := &http.Client{}

	jar := &myjar{}
	jar.jar = make(map[string][]*http.Cookie)
	client.Jar = jar

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





	/* Get Details */
	getStatus(client, "../eulerdata/status.html")

}
