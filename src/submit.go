package main

import (
	"euler"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var settings map[string]string

const permissions = 0666
const setPath = "../eulerdata/settings.dat"
const penet = "http://projecteuler.net"
const probCount = 1000 //some number > #problems

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

func say(message string, level int) {
	debugLevel, _ := strconv.Atoi(settings["debug"])
	if debugLevel >= level {
		fmt.Println(message)
	}
}

//given an authenticated client writes status.html to given path
func getStatus(client *http.Client) {

	say("Fetching progress page...", 2)
	resp, err := client.Get(penet + "/progress")
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	say("Progress page downloaded.", 1)

	say("Writing page to "+settings["statusPath"], 3)
	ioutil.WriteFile(settings["statusPath"], b, permissions)

	say(string(b), 5)

	//fmt.Println(string(b))
}

func auth(client *http.Client) {

	say("Authenticating...", 2)

	form := make(url.Values)
	form.Set("username", settings["username"])
	form.Set("password", settings["password"])
	form.Set("remember", "1")
	form.Set("login", "Login")

	// Authenticate
	_, err := client.PostForm(penet+"/login", form)
	if err != nil {
		fmt.Printf("Error Authenticating: %s", err)
	}

	say("Authenticated", 1)
}

func getData(path string) map[string]string {
	sets := euler.Import(path)

	out := make(map[string]string)

	for _, line := range sets {
		two := strings.SplitN(line, ":", 2)
		out[two[0]] = two[1]
	}

	return out
}

func putData(path string, data map[string]string) {
	out := ""
	for i := 0; i < probCount; i++ {
		word := strconv.Itoa(i)
		if ans, ok := data[word]; ok {
			out += word + ":" + ans + "\n"
		}
	}

	ioutil.WriteFile(path, proccess(out), permissions)
}

func proccess(a string) []byte {
	out := make([]byte, 0)
	for i := 0; i < len(a); i++ {
		out = append(out, a[i])
	}
	return out
}

//takes authenticated client, problem number and solution: submits answer
func submit(client *http.Client, problem int, solution string) (worked bool, message string) {
	pname := strconv.Itoa(problem)
	theURL := penet + "/problem=" + pname

	say("Fetching Problem... "+pname, 2)
	resp, err := client.Get(theURL)
	say("Problem Downloaded.", 1)

	if err != nil {
		return false, "Fetching problem failed"
	}

	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	page := string(b)

	capStart := strings.Index(page, "<img src=\"captcha")
	if capStart == -1 {
		return false, "Can't find captcha in problem page. Already Submitted?"
	}
	capEnd := strings.Index(page[capStart+10:], "\"")
	capURL := page[capStart+10 : capStart+10+capEnd]

	say("Downloading Captcha...", 2)
	resp, err = client.Get(penet + "/" + capURL)
	say("Captcha Downloaded.", 1)

	if err != nil {
		return false, "Fetching captcha failed."
	}

	b, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Println("Cracking Captcha...")
	captcha := crackCap(b)

	fmt.Println("Captcha Solved:", captcha)

	form := make(url.Values)
	form.Set("guess_"+pname, solution)
	form.Set("confirm", strconv.Itoa(captcha))

	fmt.Println("Submitting...")
	//Submit
	resp, err = client.PostForm(theURL, form)
	if err != nil {
		return false, "Trouble submitting solution"
	}

	b, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	page = string(b)

	if strings.Contains(page, "answer_wrong.png") {
		return false, "Wrong answer!"
	}

	if strings.Contains(page, "answer_correct.png") {
		return true, ""
	}

	if strings.Contains(page, "The confirmation code you entered was not valid") {

		return false, "Captcha Failed!"
	}

	return false, "wtf?"

}

func crackCap(b []byte) (crack int) {
	timeStr := strconv.FormatInt((time.Now().Unix()), 10)
	path := settings["capPath"] + timeStr + ".png"

	ioutil.WriteFile(path, b, permissions)

	do := exec.Command(settings["imageViewer"], path)
	do.Start()

	fmt.Println("Please Input Captcha:")
	fmt.Scan(&crack)

	return
}

func runProb(n int) (works bool, message, output string) {
	nstr := strconv.Itoa(n)
	for len(nstr) < 3 {
		nstr = "0" + nstr
	}
	nstr = "Problem" + nstr

	cmd := exec.Command("go", "run", nstr+".go")
	out, err := cmd.StdoutPipe()
	if err != nil {
		return false, "Trouble getting pipe.", ""
	}
	if err := cmd.Start(); err != nil {
		return false, "Trouble starting program.", ""
	}

	b, err := ioutil.ReadAll(out)
	if err != nil {
		return false, "Trouble reading output.", ""
	}

	if err := cmd.Wait(); err != nil {
		return false, "Program exitted with error.", ""
	}

	out.Close()

	programOutput := strings.Split(string(b), "\n")

	time := ""

	for i := len(programOutput) - 1; i >= 0; i-- {

		line := programOutput[i]

		say(line, 5)

		if strings.Contains(line, "Elapsed") {
			time = line

		}
		if len(line) > 0 && !strings.Contains(line, "Elapsed") {
			return true, time, line

		}
	}

	return false, "Can't find output", ""

}

func fancySubmit(client *http.Client, x int, ans string) bool {

	if worked, mess := submit(client, x, ans); worked {
		say("Correct!", 0)

		say("Adding answer to list...", 2)
		known := getData(settings["knownPath"])
		known[strconv.Itoa(x)] = ans
		putData(settings["knownPath"], known)
		say("Answer added to list.", 1)

		getStatus(client)

		return true

	} else {
		say(mess, 0)
	}
	return false

}

func main() {

	client := &http.Client{}

	jar := &myjar{}
	jar.jar = make(map[string][]*http.Cookie)
	client.Jar = jar

	settings = make(map[string]string)
	settings["capPath"] = "../eulerdata/captcha/" //trailing slash!
	settings["knownPath"] = "../eulerdata/known.txt"
	settings["statusPath"] = "../eulerdata/status.html"
	settings["imageViewer"] = "eog"
	settings["debug"] = "3"

	say("Reading settings from file...", 1)
	fileSets := getData(setPath)
	for key, val := range fileSets {
		//settings from file overwrite defaults
		settings[key] = val
	}

	if len(os.Args) == 1 {
		say("No arguments!", 0)
	} else if len(os.Args) > 3 {
		say("Too many arguments!", 0)
	} else if len(os.Args) == 2 && os.Args[1] == "R" {
		say("Updating Status:", 0)
		auth(client)
		getStatus(client)
	} else if pnumber, err := strconv.Atoi(os.Args[1]); err == nil {
		if len(os.Args) == 2 {

			say("Solving #"+strconv.Itoa(pnumber), 1)
			if works, mess, out := runProb(pnumber); works {
				say("Answer: "+out, 1)

				if mess != "" { //time
					say(mess, 2)
				}

				auth(client)
				fancySubmit(client, pnumber, out)

			} else {
				fmt.Println(mess)
			}
		}

		if len(os.Args) == 3 {
			out := os.Args[2]
			say("Submitting: "+out, 1)

			auth(client)
			fancySubmit(client, pnumber, out)

		}

	} else {
		say("Invalid Arguments", 0)
	}

}