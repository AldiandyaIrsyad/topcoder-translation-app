package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Request struct {
	Txt  string `json:"q"`
	From string `json:"sl"`
	To   string `json:"tl"`
}

func main() {
	con := "n"
	for {
		var query Request
		// var type string = "Text"

		stdin := bufio.NewReader(os.Stdin)

		fmt.Println("insert text to translate (if empty use default text)")
		query.Txt, _ = stdin.ReadString('\n')
		fmt.Println("insert source language (if empty use default source language)")
		query.From, _ = stdin.ReadString('\n')
		fmt.Println("insert target language(if empty use default text)")
		query.To, _ = stdin.ReadString('\n')

		query.Txt = strings.TrimRight(query.Txt, "\r\n")
		query.From = strings.TrimRight(query.From, "\r\n")
		query.To = strings.TrimRight(query.To, "\r\n")

		if query.Txt == "" {
			query.Txt = "Hi nama aku aldiandya, aku berasal dari Indonesia ini Challenge pertamaku untuk belajar menjadi pengembang web"
		}

		if query.From == "" {
			query.From = "id"
		}

		if query.To == "" {
			query.To = "en"
		}

		var url string = "https://translate.google.com/translate_a/single"
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("User-Agent", "AndroidTranslate/5.3.0.RC02.130475354-53000263 5.1 phone TRANSLATE_OPM5_TEST_1")

		// set query
		q := req.URL.Query()
		q.Add("iid", "1dd3b944-fa62-4b55-b330-74909a99969e")
		q.Add("client", "at")
		q.Add("dt", "t")
		q.Add("dt", "ld")
		q.Add("dt", "qca")
		q.Add("dt", "rm")
		q.Add("dt", "bd")
		q.Add("dj", "1")
		q.Add("hl", "%s")
		q.Add("ie", "UTF-8")
		q.Add("oe", "UTF-8")
		q.Add("inputm", "2")
		q.Add("otf", "2")
		q.Add("sl", query.From)
		q.Add("tl", query.To)
		q.Add("q", query.Txt)
		req.URL.RawQuery = q.Encode()

		if err != nil {
			fmt.Println("Cannot post request")
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		// fmt.Println("response Body:", string(body))

		body_output := make(map[string]interface{})
		err = json.Unmarshal([]byte(body), &body_output)
		if err != nil {
			return
		}

		fmt.Println("Translated text from " + query.From + " to " + query.To)
		fmt.Println(body_output["sentences"].([]interface{})[0].(map[string]interface{})["trans"])

		fmt.Println("\nTranslate Again? (y/n)")
		con, _ = stdin.ReadString('\n')
		con = strings.TrimRight(con, "\r\n")

		if con == "n" {
			return
		}

		fmt.Print("\n\n")
	}

}
