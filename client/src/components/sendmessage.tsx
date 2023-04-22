import React from 'react'
import { FormControl,Input,IconButton,Text,Image} from '@chakra-ui/react'


const sendmessage:React.FC=()=>{
  return (
    <div className='sticky bottom-0 left-0 right-0 px-28 py-5 bg-gradient-to-t from-[#40414f]'>
        <FormControl className='flex flex-row gap-3 bg-[#40414f] py-1 pl-5 pr-1 rounded-md'>
            <Input placeholder="Type your message here" variant='unstyled'/>
            <IconButton aria-label='Search database' icon={<Image src='/send.png' boxSize="20px"/>}  bg='#40414f'  _hover={{ bg: '#202123' }}/>
        </FormControl>
    </div>
  )
}

export default sendmessage