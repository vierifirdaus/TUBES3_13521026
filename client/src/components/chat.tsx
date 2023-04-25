import React,{useState,useEffect,useRef} from 'react'
import { useToast } from '@chakra-ui/react'
import Message from './message'
import SendMessage from './sendmessage'
import { message, chatProps } from '../interface'
import TypingAnimation from './typingMessage'


const Chat : React.FC<chatProps> = ({className,messages}) => {
    const toast = useToast()
    const refBottom = useRef<HTMLDivElement>(null)
    const [inputValue, setInputValue] = useState<string>('')
    const [chatLog, setChatLog] = useState<message[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(true)

    useEffect(() => {
        setInputValue('')
        setIsLoading(false)
        setChatLog(messages)
    }, [])

    useEffect(() => {
        refBottom.current?.scrollIntoView({ behavior: 'smooth' });
      }, [chatLog]);

    const handleInput = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        console.log("click")
        if(inputValue === ''){
            toast({
                title: 'Error Sending Message',
                description: "Message can't be empty",
                status: 'error',
                duration: 2500,
                isClosable: true,
              })
            return
        }
        setChatLog([...chatLog, {id: chatLog.length + 1, id_histori: 1, Jenis: "input", Isi: inputValue}])
        setInputValue('')
        getBotResponse()
    }
    
    const getBotResponse = async () => {
        setIsLoading(true)
        const response = await setTimeout(() => {
            setChatLog(prevChatLog => [...prevChatLog, {id: prevChatLog.length + 1, id_histori: 1, Jenis: "output", Isi: "tes response"}])
            setIsLoading(false)
        }, 8000)
    }

  return (
    <div className={className}>
        <div className='min-h-full'>
        {chatLog.map((message) => (
            <Message key={message.id} message={message}/>
            ))
        }
        {isLoading && <TypingAnimation/>}
        <span ref={refBottom}></span>
        </div>
        <SendMessage inputValue={inputValue} setInputValue={setInputValue} handleInput={handleInput}/>
    </div>
  )
}

export default Chat