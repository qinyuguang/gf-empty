package snowflake

import (
	"sync"

	"github.com/gogf/gf/frame/g"
	"github.com/sony/sonyflake"
	"github.com/sony/sonyflake/awsutil"
)

type sonyflakeType int

const (
	// NodeTypePrivate private node.
	NodeTypePrivate sonyflakeType = iota
	// NodeTypeAWS AWS node.
	NodeTypeAWS
)

var sf *sonyflake.Sonyflake
var once *sync.Once

// Init can initialize snowflake setting.
func Init(nodeType sonyflakeType) {
	var st sonyflake.Settings

	switch nodeType {
	case NodeTypeAWS:
		st.MachineID = awsutil.AmazonEC2MachineID
	default:
		sf = sonyflake.NewSonyflake(st)
	}

	sf = sonyflake.NewSonyflake(st)

	if nodeType == NodeTypeAWS && sf == nil {
		g.Log().Error("aws sonyflake not created")
		st.MachineID = nil
		sf = sonyflake.NewSonyflake(st)
	}

	if sf == nil {
		g.Log().Fatal("sonyflake not created")
	}
}

// NextID generate uint64 id.
func NextID() uint64 {
	id, err := sf.NextID()

	if err != nil {
		g.Log().Error("err:", err)
	}

	return id
}

// NextIDInt64 generate int64 id.
func NextIDInt64() int64 {
	return int64(NextID())
}
