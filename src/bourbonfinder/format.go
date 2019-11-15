// Functions for sorting and formatting the data
package bourbonfinder

import (
	"fmt"
	"sort"
	"io"
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
func PrintGroup(m map[string][]SearchResult, writer io.Writer, color bool) {
	//fmt.Printf("%-40s %-40s %-5s\n", "Location", "Product", "Quantity")
	writer.Write([]byte(fmt.Sprintf("%-40s %-40s %-5s\n", "Location", "Product", "Quantity")))
	writer.Write([]byte("-------------------------------------------------------------------------------------------\n"))
	for store, results := range m {
		writer.Write([]byte(fmt.Sprintf("%-41s", store)))
		first := true
		sort.SliceStable(results, func(i, j int) bool {
			return results[i].ProductName < results[j].ProductName
		})

		for _, result := range results {
			if result.Quantity > 0  && color {
				writer.Write([]byte(fmt.Sprintf("\033[1;32m")))
			}
			if !first {
				writer.Write([]byte(fmt.Sprintf("%-40s %-45s %-5d\n", "", result.ProductName, result.Quantity)))
			} else {
				writer.Write([]byte(fmt.Sprintf("%-45s %-5d\n", result.ProductName, result.Quantity)))
			}
			first = false
			if color {
				writer.Write([]byte(fmt.Sprintf("\033[0m")))
			}
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
