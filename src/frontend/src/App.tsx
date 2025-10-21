import { useEffect, useState } from 'react'
import axios from 'axios'
import './App.css'
import ChatScreen from './ChatScreen'

interface Agent {
  id: string
  name: string
  description: string
  system_prompt: string
}

function App() {
  const [agents, setAgents] = useState<Agent[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')
  const [selectedAgent, setSelectedAgent] = useState<Agent | null>(null)

  const API_BASE_URL = 'https://silver-giggle-q7vrx5q55573549-8080.app.github.dev/api/v1'

  const fetchAgents = async () => {
    setLoading(true)
    setError('')
    try {
      const response = await axios.get(`${API_BASE_URL}/agents`)
      setAgents(response.data.agents)
    } catch (err) {
      setError('エージェントの取得に失敗しました')
      console.error(err)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchAgents()
  }, [])

  if (selectedAgent) {
    return (
      <ChatScreen 
        agent={selectedAgent} 
        onBack={() => setSelectedAgent(null)}
        apiBaseUrl={API_BASE_URL}
      />
    )
  }

  return (
    <div className="App">
      <header>
        <h1>🤖 AI Agent Platform</h1>
        <p>AIエージェントを作成・管理・実行</p>
      </header>

      <main>
        <div className="agents-section">
          <div className="section-header">
            <h2>エージェント一覧</h2>
            <button onClick={fetchAgents} disabled={loading}>
              {loading ? '読み込み中...' : '🔄 更新'}
            </button>
          </div>

          {error && <div className="error">{error}</div>}

          <div className="agents-grid">
            {agents.map((agent) => (
              <div key={agent.id} className="agent-card">
                <h3>{agent.name}</h3>
                <p>{agent.description}</p>
                <button 
                  className="btn-primary"
                  onClick={() => setSelectedAgent(agent)}
                >
                  チャット開始
                </button>
              </div>
            ))}
          </div>

          {agents.length === 0 && !loading && !error && (
            <div className="empty-state">
              <p>エージェントがまだありません</p>
            </div>
          )}
        </div>
      </main>
    </div>
  )
}

export default App