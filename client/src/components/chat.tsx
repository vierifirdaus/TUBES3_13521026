import React from 'react'
import Message from './message'
import SendMessage from './sendmessage'

interface chatProps {
    className: string;
}

const messages = [
    {"sender": "user", "text": "Hello"},
    {"sender":"bot","text":"oke ada yang bisa dibantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
    {"sender": "user", "text": "Hello"},
    {"sender": "bot", "text": "Hello"},
    {"sender":"bot","text":"ada yang bisa saya bantu"},
    {"sender": "user", "text": "Hello"},
]

const Chat : React.FC<chatProps> = ({className}) => {
  return (
    <div className={className}>
        {messages.map((message) => (
            <Message message={message}/>
        ))}
        <SendMessage />
    </div>
  )
}

export default Chat