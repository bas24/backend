package main

import (
	"fmt"
	"time"
)

func initDB() {
	// create database schema
	db.MustExec(`CREATE DATABASE IF NOT EXISTS feedfiend;`)
	db.MustExec(`
		CREATE TABLE IF NOT EXISTS feedfiend.articles (
			id SERIAL NOT NULL PRIMARY KEY,
			source TEXT NOT NULL,
			topic TEXT NOT NULL,
			url TEXT NOT NULL,
			url_to_image TEXT NOT NULL,
			image_bytes_200 BYTES NOT NULL,
			image_bytes_600 BYTES NOT NULL,
			published_at TIMESTAMP NOT NULL,
			title_orig TEXT NOT NULL,
			title_es TEXT DEFAULT '',
			title_fr TEXT DEFAULT '',
			title_en TEXT DEFAULT '',
			title_de TEXT DEFAULT '',
			title_it TEXT DEFAULT '',
			title_pt TEXT DEFAULT '',
			description_orig TEXT NOT NULL,
			description_es TEXT DEFAULT '',
			description_fr TEXT DEFAULT '',
			description_en TEXT DEFAULT '',
			description_de TEXT DEFAULT '',
			description_it TEXT DEFAULT '',
			description_pt TEXT DEFAULT '',
			content_orig TEXT DEFAULT '',
			content_es TEXT DEFAULT '',
			content_fr TEXT DEFAULT '',
			content_en TEXT DEFAULT '',
			content_de TEXT DEFAULT '',
			content_it TEXT DEFAULT '',
			content_pt TEXT DEFAULT '',
			status_es INT DEFAULT 0,
			status_fr INT DEFAULT 0,
			status_en INT DEFAULT 0,
			status_de INT DEFAULT 0,
			status_it INT DEFAULT 0,
			status_pt INT DEFAULT 0,
			UNIQUE(url)
		);
		CREATE INDEX IF NOT EXISTS publishedAtIndex ON feedfiend.articles(published_at DESC);
		`)
	db.MustExec(`
		CREATE TABLE IF NOT EXISTS feedfiend.banners(
			id SERIAL NOT NULL PRIMAR Y KEY,
			url TEXT NOT NULL, 
			tag TEXT NOT NULL, 
			offer_url TEXT NOT NULL,
			active BOOL NOT NULL
		);
		`)
	db.MustExec(`
		CREATE TABLE IF NOT EXISTS feedfiend.subscribers(
			id SERIAL NOT NULL PRIMARY KEY,
			email TEXT NOT NULL,
			first_name TEXT,
			last_name TEXT,
			active BOOL DEFAULT true
		);
		CREATE UNIQUE INDEX IF NOT EXISTS emailUniqueIndex ON feedfiend.subscribers(email);
		`)
	db.MustExec(`
		CREATE TABLE IF NOT EXISTS feedfiend.microsoft(
			id SERIAL NOT NULL PRIMARY KEY,
			key TEXT NOT NULL,
			api_key TEXT NOT NULL,
			usage INT NOT NULL,
			free_key BOOL NOT NULL,
			month TEXT NOT NULL,
			bill_day INT NOT NULL,
			UNIQUE(key, apikey)
		);
		`)
}

// TODO need rewrite
func insertRecord(f ...string) {
	s := `
		INSERT INTO feedfiend.articles(
			published_at, 
			source, 
			title, 
			description, 
			url, 
			urlToImage, 
			topic
			)
			VALUES($1, $2, $3, $4, $5, $6, $7)
			ON CONFLICT(url, title) DO NOTHING;
		`
	vals := makeInterface(f)
	db.MustExec(s, vals...)
}

