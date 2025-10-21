import { useState } from 'react'
import axios from 'axios'
import './ChatScreen.css'

interface Agent {
  id: string
  name: string
  description: string
  system_prompt: string
}

interface Message {
  role: 'user' | 'assistant'
  content: string
}

interface ChatScreenProps {
  agent: Agent
  onBack: () => void
  apiBaseUrl: string
}

function ChatScreen({ agent, onBack, apiBaseUrl }: ChatScreenProps) {
  const [messages, setMessages] = useState<Message[]>([])
  const [inputMessage, setInputMessage] = useState('')
  const [loading, setLoading] = useState(false)

  const sendMessage = async () => {
    if (!inputMessage.trim()) return

    const userMessage = inputMessage
    setInputMessage('')
    
    setMessages(prev => [...prev, { role: 'user', content: userMessage }])
    setLoading(true)

    try {
      const response = await axios.post(`${apiBaseUrl}/chat`, {
        agent_id: agent.id,
        message: userMessage
      })

      setMessages(prev => [...prev, { role: 'assistant', content: response.data.reply }])
    } catch (error) {
      console.error('チャットエラー:', error)
      setMessages(prev => [...prev, { 
        role: 'assistant', 
        content: 'エラーが発生しました。もう一度お試しください。' 
      }])
    } finally {
      setLoading(false)
    }
  }

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault()
      sendMessage()
    }
  }

  return (
    <div className="chat-screen">
      <div className="chat-header">
        <button onClick={onBack} className="back-button">← 戻る</button>
        <div className="chat-agent-info">
          <h2>{agent.name}</h2>
          <p>{agent.description}</p>
        </div>
      </div>

      <div className="chat-messages">
        {messages.length === 0 && (
          <div className="chat-welcome">
            <h3>👋 こんにちは！</h3>
            <p>何でも聞いてください</p>
          </div>
        )}
        {messages.map((msg, index) => (
          <div key={index} className={`message ${msg.role}`}>
            <div className="message-content">
              {msg.content}
            </div>
          </div>
        ))}
        {loading && (
          <div className="message assistant">
            <div className="message-content loading">
              <span className="dot">.</span>
              <span className="dot">.</span>
              <span className="dot">.</span>
            </div>
          </div>
        )}
      </div>

      <div className="chat-input-area">
        <input
          type="text"
          value={inputMessage}
          onChange={(e) => setInputMessage(e.target.value)}
          onKeyPress={handleKeyPress}
          placeholder="メッセージを入力..."
          disabled={loading}
          className="chat-input"
        />
        <button 
          onClick={sendMessage} 
          disabled={loading || !inputMessage.trim()}
          className="send-button"
        >
          送信
        </button>
      </div>
    </div>
  )
}

export default ChatScreen