package coininfo

import (
	signproxy "github.com/NpoolPlatform/message/npool/signproxy"
)

func Int32ToCoinType(n int32) (ct signproxy.CoinType) {
	ct = signproxy.CoinType(n)
	return
}

func CoinTypeToInt32(ct signproxy.CoinType) (n int32) {
	n = int32(ct)
	return
}
