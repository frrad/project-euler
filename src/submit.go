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

var client = &http.Client{}
var authenticated bool = false

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
func getStatus() {

	if !authenticated {
		auth(client)
	}

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
	authenticated = true

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
func submit(problem int, solution string) (worked bool, message string) {
	if !authenticated {
		auth(client)
	}

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
		if strings.Contains(page, "Completed on ") || strings.Contains(page, "Go to the thread for") {
			say("Problem Already Completed", 1)
			answerStart := strings.Index(page, "Answer:")
			trunc := page[answerStart:]
			aStart := strings.Index(trunc, "<b>")
			aEnd := strings.Index(trunc, "</b>")
			correctAnswer := trunc[aStart+3 : aEnd]

			if correctAnswer == solution {
				return true, ""
			} else {
				return false, "Problem Solved: " + correctAnswer
			}

		} else {

			return false, "Can't find captcha in problem page."
		}
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
	form.Set("confirm", captcha)

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

func crackCap(b []byte) (crack string) {
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

func fancySubmit(x int, ans string) bool {

	if known, correct := check(x, ans); !known {

		if worked, mess := submit(x, ans); worked {
			say("Correct!", 0)
			say(penet+"/thread="+strconv.Itoa(x), 1)

			list(x, ans)

			getStatus()

			return true

		} else if len(mess) > 14 && mess[:14] == "Problem Solved" {
			say("Wrong answer! (Problem Already Solved)", 0)

			list(x, mess[16:])

		} else {
			say(mess, 0)
		}

	} else {
		say("Answer in list:", 1)
		if correct {
			say("Correct!", 0)
			say(penet+"/thread="+strconv.Itoa(x), 1)

		} else {
			say("Wrong answer!", 0)

		}
	}

	return false

}

func list(x int, ans string) {
	known := getData(settings["knownPath"])

	if _, ok := known[strconv.Itoa(x)]; ok {
		say("Answer already in list.", 1)
		return
	}

	say("Adding answer to list...", 2)
	known[strconv.Itoa(x)] = ans
	putData(settings["knownPath"], known)
	say("Answer added to list.", 1)
}

func check(x int, ans string) (present, correct bool) {
	known := getData(settings["knownPath"])

	if rightAnswer, ok := known[strconv.Itoa(x)]; ok {
		if ans == rightAnswer {
			return true, true
		} else {
			return true, false
		}

	}

	return false, false

}

func parse(spec string) (list []int, err bool) {
	splitted := strings.Split(spec, "-")

	if len(splitted) > 2 || len(splitted) < 1 {
		return nil, true
	}

	if len(splitted) == 1 {
		if pnumber, err := strconv.Atoi(os.Args[1]); err == nil {
			return []int{pnumber}, false
		} else {
			return nil, true
		}
	}

	a, b := 1, 1

	if splitted[0] == "" {
		if temp, err := strconv.Atoi(splitted[1]); err == nil {
			b = temp
		} else {
			return nil, true
		}
	} else {
		tempa, err1 := strconv.Atoi(splitted[0])
		tempb, err2 := strconv.Atoi(splitted[1])
		if err1 != nil || err2 != nil {
			return nil, true
		}

		a, b = tempa, tempb

	}

	list = make([]int, b-a+1)
	for i := range list {
		list[i] = i + a
	}

	return
}

func main() {

	client = &http.Client{}

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

	switch len(os.Args) {
	case 1:
		say("No arguments!", 0)

	case 2:
		//Only one argument
		switch argue := os.Args[1]; argue {

		case "R":
			say("Updating Status:", 0)
			getStatus()

		default:
			if plist, err := parse(argue); err == false {

				for _, pnumber := range plist {

					say("Solving #"+strconv.Itoa(pnumber), 1)

					if works, mess, out := runProb(pnumber); works {
						say("Answer: "+out, 1)

						if mess != "" { //time
							say(mess, 2)
						}

						fancySubmit(pnumber, out)

					} else {
						fmt.Println(mess)
					}

					fmt.Print("\n")
				}

			} else {
				say("Can't parse argument!", 0)
			}
		}
	case 3:
		if pnumber, err := strconv.Atoi(os.Args[1]); err == nil {
			out := os.Args[2]
			say("Submitting: "+out, 1)
			fancySubmit(pnumber, out)
		} else {
			say("Can't parse problem number!", 0)
		}
	default:
		say("Too many arguments!", 0)

	}

}
