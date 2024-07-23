package groupie

import (
	"html/template"
	"log"
	"net/http"
)

// Add function for template to use
func add(x, y int) int {
	return x + y
}

func Start() {
	// Trying to fetch data from the API
	var artists []Artists
	err := fetchData("/artists", &artists)
	if err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if code := PageHandle(w,r); code == 1{return}
		funcMap := template.FuncMap{
			"add": add,
		}
		t, err := template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html")
		if err != nil {
			log.Println("Error parsing template:", err)
			ServerError(w,r,err)
			return
		}

		data := map[string]interface{}{
			"artists": artists,
		}
		err = t.Execute(w, data)
		if err != nil {
			log.Println("Error executing template:", err)
			ServerError(w,r,err)
		}
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/artist", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			MethodError(w, r)
			return
		}
		// Handle the form submission here
		artistID := r.FormValue("artist_id")
		log.Printf("Artist ID submitted: %s", artistID)

		// Fetch the artist data from the API using the artistID
		var artist Artists
		err := fetchData("/artists/" + artistID, &artist)
		if err != nil {
			ServerError(w, r, err)
			return
		}

		// Fetch the relations data
		relations, err := fetchRelations(artist.ID)
		if err != nil {
			ServerError(w, r, err)
			return
		}

		tmpl, err := template.ParseFiles("templates/artist.html")
		if err != nil {
			ServerError(w, r, err)
			return
		}

		data := map[string]interface{}{
			"Name":      artist.Name,
			"Image":     artist.Image,
			"Members":   artist.Members,
			"Creation":  artist.Creation,
			"FirstAlbum": artist.FirstAlbum,
			"Relations": relations.Dates,
		}

		if err := tmpl.Execute(w, data); err != nil {
			ServerError(w, r, err)
			return
		}
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
