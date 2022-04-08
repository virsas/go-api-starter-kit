package vars

import "os"

func SetPorts() (string, string) {
	var apiPort string = "8080"
	apiPortValue, apiPortPresent := os.LookupEnv("API_PORT")
	if apiPortPresent {
		apiPort = apiPortValue
	}

	var prometheusPort string = "8081"
	prometheusPortValue, prometheusPortPresent := os.LookupEnv("PRM_PORT")
	if prometheusPortPresent {
		prometheusPort = prometheusPortValue
	}

	return apiPort, prometheusPort
}
