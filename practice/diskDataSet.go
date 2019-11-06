package practice

import "log"

type DiskDataSet struct {
	logger log.Logger

	dataSetName string
	dataSetType string
	serialNumber string

	bucketPartition string
	dataCollection *DataCollection

}
