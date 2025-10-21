package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Agent struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	SystemPrompt string `json:"system_prompt"`
}

func main() {
	// Ginのルーターを作成
	r := gin.Default()

	// CORS設定（Codespaces対応）
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:5174",
			"http://localhost:3000",
			"https://*.app.github.dev", // Codespaces用
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// GitHub Codespacesのドメインを許可
			return strings.Contains(origin, "github.dev") ||
				strings.Contains(origin, "localhost")
		},
	}))
	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "AI Agent Platform API is running!",
			"version": "1.0.0",
		})
	})

	// API v1 グループ
	v1 := r.Group("/api/v1")
	{
		// エージェント一覧取得
		v1.GET("/agents", func(c *gin.Context) {
			agents := []Agent{
				{
					ID:           "1",
					Name:         "カスタマーサポートAI",
					Description:  "顧客対応を支援するエージェント",
					SystemPrompt: "あなたは親切なカスタマーサポート担当です。",
				},
				{
					ID:           "2",
					Name:         "コーディングアシスタント",
					Description:  "プログラミングをサポートするエージェント",
					SystemPrompt: "あなたは経験豊富なソフトウェアエンジニアです。",
				},
			}
			c.JSON(http.StatusOK, gin.H{
				"agents": agents,
			})
		})

		// エージェント詳細取得
		v1.GET("/agents/:id", func(c *gin.Context) {
			id := c.Param("id")
			agent := Agent{
				ID:           id,
				Name:         "サンプルエージェント",
				Description:  "これはサンプルです",
				SystemPrompt: "あなたは役立つAIアシスタントです。",
			}
			c.JSON(http.StatusOK, agent)
		})

		// エージェント作成
		v1.POST("/agents", func(c *gin.Context) {
			var newAgent Agent
			if err := c.BindJSON(&newAgent); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			// TODO: データベースに保存
			c.JSON(http.StatusCreated, newAgent)
		})
	}

	// サーバー起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
