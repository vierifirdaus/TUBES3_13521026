import React,{useState,useEffect,useRef} from 'react'
import { useToast } from '@chakra-ui/react'
import Message from './message'
import SendMessage from './sendmessage'
import { chatProps } from '../interface'
import axios from 'axios'
import TypingAnimation from './typingMessage'


const Chat : React.FC<chatProps> = ({className,clickSide,setHistories,setClicked,count,setCount,setChatLog,chatLog,refHistori,value}) => {
    const toast = useToast()
    const refBottom = useRef<HTMLDivElement>(null)
    const [inputValue, setInputValue] = useState<string>('')
    const [isLoading, setIsLoading] = useState<boolean>(false)
    const [idHistori, setIdHistori] = useState<number>(-1)

    useEffect(() => {
        if (clickSide !== -1) {
            setIdHistori(clickSide)
            axios.get('http://localhost:1234/histori',
                {
                    params: {
                        Id_histori: clickSide
                        },
                    
                })
                .then((res) => {
                    if(res.data.isi == null){
                        setChatLog([])
                    }else{
                        setChatLog(res.data.isi)
                    }
                })
                .catch((err) => {
                    console.log(err)
                })
        }else{
            setChatLog([])
        }
        refBottom.current?.scrollIntoView({ behavior: 'smooth' });
        return () => {
            setChatLog([])
        }
    }, [clickSide])

    useEffect(() => {
        refBottom.current?.scrollIntoView({ behavior: 'smooth' });
      }, [chatLog]);

    const handleInput = async (e: React.FormEvent<HTMLFormElement> | React.KeyboardEvent<HTMLTextAreaElement>) => {
        e.preventDefault()
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
        setInputValue('')
        setChatLog([...chatLog, {id: chatLog.length + 1, id_histori: 1, jenis: "input", isi: inputValue}])
        if(count == 0 && clickSide == -1){
            const res = await updateHistory()
        }else{
            getBotResponse(idHistori)
        }
    }

    const updateHistory =async () => {
        const post = await
        axios.post('http://localhost:1234/histori',
                {
                    nama: inputValue
                })
                .then((res) => {
                    axios.get('http://localhost:1234/allhistori')
                    .then((res) => {
                        setHistories(res.data)
                        setCount(count+1)
                        setClicked(res.data[res.data.length-1].id)
                        setIdHistori(res.data[res.data.length-1].id)
                        getBotResponse(res.data[res.data.length-1].id)
                        refHistori.current?.scrollIntoView({ behavior: 'smooth' });
                    })
                    .catch((err) => {
                        console.log(err)
                    })
                })
                .catch((err) => {
                    console.log(err)
                })
    }

    const getBotResponse = async (id:number) => {
        setIsLoading(true)
        axios.post('http://localhost:1234/find',{
            pertanyaan: inputValue,
            id_histori: id,
            jenis: value
        }).then((res) => {
            console.log(res.data)
            setChatLog(prevdata => [...prevdata, {id: prevdata.length + 1, id_histori: 1, jenis: "output", isi: res.data}])
            setIsLoading(false)
        }
        ).catch((err) => {
            console.log(err)
            setIsLoading(false)
        })
    }

  return (
    <div className={className}>
        <div className='max-h-screen overflow-y-auto'>
        {chatLog.map((message,index) => (
            <Message key={index} message={message}/>
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