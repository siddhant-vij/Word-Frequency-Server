package textprocessing

import (
	"bufio"
	"os"
	"time"

	"github.com/siddhant-vij/Word-Frequency-Server/internal/worker"
)

func ReadFileAndDistributeTasks(pool *worker.Pool, filePath, word string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		task := NewWordCountTask(text, word)
		if len(pool.Tasks) > int(float64(cap(pool.Tasks))*0.8) {
			time.Sleep(time.Millisecond * 100)
		}
		pool.Tasks <- task
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
