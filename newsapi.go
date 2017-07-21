package main

import (
	"github.com/bas24/newsapi"
)

func newsapiFetch() ([]ArticleObj, error) {
	result, err := newsapi.GetArticles(NewsApiKey)
	if err != nil {
		return []ArticleObj{}, nil
	}

	articles := []ArticleObj{}
	isValid := func(art newsapi.Article) bool {
		if art.Url == "" {
			return false
		}
		if art.Title == "" {
			return false
		}
		if art.Description == "" {
			return false
		}
		if art.UrlToImage == "" {
			return false
		}
		return true
	}
	for _, r := range result {
		for j, art := range r.Articles {
			if isValid(art) {
				image := func() {

				}
				// clue Response.Articles[].Articles[].art to ArticleObj
				a := ArticleObj{
					Id:              int64(0), // db will set on insert
					Source:          r.Source,
					Topic:           newsapi.GetTopic(r.Source),
					Url:             art.Url,
					UrlToImage:      art.UrlToImage,
					ImageBytes200:   []byte{}, // TODO func to fetch, resize and encode image
					ImageBytes600:   []byte{},
					PublishedAt:     parsePublishedAt(art.PublishedAt),
					TitleORIG:       art.Title,
					TitleES:         "",
					TitleFR:         "",
					TitleEN:         "",
					TitleDE:         "",
					TitleIT:         "",
					TitlePT:         "",
					DescriptionORIG: art.Description,
					DescriptionES:   "",
					DescriptionFR:   "",
					DescriptionEN:   "",
					DescriptionDE:   "",
					DescriptionIT:   "",
					DescriptionPT:   "",
					ContentORIG:     "",
					ContentES:       "",
					ContentFR:       "",
					ContentEN:       "",
					ContentDE:       "",
					ContentIT:       "",
					ContentPT:       "",
					StatusES:        0,
					StatusFR:        0,
					StatusEN:        0,
					StatusDE:        0,
					StatusIT:        0,
					StatusPT:        0,
				}
			}
		}
	}
	return articles, nil
}
