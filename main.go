package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
	"unicode/utf8"
)

var askPage []byte
var indexPage []byte
var replyPage []byte
var db *sql.DB

func main() {

	var err error

	askPage, err = ioutil.ReadFile("./web_page/ask.html")
	if err != nil {
		return
	}
	indexPage, err = ioutil.ReadFile("./web_page/index.html")
	if err != nil {
		return
	}
	replyPage, err = ioutil.ReadFile("./web_page/reply.html")
	if err != nil {
		return
	}

	var askJs, indexJs, replyJs []byte
	askJs, err = ioutil.ReadFile("./web_page/js/ask_page.js")
	if err != nil {
		return
	}
	indexJs, err = ioutil.ReadFile("./web_page/js/index.js")
	if err != nil {
		return
	}
	replyJs, err = ioutil.ReadFile("./web_page/js/reply_page.js")
	if err != nil {
		return
	}

	var font []byte
	font, err = ioutil.ReadFile("./web_page/font/BY-Dodge-Rabbit-2.ttf")
	if err != nil {
		return
	}

	var askImg, noneImg, replyImg []byte
	askImg, err = ioutil.ReadFile("./web_page/img/Ask_Questions.png")
	if err != nil {
		return
	}
	noneImg, err = ioutil.ReadFile("./web_page/img/NONE.png")
	if err != nil {
		return
	}
	replyImg, err = ioutil.ReadFile("./web_page/img/See_Reply.png")
	if err != nil {
		return
	}

	http.HandleFunc("/ask", handleRequestAsk)
	http.HandleFunc("/", handleRequestIndex)
	http.HandleFunc("/reply", handleRequestReply)
	http.HandleFunc("/js/ask_page.js", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			_, _ = writer.Write(askJs)
		}
	})
	http.HandleFunc("/js/index.js", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			_, _ = writer.Write(indexJs)
		}
	})
	http.HandleFunc("/js/reply_page.js", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			_, _ = writer.Write(replyJs)
		}
	})
	http.HandleFunc("/font/BY-Dodge-Rabbit-2.ttf", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			_, _ = writer.Write(font)
		}
	})
	http.HandleFunc("/img/Ask_Questions.png", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			_, _ = writer.Write(askImg)
		}
	})
	http.HandleFunc("/img/NONE.png", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			_, _ = writer.Write(noneImg)
		}
	})
	http.HandleFunc("/img/See_Reply.png", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			_, _ = writer.Write(replyImg)
		}
	})

	if err != nil {
		return
	}

	db, err = sql.Open("mysql", "root:Ton@8177919@tcp(101.43.188.94:3306)/tree_hole?charset=utf8")
	if err != nil {
		fmt.Println("failed to log in the database server")
		fmt.Println(err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("failed to ping")
		fmt.Println(err.Error())
		return
	}
	defer func(db *sql.DB) {
		_ = ioutil.WriteFile("./running_info.log", []byte("progress were killed at: "+time.Now().String()), os.ModePerm)
		_ = db.Close()
	}(db)
	err = http.ListenAndServe("0.0.0.0:8080", nil)
}

func handleRequestAsk(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("header: ", request.Header)
	fmt.Println("URL: ", request.URL)
	fmt.Println("Method: ", request.Method)
	fmt.Println("Host: ", request.Host)
	fmt.Println("RequestURI: ", request.RequestURI)
	fmt.Println("Body: ", request.Body)

	if request.Method == "GET" {
		_, _ = writer.Write(askPage)
	}

	if request.Method == "POST" {
		var info fullFormat

		bytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return
		}

		var strInfo = string(bytes)
		fmt.Println(strInfo)

		err = json.Unmarshal([]byte(strInfo), &info)
		if err != nil {
			fmt.Println("unable to decode json")
			return
		}
		fmt.Println("id:" + info.UserId)
		fmt.Println("question:" + info.Question)
		exec, err := db.Exec("insert into main(user_id, question, answer) values(?, ?, 'Have not answer yet, please wait for our reply...')", info.UserId, info.Question)
		if err != nil {
			fmt.Println("data insert failed")
			return
		}
		var id int64
		id, err = exec.LastInsertId()
		if err != nil {
			return
		}
		fmt.Println("last insert id:" + strconv.FormatInt(id, 10))
	}
}

func handleRequestIndex(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		_, _ = writer.Write(indexPage)
	}
}

func handleRequestReply(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		_, _ = writer.Write(replyPage)
	} else if request.Method == "POST" {
		bytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return
		}
		var info fullFormat
		fmt.Println(utf8.DecodeRune(bytes))
		err = json.Unmarshal(bytes, &info)
		if err != nil {
			fmt.Println("unable to decode json")
			return
		}

		if info.UserId == "000" {
			fmt.Println("receive stop code, returned.")
		}

		query, err := db.Query("select * from main where user_id=?", info.UserId)
		if err != nil {
			fmt.Println("failed to query in server")
			fmt.Println(err.Error())
			return
		}
		var user_id, question, reply string
		query.Next()
		err = query.Scan(&user_id, &question, &reply)
		if err != nil {
			fmt.Println("didn't find data in database server, id: " + info.UserId)
			fmt.Println(err.Error())
			return
		}
		var responseData fullFormat
		responseData.UserId = user_id
		responseData.Question = question
		responseData.Reply = reply

		var responseJSON []byte
		responseJSON, err = json.Marshal(&responseData)
		if err != nil {
			fmt.Println("failed to generate json")
			return
		}
		fmt.Println(string(responseJSON))
		_, err = writer.Write(responseJSON)
		if err != nil {
			fmt.Println("connection closed")
			return
		}
	}
}

type fullFormat struct {
	UserId   string `json:"user_id"`
	Question string `json:"question"`
	Reply    string `json:"reply"`
}
