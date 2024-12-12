package main

import (
	"dym_model/generator"
	"dym_model/model"
)

func main() {
	generator.Add("./cmd/dym_model/model", "application_dym_gen.go", generator.Target{Model: &model.Application{}, TableName: "Application"})
	generator.Add("./cmd/dym_model/model", "berth_dym_gen.go", generator.Target{Model: &model.Berth{}, TableName: "Berth"})
	generator.Add("./cmd/dym_model/model", "map_dym_gen.go", generator.Target{Model: &model.Map{}, TableName: "Map"})
	generator.Add("./cmd/dym_model/model", "order_dym_gen.go", generator.Target{Model: &model.Order{}, TableName: "Order"})
}
