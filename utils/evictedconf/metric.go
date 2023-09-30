package evictedconf

func MetricName() string {

	env, err := GetEnvVars()
	if err != nil {
		return ""
	}

	if env.Type == "webpage" {
		metricName := "statuz_webpage"
		return metricName
	}
	if env.Type == "ip" {
		metricName := "statuz_ip"
		return metricName
	}

	return ""

}
