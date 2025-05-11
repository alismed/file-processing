package main

import (
    "strings"
    "testing"
)

func TestProcessCSV(t *testing.T) {
    csvData := `Id;Region;Name;Description
0123-4567-8910;10;product01;description of the product 01
0123-0000-8910;10;product02;description of the product 02
`
    var items []Order
    mockPutItem := func(order Order) error {
        items = append(items, order)
        return nil
    }

    err := ProcessCSV(strings.NewReader(csvData), mockPutItem)
    if err != nil {
        t.Fatalf("ProcessCSV failed: %v", err)
    }

    if len(items) != 2 {
        t.Errorf("Expected 2 items, got %d", len(items))
    }
    if items[0].Name != "product01" {
        t.Errorf("Expected first Name to be product01, got %s", items[0].Name)
    }
}