import React from 'react'

interface chatProps {
    className: string;
}


const Chat : React.FC<chatProps> = ({className}) => {
  return (
    <div className={className}>Chat</div>
  )
}

export default Chat