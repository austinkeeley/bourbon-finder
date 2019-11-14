// Functions for sorting and formatting the data
package bourbonfinder

import (
	"fmt"
	"sort"
)

// Prints the results as a table
func PrintTable(results []SearchResult) {
	fmt.Printf("%-40s %-40s %-5s\n", "Location", "Product", "Quantity")
	fmt.Printf("-------------------------------------------------------------------------------------------\n")

	for _, result := range results {
		if result.Quantity > 0 {
			fmt.Printf("\033[1;32m")
		}
		fmt.Printf("%-40s %-40s %-5d\n", result.StoreName, result.ProductName, result.Quantity)
		fmt.Printf("\033[0m")
	}

}

// Prints the results grouped by store
func PrintGroup(m map[string][]SearchResult) {
	fmt.Printf("%-40s %-40s %-5s\n", "Location", "Product", "Quantity")
	fmt.Printf("-------------------------------------------------------------------------------------------\n")
	for store, results := range m {
		fmt.Printf("%-41s", store)
		first := true
		sort.SliceStable(results, func(i, j int) bool {
			return results[i].ProductName < results[j].ProductName
		})

		for _, result := range results {
			if result.Quantity > 0 {
				fmt.Printf("\033[1;32m")
			}
			if !first {
				fmt.Printf("%-40s %-40s %-5d\n", "", result.ProductName, result.Quantity)
			} else {
				fmt.Printf("%-40s %-5d\n", result.ProductName, result.Quantity)
			}
			first = false
			fmt.Printf("\033[0m")
		}
	}

}

func SortByStore(results []SearchResult) []SearchResult {
	sort.SliceStable(results, func(i, j int) bool {
		return results[i].StoreName < results[j].StoreName
	})
	return results
}

func GroupByStore(results []SearchResult) map[string][]SearchResult {
	m := make(map[string][]SearchResult)

	for _, result := range results {
		if len(m[result.StoreName]) == 0 {
			m[result.StoreName] = make([]SearchResult, 0)
		}
		m[result.StoreName] = append(m[result.StoreName], result)
	}

	return m
}
