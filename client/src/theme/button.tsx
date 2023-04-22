import { defineStyle, defineStyleConfig } from '@chakra-ui/react'

const sideButtonHover= defineStyle({
    color:'white',
    background:'#202123',
    fontWeight:'normal',
    fontSize:'15px',
    _hover:{
        background:'#343541',
    }
})

const sideButtonAdd = defineStyle({
    color:'white',
    background:'#202123',
    fontWeight:'normal',
    fontSize:'15px',
    border:'1px solid #4d4d4f',
    _hover:{
        background:'#343541'
    }
  });
  

export const buttonTheme = defineStyleConfig({
    variants: {
        sideButtonHover,
        sideButtonAdd
    }
})
