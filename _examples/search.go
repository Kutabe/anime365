package main

import (
	"fmt"
	"log"

	"github.com/kutabe/anime365"
)

func main() {
	parameters := make(map[string]string)
	parameters["query"] = "Bakemonogatari"
	respSeries, err := anime365.GetSeries(parameters)
	if err != nil {
		log.Panic(err)
	}
	for _, series := range respSeries {
		fmt.Printf("Title: %s\nDescription:\n", series.Title)
		for _, description := range series.Descriptions {
			fmt.Printf("%s:\n\t%s\n", description.Source, description.Value)
		}
		fmt.Print("Episodes:\n")
		for _, episode := range series.Episodes {
			fmt.Printf("#%s\n", episode.EpisodeInt)
			fmt.Printf("Translations:\n")
			respEpisode, err := anime365.GetEpisodeByID(episode.ID, nil)
			if err != nil {
				log.Println(err)
			}
			for _, translation := range respEpisode.Translations {
				fmt.Printf("\tLang: %s \n\tAuthors: %s \n\tDownload link: http://smotret-anime.ru/translations/mp4/%d\n", translation.TypeLang, translation.AuthorsSummary, translation.ID)
			}

		}
	}
}