func getLastArticles(date time.Time) []ArticleObj {
	// getting rows from time specified until Now
	stmt := fmt.Sprintf(`
		SELECT 
			id, 
			source, 
			topic, 
			url,
			url_to_image,
			image_bytes_200,
			image_bytes_600,
			published_at,
			title_orig,
			title_es,
			title_fr,
			title_en,
			title_de,
			title_it,
			title_pt,
			description_orig,
			description_es,
			description_fr,
			description_en,
			description_de,
			description_it,
			description_pt,
			content_orig,
			content_es,
			content_fr,
			content_en,
			content_de,
			content_it,
			content_pt,
			status_es,
			status_fr,
			status_en,
			status_de,
			status_it,
			status_pt
			FROM feedfiend.articles@publishedAtIndex
				WHERE publishedAt > %v`, date)

	articles := []ArticleObj{}

	db.Select(&articles, stmt)

	return articles
}

func loadSpinnedData(a ArticleObj) error {
	stmt := fmt.Sprintf(`
		UPDATE feedfiend.articles
			SET title_es = '%v',
				title_fr = '%v',
				title_en = '%v',
				title_de = '%v',
				title_it = '%v',
				title_pt = '%v',
				description_es = '%v',
				description_fr = '%v', 
				description_en = '%v', 
				description_de = '%v',
				description_it = '%v',
				description_pt = '%v',
				content_orig = '%v', 
				content_es = '%v', 
				content_fr = '%v',
				content_en = '%v',
				content_de = '%v',
				content_it = '%v',
				content_pt = '%v'
			WHERE id = $1;
	`, rQ(a.TitleES),
		rQ(a.TitleFR),
		rQ(a.TitleEN),
		rQ(a.TitleDE),
		rQ(a.TitleIT),
		rQ(a.TitlePT),
		rQ(a.DescriptionES),
		rQ(a.DescriptionFR),
		rQ(a.DescriptionEN),
		rQ(a.DescriptionDE),
		rQ(a.DescriptionIT),
		rQ(a.DescriptionPT),
		rQ(a.ContentORIG),
		rQ(a.ContentES),
		rQ(a.ContentFR),
		rQ(a.ContentEN),
		rQ(a.ContentDE),
		rQ(a.ContentIT),
		rQ(a.ContentPT),
	)
	db.MustExec(stmt, a.Id)
	return nil
}

func loadMicrosoftApiKeyToDb(mK msKey) {
	stmt := fmt.Sprintf(`
		INSERT INTO feedfiend.microsoft(
			key, 
			api_key, 
			usage, 
			free_key, 
			month, 
			bill_day
			) 
			VALUES (%v, %v, %v, %v, %v, %v)
		ON CONFLICT(key, api_key) DO NOTHING;
	`, mK.Key,
		mK.ApiKey,
		mK.Usage,
		mK.FreeKey,
		mK.Month,
		mK.BillDay,
	)
	db.MustExec(stmt)
}

func getMicrosoftApiKeysFromDb() ([]msKey, error) {
	t := time.Now()
	m := t.Month().String()
	keys := []msKey{}
	stmt := `
		SELECT 
			id, 
			key, 
			api_key, 
			usage, 
			free_key, 
			month, 
			bill_day
			FROM feedfiend.microsoft;
	`
	db.Select(&keys, stmt)

	isValid := func(key msKey) bool {
		if key.Month != m && t.Day() == key.BillDay && t.Hour() > 10 {
			return true
		}
		return false
	}
	i := 0 // index
	for _, key := range keys {
		// check if we need to reset usage
		// reseting on billing day of current month
		if isValid(key) {
			keys[i] = key
			i++
		} else {
			resetMSApiKeysUsage(m, key.ApiKey)
		}
	}
	return keys, nil
}

func addToMSApiKey(newUsage int, msApiKey string) error {
	s := `
		UPDATE feedfiend.microsoft 
			SET usage = $1 
			WHERE api_key = $2;
	`
	db.MustExec(s, newUsage, msApiKey)
	return nil
}

// Reseting once per month api usage statistics
func resetMSApiKeysUsage(month, apiKey string) {
	s := `
		UPDATE feedfiend.microsoft 
			SET usage = $1, month = $2
			WHERE api_key = $3;
	`
	db.MustExec(s, 0, month, apiKey)
}
