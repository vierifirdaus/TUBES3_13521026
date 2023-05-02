import React,{useRef} from 'react'
import { FormControl,Input,IconButton,Image, Textarea } from '@chakra-ui/react'
import { sendMessageProps } from 'src/interface'


const SendMessage:React.FC<sendMessageProps> = ({inputValue,setInputValue,handleInput}) =>{
  const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      handleInput(e);
    }
  };

  return (
    <div className="px-28 py-5 bg-gradient-to-t from-[#40414f] m-0">
      <form onSubmit={handleInput} >
        <FormControl className='flex flex-row gap-3 bg-[#40414f] py-1 pl-5 pr-1 rounded-md'>
            <Textarea onKeyDown={handleKeyDown} rows={1} className="mt-2 h-10"autoComplete="off" placeholder="Type your message here" variant='outline' value={inputValue} onChange={(e)=>setInputValue(e.target.value)}/>
            <IconButton aria-label='Search database' icon={<Image src='/send.png' boxSize="20px"/>}  className="mt-2" bg='#40414f'  _hover={{ bg: '#202123' }} type='submit'/>
        </FormControl>
      </form>
    </div>
  )
}

export default SendMessage