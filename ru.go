package statCovid

import (
	"encoding/json"
	. "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Stat структура общей статистики
type Stat struct {
	Sick         string `json:"sick"`
	SickChange   string `json:"sickChange"`
	Healed       string `json:"healed"`
	HealedChange string `json:"healedChange"`
	Died         string `json:"died"`
	DiedChange   string `json:"diedChange"`
	OldStat      []struct {
		Date   string `json:"date"`
		Sick   int    `json:"sick"`
		Healed int    `json:"healed"`
		Died   int    `json:"died"`
	}
}

//RegionsStat струтура статистики по регионам
type RegionStat struct {
	Title      string `json:"title"`
	Code       string `json:"code"`
	IsCity     bool   `json:"is_city"`
	CoordX     string `json:"coord_x"`
	CoordY     string `json:"coord_y"`
	Sick       int    `json:"sick"`
	Healed     int    `json:"healed"`
	Died       int    `json:"died"`
	SickIncr   int    `json:"sick_incr"`
	HealedIncr int    `json:"healed_incr"`
	DiedIncr   int    `json:"died_incr"`
	Isolation  struct {
		StartDate string `json:"start_date"`
		Descr     string `json:"descr"`
		StateID   int    `json:"state_id"`
		Level     int    `json:"level"`
	} `json:"isolation"`
}

type RegionsStat []RegionStat

const URL = "https://стопкоронавирус.рф/information/"

//GetRuTotal возвращает общую статистику на момент запроса
func GetRuTotal() (*Stat, error) {
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("ERROR CLOSE BODY: %v \n", err)
		}
	}()

	if res.StatusCode != 200 {
		return nil, Errorf("STATUS ERROR CODE: %d %s", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	bodyString := string(body)

	a := strings.Index(bodyString, ":stats-data='")
	b := strings.Index(bodyString, "'></cv-stats-virus>")
	c := strings.Index(bodyString, "' :charts-data='")

	j := strings.Replace(strings.Replace(bodyString[a:c], ":stats-data='", "", -1)[:len(strings.Replace(bodyString[a:c], ":stats-data='", "", -1))-1]+",\"OldStat\":"+
		strings.Replace(bodyString[c:b],
			"' :charts-data='",
			"", -1)+"}",
		" ", "", -1,
	)

	st := &Stat{}

	if err := json.Unmarshal([]byte(j), st); err != nil {
		return nil, Errorf("ERROR JSON UNMARSHAL: %v", err)
	}

	return st, nil

}

//GetRuRegions возвращает статистику всех регионов
func GetRuRegions() (*RegionsStat, error) {

	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("ERROR CLOSE BODY: %v \n", err)
		}
	}()

	if res.StatusCode != 200 {
		return nil, Errorf("STATUS ERROR CODE: %d %s", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	bodyString := string(body)

	a := strings.Index(bodyString, ":spread-data='")
	b := strings.Index(bodyString, "' :isolation-data='")

	j := strings.Replace(bodyString[a:b], ":spread-data='", "", -1)
	st := &RegionsStat{}

	if err := json.Unmarshal([]byte(j), st); err != nil {
		return nil, Errorf("ERROR JSON UNMARSHAL: %v", err)
	}

	return st, nil
}

/*
SearchRegion поиск по буквенному коду региона России (ISO 3166-2)
Пример: поиск региона Москва(SearchRegion([]RegionsStat, "MOW"))
*/
func (rs *RegionsStat) SearchRuRegion(region string) (*RegionStat, error) {

	for _, r := range *rs {
		if r.Code[3:] == strings.ToUpper(region) {
			return &r, nil
		}
	}
	return nil, Errorf("Регион %s не найден! ", region)
}

//GetCodes возвращает список регионов
func (rs *RegionsStat) GetCodes() [][]string {
	var regions [][]string

	for _, r := range *rs {
		regions = append(regions, []string{r.Title, r.Code[3:]})
	}
	return regions
}
