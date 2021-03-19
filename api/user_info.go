package api

import (
	"encoding/json"
	"firebase.google.com/go/db"
	"fmt"
	"github.com/vshnv/messagelogger/firebase"
	"net/http"
)

type UserInfo struct {
	Ip       string `json:"ip"`
	Username string `json:"user"`
	Epoch    int64  `json:"time"`
}

type IpInfo struct {
	Ip    string `json:"ip"`
	Epoch int64  `json:"time"`
}

func (u *UserInfo) ExtractIpInfo() *IpInfo {
	return &IpInfo{
		Ip:    u.Ip,
		Epoch: u.Epoch,
	}
}

func HandleUserInfo(w http.ResponseWriter, r *http.Request, client *db.Client) {
	fmt.Println("Endpoint Hit: userInfo")
	var p UserInfo
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	go client.NewRef("userinfo").Child(p.Username).Push(firebase.Ctx, p.ExtractIpInfo())
	fmt.Println("User Added!")
}
