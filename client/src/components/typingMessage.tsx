import React from 'react';
import { Image } from '@chakra-ui/react';
import { BeatLoader } from 'react-spinners';

const TypingAnimation: React.FC = () => {
    return (
        <div className={'flex bg-[#343541] max-w-screen gap-x-5 justify-start px-36 py-5'}>
                <Image src={ '/logo.png'} alt="user" boxSize="30px"/>
                <div className='flex items-center space-x-2'>
                    <BeatLoader color='#ffffff' size={10} loading={true} />
                </div>
        </div>
    );
}

export default TypingAnimation;