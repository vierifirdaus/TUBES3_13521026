import React,{useState} from 'react'
import Sidebar from '../components/sidebar'
import Chat from '../components/chat'
import { history, message, } from '../interface'


const home : React.FC = () => {
    const [clickSide, setClickSide] = useState<number>(-1)
    const [value, setValue] = useState<string>('1');
    const [history, setHistories] = useState<history[]>([]);
    const [clicked, setClicked] = useState<number>(-1);


  return (
    <div className='flex flex-row w-full p-0 m-0 min-h-screen'>
        <Sidebar className='bg-dark-grey basis-1/5 pl-3 pr-6 min-h-screen  max-h-screen flex flex-col py-3' setClickSide={setClickSide} value={value} setValue={setValue} setHistories={setHistories} history={history} setClicked={setClicked} clicked={clicked}/>
        <Chat className='bg-grey-chat basis-4/5 flex-col justify-center relative flex max-h-screen overflow-y-auto' clickSide={clickSide}  setHistories={setHistories} history={history} setClicked={setClicked}/>
    </div>
  )
}

export default home