import React from 'react'
import Sidebar from '../components/sidebar'
import Chat from '../components/chat'


const home : React.FC = () => {
    const dummyData:string[] = ["siapa kamu??"]
  return (
    <div className='flex flex-row w-full p-0 m-0 min-h-full'>
        <Sidebar className='bg-dark-grey basis-64 p-3 max-h-screen overflow-y-auto' questions={dummyData}/>
        <Chat className='bg-grey-chat grow min-h-screen flex-col justify-center relative max-h-screen overflow-y-auto'/>
    </div>
  )
}

export default home