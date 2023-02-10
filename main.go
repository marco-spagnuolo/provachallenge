package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

var wg sync.WaitGroup

func main() {
	// Apri il file per la lettura
	file, err := os.Open("links.txt")
	resultFile, errResult := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	//check-up err
	if err != nil {
		fmt.Println("Errore durante l'apertura del file:", err)
		return
	}
	if errResult != nil {
		fmt.Println("Errore durante l'apertura del file:", err)
		return
	}

	defer file.Close()
	defer resultFile.Close()

	// Crea uno scanner per leggere il file riga per riga
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		//se siamo arrivati a fine file => fine ciclo for
		if line == "" {
			break
		}

		wg.Add(1)
		go check(line, resultFile)
	}

	wg.Wait()
}

func check(line string, fp *os.File) int {
	//richiesta http
	resp, err := http.Get(line)
	if err != nil {
		log.Fatal().Err(err).Msg("Errore durante la richiesta HTTP")
	}
	//stampa valori
	if resp.StatusCode != 200 {
		log.Warn().Str("url", line).Int("status_code", resp.StatusCode).Msg("URL non raggiungibile")
	} else {
		log.Info().Str("url", line).Int("status_code", resp.StatusCode).Msg("URL raggiungibile")
	}
	//scrivi valori in file
	_, err = fp.WriteString(line + " " + resp.Status + " " + time.Now().String() + "\n")
	if err != nil {
		log.Fatal().Err(err).Msg("Errore durante la scrittura sul file")
	}

	wg.Done()
	return resp.StatusCode
}
