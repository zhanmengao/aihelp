package database

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var metricNamePrefix = ""

var (
	labelRedisCMD     = "rdb_command"
	labelRedisKey     = "key_fmt"
	labelRedisField   = "field_fmt"
	labelRedisResult  = "rdb_result"
	labelRedisRpcName = "cmd"
)

var (
	metricOpCounter    *prometheus.CounterVec
	metricOpDurationMS *prometheus.HistogramVec
	metricOpDataBytes  *prometheus.HistogramVec
)

func initMetrics(prefix string) []prometheus.Collector {
	metricNamePrefix = prefix
	metricOpCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metricNamePrefix + "_rdb_operation_count",
			Help: "redis型数据库操作的次数",
		},
		[]string{
			labelRedisCMD,
			labelRedisKey,
			labelRedisField,
			labelRedisRpcName,
		},
	)
	metricOpDurationMS = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    metricNamePrefix + "_rdb_duration_ms",
			Help:    "redis型数据库操作的耗时(ms)",
			Buckets: []float64{1, 100, 1000},
		},
		[]string{
			labelRedisCMD,
			labelRedisKey,
			labelRedisField,
			labelRedisResult,
		},
	)
	metricOpDataBytes = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    metricNamePrefix + "_rdb_data_bytes",
			Help:    "redis型数据库操作的数据大小(bytes)",
			Buckets: []float64{64, 1024, 10240},
		},
		[]string{
			labelRedisCMD,
			labelRedisKey,
			labelRedisField,
		},
	)
	return []prometheus.Collector{
		metricOpCounter,
		metricOpDurationMS,
		metricOpDataBytes,
	}
}

// IncrOperationCount
//
//	@Description:DB QPS。使用生成的gendb无需手动调用
//	@param cmd：HSET，HGET等redis命令字
//	@param key redis key format，不可上报完整Key
//	@param field redis hash field format，不可上报format后的完整field
func IncrOperationCount(cmd, key, field, rpcName string) {
	if metricOpCounter == nil {
		return
	}
	metricOpCounter.With(prometheus.Labels{
		labelRedisCMD:     cmd,
		labelRedisKey:     key,
		labelRedisField:   field,
		labelRedisRpcName: rpcName,
	}).Inc()
}

// ObserveDurationMS
//
//	@Description: 上报DB读写耗时，其中cmd、key、field的数量必须为可控。使用生成的gendb无需手动调用
//	@param cmd：HSET，HGET等redis命令字
//	@param key redis key format，不可上报完整Key
//	@param field redis hash field format，不可上报format后的完整field
//	@param dur	DB操作耗时
//	@param opErr succ/fail
func ObserveDurationMS(cmd, key, field string, dur time.Duration, opErr error) {
	if metricOpDurationMS == nil {
		return
	}
	result := "succ"
	if opErr != nil {
		result = "fail"
	}
	metricOpDurationMS.With(prometheus.Labels{
		labelRedisCMD:    cmd,
		labelRedisKey:    key,
		labelRedisField:  field,
		labelRedisResult: result,
	}).Observe(float64(dur.Milliseconds()))
}

// ObserveDataBytes
//
//	@Description: 上报读DB字节数。使用生成的gendb无需手动调用
//	@param cmd：HSET，HGET等redis命令字
//	@param key redis key format，不可上报完整Key
//	@param field redis hash field format，不可上报format后的完整field
//	@param size DB操作字节数
func ObserveDataBytes(cmd, key, field string, size int) {
	if metricOpDataBytes == nil {
		return
	}
	metricOpDataBytes.With(prometheus.Labels{
		labelRedisCMD:   cmd,
		labelRedisKey:   key,
		labelRedisField: field,
	}).Observe(float64(size))
}
