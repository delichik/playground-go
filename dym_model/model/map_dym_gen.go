package model

import (
	"database/sql"
	"strconv"
)

func (m *Map) Create(db *sql.DB) error {
	columnString := ""
	columnValueString := ""

	columnString += "ID"
	columnString += ","
	columnValueString += strconv.Itoa(m.ID)
	columnValueString += ","

	if m.Version != "" {
		columnString += "Version"
		columnString += ","
		columnValueString += "\""
		columnValueString += m.Version
		columnValueString += "\","
	}

	if m.Content != "" {
		columnString += "Content"
		columnString += ","
		columnValueString += "\""
		columnValueString += m.Content
		columnValueString += "\","
	}

	columnString += "Status"
	columnString += ","
	columnValueString += strconv.Itoa(m.Status)
	columnValueString += ","

	if m.Remark != "" {
		columnString += "Remark"
		columnString += ","
		columnValueString += "\""
		columnValueString += m.Remark
		columnValueString += "\","
	}

	if m.CreatedAt.String() != "" {
		columnString += "CreatedAt"
		columnString += ","
		columnValueString += "\""
		columnValueString += m.CreatedAt.String()
		columnValueString += "\","
	}

	if m.UpdatedAt.String() != "" {
		columnString += "UpdatedAt"
		columnString += ","
		columnValueString += "\""
		columnValueString += m.UpdatedAt.String()
		columnValueString += "\","
	}

	query := "insert into Map ("
	query += columnString
	query += ") values ("
	query += columnValueString
	query += ")"
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
