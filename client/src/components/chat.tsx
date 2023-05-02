import React,{useState,useEffect,useRef} from 'react'
import { useToast } from '@chakra-ui/react'
import Message from './message'
import SendMessage from './sendMessage'
import { message, chatProps } from '../interface'
import axios from 'axios'
import TypingAnimation from './typingMessage'


const Chat : React.FC<chatProps> = ({className,clickSide,setHistories,setClicked,history}) => {
    const toast = useToast()
    const refBottom = useRef<HTMLDivElement>(null)
    const [inputValue, setInputValue] = useState<string>('')
    const [chatLog, setChatLog] = useState<message[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(false)
    const [count,setCount] = useState<number>(0)

    useEffect(() => {
        if (clickSide !== -1) {
            axios.get('http://localhost:1234/chat',
                {
                    params: {
                        id_histori: clickSide
                        },
                    
                })
                .then((res) => {
                    console.log("data")
                    console.log(res.data)
                })
                .catch((err) => {
                    console.log(err)
                })
        }
        console.log(clickSide)
        const messages:message[] = [
            {
              id : 1,
              id_histori : 2,
              jenis : "input",
              isi : "siapa kamu??"
           },
            // {
            //   id : 2,
            //   id_histori : 1,
            //   Jenis : "output",
            //   Isi : "saya bot"
            // },
            // {
            //   id : 3,
            //   id_histori : 1,
            //   Jenis : "output",
            //   Isi : "saya bot"
            // },
            // {
            //   id : 4,
            //   id_histori : 1,
            //   Jenis : "output",
            //   Isi : "saya bot"
            // },
            // {
            //   id : 5,
            //   id_histori : 1,
            //   Jenis : "output",
            //   Isi : "saya bot"
            // },
            // {
            //   id : 6,
            //   id_histori : 1,
            //   Jenis : "output",
            //   Isi : "saya bot"
            // },
            // {
            //   id : 7,
            //   id_histori : 1,
            //   Jenis : "output",
            //   Isi : "saya bot"
            // },
            // {
            //   id : 8,
            //   id_histori : 1,
            //   Jenis : "output",
            //   Isi : "saya bot"
            // },
            // {
            //   id : 9,
            //   id_histori : 1,
            //   Jenis : "output",
            //   Isi : "saya bot"
            // },
            // {
            //   id : 10,
            //   id_histori : 1,
            //   Jenis : "output",
            //   Isi : "saya bot"
            // },
            // {
            //   id : 11,
            //   id_histori : 1,
            //   Jenis : "output",
            //   Isi : "In the example above, the sendIcon.svg image is imported and passed as a component to the icon prop of the IconButton component, making it the icon for the button. Please make sure that the sendIcon.svg image file is located in the correct path and that it's properly imported into your component. You can adjust the styling and positioning of the IconButton component and the Input component using the respective Chakra UI props and Tailwind CSS classes to achieve the desired look for your send message input component."
            // },
            // {
            //   id : 12,
            //   id_histori : 1,
            //   Jenis : "output",
            //   Isi : "saya bot"
            // }
        ]
        setChatLog(messages)
    }, [clickSide])

    useEffect(() => {
        refBottom.current?.scrollIntoView({ behavior: 'smooth' });
      }, [chatLog]);

    const handleInput = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        console.log("click")
        if( inputValue.trim() === ''){
            toast({
                title: 'Error Sending Message',
                description: "Message can't be empty",
                status: 'error',
                duration: 2500,
                isClosable: true,
              })
            return
        }

        if(count == 0 && clickSide == -1){
            setHistories([{id: 5, nama: inputValue}, ...history])
            setClicked(5)
            setCount(count+1)
        }
        setInputValue('')
        setChatLog([...chatLog, {id: chatLog.length + 1, id_histori: 1, jenis: "input", isi: inputValue}])
        getBotResponse()
    }
    const getBotResponse = async () => {
        setIsLoading(true)
        const response = await setTimeout(() => {
            setChatLog(prevChatLog => [...prevChatLog, {id: prevChatLog.length + 1, id_histori: 1, jenis: "output", isi: "tes response"}])
            setIsLoading(false)
        }, 1000)
    }

  return (
    <div className={className}>
        <div className='max-h-screen overflow-y-auto'>
        {chatLog.map((message) => (
            <Message key={message.id} message={message}/>
            ))
        }
        {isLoading && <TypingAnimation/>}
        <div ref={refBottom}></div>
        </div>
        <span className='flex-auto'></span>
        <SendMessage inputValue={inputValue} setInputValue={setInputValue} handleInput={handleInput}/>
    </div>
  )
}

export default Chat