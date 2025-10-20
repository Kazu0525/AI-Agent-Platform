import { useEffect, useState } from 'react'
import axios from 'axios'
import './App.css'
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
  const API_BASE_URL = 'http://localhost:8080/api/v1'
