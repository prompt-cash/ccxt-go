package utils

import "net/http"

func HttpDumpForm(r *http.Request) string {
	r.ParseForm()
	var all = ""
	for k, v := range r.PostForm {
		all += k + "=" + v[0] + "\r\n"
	}
	return all
}
