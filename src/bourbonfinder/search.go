package bourbonfinder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"bourboncommon"
)

const (
	API_ENDPOINT = "https://www.abc.virginia.gov/webapi/inventory/mystore"
)

type SearchResult struct {
	StoreName   string // Name of the store where the product is
	ProductName string // Name of the product
	Quantity    int    // Quantity in stock
}

/*
type Store struct {
	Name    string `json:"name"`
	StoreID int    `json:"storeID"`
}

type Product struct {
	Name        string `json:"name"`
	ProductCode string `json:"productCode"`
}

type Config struct {
	Stores   []Store   `json:"stores"`
	Wishlist []Product `json:"wishlist"`
}
*/

// Data returned from the endpoint

type wsResult struct {
	Products []wsProduct `json:"products"`
}

type wsProduct struct {
	ProductID string    `json:"productId"`
	StoreInfo storeInfo `json:"storeInfo"`
}

type storeInfo struct {
	Quantity int `json:"quantity"`
}

// Performs a serach using the stores and wishlist from a config.json file
/*func Search(configFileName string) ([]SearchResult, error) {
log.Println("Opening file " + configFileName)
file, err := os.Open(configFileName)
if err != nil {
	log.Println("Could not open config file " + configFileName)
	return nil, err
}
defer file.Close()

var config *Config
err = json.NewDecoder(file).Decode(&config)
if err != nil {
	log.Println("Error parsing config file " + configFileName)
	return nil, err
}*/
func Search(config *bourboncommon.Config) ([]SearchResult, error) {

	client := &http.Client{}
	var wg sync.WaitGroup

	/*
		fmt.Printf("%-40s %-40s %-5s\n", "Location", "Product", "Quantity")
		fmt.Printf("-------------------------------------------------------------------------------------------\n")
	*/

	var results []SearchResult
	for _, store := range config.Stores {
		log.Println("Searching " + store.Name)
		for _, product := range config.Wishlist {
			wg.Add(1)
			go ProductStoreSearch(client, store, product, func(result SearchResult, err error) {
				/*
					if result.Quantity > 0 {
						fmt.Printf("\033[1;32m")
					}
					fmt.Printf("%-40s %-40s %-5d\n", result.StoreName, result.ProductName, result.Quantity)
					fmt.Printf("\033[0m")
				*/
				results = append(results, result)
				wg.Done()
			})
		}
	}

	wg.Wait()
	return results, nil
}

// Searches for a product at a store
func ProductStoreSearch(client *http.Client, store bourboncommon.Store, product bourboncommon.Product, cb func(result SearchResult, err error)) {
	url := fmt.Sprintf("%s?storeNumbers=%d&productCodes=%s", API_ENDPOINT, store.StoreID, product.ProductCode)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	resp, _ := client.Do(req)

	var r wsResult
	err := json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range r.Products {
		r := SearchResult{}
		r.StoreName = store.Name
		r.ProductName = product.Name
		r.Quantity = v.StoreInfo.Quantity
		cb(r, nil)
	}

}
