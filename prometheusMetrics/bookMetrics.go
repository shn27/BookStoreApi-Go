package prometheusMetrics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(BookCreateCounter)
}

var BookCreateCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "book_create_count",
		Help: "Number of books successfully added to database",
	},
)
