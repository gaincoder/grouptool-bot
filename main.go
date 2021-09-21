package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	database := connectToDatabase()
	defer database.close()

	bot := createBot(database)

	go bot.api.Start()

	// c := cron.New()
	// c.AddFunc("0 * * * *", func() { bot.ApiInput("minutely cron") })

	// go c.Start()

	mux := http.NewServeMux()

	// mux.Handle("/", apiAuth(finalHandler))
	mux.Handle("/eventCreated", apiAuth(handleEventCreated(bot)))
	mux.Handle("/infoCreated", apiAuth(handleInfoCreated(bot)))
	mux.Handle("/pollCreated", apiAuth(handlePollCreated(bot)))

	http.ListenAndServe(":8088", mux)
}

func apiAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.Header.Get("X-TOKEN") != os.Getenv("API_TOKEN") {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
