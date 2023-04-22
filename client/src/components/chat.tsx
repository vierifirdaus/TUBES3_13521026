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
        text: "The error message you are encountering is a TypeScript compiler warning indicating that there are two files with similar names in your project, but with different casings. In this case, the file /Users/fajarherawan/Documents/TUBES3_STIMA/client/src/App.tsx is importing a file named ./pages/home but the actual file name on the file system is /Users/fajarherawan/Documents/TUBES3_STIMA/client/src/pages/Home.tsx, and the only difference is the casing of the filename This warning is raised because file systems in some operating systems (such as macOS and Windows) are case-insensitive, while TypeScript is case-sensitive. This can cause issues with module resolution in TypeScript, as the compiler may not be able to correctly resolve the correct file to import. To fix this issue, you can make sure that the casing of the file name in your import statement in /Users/fajarherawan/Documents/TUBES3_STIMA/client/src/App.tsx matches the actual file name on the file system. In this case, you can update the import statement to use the correct casing:"
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
        { messages.map((message) => (
            <Message message={message}/>
        ))}
        </div>
        <SendMessage/>
    </div>
  )
}

export default Chat