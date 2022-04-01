package personality

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type WelcomeFile struct {
	Welcomes []string `json:"welcomes"`
}

var welcomes []string

func init() {
	file, err := ioutil.ReadFile("assets/welcomes.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var welcomefile WelcomeFile

	err = json.Unmarshal(file, &welcomefile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	welcomes = welcomefile.Welcomes

	fmt.Println("welcomes loaded")
}

func RandomWelcome() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return welcomes[r.Intn(len(footers))]
}
