import React from 'react'
import Sidebar from '../components/sidebar'
import Chat from '../components/chat'


const home : React.FC = () => {
    const dummyData = ["apakah ini?", "apakaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?", "apakah itu?"]
  return (
    <div className='flex flex-row w-full p-0 m-0 min-h-full'>
        <Sidebar className='bg-dark-grey basis-64 min-h-screen p-3 max-h-screen overflow-y-auto' questions={dummyData}/>
        <Chat className='bg-grey-chat grow min-h-screen flex-col justify-center relative max-h-screen overflow-y-auto'/>
    </div>
  )
}

export default home