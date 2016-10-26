package cocapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var urlClan = "https://api.clashofclans.com/v1/clans/%s"
var urlMembers = "https://api.clashofclans.com/v1/clans/%s/members"
var myKey, myClanTag string

type ServerError struct {
	msg       string
	ErrorCode int
}

func (e *ServerError) Error() string {
	return e.msg
}

func init() {
	myKey = os.Getenv("COC_KEY")
	myClanTag = os.Getenv("COC_CLANTAG")
}

type ClanInfo struct {
	BadgeUrls struct {
		Large  string `json:"large"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
	} `json:"badgeUrls"`
	ClanLevel   int    `json:"clanLevel"`
	ClanPoints  int    `json:"clanPoints"`
	Description string `json:"description"`
	Location    struct {
		ID        int    `json:"id"`
		IsCountry bool   `json:"isCountry"`
		Name      string `json:"name"`
	} `json:"location"`
	MemberList []struct {
		ClanRank          int `json:"clanRank"`
		Donations         int `json:"donations"`
		DonationsReceived int `json:"donationsReceived"`
		ExpLevel          int `json:"expLevel"`
		League            struct {
			IconUrls struct {
				Medium string `json:"medium"`
				Small  string `json:"small"`
				Tiny   string `json:"tiny"`
			} `json:"iconUrls"`
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"league"`
		Name             string `json:"name"`
		PreviousClanRank int    `json:"previousClanRank"`
		Role             string `json:"role"`
		Trophies         int    `json:"trophies"`
	} `json:"memberList"`
	Members          int    `json:"members"`
	Name             string `json:"name"`
	RequiredTrophies int    `json:"requiredTrophies"`
	Tag              string `json:"tag"`
	Type             string `json:"type"`
	WarFrequency     string `json:"warFrequency"`
	WarWins          int    `json:"warWins"`
}

type Members struct {
	Items []struct {
		ClanRank          int `json:"clanRank"`
		Donations         int `json:"donations"`
		DonationsReceived int `json:"donationsReceived"`
		ExpLevel          int `json:"expLevel"`
		League            struct {
			IconUrls struct {
				Medium string `json:"medium"`
				Small  string `json:"small"`
				Tiny   string `json:"tiny"`
			} `json:"iconUrls"`
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"league"`
		Tag              string `json:"tag"`
		Name             string `json:"name"`
		PreviousClanRank int    `json:"previousClanRank"`
		Role             string `json:"role"`
		Trophies         int    `json:"trophies"`
	} `json:"items"`
}

func GetMemberInfo() (members Members, err error) {
	body, err := getUrl(fmt.Sprintf(urlMembers, url.QueryEscape(myClanTag)), myKey)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &members)
	return
}

func getClanInfo(clanTag string) (clan ClanInfo, err error) {
	body, err := getUrl(fmt.Sprintf(urlClan, url.QueryEscape(clanTag)), myKey)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &clan)
	return
}

func getUrl(url, key string) (b []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Add("authorization", "Bearer "+key)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		b = []byte{}
		//err = errors.New("Error from server: " + strconv.Itoa(resp.StatusCode))
		err = &ServerError{msg: "Error from server: " + strconv.Itoa(resp.StatusCode), ErrorCode: resp.StatusCode}
	}
	return
}
