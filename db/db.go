package db

import (
	"amcollie/go-htmx/models"
	"log"
	"net/http"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("movies.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func InitDB() {
	db, err := DB()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.MovieDB{})
	films := []*models.Film{
		{
			Title:     "Star Wars: Episode IV - A New Hope",
			Director:  "George Lucas",
			Year:      1977,
			BoxOffice: 775800000,
		},
		{
			Title:     "Star Wars: Episode V - The Empire Strikes Back",
			Director:  "Irvin Kershner",
			Year:      1980,
			BoxOffice: 549000000,
		},
		{
			Title:     "Star Wars: Episode VI - Return of the Jedi",
			Director:  "Richard Marquand",
			Year:      1983,
			BoxOffice: 482000000,
		},
		{
			Title:     "Star Wars: Episole I - The Phantom Menace",
			Director:  "George Lucas",
			Year:      1999,
			BoxOffice: 1027000000,
		},
		{
			Title:     "Star Wars: Episole II - Attack of the Clones",
			Director:  "George Lucas",
			Year:      2002,
			BoxOffice: 653800000,
		},
		{
			Title:     "Star Wars: Episole III - Revenge of the Sith",
			Director:  "George Lucas",
			Year:      2005,
			BoxOffice: 868400000,
		},
		{
			Title:     "Star Wars: Episole VII - The Force Awakens",
			Director:  "J. J. Abrams",
			Year:      2015,
			BoxOffice: 2071100000,
		},
		{
			Title:     "Star Wars: Episole VIII - The Last Jedi",
			Director:  "Rian Johnson",
			Year:      2017,
			BoxOffice: 1334000000,
		},
		{
			Title:     "Star Trek: The Motion Picture",
			Director:  "Robert Wise",
			Year:      1979,
			BoxOffice: 139000000,
		},
		{
			Title:     "Star Trek II: The Wrath of Khan",
			Director:  "Nicholas Meyer",
			Year:      1982,
			BoxOffice: 97000000,
		},
		{
			Title:     "Star Trek III: The Search for Spock",
			Director:  "Leonard Nimoy",
			Year:      1984,
			BoxOffice: 87000000,
		},
		{
			Title:     "Star Trek IV: The Voyage Home",
			Director:  "Leonard Nimoy",
			Year:      1986,
			BoxOffice: 133000000,
		},
		{
			Title:     "Star Trek V: The Final Frontier",
			Director:  "William Shatner",
			Year:      1989,
			BoxOffice: 63000000,
		},
		{
			Title:     "Star Trek VI: The Undiscovered Country",
			Director:  "Nicholas Meyer",
			Year:      1991,
			BoxOffice: 96800000,
		},
		{
			Title:     "Star Trek Generations",
			Director:  "David Carson",
			Year:      1994,
			BoxOffice: 118000000,
		},
		{
			Title:     "Star Trek: First Contact",
			Director:  "Jonathan Frakes",
			Year:      1996,
			BoxOffice: 146000000,
		},
		{
			Title:     "Star Trek: Insurrection",
			Director:  "Jonathan Frakes",
			Year:      1998,
			BoxOffice: 117800000,
		},
		{
			Title:     "Star Trek: Nemesis",
			Director:  "Stuart Baird",
			Year:      2002,
			BoxOffice: 67300000,
		},
		{
			Title:     "Star Trek",
			Director:  "J. J. Abrams",
			Year:      2009,
			BoxOffice: 385700000,
		},
		{
			Title:     "Star Trek Into Darkness",
			Director:  "J. J. Abrams",
			Year:      2013,
			BoxOffice: 467400000,
		},
		{
			Title:     "Star Trek Beyond",
			Director:  "Justin Lin",
			Year:      2016,
			BoxOffice: 345500000,
		},
	}
	result := db.Create(films)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
