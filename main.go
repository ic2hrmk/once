package main

import (
	"once/once"

	"flag"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"strings"
)

var filename string

func init() {
	flag.StringVar(&filename, "conf", "once.conf", "File path to configuration file")
	flag.Parse()
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		r.ParseForm()
		shortenLink, err := once.GenerateShortLink(r.Form.Get("url"))
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Println(err.Error())
		} else {
			jsonResponse, _ := json.Marshal(struct {
				Link string `json:"short_link"`
			}{
				Link: shortenLink,
			})

			w.Write(jsonResponse)
			log.Println("added new link: " + shortenLink)
		}
	case http.MethodGet:
		token := strings.Replace(r.URL.Path, "/", "", 1)

		if once.IsShortLinkUsed(token) {
			http.Error(w, "", http.StatusNotFound)
			log.Println("token " + token + " is used or not found")
		} else {
			longUrl, err := once.GetShortLinkValue(token)
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				log.Println(err.Error())
			} else {
				once.SetShortLinkAsUsed(token)
				http.Redirect(w, r, longUrl, http.StatusFound)
				log.Println("token " + token + " was used")
			}
		}
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func main() {
	configuration, err := NewConfiguration(filename)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = once.InitOnce(configuration.OnceConfiguration)
	if err != nil {
		log.Fatal(err.Error())
	}

	port := ":" + strconv.Itoa(configuration.Port)
	log.Println("Short link server started at port " + port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))
}