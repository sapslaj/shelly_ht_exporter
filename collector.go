package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type ShellyHTCollector struct {
	recorder     *Recorder
	temperature  *prometheus.Desc
	humiditiy    *prometheus.Desc
	lastReported *prometheus.Desc
}

var deviceLabels = []string{
	"id",
}

func NewShellyHTCollector(recorder *Recorder) *ShellyHTCollector {
	return &ShellyHTCollector{
		recorder:     recorder,
		temperature:  prometheus.NewDesc("shellyht_temperature", "Temperature", deviceLabels, nil),
		humiditiy:    prometheus.NewDesc("shellyht_humidity", "Humidity", deviceLabels, nil),
		lastReported: prometheus.NewDesc("shellyht_last_report", "Timestamp of last report", deviceLabels, nil),
	}
}

func (c *ShellyHTCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.temperature
	ch <- c.humiditiy
	ch <- c.lastReported
}

func (c *ShellyHTCollector) Collect(ch chan<- prometheus.Metric) {
	for id, record := range c.recorder.Store {
		ch <- prometheus.MustNewConstMetric(c.humiditiy, prometheus.GaugeValue, record.Humidity, id)
		ch <- prometheus.MustNewConstMetric(c.temperature, prometheus.GaugeValue, record.Temperature, id)
		ch <- prometheus.MustNewConstMetric(c.lastReported, prometheus.GaugeValue, float64(record.LastReport.Unix()), id)
	}
}
