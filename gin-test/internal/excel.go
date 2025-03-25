package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func ExportDataToExcel(filename string) error {
	cursor, err := MongoCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return fmt.Errorf("failed to fetch data from MongoDB: %v", err)
	}
	defer cursor.Close(context.TODO())

	f := excelize.NewFile()
	sheet := "Sheet1"
	f.NewSheet(sheet)
	f.SetSheetRow(sheet, "A1", &[]string{"Name", "Age", "Email"})

	row := 2
	for cursor.Next(context.TODO()) {
		var data map[string]interface{}
		if err := cursor.Decode(&data); err != nil {
			log.Printf("Failed to decode data: %v", err)
			continue
		}
		f.SetSheetRow(sheet, fmt.Sprintf("A%d", row), &[]interface{}{data["name"], data["age"], data["email"]})
		row++
	}

	if err := f.SaveAs(filename); err != nil {
		return fmt.Errorf("failed to save Excel file: %v", err)
	}
	return nil
}
