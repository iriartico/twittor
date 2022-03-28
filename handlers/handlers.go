package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/iriartico/twittor/middlewares"
	"github.com/iriartico/twittor/routers"
	"github.com/rs/cors"
)

/*Handlers seteo el puerto, el Handler y escucha del servidor*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/view-profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modify-profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.RecordTweet))).Methods("POST")
	router.HandleFunc("/read-tweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/remove-tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.RemoveTweet))).Methods("DELETE")

	router.HandleFunc("/upload-avatar", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/get-avatar", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/upload-banner", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/get-banner", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetBanner))).Methods("GET")

	router.HandleFunc("/high-relationship", middlewares.CheckDB(middlewares.ValidateJWT(routers.HigRelationship))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
