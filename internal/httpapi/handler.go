package httpapi

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/siddhant-vij/Word-Frequency-Server/internal/textprocessing"
	"github.com/siddhant-vij/Word-Frequency-Server/internal/worker"
)

func HandleFrequencyRequest(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")
	if word == "" {
		http.Error(w, "Missing 'word' parameter", http.StatusBadRequest)
		return
	}

	numServerThreads, err := strconv.Atoi(r.URL.Query().Get("numServerThreads"))
	if err != nil || numServerThreads <= 0 {
		http.Error(w, "Invalid 'numServerThreads' parameter", http.StatusBadRequest)
		return
	}

	pool := worker.NewPool(numServerThreads)
	pool.Start()

	var frequency int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range pool.Results {
			frequency += result
		}
	}()

	filePath := filepath.Join("resources", "war_and_peace.txt")
	if err := textprocessing.ReadFileAndDistributeTasks(pool, filePath, word); err != nil {
		http.Error(w, fmt.Sprintf("Error processing file: %v", err), http.StatusInternalServerError)
		return
	}

	close(pool.Tasks)
	pool.Wg.Wait()

	close(pool.Results)
	wg.Wait()

	fmt.Fprintf(w, "Frequency of '%s': %d", word, frequency)
}
