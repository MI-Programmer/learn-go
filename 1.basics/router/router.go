package router

func Get(path string) map[string]string {
	return map[string]string{"method": "Get", "path": path}
}

func Post(path string) map[string]string {
	return map[string]string{"method": "Post", "path": path}
}
