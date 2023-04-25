import React,{useState,useEffect} from 'react';
import { 
    Button
} from '@chakra-ui/react';
import {
    AddIcon
} from '@chakra-ui/icons';
import SideButton from './sideButton';
import { history, sidebarProps } from '../interface';

const Sidebar: React.FC<sidebarProps> = ({ className,histories }) => {
  const [history, setHistories] = useState<history[]>([]);
  const [clicked, setClicked] = useState<number>(-1);
  useEffect(() => {
    setHistories(histories);
  }, []);
  const handleClick = (i:number) => {
    setClicked(i);
  }

  const handleDelete = (i : number) => {
    /* sementara */
    setHistories(history.filter((history) => history.id !== i));
  }
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
