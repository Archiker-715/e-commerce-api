package app

import "github.com/Archiker-715/e-commerce-api/internal/elastic"

func startElastic() *elastic.Elastic {
	return elastic.NewElastic()
}
