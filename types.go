package main

import ()

type subscriber struct {
	Id        int64
	Email     string
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Active    bool
}

// microsoft api keys
type msKey struct {
	Id      int64
	Key     string
	ApiKey  string `db:"api_key"`
	Usage   int
	FreeKey bool `db:"free_key"`
	Month   string
	BillDay int `db:"bill_day"`
}

type resultLang struct {
	Title       string
	Description string
	Content     string
}

type pagesBetween struct {
	Second string
	Third  string
	Fourth string
	Fifth  string
	Sixth  string
}

type Subscribe struct {
	Email string `form:"email"`
}

type banner struct {
	Id       int
	Url      string
	Tag      string
	OfferUrl string `db:"offer_url"`
	Active   bool
}

type ArticleObj struct {
	Id              int64
	Source          string
	Topic           string
	Url             string
	UrlToImage      string `db:"url_to_image"`
	ImageBytes200   []byte `db:"image_bytes_200"`
	ImageBytes600   []byte `db:"image_bytes_600"`
	PublishedAt     string `db:"published_at"`
	TitleORIG       string `db:"title_orig"`
	TitleES         string `db:"title_es"`
	TitleFR         string `db:"title_fr"`
	TitleEN         string `db:"title_en"`
	TitleDE         string `db:"title_de"`
	TitleIT         string `db:"title_it"`
	TitlePT         string `db:"title_pt"`
	DescriptionORIG string `db:"description_orig"`
	DescriptionES   string `db:"description_es"`
	DescriptionFR   string `db:"description_fr"`
	DescriptionEN   string `db:"description_en"`
	DescriptionDE   string `db:"description_de"`
	DescriptionIT   string `db:"description_it"`
	DescriptionPT   string `db:"description_pt"`
	ContentORIG     string `db:"content_orig"`
	ContentES       string `db:"content_es"`
	ContentFR       string `db:"content_fr"`
	ContentEN       string `db:"content_en"`
	ContentDE       string `db:"content_de"`
	ContentIT       string `db:"content_it"`
	ContentPT       string `db:"content_pt"`
	StatusES        int    `db:"status_es"`
	StatusFR        int    `db:"status_fr"`
	StatusEN        int    `db:"status_en"`
	StatusDE        int    `db:"status_de"`
	StatusIT        int    `db:"status_it"`
	StatusPT        int    `db:"status_pt"`
}
