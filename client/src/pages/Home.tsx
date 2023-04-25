import React from 'react'
import Sidebar from '../components/sidebar'
import Chat from '../components/chat'
import { history, message, } from '../interface'


const home : React.FC = () => {
    const dummyData: history[] = [{id:1,nama:"siapaaaaaaaa kamu??"},{id:2,nama:"kamu??"},{id:3,nama:"siapa kamu??"}]
    const messages: message[]= [
      {
        id : 1,
        id_histori : 2,
        Jenis : "input",
        Isi : "siapa kamu??"
      },
      {
        id : 2,
        id_histori : 1,
        Jenis : "output",
        Isi : "saya bot"
      },
      {
        id : 3,
        id_histori : 1,
        Jenis : "output",
        Isi : "saya bot"
      },
      {
        id : 4,
        id_histori : 1,
        Jenis : "output",
        Isi : "saya bot"
      },
      {
        id : 5,
        id_histori : 1,
        Jenis : "output",
        Isi : "saya bot"
      },
      {
        id : 6,
        id_histori : 1,
        Jenis : "output",
        Isi : "saya bot"
      },
      {
        id : 7,
        id_histori : 1,
        Jenis : "output",
        Isi : "saya bot"
      },
      {
        id : 8,
        id_histori : 1,
        Jenis : "output",
        Isi : "saya bot"
      },
      {
        id : 9,
        id_histori : 1,
        Jenis : "output",
        Isi : "saya bot"
      },
      {
        id : 10,
        id_histori : 1,
        Jenis : "output",
        Isi : "saya bot"
      },
      {
        id : 11,
        id_histori : 1,
        Jenis : "output",
        Isi : "In the example above, the sendIcon.svg image is imported and passed as a component to the icon prop of the IconButton component, making it the icon for the button. Please make sure that the sendIcon.svg image file is located in the correct path and that it's properly imported into your component. You can adjust the styling and positioning of the IconButton component and the Input component using the respective Chakra UI props and Tailwind CSS classes to achieve the desired look for your send message input component."
      },
      {
        id : 12,
        id_histori : 1,
        Jenis : "output",
        Isi : "saya bot"
      }
  ]
  return (
    <div className='flex flex-row w-full p-0 m-0 min-h-screen'>
        <Sidebar className='bg-dark-grey basis-1/5 p-3 max-h-screen overflow-y-auto' histories={dummyData}/>
        <Chat className='bg-grey-chat basis-4/5 flex-col justify-center relative max-h-screen overflow-y-auto' messages={messages}/>
    </div>
  )
}

export default home