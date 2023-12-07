package banta

var (
	DataKline = []Kline{
		{1688169600000, 30460.2, 30668.2, 30311.3, 30573.6, 135520.246},
		{1688256000000, 30573.6, 30800, 30149.9, 30612.7, 231866.18800000002},
		{1688342400000, 30612.7, 31395.2, 30559.4, 31149, 370293.2360000001},
		{1688428800000, 31149, 31319.4, 30600, 30756.1, 300832.26},
		{1688515200000, 30756.1, 30875.2, 30175.8, 30488.4, 294896.578},
		{1688601600000, 30488.4, 31568, 29818, 29874.4, 721267.2189999999},
		{1688688000000, 29874.3, 30443.6, 29680, 30327.9, 337801.65400000004},
		{1688774400000, 30328, 30380, 30026.8, 30269.3, 138734.49300000002},
		{1688860800000, 30269.2, 30443.4, 30042, 30147.8, 162296.263},
		{1688947200000, 30147.8, 31040, 29928.8, 30396.9, 429115.53699999995},
		{1689033600000, 30396.9, 30804.9, 30261.4, 30608.4, 298904.7469999999},
		{1689120000000, 30608.3, 30980.5, 30186, 30368.9, 425058.257},
		{1689206400000, 30368.9, 31850, 30233, 31441.7, 696023.5100000001},
		{1689292800000, 31441.6, 31640, 29876.6, 30293.3, 538692.2459999999},
		{1689379200000, 30293.3, 30380, 30220.2, 30276.4, 111622.47400000002},
		{1689465600000, 30276.5, 30441.6, 30050, 30216.8, 173805.94},
		{1689552000000, 30216.9, 30329.7, 29630, 30126.1, 344113.42699999997},
		{1689638400000, 30126, 30227.5, 29400, 29845.6, 337396.417},
		{1689724800000, 29845.7, 30178.5, 29745, 29895.5, 298418.854},
		{1689811200000, 29895.5, 30420, 29468.8, 29791, 410263.76399999997},
		{1689897600000, 29791, 30056.2, 29705.9, 29891.4, 221145.07399999994},
		{1689984000000, 29891.5, 29986, 29602.2, 29783.5, 143375.95600000003},
		{1690070400000, 29783.5, 30368.2, 29718.9, 30070.8, 224713.84599999996},
		{1690156800000, 30070.9, 30091.7, 28830, 29163.8, 434182.417},
		{1690243200000, 29163.8, 29379.5, 29033, 29216.3, 196438.618},
		{1690329600000, 29216.2, 29670.5, 29083, 29336, 318527.58199999994},
		{1690416000000, 29336, 29558.4, 29065.5, 29209.7, 223376.53400000007},
		{1690502400000, 29209.8, 29535.7, 29112.8, 29299.9, 221299.86699999997},
		{1690588800000, 29300, 29397.9, 29237.5, 29339.1, 87100.719},
		{1690675200000, 29339.1, 29442.5, 29006, 29271.1, 169901.253},
		{1690761600000, 29271.2, 29525.8, 29100, 29220.8, 230381.72199999998},
		{1690848000000, 29220.8, 29735, 28550, 29701.2, 442328.643},
		{1690934400000, 29701.2, 30059.9, 28906.3, 29170.1, 491725.49199999997},
		{1691020800000, 29170.2, 29440, 28955.4, 29180.2, 251967.94199999998},
		{1691107200000, 29180.1, 29323.8, 28780, 29101.1, 229395.67300000004},
		{1691193600000, 29101.1, 29145, 28960, 29057.7, 82508.624},
		{1691280000000, 29057.7, 29199, 28978.5, 29075.9, 103094.43999999999},
		{1691366400000, 29075.9, 29274.5, 28682.3, 29202.7, 297752.4390000001},
		{1691452800000, 29202.7, 30250, 29132.4, 29759, 459269.80100000004},
		{1691539200000, 29758.9, 30149.7, 29362, 29572.8, 362009.70100000006},
		{1691625600000, 29572.9, 29729.8, 29303.7, 29443.7, 234546.03800000003},
		{1691712000000, 29443.7, 29565.1, 29220, 29415.5, 179044.008},
		{1691798400000, 29415.5, 29470, 29360, 29420.7, 56238.034999999996},
		{1691884800000, 29420.8, 29463, 29256.6, 29293.3, 84190.505},
		{1691971200000, 29293.3, 29686.7, 29070, 29419.5, 295035.38300000003},
		{1692057600000, 29419.5, 29492.1, 29050, 29188.8, 191289.959},
		{1692144000000, 29188.9, 29257.4, 28705.1, 28714.4, 280543.0889999999},
		{1692230400000, 28714.4, 28775.9, 24581, 26609.7, 868508.2199999996},
		{1692316800000, 26609.7, 26818, 25600, 26042.1, 522375.31599999993},
		{1692403200000, 26042, 26269.4, 25783.4, 26088.3, 202219.153},
		{1692489600000, 26088.4, 26285, 25948.5, 26175.9, 141089.85},
		{1692576000000, 26175.9, 26280, 25800, 26115.4, 232505.26499999993},
		{1692662400000, 26115.4, 26128.2, 25280, 26044.4, 313113.576},
		{1692748800000, 26044.5, 26806, 25800, 26419.2, 412969.02800000005},
		{1692835200000, 26419.2, 26568.3, 25835, 26164.6, 286753.33699999994},
		{1692921600000, 26164.6, 26300, 25754.4, 26051.7, 274830.49399999995},
		{1693008000000, 26051.7, 26129.4, 25969, 26004.3, 63925.861999999994},
		{1693094400000, 26004.3, 26173.6, 25955.6, 26087.7, 86505.398},
	}
)

func RunFakeEnv(env *BarEnv, barCb func(int, Kline)) {
	env.Reset()
	for i, bar := range DataKline {
		env.OnBar(&bar)
		if barCb != nil {
			barCb(i, bar)
		}
	}
}