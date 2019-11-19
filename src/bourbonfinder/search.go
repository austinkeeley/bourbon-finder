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

func Search(config *bourboncommon.Config) ([]SearchResult, error) {

	client := &http.Client{}
	var wg sync.WaitGroup

	var results []SearchResult
	for _, store := range config.Stores {
		log.Println("Searching " + store.Name)
		for _, product := range config.Wishlist {
			wg.Add(1)
			go ProductStoreSearch(client, store, product, func(result SearchResult, err error) {
				if err != nil {
					log.Println("Error trying to search " + store.Name + " and product " + product.Name)
				} else {
					results = append(results, result)
				}
				wg.Done()
			})
		}
	}

	wg.Wait()
	return results, nil
}

// Searches for a product at a store
func ProductStoreSearch(client *http.Client, store bourboncommon.Store,
	product bourboncommon.Product, cb func(result SearchResult, err error)) {
	url := fmt.Sprintf("%s?storeNumbers=%d&productCodes=%s", API_ENDPOINT, store.StoreID, product.ProductCode)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	resp, webErr := client.Do(req)
	if webErr != nil {
		cb(SearchResult{}, webErr)
	}

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
