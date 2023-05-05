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

const Sidebar: React.FC<sidebarProps> = ({ className,setClickSide,value,setValue,history,setHistories,setClicked,clicked,setCount,setChatLog,refHistori}) => {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const handleClick = (i:number) => {
    setClicked(i);
    setClickSide(i);
  }

  const handleDelete = (i : number) => {
    setChatLog([]);
    axios.delete('http://104.248.157.133:1234/histori',
    {
        params: {
            Id_histori: i
        }
      })
      .then((res) => {
        console.log(res.data);
        axios.get('http://104.248.157.133:1234/allhistori').then((res) => {
          setIsLoading(false);
          if(res.data == null){
            setClickSide(-1);
            setClicked(-1);
            setCount(0);
            setHistories([]);
          }else{
            setHistories(res.data);
          }

        }
        ).catch((err) => {
          console.log(err);
          setIsLoading(false);
        })
      })
      .catch((err) => {
        console.log(err);
      }
    )
    
  }
  useEffect(() => {
    setIsLoading(true);
    axios.get('http://104.248.157.133:1234/allhistori').then((res) => {
      setIsLoading(false);
      if(res.data == null){
        setClickSide(-1);
        setClicked(-1);
        setCount(0);
        setHistories([]);
      }else{
        setHistories(res.data);
      }
    }).catch((err) => {
        console.log(err);
        setIsLoading(false);
    })
    return () => {
      // cleanup
    }
  }, []);

  

  return (
      <div className={className}>
        <Button className='w-full p-3' size="lg" m="1" variant="sideButtonAdd" leftIcon={<AddIcon/>} justifyContent="flex-center" onClick={()=>{handleClick(-1);setCount(0);setChatLog([])}}>New Chat</Button>
        <div className='flex-col max-h-screen overflow-y-auto'>
                <div>
                  {!isLoading && history.map((i:history) => (
                      <SideButton key={i.id} history={i} clicked={clicked} handleClick={handleClick} handleDelete={handleDelete} setHistories={setHistories}/>
                  ))}
                </div>
                <div ref={refHistori}></div>
        </div>
        <span className='flex-auto flex justify-center items-center'>
          {isLoading && <Spinner color="white" />}
        </span>
        <RadioGroup defaultValue={value} className='pt-6 mx-6 border-t mt-2 mb-10' onChange={setValue} value={value} sx={{color: 'white'}}>
          <Stack direction='column'>
            <Radio value='1'>KMP</Radio>
            <Radio value='2'>BM</Radio>
          </Stack>
        </RadioGroup>
      </div>
  );
}

export default Sidebar;
