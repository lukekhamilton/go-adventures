package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sdcoffey/big"

	"github.com/sdcoffey/techan"
)

func main() {
	series := techan.NewTimeSeries()

	dataSet := [][]string{
		//Timestamp, Open, Close, High, Low, volume
		{"1234567", "1", "2", "3", "5", "6"},
	}

	for _, data := range dataSet {
		start, _ := strconv.ParseInt(data[0], 10, 64)
		period := techan.NewTimePeriod(time.Unix(start, 0), time.Hour*24)

		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(data[1])
		candle.ClosePrice = big.NewFromString(data[2])
		candle.MaxPrice = big.NewFromString(data[3])
		candle.MinPrice = big.NewFromString(data[4])

		series.AddCandle(candle)
	}

	closePrices := techan.NewClosePriceIndicator(series)
	movingAvg := techan.NewEMAIndicator(closePrices, 10)

	fmt.Println(movingAvg.Calculate(0).FormattedString(2))
}
