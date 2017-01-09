package anime365

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const apiURL = "https://smotret-anime.ru/api/"

type genre struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

type link struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type title struct {
	RU     string `json:"ru"`
	Romaji string `json:"romaji"`
	JA     string `json:"ja"`
	EN     string `json:"en"`
	Short  string `json:"short"`
}

type fansubsTranslation struct {
	ID              uint32 `json:"id"`
	FansubsSeriesID uint32 `json:"fansubsSeriesId"`
	Count           string `json:"count"`
	Date            string `json:"date"`
	Comment         string `json:"comment"`
	File            string `json:"file"`
	FileURL         string `json:"fileUrl"`
}

type description struct {
	Source          string `json:"source"`
	Value           string `json:"value"`
	UpdatedDateTime string `json:"updatedDateTime"`
}

type Series struct {
	ID                 uint32      `json:"id"`
	AniDbID            uint32      `json:"aniDbId"`
	AnimeNewsNetworkID uint32      `json:"animeNewsNetworkId"`
	FansubsID          uint32      `json:"fansubsId"`
	ImdbID             uint32      `json:"imdbId"`
	WorldArtID         uint32      `json:"worldArtId"`
	IsActive           uint32      `json:"isActive"`
	IsAiring           uint32      `json:"isAiring"`
	IsHentai           uint32      `json:"isHentai"`
	Links              []link      `json:"link"`
	MyAnimeListID      uint32      `json:"myAnimeListId"`
	MyAnimeListScore   string      `json:"myAnimeListScore"`
	WorldArtListScore  string      `json:"worldArtListScore"`
	WorldArtTopPlace   uint32      `json:"worldArtTopPlace"`
	NumberOfEpisodes   uint32      `json:"numberOfEpisodes"`
	PosterURL          string      `json:"posterUrl"`
	PosterURLSmall     string      `json:"posterUrlSmall"`
	Season             string      `json:"string"`
	Year               uint32      `json:"year"`
	Type               string      `json:"type"`
	TypeTitle          string      `json:"typeTitle"`
	CountViews         uint32      `json:"countViews"`
	Titles             title       `json:"titles"`
	TitleLines         []string    `json:"titleLines"`
	AllTitles          []string    `json:"allTitles"`
	Title              string      `json:"title"`
	URL                string      `json:"url"`
	Descriptions       description `json:"descriptions"`
	Episodes           []Episode   `json:"episodes"`
	Genres             []genre     `json:"genres"`
}

type Translation struct {
	ID                   uint32             `json:"id"`
	AddedDateTime        string             `json:"addedDateTime"`
	ActiveDateTime       string             `json:"activeDateTime"`
	AuthorsList          []string           `json:"authorsList"`
	FansubsTranslationID uint32             `json:"fansubsTranslationId"`
	IsActive             uint32             `json:"isActive"`
	Priority             uint32             `json:"priority"`
	QualityType          string             `json:"qualityType"`
	Type                 string             `json:"type"`
	TypeKind             string             `json:"typeKind"`
	TypeLang             string             `json:"typeLang"`
	UpdatedDateTime      string             `json:"updatedDateTime"`
	Title                string             `json:"title"`
	SeriesID             uint32             `json:"seriesId"`
	EpisodeID            uint32             `json:"episodeId"`
	CountViews           uint32             `json:"countViews"`
	URL                  string             `json:"url"`
	EmbedURL             string             `json:"embedUrl"`
	AuthorsSummary       string             `json:"authorsSummary"`
	Episode              Episode            `json:"episode"`
	Series               Series             `json:"series"`
	FansubsTranslation   fansubsTranslation `json:"url"`
	Duration             string             `json:"duration"`
	Width                uint32             `json:"width"`
	Height               uint32             `json:"height"`
}

type Episode struct {
	ID                    uint32        `json:"id"`
	EpisodeFull           string        `json:"episodeFull"`
	EpisodeInt            string        `json:"episodeInt"`
	EpisodeTitle          string        `json:"episodeTitle"`
	EpisodeType           string        `json:"episodeType"`
	FirstUploadedDateTime string        `json:"firstUploadedDateTime"`
	IsActive              uint32        `json:"isActive"`
	SeriesID              uint32        `json:"seriesId"`
	CountViews            uint32        `json:"countViews"`
	Translations          []Translation `json:"translations"`
}

func GetTranslations(parameters map[string]string) (*[]Translation, error) {
	requestURL, err := url.Parse(apiURL + "translations")
	if err != nil {
		return nil, err
	}
	for key, value := range parameters {
		requestURL.Query().Set(key, value)
	}
	requestURL.RawQuery = requestURL.Query().Encode()
	response, err := http.Get(requestURL.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var jsonResponse map[string]*[]Translation
	if err := json.Unmarshal(content, &jsonResponse); err != nil {
		return nil, err
	}
	return jsonResponse["data"], nil
}

func GetTranslationByID(id uint32, parameters map[string]string) (*Translation, error) {
	requestURL, err := url.Parse(apiURL + "translations/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	for key, value := range parameters {
		requestURL.Query().Set(key, value)
	}
	requestURL.RawQuery = requestURL.Query().Encode()
	response, err := http.Get(requestURL.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var jsonResponse map[string]*Translation
	if err := json.Unmarshal(content, &jsonResponse); err != nil {
		return nil, err
	}
	return jsonResponse["data"], nil
}

func GetSeries(parameters map[string]string) (*[]Series, error) {
	requestURL, err := url.Parse(apiURL + "series")
	if err != nil {
		return nil, err
	}
	for key, value := range parameters {
		requestURL.Query().Set(key, value)
	}
	requestURL.RawQuery = requestURL.Query().Encode()
	response, err := http.Get(requestURL.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var jsonResponse map[string]*[]Series
	if err := json.Unmarshal(content, &jsonResponse); err != nil {
		return nil, err
	}
	return jsonResponse["data"], nil
}

func GetSeriesById(id uint32, parameters map[string]string) (*Series, error) {
	requestURL, err := url.Parse(apiURL + "series/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	for key, value := range parameters {
		requestURL.Query().Set(key, value)
	}
	requestURL.RawQuery = requestURL.Query().Encode()
	response, err := http.Get(requestURL.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var jsonResponse map[string]*Series
	if err := json.Unmarshal(content, &jsonResponse); err != nil {
		return nil, err
	}
	return jsonResponse["data"], nil
}

func GetEpisodeByID(id uint32, parameters map[string]string) (*Episode, error) {
	requestUrlRaw := apiURL + "episodes/" + strconv.Itoa(id)
	requestURL, err := url.Parse(requestUrlRaw)
	if err != nil {
		return nil, err
	}
	for key, value := range parameters {
		requestURL.Query().Set(key, value)
	}
	requestURL.RawQuery = requestURL.Query().Encode()
	response, err := http.Get(requestURL.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var jsonResponse map[string]*Episode
	if err := json.Unmarshal(content, &jsonResponse); err != nil {
		return nil, err
	}
	return jsonResponse["data"], nil
}
