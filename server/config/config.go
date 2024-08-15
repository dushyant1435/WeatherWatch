package config

type AlertThresholds struct {
    Temperature float64
    Condition   string
    Consecutive int
}

// Example thresholds
var Thresholds = AlertThresholds{
    Temperature: 20.0,
    Condition:   "",  // Empty if no specific condition is tracked
    Consecutive: 2,
}
