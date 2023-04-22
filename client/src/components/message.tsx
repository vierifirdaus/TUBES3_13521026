import React from 'react';
import { Image, Text } from '@chakra-ui/react';

interface messageProps {
    message: {
        sender: string,
        text: string
    }
}

const Message: React.FC<messageProps> = ({ message }) => {
    return (
        <div className={message.sender === 'user' ? 'flex gap-x-5 bg-[#434654] max-w-screen justify-start px-36 py-5' : 'flex bg-[#343541] max-w-screen gap-x-5 justify-start px-36 py-5'}>
                <Image src={message.sender === 'user' ? '/user.png' : '/logo.png'} alt="user" boxSize="30px"/>
                <Text className='text-white '>{message.text}</Text>
        </div>
    );
}

export default Message;
