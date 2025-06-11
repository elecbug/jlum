package main

import (
	"flag"
	"jlum-server/internal/db"
	"log"
)

func main() {
	dbPassword := flag.String("db-password", "", "MySQL password for jlum_user")
	flag.Parse()

	db.InitDB(*dbPassword)
	defer db.DB.Close()

	testDBOperations()
}

func testDBOperations() {
	// Test Insert
	affiliation := &db.Affiliation{Name: "Test University", Link: "http://testuniversity.edu"}
	id, err := db.InsertAffiliation(affiliation)
	if err != nil {
		log.Fatalf("InsertAffiliation failed: %v", err)
	}
	log.Printf("Inserted Affiliation with ID: %d", id)

	// Test Get
	retrievedAffiliation, err := db.GetAffiliation("id", id)
	if err != nil {
		log.Fatalf("GetAffiliation failed: %v", err)
	}
	log.Printf("Retrieved Affiliation: %+v", retrievedAffiliation)

	// Test Update
	if err := db.UpdateAffiliation("id", id, "name", "Updated University"); err != nil {
		log.Fatalf("UpdateAffiliation failed: %v", err)
	}
	log.Println("Updated Affiliation name")

	// Test Delete
	if err := db.DeleteAffiliation("id", id); err != nil {
		log.Fatalf("DeleteAffiliation failed: %v", err)
	}
	log.Println("Deleted Affiliation")
}
