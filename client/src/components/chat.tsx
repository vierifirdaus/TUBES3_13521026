import React,{useState,useEffect,useRef} from 'react'
import { useToast } from '@chakra-ui/react'
import Message from './message'
import SendMessage from './sendMessage'
import { message, chatProps } from '../interface'
import axios from 'axios'
import TypingAnimation from './typingMessage'


const Chat : React.FC<chatProps> = ({className,clickSide,setHistories,setClicked,count,setCount}) => {
    const toast = useToast()
    const refBottom = useRef<HTMLDivElement>(null)
    const [inputValue, setInputValue] = useState<string>('')
    const [chatLog, setChatLog] = useState<message[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(false)
    const [idHistori, setIdHistori] = useState<number>(-1)

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
                    setIdHistori(clickSide)
                })
                .catch((err) => {
                    console.log(err)
                })
        }
        console.log(clickSide)
        const messages:message[] = [
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

    const handleInput = async (e: React.FormEvent<HTMLFormElement> | React.KeyboardEvent<HTMLTextAreaElement>) => {
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
        console.log(clickSide)
        console.log(count)
        if(count == 0 && clickSide == -1){
            const res = await updateHistory()
            //
        }
        setInputValue('')
        setChatLog([...chatLog, {id: chatLog.length + 1, id_histori: 1, jenis: "input", isi: inputValue}])
        getBotResponse()
    }

    const updateHistory =async () => {
        axios.post('http://localhost:1234/histori',
                {
                    nama: inputValue
                })
                .then((res) => {
                    console.log("data")
                    console.log(res.data)
                    axios.get('http://localhost:1234/allhistori')
                    .then((res) => {
                        console.log("data")
                        console.log(res.data)
                        setHistories(res.data)
                        setClicked(res.data[res.data.length-1].id)
                        setCount(count+1)
                        setIdHistori(res.data[res.data.length-1].id)
                    })
                    .catch((err) => {
                        console.log(err)
                    })
                })
                .catch((err) => {
                    console.log(err)
                })
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