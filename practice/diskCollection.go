package practice

import (
	"flag"
	"log"
)

var (
	collectionTaskName string
	collectionDeviceName string
)

func init() {
	flag.StringVar(&collectionTaskName, "name", "", "get task")
	flag.StringVar(&collectionDeviceName, "type", "anc", "co device type")
}

type DiskCollection struct {
	logger log.Logger

	collectionName string
	CollectionType string
	mountPoint string
	bucketPartition string

	diskInfoFile string


	infoDataSets map[string]*DiskDataSet
	diskDataSets map[string]*DiskDataSet

}
