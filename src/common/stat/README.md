# stat
	api接口统计

sql 记录 --- 原数据
	api name
	req ip
	timestamp --- 毫秒
	
sql 记录 --- 最小单位分钟
	api name
	count
	time --- 秒 / 分		timestamp/60
	
读取的时候 按照分钟直接读取后 在内存处理输出 相应数据

当前模块提供接口数据输出

在应用场景增加界面输出等等

增加 https://www.echartsjs.com/examples/ 输出
