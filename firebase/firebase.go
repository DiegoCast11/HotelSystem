package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// FirebaseApp inicializa la aplicación de Firebase.
func FirebaseApp() (*firebase.App, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("sanfelipehotel-83ae3-firebase-adminsdk-2mpje-e444d2af98.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error al inicializar la aplicación de Firebase: %v\n", err)
		return nil, err
	}
	return app, nil
}
