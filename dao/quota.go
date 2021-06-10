package dao

import (
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"innogrid.com/hcloud-classic/pb"
)

// ReadQuota : Get the quota of the group
func ReadQuota(groupID int64) (*pb.GroupQuota, error) {
	var quota pb.GroupQuota

	var groupName string
	var limitCPUCores int
	var limitMemoryGB int
	var limitSubnetCnt int
	var limitAdaptiveIPCnt int
	var poolName string
	var limitSSDGB int
	var limitHDDGB int

	sql := "select piccolo.quota.group_id, piccolo.group.name as group_name, " +
		"piccolo.quota.limit_cpu_cores, piccolo.quota.limit_memory_gb, " +
		"piccolo.quota.limit_subnet_cnt, piccolo.quota.limit_adaptive_ip_cnt, " +
		"piccolo.quota.pool_name, piccolo.quota.limit_ssd_gb, piccolo.quota.limit_hdd_gb" +
		" from piccolo.quota, piccolo.group where piccolo.quota.group_id = piccolo.group.id" +
		" and piccolo.quota.group_id = ?"
	row := mysql.Db.QueryRow(sql, groupID)
	err := mysql.QueryRowScan(row,
		&groupID,
		&groupName,
		&limitCPUCores,
		&limitMemoryGB,
		&limitSubnetCnt,
		&limitAdaptiveIPCnt,
		&poolName,
		&limitSSDGB,
		&limitHDDGB)
	if err != nil {
		errStr := "ReadGroup(): " + err.Error()
		logger.Logger.Println(errStr)

		return nil, err
	}

	quota = pb.GroupQuota{
		GroupID:            groupID,
		GroupName:          groupName,
		LimitCPUCores:      int32(limitCPUCores),
		LimitMemoryGB:      int32(limitMemoryGB),
		LimitSubnetCnt:     int32(limitSubnetCnt),
		LimitAdaptiveIPCnt: int32(limitAdaptiveIPCnt),
		PoolName:           poolName,
		LimitSSDGB:         int32(limitSSDGB),
		LimitHDDGB:         int32(limitHDDGB),
	}

	return &quota, nil
}
