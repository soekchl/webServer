package memory

import (
	"sync"
	"time"

	. "github.com/soekchl/myUtils"
)

var (
	saveMem              sync.Map                 // 保存主要数据
	saveSecond           = int64(30 * 60)         // 数据默认保留时常
	delList              []string                 // 删除列表， 越靠前先删除
	activeSecondMap      = make(map[string]int64) // 删除列表 记录删除时间 每次使用数据就从新记录时间
	activeSecondMapMutex sync.Mutex
)

func init() {
	go startCheck()
}

func startCheck() {
	defer func() {
		if err := recover(); err != nil {
			Error(err)
			startCheck()
		}
	}()
	Warn("[Memory] startCheck  ---> ")
	n := 0
	i := 0
	leaveTime := int64(0)
	activeTime := int64(0)
	for ; ; time.Sleep(time.Second) {
		n = 128
		if len(delList) < n {
			n = len(delList)
		}
		//Debugf("\t\tn=%v delList=%#v", n, delList) // n is check count
		if n < 1 {
			continue
		}

		activeSecondMapMutex.Lock()
		for i = 0; i < n; i++ {
			activeTime = activeSecondMap[delList[i]]
			leaveTime = activeTime - time.Now().Unix()
			if leaveTime > saveSecond/10*7 {
				//Debugf("move i=%v key=%v", i, delList[i])
				f := moveLast(i)
				if f {
					i--
					n--
				}
			} else if leaveTime < 1 {
				//Debugf("del i=%v key=%v", i, delList[i])
				delete(activeSecondMap, delList[i])
				delValue(i)
				i--
				n--
			}
		}
		activeSecondMapMutex.Unlock()
	}
	Warn("[Memory] startCheck <--- ")
}

func moveLast(i int) bool {
	n := len(delList)
	if n <= 1 || i+1 == n { // one len or last index
		return false
	}
	v := delList[i]
	if i == 0 {
		delList = delList[1:]
	} else {
		delList = append(delList[:i], delList[i+1:]...)
	}
	delList = append(delList, v)
	return true
}

func delValue(i int) {
	n := len(delList)
	if n <= 1 { // one len
		delList = []string{}
		return
	}
	if n == i+1 {
		delList = delList[:n-1]
		return
	}

	if i == 0 {
		delList = delList[1:]
	} else {
		delList = append(delList[:i], delList[i+1:]...)
	}
}

func Delete(key string) {
	saveMem.Delete(key)
	deleteLeave(key)
}

func deleteLeave(key string) {
	activeSecondMapMutex.Lock()
	defer activeSecondMapMutex.Unlock()
	if _, ok := activeSecondMap[key]; ok {
		activeSecondMap[key] = 0
	}
}

func Load(key string) interface{} {
	value, ok := saveMem.Load(key)
	if ok {
		activeLeave(key)
	}
	return value
}

func activeLeave(key string) {
	activeSecondMapMutex.Lock()
	defer activeSecondMapMutex.Unlock()
	if _, ok := activeSecondMap[key]; !ok {
		delList = append(delList, key)
	}
	activeSecondMap[key] = time.Now().Unix() + saveSecond
	//Debugf("active key=%v", key)
}

func Store(key string, value interface{}) {
	//Debugf("add key=%v", key)
	saveMem.Store(key, value)
	activeLeave(key)
}
