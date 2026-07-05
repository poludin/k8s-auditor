package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Сканирование кластера на соответствие Золотому стандарту",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Инициализация сканера...")
		fmt.Printf("Используем kubeconfig: %s\n", kubeconfig)

		// 1. Собираем конфигурацию из файла
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatalf("Ошибка чтения kubeconfig: %v", err)
		}

		// 2. Создаем клиента (Clientset) для общения с K8s API
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalf("Ошибка создания K8s клиента: %v", err)
		}

		// 3. Делаем тестовый запрос: получаем версию сервера
		serverVersion, err := clientset.Discovery().ServerVersion()
		if err != nil {
			log.Fatalf("Не удалось подключиться к кластеру: %v", err)
		}

		fmt.Printf("Успешное подключение! Версия Kubernetes: %s\n", serverVersion.String())
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}