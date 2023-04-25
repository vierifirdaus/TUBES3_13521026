import React,{useState,useEffect} from 'react';
import { 
    Button
} from '@chakra-ui/react';
import {
    AddIcon
} from '@chakra-ui/icons';
import SideButton from './sideButton';
import { history, sidebarProps } from '../interface';

const Sidebar: React.FC<sidebarProps> = ({ className,setClickSide}) => {
  const [history, setHistories] = useState<history[]>([]);
  const [clicked, setClicked] = useState<number>(-1);
  const handleClick = (i:number) => {
    setClicked(i);
    setClickSide(i);
  }

  const handleDelete = (i : number) => {
    /* sementara */
    setHistories(history.filter((history) => history.id !== i));
  }
  useEffect(() => {
    const histories: history[] = [{id:1,nama:"siapaaaaaaaa kamu??"},{id:2,nama:"kamu??"},{id:3,nama:"siapa kamu??"}]
    setHistories(histories);
  }, []);
  return (
    <div className={className}>
        <div className='flex flex-col'>
                <Button size="lg" m="1" variant="sideButtonAdd" leftIcon={<AddIcon/>} justifyContent="flex-center">New Chat</Button>
                {history.map((i:history) => (
                    <SideButton key={i.id} history={i} clicked={clicked} handleClick={handleClick} handleDelete={handleDelete}/>
                ))}
        </div>
    </div>
  );
}

export default Sidebar;
