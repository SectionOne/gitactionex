package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	// Obtenir els últims 3 commits
	cmd := exec.Command("git", "log", "-n", "3", "--pretty=format:%h - %an, %ar : %s")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executant git log: %v\n", err)
		os.Exit(1)
	}

	// Crear directori log si no existeix
	logDir := "log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.Mkdir(logDir, 0755)
		if err != nil {
			fmt.Printf("Error creant directori %s: %v\n", logDir, err)
			os.Exit(1)
		}
	}

	// Generar nom d'arxiu amb data i hora
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	logFile := filepath.Join(logDir, fmt.Sprintf("commits_%s.txt", currentTime))

	// Escriure a l'arxiu
	content := fmt.Sprintf("Últims 3 commits del repositori:\n\n%s", string(out))
	err = ioutil.WriteFile(logFile, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error escrivint a l'arxiu %s: %v\n", logFile, err)
		os.Exit(1)
	}

	fmt.Printf("Arxiu de log creat correctament a %s\n", logFile)
}
