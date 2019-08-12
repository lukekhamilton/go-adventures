package main

import (
	"expvar"
	"log"
	"net/http"
	"time"

	//_ "expvar"
	"fmt"
)

// UpTimeVar
type UpTimeVar struct {
	value time.Time
}

func (v *UpTimeVar) Set(date time.Time) {
	v.value = date
}

func (v *UpTimeVar) String() string {
	//return v.value.Format(time.UnixDate)
	return fmt.Sprintf("%s", time.Since(v.value))
}

// TimeVar
type TimeVar struct {
	value time.Time
}

func (v *TimeVar) Set(date time.Time) {
	v.value = date
}

func (v *TimeVar) Add(duration time.Duration) {
	v.value = v.value.Add(duration)
}

func (v *TimeVar) String() string {
	return v.value.Format(time.UnixDate)
}

var (
	orderCounter       *expvar.Int
	balanceCounter     *expvar.Float
	transactionMetrics *expvar.String
	startTimeMetrics   *TimeVar
	upTimeMetrics      *UpTimeVar
)

func init() {
	orderCounter = expvar.NewInt("counter")
	balanceCounter = expvar.NewFloat("balance")
	transactionMetrics = expvar.NewString("transaction")
	startTimeMetrics = &TimeVar{value: time.Now()}
	upTimeMetrics = &UpTimeVar{value: time.Now()}
	expvar.Publish("StartTime", startTimeMetrics)
	expvar.Publish("UpTime", upTimeMetrics)
}

func main() {
	fmt.Println("booting....")
	go http.ListenAndServe(":8080", http.DefaultServeMux)
	go processOrder()

	//memstatsFunc := expvar.Get("memstats").(expvar.Func)
	//memstats := memstatsFunc().(runtime.MemStats)
	//fmt.Println(memstats.Alloc)

	//expvar.Do(func(variable expvar.KeyValue) {
	//	fmt.Printf("expvar.Key: %s expvar.Value: %s\n\n", variable.Key, variable.Value)
	//})

	// Sets a float value to expvar.Float metrics
	balanceCounter.Set(1000)
	// Sets a string to expvar.String metrics
	transactionMetrics.Set("this is my transaction")

	for i := 0; i < 10; i++ {
		go doStuff(fmt.Sprintf("rid-%v", i))
	}

	fmt.Scanln()
}

func processOrder() {
	for {
		time.Sleep(time.Second)
		// Adds an integer value to expvar.Int counter
		orderCounter.Add(1)
	}
}

type Transaction struct {
	RID       string
	startTime time.Time
	endTime   time.Time
	totalTime time.Duration
}

func StartTransaction(rid string) *Transaction {
	log.Printf("Staring: %v\n", rid)
	return &Transaction{
		startTime: time.Now(),
		RID:       rid,
	}
}

func (t *Transaction) EndTransaction() {
	t.endTime = time.Now()
	t.totalTime = time.Since(t.startTime)
	log.Printf("End: %v\n", t.RID)
}

func doStuff(rid string) {
	t := StartTransaction(rid)

	time.Sleep(time.Second * 3)

	t.EndTransaction()
}
