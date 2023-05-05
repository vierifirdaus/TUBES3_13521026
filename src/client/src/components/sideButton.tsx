import React,{useEffect, useState} from 'react'
import { Button,ButtonGroup,IconButton,Input,Box } from '@chakra-ui/react'
import { EditIcon, DeleteIcon, ChatIcon,CheckIcon,CloseIcon } from '@chakra-ui/icons'
import {buttonProps} from '../interface'
import axios from 'axios'


const SideButton :React.FC<buttonProps> = ({history,clicked,handleClick,handleDelete,setHistories}) => {
  const [editing, setEditing] = useState(false);
  const [editedName, setEditedName] = useState<string>('')
  const [name, setName] = useState('');

  useEffect(()=>{
    setName(history.nama)
    return () => {
      setEditing(false);
    }
  },[clicked])


  const handleEditButtonClick = () => {
    // Set editing state to true when "Edit name" button is clicked
    // Set editedName state to current history.nama value
    // setName(history.nama);
    setEditedName(history.nama)
    setEditing(true);
  };

  const handleSaveButtonClick = () => {
    // Perform save logic here, e.g., update history.nama with editedName
    // // After save, set editing state to false to hide the text field
    // setName(history.nama);
    // console.log('history.nama changed to: ', history.nama)
    axios.put('http://localhost:1234/histori',
    {
      new_name: editedName,  
      id_histori: history.id
    }).then((res) => {
      console.log(res.data);
      setName(editedName);
      setEditing(false);
      axios.get('http://localhost:1234/allhistori').then((res) => {
        setHistories(res.data);
      }
      ).catch((err) => {
        console.log(err);
      }
      )
    }).catch((err) => {
      console.log(err);
    })
  };
      

  const handleCancelButtonClick = () => {
    setEditing(false);
  };

  return clicked === history.id ? 
  (
      <>
      {editing ? (
        <>
          <Box className='w-full flex flex-row bg-[#40414f] rounded items-center min-w-full px-3 my-2'>
            <ChatIcon color={"white"}/>
            <Input
              size="sm"
              value={editedName}
              onChange={(e) => setEditedName(e.target.value)}
              placeholder="Edit name"
              className='ml-2 rounded-lg'
            />
            <IconButton
              size="lg"
              aria-label="Save"
              variant="sideButtonClick"
              m={0}
              p={0}
              icon={<CheckIcon color="#bcbcc9" />}
              onClick={handleSaveButtonClick}
            />
            <IconButton
              size="lg"
              aria-label="Cancel"
              variant="sideButtonClick"
              m={0}
              p={0}
              icon={<CloseIcon color="#bcbcc9" />}
              onClick={handleCancelButtonClick}
            />
          </Box>
        </>
      ) : (
          <ButtonGroup className="w-full my-2" isAttached >
            <Button className="w-full" size="lg" m={0} variant="sideButtonClick" leftIcon={<ChatIcon />} justifyContent="flex-start">
              {name.length > 10 ? name.substring(0, 10) + "..." : name}
            </Button>
            <IconButton
              size="lg"
              aria-label="Edit Name"
              icon={<EditIcon color="#bcbcc9" _hover={{ color: "white" }} />}
              onClick={handleEditButtonClick}
              variant="sideButtonClick"
              bg="#40414f"
              p={0}
              m={0}
            />
            <IconButton
              size="lg"
              aria-label="Delete Chat"
              icon={<DeleteIcon color="#bcbcc9" _hover={{ color: "white" }} />}
              onClick={() => handleDelete(history.id)}
              variant="sideButtonClick"
              bg="#40414f"
              p={0}
              m={0}
            />
          </ButtonGroup>
      )}
      </>
  )
  :
  (
    <Button className="w-full my-2" size="lg"  variant="sideButtonHover" leftIcon={<ChatIcon/>} justifyContent="flex-start" onClick={() => handleClick(history.id)}>
                        {
                          name.length > 18 ? name.substring(0,18) + "..." : name
                        }
    </Button>
  )
}

export default SideButton