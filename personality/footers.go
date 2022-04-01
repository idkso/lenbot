package personality

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type FooterFile struct {
	Footers []string `json:"footers"`
}

var footers []string

func init() {
	file, err := ioutil.ReadFile("assets/footers.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var footerfile FooterFile

	err = json.Unmarshal(file, &footerfile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	footers = footerfile.Footers

	fmt.Println("footers loaded")
}

func RandomFooter() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return footers[r.Intn(len(footers))]
}
