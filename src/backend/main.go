package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

type Agent struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	SystemPrompt string `json:"system_prompt"`
}

type ChatRequest struct {
	AgentID string `json:"agent_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}

var openaiClient *openai.Client

func main() {
	// .envファイルを読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// OpenAI クライアントを初期化
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}
	openaiClient = openai.NewClient(apiKey)
	log.Println("✅ OpenAI client initialized")

	r := gin.Default()

	// CORS設定
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
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
		v1.GET("/agents", getAgents)
		v1.GET("/agents/:id", getAgent)
		v1.POST("/agents", createAgent)
		v1.POST("/chat", chatWithAgent)
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

func getAgents(c *gin.Context) {
	agents := []Agent{
		{
			ID:           "1",
			Name:         "カスタマーサポートAI",
			Description:  "顧客対応を支援するエージェント",
			SystemPrompt: "あなたは親切で丁寧なカスタマーサポート担当です。お客様の質問に対して、分かりやすく、親身になって回答してください。",
		},
		{
			ID:           "2",
			Name:         "コーディングアシスタント",
			Description:  "プログラミングをサポートするエージェント",
			SystemPrompt: "あなたは経験豊富なソフトウェアエンジニアです。コードの質問に対して、具体的で実践的なアドバイスを提供してください。",
		},
		{
			ID:           "3",
			Name:         "営業サポートAI",
			Description:  "セールス活動を支援するエージェント",
			SystemPrompt: "あなたは優秀な営業担当です。商品の魅力を伝え、顧客のニーズに合った提案をしてください。",
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"agents": agents,
	})
}

func getAgent(c *gin.Context) {
	id := c.Param("id")
	agents := []Agent{
		{
			ID:           "1",
			Name:         "カスタマーサポートAI",
			Description:  "顧客対応を支援するエージェント",
			SystemPrompt: "あなたは親切で丁寧なカスタマーサポート担当です。お客様の質問に対して、分かりやすく、親身になって回答してください。",
		},
		{
			ID:           "2",
			Name:         "コーディングアシスタント",
			Description:  "プログラミングをサポートするエージェント",
			SystemPrompt: "あなたは経験豊富なソフトウェアエンジニアです。コードの質問に対して、具体的で実践的なアドバイスを提供してください。",
		},
		{
			ID:           "3",
			Name:         "営業サポートAI",
			Description:  "セールス活動を支援するエージェント",
			SystemPrompt: "あなたは優秀な営業担当です。商品の魅力を伝え、顧客のニーズに合った提案をしてください。",
		},
	}

	for _, agent := range agents {
		if agent.ID == id {
			c.JSON(http.StatusOK, agent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
}

func createAgent(c *gin.Context) {
	var newAgent Agent
	if err := c.BindJSON(&newAgent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newAgent)
}

func chatWithAgent(c *gin.Context) {
	var req ChatRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("💬 Chat request - Agent: %s, Message: %s", req.AgentID, req.Message)

	// エージェントのシステムプロンプトを取得
	agents := []Agent{
		{
			ID:           "1",
			SystemPrompt: "あなたは親切で丁寧なカスタマーサポート担当です。お客様の質問に対して、分かりやすく、親身になって回答してください。",
		},
		{
			ID:           "2",
			SystemPrompt: "あなたは経験豊富なソフトウェアエンジニアです。コードの質問に対して、具体的で実践的なアドバイスを提供してください。",
		},
		{
			ID:           "3",
			SystemPrompt: "あなたは優秀な営業担当です。商品の魅力を伝え、顧客のニーズに合った提案をしてください。",
		},
	}

	var systemPrompt string
	for _, agent := range agents {
		if agent.ID == req.AgentID {
			systemPrompt = agent.SystemPrompt
			break
		}
	}

	if systemPrompt == "" {
		systemPrompt = "あなたは役立つAIアシスタントです。"
	}

	// OpenAI APIを呼び出し
	resp, err := openaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: req.Message,
				},
			},
			Temperature: 0.7,
			MaxTokens:   500,
		},
	)

	if err != nil {
		log.Printf("❌ OpenAI API error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response from AI"})
		return
	}

	reply := resp.Choices[0].Message.Content
	log.Printf("✅ AI reply: %s", reply)

	c.JSON(http.StatusOK, ChatResponse{
		Reply: reply,
	})
}
