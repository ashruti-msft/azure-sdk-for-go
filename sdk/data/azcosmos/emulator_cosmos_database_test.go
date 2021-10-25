// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"context"
	"testing"
)

func TestDatabaseCRUD(t *testing.T) {
	emulatorTests := newEmulatorTests(t)
	client := emulatorTests.getClient(t)

	database := DatabaseProperties{ID: "baseDbTest"}

	resp, err := client.CreateDatabase(context.TODO(), database, nil)
	if err != nil {
		t.Fatalf("Failed to create database: %v", err)
	}

	if resp.DatabaseProperties.ID != database.ID {
		t.Errorf("Unexpected id match: %v", resp.DatabaseProperties)
	}

	db, _ := client.NewDatabase("baseDbTest")
	resp, err = db.Read(context.TODO(), nil)
	if err != nil {
		t.Fatalf("Failed to read database: %v", err)
	}

	if resp.DatabaseProperties.ID != database.ID {
		t.Errorf("Unexpected id match: %v", resp.DatabaseProperties)
	}

	throughputResponse, err := db.ReadThroughput(context.TODO(), nil)
	if err == nil {
		t.Fatalf("Expected not finding throughput but instead got : %v", throughputResponse)
	}

	resp, err = db.Delete(context.TODO(), nil)
	if err != nil {
		t.Fatalf("Failed to delete database: %v", err)
	}
}

func TestDatabaseWithOfferCRUD(t *testing.T) {
	emulatorTests := newEmulatorTests(t)
	client := emulatorTests.getClient(t)

	database := DatabaseProperties{ID: "baseDbTest"}
	tp := NewManualThroughputProperties(400)
	resp, err := client.CreateDatabase(context.TODO(), database, &CreateDatabaseOptions{ThroughputProperties: &tp})
	if err != nil {
		t.Fatalf("Failed to create database: %v", err)
	}

	if resp.DatabaseProperties.ID != database.ID {
		t.Errorf("Unexpected id match: %v", resp.DatabaseProperties)
	}

	db, _ := client.NewDatabase("baseDbTest")
	resp, err = db.Read(context.TODO(), nil)
	if err != nil {
		t.Fatalf("Failed to read database: %v", err)
	}

	if resp.DatabaseProperties.ID != database.ID {
		t.Errorf("Unexpected id match: %v", resp.DatabaseProperties)
	}

	throughputResponse, err := db.ReadThroughput(context.TODO(), nil)
	if err != nil {
		t.Fatalf("Failed to read throughput: %v", err)
	}

	mt, err := throughputResponse.ThroughputProperties.ManualThroughput()
	if err != nil {
		t.Errorf("Failed to read throughput: %v", err)
	}

	if mt != 400 {
		t.Errorf("Unexpected throughput: %v", mt)
	}

	newScale := NewManualThroughputProperties(500)
	_, err = db.ReplaceThroughput(context.TODO(), newScale, nil)
	if err != nil {
		t.Errorf("Failed to read throughput: %v", err)
	}

	resp, err = db.Delete(context.TODO(), nil)
	if err != nil {
		t.Fatalf("Failed to delete database: %v", err)
	}
}
