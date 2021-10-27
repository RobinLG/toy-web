package v1

import (
	"fmt"
	"time"
)

type Filter func(c *Context)

type FilterBuilder func(next Filter) Filter

func MetricFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		startTime := time.Now().UnixNano()
		next(c)
		endTime := time.Now().UnixNano()
		fmt.Printf("run time: %d \n", endTime-startTime)
	}
}
