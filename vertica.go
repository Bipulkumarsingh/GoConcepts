package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/vertica/vertica-sql-go"
	"github.com/vertica/vertica-sql-go/logger"
)

func mainvertica() {
	//Have our logger output INFO and above
	logger.SetLogLevel(logger.INFO)

	var testLogger = logger.New("Samplecode")

	ctx := context.Background()

	// Create a connection to our database, connection is lazzy and won't happen until it's used.
	connDB, err := sql.Open("vertica", "credentials")

	if err != nil {
		testLogger.Fatal(err.Error())
		os.Exit(1)
	}
	defer connDB.Close() // Close the connection before leaving main function

	// Ping the database connection to force it to attempt to connect.
	if err = connDB.PingContext(ctx); err != nil {
		testLogger.Fatal(err.Error())
		os.Exit(1)
	}

	// query a standard metric table in vertica.
	rows, err := connDB.QueryContext(ctx, "SELECT column from table limit 10")

	if err != nil {
		testLogger.Fatal(err.Error())
		os.Exit(1)
	}
	// check all the columns of the rows
	//If you need to examine the names of the columns, simply access the Columns() operator of the rows object.
	columnNames, _ := rows.Columns()
	fmt.Println("columns Name: ", columnNames)
	// for _, columnName := range columnNames {
	// 		// use the column name here.
	// }

	// Iterate over the results and print them out.
	for rows.Next() {
		var question string
		if err = rows.Scan(&question); err != nil {
			testLogger.Fatal(err.Error())
			os.Exit(1)
		}
		testLogger.Info("%s", question)
	}
	testLogger.Info("Test complete")
	os.Exit(0)

	defer rows.Close()

}
