package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error)  {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}

	snowflake.Epoch = st.UnixNano() / 10000000
	node, err = snowflake.NewNode(machineID)
	return
}


func GenID() int64  {
	return node.Generate().Int64()
}

//func main() {
//	if err := Init("2021-05-14", 1); err != nil {
//		return
//	}
//	ID := GenID()
//
//	fmt.Println(ID)
//}