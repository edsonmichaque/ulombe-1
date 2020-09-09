package ulombe

type Task struct {
        Description string `json:"description"`
        Resource string `json:"resource"`
        Operation string `json:"operation"`
        Data map[string]interface{} `json:"data"`
        Hosts []map[string]string `json:"hosts"`
}

