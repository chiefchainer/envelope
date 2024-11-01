package main

import (
	"testing"
	"time"
)

func TestMainExecution(t *testing.T) {
	const repeatCount = 10
	var totalDuration time.Duration

	for i := 0; i < repeatCount; i++ {
		start := time.Now()
		main()

		duration := time.Since(start)
		t.Logf("Тест %d занял %.3f секунд", i+1, duration.Seconds())
		totalDuration += duration
	}
	averageDuration := totalDuration / time.Duration(repeatCount)
	t.Logf("Среднее время выполнения %d тестов: %.3f секунд", repeatCount, averageDuration.Seconds())
}
