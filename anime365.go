package anime365

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

// Series structure contains parsed JSON response with all fields
type Series struct {
	ID                 uint32        `json:"id"`
	AniDbID            uint32        `json:"aniDbId"`
	AnimeNewsNetworkID uint32        `json:"animeNewsNetworkId"`
	FansubsID          uint32        `json:"fansubsId"`
	ImdbID             uint32        `json:"imdbId"`
	WorldArtID         uint32        `json:"worldArtId"`
	Links              []link        `json:"link"`
	MyAnimeListID      uint32        `json:"myAnimeListId"`
	MyAnimeListScore   string        `json:"myAnimeListScore"`
	WorldArtListScore  string        `json:"worldArtListScore"`
	WorldArtTopPlace   uint32        `json:"worldArtTopPlace"`
	NumberOfEpisodes   uint32        `json:"numberOfEpisodes"`
	PosterURL          string        `json:"posterUrl"`
	PosterURLSmall     string        `json:"posterUrlSmall"`
	Season             string        `json:"string"`
	Year               uint32        `json:"year"`
	Type               string        `json:"type"`
	TypeTitle          string        `json:"typeTitle"`
	CountViews         uint32        `json:"countViews"`
	Titles             title         `json:"titles"`
	TitleLines         []string      `json:"titleLines"`
	AllTitles          []string      `json:"allTitles"`
	Title              string        `json:"title"`
	URL                string        `json:"url"`
	Descriptions       []description `json:"descriptions"`
	Episodes           []Episode     `json:"episodes"`
	Genres             []genre       `json:"genres"`
	isActive           uint8
	isAiring           uint8
	isHentai           uint8
}

// IsActive returns boolean value whether series is active (true) or not (false)
func (s *Series) IsActive() bool {
	if s.isActive == 1 {
		return true
	}
	return false
}

// IsAiring returns boolean value whether series is airing (true) or not (false)
func (s *Series) IsAiring() bool {
	if s.isAiring == 1 {
		return true
	}
	return false
}

// IsHentai returns boolean value whether series is hentai (true) or not (false)
func (s *Series) IsHentai() bool {
	if s.isHentai == 1 {
		return true
	}
	return false
}

// Translation structure contains parsed JSON response with all fields
type Translation struct {
	ID                   uint32             `json:"id"`
	AddedDateTime        string             `json:"addedDateTime"`
	ActiveDateTime       string             `json:"activeDateTime"`
	AuthorsList          []string           `json:"authorsList"`
	FansubsTranslationID uint32             `json:"fansubsTranslationId"`
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
	isActive             uint8
}

// IsActive returns boolean value whether translation is active (true) or not (false)
func (t *Translation) IsActive() bool {
	if t.isActive == 1 {
		return true
	}
	return false
}

// Episode structure contains parsed JSON response with all fields
type Episode struct {
	ID                    uint32        `json:"id"`
	EpisodeFull           string        `json:"episodeFull"`
	EpisodeInt            string        `json:"episodeInt"`
	EpisodeTitle          string        `json:"episodeTitle"`
	EpisodeType           string        `json:"episodeType"`
	FirstUploadedDateTime string        `json:"firstUploadedDateTime"`
	SeriesID              uint32        `json:"seriesId"`
	CountViews            uint32        `json:"countViews"`
	Translations          []Translation `json:"translations"`
	isActive              uint8
}

// IsActive returns boolean value whether episode is active (true) or not (false)
func (e *Episode) IsActive() bool {
	if e.isActive == 1 {
		return true
	}
	return false
}

// GetTranslations makes request to API with followed parameters and returns list of Translation structures
func GetTranslations(parameters map[string]string) ([]Translation, error) {
	requestURL, err := url.Parse(apiURL + "translations")
	if err != nil {
		return nil, err
	}
	requestQuery := requestURL.Query()
	for key, value := range parameters {
		requestQuery.Set(key, value)
	}
	requestURL.RawQuery = requestQuery.Encode()
	response, err := http.Get(requestURL.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var jsonResponse map[string][]Translation
	if err := json.Unmarshal(content, &jsonResponse); err != nil {
		return nil, err
	}
	return jsonResponse["data"], nil
}

// GetTranslationByID makes request by ID with followed parameters to API and returns a Translation structure
func GetTranslationByID(id uint32, parameters map[string]string) (*Translation, error) {
	requestURL, err := url.Parse(apiURL + "translations/" + fmt.Sprint(id))
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

// GetSeries makes request to API with followed parameters and returns list of Series structures
func GetSeries(parameters map[string]string) ([]Series, error) {
	requestURL, err := url.Parse(apiURL + "series")
	if err != nil {
		return nil, err
	}
	requestQuery := requestURL.Query()
	for key, value := range parameters {
		requestQuery.Set(key, value)
	}
	requestURL.RawQuery = requestQuery.Encode()
	response, err := http.Get(requestURL.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var jsonResponse map[string][]Series
	if err := json.Unmarshal(content, &jsonResponse); err != nil {
		return nil, err
	}
	return jsonResponse["data"], nil
}

// GetSeriesByID makes request by ID with followed parameters to API and returns a Series structure
func GetSeriesByID(id uint32, parameters map[string]string) (*Series, error) {
	requestURL, err := url.Parse(apiURL + "series/" + fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	requestQuery := requestURL.Query()
	for key, value := range parameters {
		requestQuery.Set(key, value)
	}
	requestURL.RawQuery = requestQuery.Encode()
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

// GetEpisodeByID makes request by ID with followed parameters to API and returns an Episode structure
func GetEpisodeByID(id uint32, parameters map[string]string) (*Episode, error) {
	requestURL, err := url.Parse(apiURL + "episodes/" + fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	requestQuery := requestURL.Query()
	for key, value := range parameters {
		requestQuery.Set(key, value)
	}
	requestURL.RawQuery = requestQuery.Encode()
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
