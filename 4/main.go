package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	initDB()

	r := gin.Default()

	r.POST("/api/transaction", handleTransaction)

	r.Run(":8080")
}

func handleTransaction(c *gin.Context) {
	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}
	defer tx.Rollback()

	// Perform database operations within the transaction
	// For simplicity, let's assume we are inserting into 10 different tables
	if err := insertDataIntoTable1(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into Table1"})
		return
	}
	if err := insertDataIntoTable2(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into Table2"})
		return
	}
	if err := insertDataIntoTable3(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into Table2"})
		return
	}
	if err := insertDataIntoTable4(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into Table2"})
		return
	}
	if err := insertDataIntoTable5(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into Table2"})
		return
	}
	if err := insertDataIntoTable6(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into Table2"})
		return
	}
	if err := insertDataIntoTable7(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into Table2"})
		return
	}
	if err := insertDataIntoTable8(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into Table2"})
		return
	}
	if err := insertDataIntoTable9(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into Table2"})
		return
	}
	if err := insertDataIntoTable10(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into Table2"})
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction successful"})
}

func insertDataIntoTable1(tx *sql.Tx) error {
	// Perform insert operation for Table1
	// Example: test , 1
	_, err := tx.Exec("INSERT INTO Table1 (column1, column2) VALUES (?, ?)", "test", 1)
	return err
}

func insertDataIntoTable2(tx *sql.Tx) error {
	// Perform insert operation for Table2
	// Example: test , 2
	_, err := tx.Exec("INSERT INTO Table2 (column1, column2) VALUES (?, ?)", "test", 2)
	return err
}

func insertDataIntoTable3(tx *sql.Tx) error {
	// Perform insert operation for Table3
	// Example: test , 3
	_, err := tx.Exec("INSERT INTO Table3 (column1, column2) VALUES (?, ?)", "test", 2)
	return err
}
func insertDataIntoTable4(tx *sql.Tx) error {
	// Perform insert operation for Table4
	// Example: test , 4
	_, err := tx.Exec("INSERT INTO Table4 (column1, column2) VALUES (?, ?)", "test", 4)
	return err
}
func insertDataIntoTable5(tx *sql.Tx) error {
	// Perform insert operation for Table5
	// Example: test , 5
	_, err := tx.Exec("INSERT INTO Table5 (column1, column2) VALUES (?, ?)", "test", 5)
	return err
}
func insertDataIntoTable6(tx *sql.Tx) error {
	// Perform insert operation for Table6
	// Example: test , 6
	_, err := tx.Exec("INSERT INTO Table6 (column1, column2) VALUES (?, ?)", "test", 6)
	return err
}
func insertDataIntoTable7(tx *sql.Tx) error {
	// Perform insert operation for Table7
	// Example: test , 7
	_, err := tx.Exec("INSERT INTO Table7 (column1, column2) VALUES (?, ?)", "test", 7)
	return err
}
func insertDataIntoTable8(tx *sql.Tx) error {
	// Perform insert operation for Table8
	// Example: test , 8
	_, err := tx.Exec("INSERT INTO Table8 (column1, column2) VALUES (?, ?)", "test", 8)
	return err
}
func insertDataIntoTable9(tx *sql.Tx) error {
	// Perform insert operation for Table9
	// Example: test , 9
	_, err := tx.Exec("INSERT INTO Table9 (column1, column2) VALUES (?, ?)", "test", 9)
	return err
}
func insertDataIntoTable10(tx *sql.Tx) error {
	// Perform insert operation for Table10
	// Example: test , 2
	_, err := tx.Exec("INSERT INTO Table10 (column1, column2) VALUES (?, ?)", "test", 10)
	return err
}