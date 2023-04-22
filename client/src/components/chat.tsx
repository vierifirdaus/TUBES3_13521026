import React from 'react'
import Message from './message'
import SendMessage from './sendmessage'

interface chatProps {
    className: string;
}

interface message {
    sender: string,
    text: string
}

const messages: message[]= [
    {
        sender: 'user',
        text: 'Hello'
    },
    {
        sender: 'bot',
        text: 'Hi'
    },
    {
        sender: 'user',
        text: 'How are you?'
    },
    {
        sender: 'bot',
        text: 'I am fine'
    },
    {
        sender: 'user',
        text: 'What are you doing?'
    },
    {
        sender: 'bot',
        text: 'I am chatting with you'
    },
    {
        sender: 'user',
        text: 'What is your name?'
    },
    {
        sender: 'bot',
        text: 'My name is Chatbot'
    },
    {
        sender: 'user',
        text: 'What is your age?'
    },
    {
        sender: 'bot',
        text: 'I am 1 year old'
    },
    {
        sender: 'user',
        text:'What is your favourite food?'
    },
    {
        sender: 'bot',
        text: 'I like to eat data'
    },
    {
        sender: 'user',
        text: 'What is your favourite color?'
    },
    {
        sender: 'bot',
        text: 'I like to be black'
    },
    {
        sender: 'user',
        text: 'What is your favourite movie?'
    },
    {
        sender: 'bot',
        text: 'I like to watch The Matrix'
    }
]

const Chat : React.FC<chatProps> = ({className}) => {
  return (
    <div className={className}>
        <div className='min-h-full'>
        { messages.length !== 0 && messages.map((message) => (
            <Message message={message}/>
        ))}
        </div>
        <SendMessage/>
    </div>
  )
}

export default Chat