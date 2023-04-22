import React from 'react';
import { 
    Box,
    Button
} from '@chakra-ui/react'
import {
    ChatIcon,
    AddIcon
} from '@chakra-ui/icons'




interface SidebarProps {
  className: string,
  questions: string[]
}

const Sidebar: React.FC<SidebarProps> = ({ className,questions }) => {
  return (
    <div className={className}>
        <div className='flex flex-col'>
                <Button size="lg" m="1" variant="sideButtonAdd" leftIcon={<AddIcon/>} justifyContent="flex-start">New Chat</Button>
                {questions.map((question) => (
                    <Button size="lg" m="1" variant="sideButtonHover" leftIcon={<ChatIcon/>} justifyContent="flex-start">
                        <Box textAlign="left">
                        {
                        question.length > 20 ? question.substring(0,20) + "..." : question
                        }
                        </Box>
                    </Button>
                ))}
        </div>
    </div>
  );
}

export default Sidebar;
