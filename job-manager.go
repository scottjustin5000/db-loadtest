package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var db *sql.DB
var queryResults []ResponseResult
var queries []string
var wg sync.WaitGroup

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func initDatabase(connection string) *sql.DB {
	dbInfo := GetConnectionInfo(connection)
	fmt.Println(dbInfo)
	db, err := sql.Open(dbInfo.Type, dbInfo.ParsedConnection)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected")

	db.SetMaxIdleConns(100)
	return db
}

func runJob(index int, query string) {
	defer wg.Done()
	fmt.Println("starting job...")
	results := Execute(db, index, query)
	fmt.Println("job complete...")
	queryResults = append(queryResults, results)
}

func getAverageExecutionTime() int {
	var totalvalues int

	for _, qr := range queryResults {
		totalvalues += qr.ResponseTimeMs
	}

	var count = len(queryResults)
	return totalvalues / count
}

func printResults() {
	for _, qr := range queryResults {
		fmt.Printf("Index: %d, Time: %d ms, Error: %s\n", qr.Index, qr.ResponseTimeMs, qr.ErrorMessage)
	}
	avg := getAverageExecutionTime()
	fmt.Printf("Average time: %d ms", avg)

}

// Launch starts the job manager to condunt db load tests
func Launch(config Config) {
	if config.Concurrency > 0 {
		return
	}
	db = initDatabase(config.ConnectionString)
	defer db.Close()
	fmt.Println("Initialized...")
	queries = config.Queries()
	if config.Duration > 0 {
		launchDurationBased(config)
	} else {
		launchFileBased(config)
	}
	printResults()
}

func launchSequential(config Config) {
	start := time.Now()
	counter := 0
	shouldEnd := start.Add(time.Second * time.Duration(config.Duration))
	for time.Now().Before(shouldEnd) {
		for i := 0; i < config.RequestsPerSecond; i++ {
			if counter >= len(queries) && config.EOF == false {
				//reset
				counter = 0
			}
			//if we have no more queries and want to end once we have exhausted our queries we stop
			if counter >= len(queries) {
				break
			}
			wg.Add(1)
			go runJob(counter, queries[counter])
			counter++
		}
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}

func launchRandom(config Config) {
	start := time.Now()
	rand.Seed(time.Now().Unix())
	var index int
	min := 0
	max := len(queries)
	shouldEnd := start.Add(time.Second * time.Duration(config.Duration))
	for time.Now().Before(shouldEnd) {
		for i := 0; i < config.RequestsPerSecond; i++ {
			index = rand.Intn(max-min) + min
			wg.Add(1)
			go runJob(index, queries[index])
		}
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}

func launchDurationBased(config Config) {
	fmt.Println("Starting run")
	if config.Random == true && config.EOF == false {
		launchRandom(config)
	} else {
		launchSequential(config)
	}
}

func launchConcurrencyFileBased(config Config) {
	concurrency := config.Concurrency
	limit := len(queries)
	var iterator = 0
	for iterator < limit {
		indexes := makeRange(iterator, concurrency)
		for i := range indexes {
			wg.Add(1)
			go runJob(i, queries[i])
		}
		iterator++
		wg.Wait()
	}
}

func launchRpsFileBased(config Config) {
	leng := len(queries)
	fmt.Println("X", len(queries))
	i := 0
	for i < leng {
		for j := 0; j < config.RequestsPerSecond; j++ {
			if i >= leng {
				fmt.Println("BREAK!")
				break
			}
			wg.Add(1)
			go runJob(i, queries[i])
			i++
		}
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}

func launchFileBased(config Config) {
	if config.Concurrency > 0 {
		launchConcurrencyFileBased(config)
	} else {
		launchRpsFileBased(config)
	}

}
