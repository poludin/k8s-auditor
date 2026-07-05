/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k8s-auditor",
	Short: "CLI-сканер для проверки K8s-инфраструктуры",
	Long: `k8s-auditor — это инструмент для автоматического аудита кластеров.
Он проверяет манифесты на соответствие 14 правилам "Золотого стандарта", 
находит "мусорные" ресурсы (например, оставшиеся после сетапа сервисы), 
валидирует ConfigMaps и проверяет логику формирования URL-адресов для разных окружений.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var kubeconfig string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Если мы запускаем утилиту на локальном ноуте, по умолчанию ищем конфиг в ~/.kube/config
	home := os.Getenv("HOME")
	defaultKubeconfig := ""
	if home != "" {
		defaultKubeconfig = home + "/.kube/config"
	}

	// PersistentFlags означает, что этот флаг будет доступен для команд scan и fix
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", defaultKubeconfig, "путь к файлу конфигурации Kubernetes")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


