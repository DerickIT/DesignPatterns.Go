package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var r *rand.Rand

type RedPacket struct {
	TotalAmount float64
	Count       int
	Expiry      time.Time
	Amounts     []float64
	mu          sync.Mutex
}

type GrabRecord struct {
	UserName string
	Amount   float64
	Time     time.Time
}

func CreateRedPacket(totalAmount float64, count int, expiryMinutes int) *RedPacket {
	rp := &RedPacket{
		TotalAmount: totalAmount,
		Count:       count,
		Expiry:      time.Now().Add(time.Duration(expiryMinutes) * time.Minute),
		Amounts:     make([]float64, count),
	}
	rp.distributeAmount()
	return rp
}

func (rp *RedPacket) distributeAmount() {
	remaining := rp.TotalAmount
	for i := 0; i < rp.Count-1; i++ {
		maxAmount := remaining / 2
		// randFloat := rand.Float64()
		remainingSum := remaining - 0.01*float64(rp.Count-i-1)
		if maxAmount > remainingSum {
			maxAmount = remainingSum
		}
		// amount := 0.01 + randFloat*(maxAmount-0.01)
		// amount = float64(int(amount*100)) / 100
		amount := 0.01 + r.Float64()*(maxAmount-0.01)
		amount = math.Floor(amount*100) / 100 // 向下取整到分
		rp.Amounts[i] = amount
		remaining -= amount
	}
	// rp.Amounts[rp.Count-1] = float64(int(remaining*100)) / 100
	rp.Amounts[rp.Count-1] = math.Round(remaining*100) / 100
}

func (rp *RedPacket) GrabRedPacket(userName string) (GrabRecord, error) {
	rp.mu.Lock()
	defer rp.mu.Unlock()

	if time.Now().After(rp.Expiry) {
		return GrabRecord{}, fmt.Errorf("红包已过期")
	}
	if len(rp.Amounts) == 0 {
		return GrabRecord{}, fmt.Errorf("红包已被抢完")
	}
	amount := rp.Amounts[0]
	rp.Amounts = rp.Amounts[1:]

	record := GrabRecord{
		UserName: userName,
		Amount:   amount,
		Time:     time.Now(),
	}
	return record, nil
}

func grabRedPacketWorker(rp *RedPacket, userName string, resultChan chan<- GrabRecord, wg *sync.WaitGroup) {
	defer wg.Done()
	record, err := rp.GrabRedPacket(userName)
	if err != nil {
		fmt.Printf("用户 %s 抢红包失败: %s\n", userName, err)
		return
	}
	resultChan <- record
}

func main() {
	// rand.Seed(time.Now().UnixNano())

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成15个用户名
	userNames := []string{
		"张三", "李四", "王五", "赵六", "钱七",
		"孙八", "周九", "吴十", "郑十一", "王十二",
		"冯十三", "陈十四", "楮十五", "魏十六", "蒋十七",
	}

	redPacket := CreateRedPacket(100, 5, 5)

	var wg sync.WaitGroup
	resultChan := make(chan GrabRecord, len(userNames))

	for _, userName := range userNames {
		wg.Add(1)
		go grabRedPacketWorker(redPacket, userName, resultChan, &wg)
	}

	// 等待所有goroutine完成
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集并打印结果
	var totalGrabbed float64
	var successfulGrabs int
	for record := range resultChan {
		fmt.Printf("用户 %s 抢到红包: %.2f 元, 时间: %v\n", record.UserName, record.Amount, record.Time)
		totalGrabbed += record.Amount
		successfulGrabs++
	}

	fmt.Printf("\n总共 %d 人成功抢到红包, 总金额: %.2f 元\n", successfulGrabs, totalGrabbed)
}
