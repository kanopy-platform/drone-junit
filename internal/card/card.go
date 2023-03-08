package card

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"os"

	"github.com/drone/drone-go/drone"
	"github.com/kanopy-platform/drone-junit/internal/junit"
)

func WriteCard(suite *junit.Testsuites) error {

	cardData, err := json.Marshal(suite)
	if err != nil {
		return err
	}

	card := &drone.CardInput{
		Schema: "https://kanopy-platform.github.io/drone-junit/card.json",
		Data:   cardData,
	}

	return write(card)
}

func write(card *drone.CardInput) error {
	data, err := json.Marshal(card)
	if err != nil {
		return err
	}
	out := os.Stdout
	encoded := base64.StdEncoding.EncodeToString(data)
	io.WriteString(out, "\u001B]1338;")
	io.WriteString(out, encoded)
	io.WriteString(out, "\u001B]0m")
	io.WriteString(out, "\n")
	return nil
}
