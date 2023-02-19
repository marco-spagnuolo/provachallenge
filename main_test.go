//testing
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestCheck(t *testing.T) {

	// Apri il file per la lettura
	file, err := os.Open("links.txt")

	if err != nil {
		fmt.Println("Errore durante l'apertura del file:", err)
		return
	}

	//chiudi file
	defer file.Close()

	// Crea uno scanner per leggere il file riga per riga
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		//se siamo arrivati a fine file => fine ciclo for
		if line == "" {
			break
		}
		//prendi response da get(url)
		resp, err := http.Get(line)
		if err != nil {
			log.Fatal(err)
		}
		//se la response non è 200 errore
		if resp.StatusCode != 200 {
			t.Errorf("error on %s", line)
		}
	}
}
