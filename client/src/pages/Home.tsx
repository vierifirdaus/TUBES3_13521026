import React from 'react'
import Sidebar from '../components/sidebar'
import Chat from '../components/chat'


const home : React.FC = () => {
    const dummyData:string[] = ["siapa kamu??"]
  return (
    <div className='flex flex-row w-full p-0 m-0 min-h-screen'>
        <Sidebar className='bg-dark-grey basis-1/5 p-3 max-h-screen overflow-y-auto' questions={dummyData}/>
        <Chat className='bg-grey-chat basis-4/5 flex-col justify-center relative max-h-screen overflow-y-auto'/>
    </div>
  )
}

export default home