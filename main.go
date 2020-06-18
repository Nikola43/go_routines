package main

import (
	"fmt"
	"github.com/nikola43/keysklub/utils"
	"net/http"

	"github.com/panjf2000/ants"
)

type Request struct {
	Param  []byte
	Result chan []byte
}

type User struct {
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
}

type UserResult struct {
	Result chan User `json:"result"`
}

var pool *ants.Pool

func main() {
	pool, _ = ants.NewPool(100000)
	defer pool.Release()

	http.HandleFunc("/reverse", hand)

	http.ListenAndServe(":8080", nil)
}

func hand(w http.ResponseWriter, r *http.Request) {
	// create empty user and user result
	user := &User{}
	userResult := UserResult{
		Result: make(chan User),
	}

	// decode request
	err := utils.DecodeHttpRequestPayload(w, r, user)
	if err != nil {
		http.Error(w, "request error", http.StatusInternalServerError)
	}
	fmt.Println(user)

	// execute task
	poolError := pool.Submit(func() {
		u := User{
			Name:     "jursd",
			LastName: "pape",
		}

		userResult.Result <- u
	})
	if poolError != nil {
		fmt.Println(poolError)
	}

	// print msg

	msg := <-userResult.Result
	fmt.Println(msg)

	//w.Write([]byte(msg))

	utils.RespondHttpRequest(w, nil, msg)
}

/*

package main

import (
	"github.com/nikola43/keysklub/utils"
	"net/http"

	"github.com/panjf2000/ants"
)

type Request struct {
	Param  []byte
	Result chan []byte
}

type User struct {
	Name     string      `json:"name"`
	LastName string      `json:"last_name"`
	Result   chan []byte `json:"result"`
}

var pool *ants.PoolWithFunc

func main() {
	pool, _ = ants.NewPoolWithFunc(100000, reverseString)
	defer pool.Release()

	http.HandleFunc("/reverse", hand)

	http.ListenAndServe(":8080", nil)
}

func hand(w http.ResponseWriter, r *http.Request) {
	user := &User{
		Result: make(chan []byte),
	}
	err := utils.DecodeHttpRequestPayload(w, r, user)
	if err != nil {
		http.Error(w, "request error", http.StatusInternalServerError)
	}

	// Throttle the requests traffic with ants pool. This process is asynchronous and
	// you can receive a result from the channel defined outside.
	if err := pool.Invoke(user); err != nil {
		http.Error(w, "throttle limit error", http.StatusInternalServerError)
	}

	//w.Write(<-user.Result)

	utils.RespondHttpRequest(w,nil, user.Result)
}

func reverseString(payload interface{}) {
	user, ok := payload.(*User)
	if !ok {
		return
	}

	user.Result <- []byte("hola juan")
}

*/
