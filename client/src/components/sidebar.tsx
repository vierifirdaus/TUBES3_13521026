import React,{useState,useEffect} from 'react';
import { 
    Button,
    Radio, 
    RadioGroup,
    Stack,
    Spinner
} from '@chakra-ui/react';
import {
    AddIcon
} from '@chakra-ui/icons';
import SideButton from './sideButton';
import { history, sidebarProps } from '../interface';
import axios from 'axios';

const Sidebar: React.FC<sidebarProps> = ({ className,setClickSide,value,setValue}) => {
  const [history, setHistories] = useState<history[]>([]);
  const [clicked, setClicked] = useState<number>(-1);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const handleClick = (i:number) => {
    setClicked(i);
    setClickSide(i);
  }

  const handleDelete = (i : number) => {
    /* sementara  nanti disini update history */
    setHistories(history.filter((history) => history.id !== i));
  }
  useEffect(() => {
    setIsLoading(true);
    axios.get('http://localhost:1234/allhistori').then((res) => {
      setIsLoading(false);
      setHistories(res.data);
    }).catch((err) => {
        console.log(err);
        setIsLoading(false);
    })
    // const histories: history[] = [{id:1,nama:"siapaaaaaaaa kamu??"},{id:2,nama:"siapaaaaaaaa kamu??"},{id:3,nama:"siapaaaaaaaa kamu??"},{id:4,nama:"siapaaaaaaaa kamu??"},{id:5,nama:"siapaaaaaaaa kamu??"},{id:6,nama:"siapaaaaaaaa kamu??"},{id:7,nama:"siapaaaaaaaa kamu??"},{id:8,nama:"siapaaaaaaaa kamu??"},{id:9,nama:"siapaaaaaaaa kamu??"},{id:10,nama:"siapaaaaaaaa kamu??"},{id:11,nama:"siapaaaaaaaa kamu??"},{id:12,nama:"siapaaaaaaaa kamu??"},{id:13,nama:"siapaaaaaaaa kamu??"},{id:14,nama:"siapaaaaaaaa kamu??"},{id:15,nama:"siapaaaaaaaa kamu??"},{id:16,nama:"siapaaaaaaaa kamu??"},{id:17,nama:"siapaaaaaaaa kamu??"},{id:18,nama:"siapaaaaaaaa kamu??"},{id:19,nama:"siapaaaaaaaa kamu??"},{id:20,nama:"siapaaaaaaaa kamu??"}]
    // setHistories(histories);
    return () => {
      // cleanup
    }
  }, []);
  return (
      <div className={className}>
        <Button className='w-full' size="lg" m="1" variant="sideButtonAdd" leftIcon={<AddIcon/>} justifyContent="flex-center" onClick={()=>handleClick(-1)}>New Chat</Button>
        <div className='flex-col max-h-screen overflow-y-auto'>
                <div>
                  {!isLoading && history.map((i:history) => (
                      <SideButton key={i.id} history={i} clicked={clicked} handleClick={handleClick} handleDelete={handleDelete}/>
                  ))}
                </div>
        </div>
        <span className='flex-auto flex justify-center items-center'>
          {isLoading && <Spinner color="white" />}
        </span>
        <RadioGroup className='pt-6 mx-6 border-t mt-2 mb-10' onChange={setValue} value={value} sx={{color: 'white'}}>
          <Stack direction='column'>
            <Radio value='1'>KMP</Radio>
            <Radio value='2'>BM</Radio>
          </Stack>
        </RadioGroup>
      </div>
  );
}

export default Sidebar;
