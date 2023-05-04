import React from 'react';
import { Image, Text } from '@chakra-ui/react';
import { messageProps } from '../interface';

const Message: React.FC<messageProps> = ({ message }) => {
    // Split the text into an array of substrings based on the line break character
const messageLines = message.isi.split('\n');


// Render the Text components inside the div element
return (
  <div className={message.jenis === 'input' ? 'flex gap-x-5 bg-[#434654] max-w-screen justify-start px-36 py-5' : 'flex bg-[#343541] max-w-screen gap-x-5 justify-start px-36 py-5'}>
    <Image src={message.jenis === 'input' ? '/user.png' : '/logo.png'} alt="user" boxSize="30px" />
    <div className='flex flex-col'>
        {messageLines.map((line, index) => (
            <Text key={index} color='#ffffff' fontSize='md' fontWeight='semi-bold'>
                {line}
            </Text>
        ))}
    </div>
  </div>
    );
}

export default Message;
