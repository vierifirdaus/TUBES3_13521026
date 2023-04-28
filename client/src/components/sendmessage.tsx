import React from 'react'
import { FormControl,Input,IconButton,Image } from '@chakra-ui/react'
import { sendMessageProps } from 'src/interface'


const SendMessage:React.FC<sendMessageProps> = ({inputValue,setInputValue,handleInput}) =>{
  return (
    <div className="px-28 py-5 bg-gradient-to-t from-[#40414f]">
      <form onSubmit={handleInput}>
        <FormControl className='flex flex-row gap-3 bg-[#40414f] py-1 pl-5 pr-1 rounded-md'>
            <Input   autoComplete="off" placeholder="Type your message here" variant='unstyled' value={inputValue} onChange={(e)=>setInputValue(e.target.value)}/>
            <IconButton aria-label='Search database' icon={<Image src='/send.png' boxSize="20px"/>}  bg='#40414f'  _hover={{ bg: '#202123' }} type='submit'/>
        </FormControl>
      </form>
    </div>
  )
}

export default SendMessage