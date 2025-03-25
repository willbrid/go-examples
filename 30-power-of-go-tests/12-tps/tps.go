package tps

import "os"

func WriteReportFile(path, content string) {
	err := os.WriteFile(path, []byte(content), 0o766)
	if err != nil {
		panic(err)
	}
}
