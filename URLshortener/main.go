package main

import (
	"fmt"

	"sync"
	"utils/helpers"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

// All in all To-do:

// -
//
// - Simple Database intergration
//
// - In the search engine:
// 		-> Keyword: URL redirect mapping
// 		- Hey SearchEngine, this links that you have put in leads actually
// 		  to a longer URL
//		- HTTP-Protocol-Statuscode: 301

// Process:

// 1.Try URL 						done. test no
// 2.Shorten URL 					done, test no
// 3.Put in on DB
// 4.Read it from DB
// 5.Delete it manually from DB
//
// 6. Create with localhost:8080, http.HandleFunc, http.HandleRedirect ...
// 7. Test in the browser
var (
	wg sync.WaitGroup
	//db  *sql.DB //this is nil
	url string
)

func main() {
	url = "https://google.com"
	ch := make(chan string)

	if err := helpers.Connection(); err != nil {
		fmt.Println("holly molly")
	}

	wg.Add(1)
	go func() {
		ch <- helpers.ShortURL(url)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		helpers.IsURL(url)
		wg.Done()
	}()
	wg.Wait()

	// These two do not work:

	//helpers.TableCreate()
	//helpers.CreateData(surl, url)

	/*

		if err := helpers.ReadShortURL(surl); err != nil {
			fmt.Println("no")
		} else if err := helpers.DeleteRow(url); err != nil {
			fmt.Println("oh man:(")
		} else {
			fmt.Println("all done mate!")
		}

		defer os.Remove("./helpers/local.db") // for testing purposes
		defer db.Close()
	*/
}
