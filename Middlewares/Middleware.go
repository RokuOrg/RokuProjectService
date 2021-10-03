package Middlewares

import (
	"RokuProject-Back-End/Models"
	"encoding/json"
	"fmt"
	"github.com/stretchr/objx"
	"io/ioutil"
	"net/http"
)

func VerifyUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-JWT-Token")

		if token == "" {
			http.Error(w, "no X-JWT-Token header found", 400)
			return
		}

		//Send Token to Auth API for verification
		client := &http.Client{}

		req, err := http.NewRequest("GET", "http://157.90.164.50/api/user/verify", nil)

		if err != nil {
			http.Error(w, "err", 400)
			return
		}

		req.Header.Add("X-JWT-Token", token)

		res, err := client.Do(req)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		if res.StatusCode >= 400 {
			http.Error(w, res.Status, res.StatusCode)
			return
		}

		defer res.Body.Close()

		//Parse Response body dynamically kinda
		var f interface{}

		bodyBytes, err := ioutil.ReadAll(res.Body)

		err = json.Unmarshal(bodyBytes, &f)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		m := f.(map[string]interface{})

		var response Models.ResponseMessage

		for k, v := range m {
			if k == "succes" {
				o := objx.New(v)
				fmt.Println(o)
				response.Succes = o.Get("v").Bool()
			}
			if k == "error" {
				o := objx.New(v)
				response.Object.Error = o.Get("v").String()
			}
			if k == "message" {
				o := objx.New(v)
				response.Object.Message = o.Get("v").String()
			}
			if k == "token" {
				o := objx.New(v)
				response.Object.Token = o.Get("v").String()
			}
			if k == "username" {
				o := objx.New(v)
				response.Object.Username = o.Get("v").String()
			}
			if k == "id" {
				o := objx.New(v)
				response.Object.Id = o.Get("v").Int()
			}
		}

		if response.Succes {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, response.Object.Message, 400)
			return
		}
	})
}
