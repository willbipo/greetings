package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("nombre vacio")
	}

	mensages := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Saludos %v!, ¡Bienvenido!",
		"¡Que tal %v!, ¡Bienvenido!",
	}
	texto := mensages[rand.Intn(len(mensages))]
	message := fmt.Sprintf(texto, name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	mapa := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		mapa[name] = message

	}

	return mapa, nil
}
