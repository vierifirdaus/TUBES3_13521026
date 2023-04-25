import React from 'react'
import Sidebar from '../components/sidebar'
import Chat from '../components/chat'
import { history, message, } from '../interface'


const home : React.FC = () => {
    const [clickSide, setClickSide] = React.useState<number>(-1)
  return (
    <div className='flex flex-row w-full p-0 m-0 min-h-screen'>
        <Sidebar className='bg-dark-grey basis-1/5 p-3 max-h-screen overflow-y-auto' setClickSide={setClickSide}/>
        <Chat className='bg-grey-chat basis-4/5 flex-col justify-center relative max-h-screen overflow-y-auto' clickSide={clickSide}/>
    </div>
  )
}

export default home