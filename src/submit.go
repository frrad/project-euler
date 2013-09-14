package main

import (
	//"euler"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
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

func main() {

	client := &http.Client{}

	jar := &myjar{}
	jar.jar = make(map[string][]*http.Cookie)
	client.Jar = jar

	/* Get Details */
	req, _ := http.NewRequest("GET", "http://projecteuler.net/login", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.22 Safari/537.36")
	resp, err := client.Do(req)

	fmt.Println(err)
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	fmt.Println(resp.ContentLength)

	d := time.Duration(10000000000)
	fmt.Println("seconds:", d.Seconds())
	time.Sleep(d)

	b, err := ioutil.ReadAll(resp.Body)

	fmt.Println("ERROR : ", err)

	resp.Body.Close()
	fmt.Println("Login Page:", string(b))

	form := make(url.Values)
	form.Set("username", "antest")
	form.Set("password", "password")
	form.Set("remember", "1")
	fmt.Println(form)
	/* Authenticate */
	resp, err = client.PostForm("http://projecteuler.net/login", form)

	b, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(b))

	if err != nil {
		fmt.Printf("Error Authenticating: %s", err)
	}

	/* Get Details */
	resp, err = client.Get("http://projecteuler.net/problems")
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	b, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(b))

}
