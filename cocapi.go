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

var urlPlayers = "https://api.clashofclans.com/v1/players/%s"
var urlClan = "https://api.clashofclans.com/v1/clans/%s"
var urlMembers = "https://api.clashofclans.com/v1/clans/%s/members"

var myKey /*, myClanTag*/ string

type ServerError struct {
	msg       string
	ErrorCode int
}

func (e *ServerError) Error() string {
	return e.msg
}

func init() {
	myKey = os.Getenv("COC_KEY")
	//myClanTag = os.Getenv("COC_CLANTAG")
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
	MemberList       []Member `json:"memberList"`
	Members          int      `json:"members"`
	Name             string   `json:"name"`
	RequiredTrophies int      `json:"requiredTrophies"`
	Tag              string   `json:"tag"`
	Type             string   `json:"type"`
	WarFrequency     string   `json:"warFrequency"`
	WarWins          int      `json:"warWins"`
}

type Member struct {
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
	Tag              string `json:"tag"`
	Trophies         int    `json:"trophies"`
}

type Members struct {
	Items []Member `json:"items"`
}

type Player struct {
	Tag               string `json:"tag"`
	Name              string `json:"name"`
	TownHallLevel     int    `json:"townHallLevel"`
	ExpLevel          int    `json:"expLevel"`
	Trophies          int    `json:"trophies"`
	BestTrophies      int    `json:"bestTrophies"`
	WarStars          int    `json:"warStars"`
	AttackWins        int    `json:"attackWins"`
	DefenseWins       int    `json:"defenseWins"`
	Role              string `json:"role"`
	Donations         int    `json:"donations"`
	DonationsReceived int    `json:"donationsReceived"`
	Clan              struct {
		Tag       string `json:"tag"`
		Name      string `json:"name"`
		ClanLevel int    `json:"clanLevel"`
		BadgeUrls struct {
			Small  string `json:"small"`
			Large  string `json:"large"`
			Medium string `json:"medium"`
		} `json:"badgeUrls"`
	} `json:"clan"`
	League struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		IconUrls struct {
			Small  string `json:"small"`
			Tiny   string `json:"tiny"`
			Medium string `json:"medium"`
		} `json:"iconUrls"`
	} `json:"league"`
	Achievements []struct {
		Name           string `json:"name"`
		Stars          int    `json:"stars"`
		Value          int    `json:"value"`
		Target         int    `json:"target"`
		Info           string `json:"info"`
		CompletionInfo string `json:"completionInfo,omitempty"`
	} `json:"achievements"`
	Troops []struct {
		Name     string `json:"name"`
		Level    int    `json:"level"`
		MaxLevel int    `json:"maxLevel"`
	} `json:"troops"`
	Heroes []struct {
		Name     string `json:"name"`
		Level    int    `json:"level"`
		MaxLevel int    `json:"maxLevel"`
	} `json:"heroes"`
	Spells []struct {
		Name     string `json:"name"`
		Level    int    `json:"level"`
		MaxLevel int    `json:"maxLevel"`
	} `json:"spells"`
}

func GetMemberInfo(clan string) (members Members, err error) {
	body, err := getUrl(fmt.Sprintf(urlMembers, url.QueryEscape(clan)), myKey)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &members)
	return
}

func GetClanInfo(clanTag string) (clan ClanInfo, err error) {
	body, err := getUrl(fmt.Sprintf(urlClan, url.QueryEscape(clanTag)), myKey)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &clan)
	return
}

func GetPlayerInfo(memberTag string) (player Player, err error) {
	body, err := getUrl(fmt.Sprintf(urlPlayers, url.QueryEscape(memberTag)), myKey)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &player)
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
