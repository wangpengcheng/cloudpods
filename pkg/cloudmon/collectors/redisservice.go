package collectors

import (
	"yunion.io/x/onecloud/pkg/cloudmon/collectors/common"
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/util/shellutils"
)

func init() {
	shellutils.R(&common.ReportOptions{}, "report-redis", "Report Redis", reportRedis)
}

func reportRedis(session *mcclient.ClientSession, args *common.ReportOptions) error {
	return common.ReportCloudMetricOfoperatorType(string(common.REDIS), session, args)
}