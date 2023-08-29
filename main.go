package main

import (
	"amcollie/go-htmx/db"
	"amcollie/go-htmx/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/dustin/go-humanize"
)

func main() {
	fmt.Println("Hello, World!")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("static/templates/index.html", "static/templates/partials/table-data.html"))

		db, dbConnectionError := db.DB()
		if dbConnectionError != nil {
			log.Fatal(dbConnectionError)
		}

		var filmList []models.Film
		result := db.Find(&filmList)
		if result.Error != nil {
			log.Println(result.Error)
		}

		films := make(map[string][]models.FilmDsp)

		for _, film := range filmList {
			films["Films"] = append(
				films["Films"],
				models.FilmDsp{Title: film.Title, Director: film.Director, Year: film.Year, BoxOffice: humanize.Commaf(film.BoxOffice)},
			)
		}

		tmpl.Execute(w, films)
	})

	http.HandleFunc("/add-film", func(w http.ResponseWriter, r *http.Request) {

		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		year, convIntErr := strconv.Atoi(r.PostFormValue("year"))
		if convIntErr != nil {
			log.Fatal(convIntErr)
		}
		boxOffice, convBoxOfficeErr := strconv.ParseFloat(r.PostFormValue("boxOffice"), 64)
		if convBoxOfficeErr != nil {
			log.Fatal(convBoxOfficeErr)
		}

		tmpl := template.Must(template.ParseFiles("static/templates/index.html", "static/templates/partials/table-data.html"))

		db, dbConnectionError := db.DB()
		if dbConnectionError != nil {
			log.Fatal(dbConnectionError)
		}

		film := models.Film{Title: title, Director: director, Year: year, BoxOffice: boxOffice}
		result := db.Create(&film)
		if result.Error != nil {
			log.Println(result.Error)
		}

		filmDsp := models.FilmDsp{Title: title, Director: director, Year: year, BoxOffice: humanize.Commaf(boxOffice)}
		tmpl.ExecuteTemplate(w, "film-table-element", filmDsp)
	})

	http.HandleFunc("/edit-film", func(w http.ResponseWriter, r *http.Request) {
		db, errConnectionError := db.DB()
		if errConnectionError != nil {
			log.Fatal(errConnectionError)
		}

		id := r.PostFormValue("id")
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		year, convIntErr := strconv.Atoi(r.PostFormValue("year"))
		if convIntErr != nil {
			log.Fatal(convIntErr)
		}
		boxOffice, convBoxOffficeErr := strconv.ParseFloat(r.PostFormValue("boxOffice"), 64)
		if convBoxOffficeErr != nil {
			log.Fatal(convBoxOffficeErr)
		}

		var film models.Film
		db.Find(&film, id)
		db.Model(&film).Updates(models.Film{Title: title, Director: director, Year: year, BoxOffice: boxOffice})

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
