package restful

import (
	"fmt"
	//
	"net/http"
)

func Getuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are Get modify user %s", uid)
}

func Modifyuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are Post modify user %s", uid)

}

func Deleteuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are Del delete user %s", uid)
}

func Adduser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprint(w, "you are PUT add user %s", uid)
}
