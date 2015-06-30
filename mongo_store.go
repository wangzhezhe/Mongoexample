package main

import (
	"bufio"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"os"
	"strings"
)

var (
	username string
	password string
	host     string
	port     string
	instance string
)

type Command struct {
	Type  string `bson:"type"`
	Intro string `bson:"intro"`
}

type Dockerfile struct {
	Template string
	Type     string
}

func Config() {

	username = os.Getenv("MONGODB_USERNAME")
	password = os.Getenv("MONGODB_PASSWORD")
	host = os.Getenv("MONGODB_PORT_27017_TCP_ADDR")
	port = os.Getenv("MONGODB_PORT_27017_TCP_PORT")

	if len(host) == 0 {
		host = "localhost"
	}

	fmt.Println(port)
	if len(port) == 0 {
		port = "27017"
	}

	instance = os.Getenv("MONGODB_INSTANCE_NAME")
}

func main() {

	//os.Setenv("MONGODB_PORT_27017_TCP_ADDR", "10.10.105.204")
	//os.Setenv("MONGODB_PORT_27017_TCP_PORT", "27017")

	Config()

	conn := ""
	if len(username) > 0 {
		conn += username

		if len(password) > 0 {
			conn += ":" + password
		}

		conn += "@"
	}

	conn += fmt.Sprintf("%s:%s", host, port)

	fmt.Println("conn info:", conn)

	session, err := mgo.Dial(conn)

	//session, err := mgo.Dial("10.10.72.139:27017")
	//session, err := mgo.Dial("10.10.105.204:27017")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	//切换到对应的数据库
	c := session.DB("toturial").C("cmds")

	//read from the file
	//fmt.Println(os.Getwd())

	f, err := os.Open("tuto_a.md")

	if err != nil {
		panic(err)
	}
	readf := bufio.NewReader(f)

	Type := ""
	Intro := ""
	for {
		//读出内容保存为string 每次读到以'\n'为标记的位置
		line, err := readf.ReadString('\n')
		//fmt.Print(line)
		if err == io.EOF {
			break
		}

		if strings.Contains(line, "####") {

			if Type != "" && Intro != "" {
				fmt.Println("Type:", Type, "Intro:", Intro)
				err = c.Insert(&Command{Type, Intro})
				//清零
				Intro = ""
			}

			Type = strings.TrimLeft(line, "#### ")
			Type = strings.TrimRight(Type, "\n")
			fmt.Println(Type)
			continue

		}

		Intro = Intro + line

	}
	result := Command{}
	//注意存在mongodb中的 type 有一个\n
	err = c.Find(bson.M{"type": "FROM"}).One(&result)
	if err != nil {
		panic(err)
	}

}
