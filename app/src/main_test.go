package main

import (
    "errors"
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

func TestProcessCSV_EmptyFile(t *testing.T) {
    csvData := ""
    var items []Order
    err := ProcessCSV(strings.NewReader(csvData), func(order Order) error {
        items = append(items, order)
        return nil
    })
    if err == nil {
        t.Error("Expected error for empty CSV, got nil")
    }
}

func TestProcessCSV_HeaderOnly(t *testing.T) {
    csvData := "Id;Region;Name;Description\n"
    var items []Order
    err := ProcessCSV(strings.NewReader(csvData), func(order Order) error {
        items = append(items, order)
        return nil
    })
    if err != nil {
        t.Errorf("Did not expect error for header-only CSV, got %v", err)
    }
    if len(items) != 0 {
        t.Errorf("Expected 0 items, got %d", len(items))
    }
}

func TestProcessCSV_InvalidLine(t *testing.T) {
    csvData := "Id;Region;Name;Description\n1;2;3\n"
    var items []Order
    err := ProcessCSV(strings.NewReader(csvData), func(order Order) error {
        items = append(items, order)
        return nil
    })
    if err == nil {
        t.Error("Expected error for invalid line, got nil")
    }
}

func TestProcessCSV_CallbackError(t *testing.T) {
    csvData := "Id;Region;Name;Description\n1;2;3;4\n"
    err := ProcessCSV(strings.NewReader(csvData), func(order Order) error {
        return errors.New("mock error")
    })
    if err == nil {
        t.Error("Expected error from callback, got nil")
    }
}
