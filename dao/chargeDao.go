package dao

import (
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"innogrid.com/hcloud-classic/pb"
)

// ReadCharge : Get infos of the charge
func ReadCharge(groupID int64) (*pb.Charge, error) {
	var charge pb.Charge

	var chargeCPUPerCore int64
	var chargeMemoryPerGB int64
	var chargeNICList string
	var chargeSubnetPerCnt int64
	var chargeAdaptiveIPPerCnt int64
	var chargeSSDPerGB int64
	var chargeHDDPerGB int64
	var chargeTrafficPerKB float32

	sql := "select * from piccolo.charge where group_id = ?"
	row := mysql.Db.QueryRow(sql, groupID)
	err := mysql.QueryRowScan(row,
		&groupID,
		&chargeCPUPerCore,
		&chargeMemoryPerGB,
		&chargeNICList,
		&chargeSubnetPerCnt,
		&chargeAdaptiveIPPerCnt,
		&chargeSSDPerGB,
		&chargeHDDPerGB,
		&chargeTrafficPerKB)
	if err != nil {
		errStr := "ReadCharge(): " + err.Error()
		logger.Logger.Println(errStr)

		return nil, err
	}

	charge.GroupID = groupID
	charge.ChargeCPUPerCore = chargeCPUPerCore
	charge.ChargeMemoryPerGB = chargeMemoryPerGB
	charge.ChargeNicList = chargeNICList
	charge.ChargeSubnetPerCnt = chargeSubnetPerCnt
	charge.ChargeAdaptiveIPPerCnt = chargeAdaptiveIPPerCnt
	charge.ChargeSSDPerGB = chargeSSDPerGB
	charge.ChargeHDDPerGB = chargeHDDPerGB
	charge.ChargeTrafficPerKB = chargeTrafficPerKB

	return &charge, nil
}
