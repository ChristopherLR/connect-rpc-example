import { useState } from 'react'
import './App.css'

import { createClient, ConnectError } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

import { GreetService } from "./gen/greet/v1/greet_connect";
import { GreetResponse } from "./gen/greet/v1/greet_pb";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

const client = createClient(GreetService, transport);

function App() {
  const [name, setName] = useState("");
  const [messages, setMessages] = useState<string[]>([]);

  const handleGreet = async () => {
    setMessages([]);
    try {
      const res: GreetResponse = await client.greet({ name });
      setMessages([res.greeting]);
    } catch (err) {
      console.error(err);
      const connectErr = ConnectError.from(err);
      setMessages((prev) => [...prev, `Error: ${connectErr.message}`]);
    }
  };

  const handleStream = async () => {
    setMessages([]);
    try {
      for await (const res of client.streamGreetings({ name })) {
        setMessages((prev) => [...prev, res.greeting]);
      }
    } catch (err) {
      console.error(err);
      const connectErr = ConnectError.from(err);
      setMessages((prev) => [...prev, `Error: ${connectErr.message}`]);
    }
  };

  return (
    <>
      <h1>Connect-RPC Example</h1>
      <div className="card">
        <input
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="Enter your name"
          style={{ padding: '8px', marginRight: '8px' }}
        />
        <div style={{ marginTop: '16px' }}>
          <button onClick={handleGreet}>
            Unary Greet
          </button>
          <button onClick={handleStream} style={{ marginLeft: '8px' }}>
            Stream Greetings
          </button>
        </div>
      </div>
      <div className="output">
        {messages.map((msg, i) => (
          <p key={i}>{msg}</p>
        ))}
      </div>
    </>
  )
}

export default App
