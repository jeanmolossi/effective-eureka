package tests

import (
	"strings"

	"github.com/cucumber/godog"
	"github.com/jeanmolossi/effective-eureka/src/core/shared"
)

func (a *ApiFeature) ThereAreAny(tableName string, data *godog.Table) error {
	var fields []string
	var marks []string

	head := data.Rows[0].Cells
	for _, cell := range head {
		fields = append(fields, cell.Value)
		marks = append(marks, "?")
	}

	dbConn := shared.NewDbConnection()
	err := dbConn.Connect()
	if err != nil {
		return err
	}

	for i := 1; i < len(data.Rows); i++ {
		var vals []interface{}
		for _, cell := range data.Rows[i].Cells {
			vals = append(vals, cell.Value)
		}

		stmt := dbConn.DB().Exec(
			`INSERT INTO `+tableName+` (`+strings.Join(fields, ",")+`) VALUES (`+strings.Join(marks, ",")+`)`,
			vals...,
		)

		if err := stmt.Error; err != nil {
			return err
		}
	}

	return nil
}
