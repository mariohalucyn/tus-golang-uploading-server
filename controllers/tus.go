package controllers

import (
	"log"
	"net/http"

	"github.com/tus/tusd/v2/pkg/filelocker"
	"github.com/tus/tusd/v2/pkg/filestore"
	tusd "github.com/tus/tusd/v2/pkg/handler"
)

func Tus() http.Handler {
	store := filestore.New("./storage")
	locker := filelocker.New("./storage")
	composer := tusd.NewStoreComposer()
	store.UseIn(composer)
	locker.UseIn(composer)
	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:              "/files/",
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
	})
	if err != nil {
		log.Fatalf("unable to create handler: %s", err)
	}
	go func() {
		for {
			event := <-handler.CompleteUploads
			log.Printf("Upload %s finished\n", event.Upload.ID)
		}
	}()

	return handler
}
