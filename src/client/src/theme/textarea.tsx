import { defineStyle, defineStyleConfig } from '@chakra-ui/react'

const outline = defineStyle({
    color: 'white',
    bg: 'transparent',
    border: '1px solid #ffffff',
    height: '1',
})


const sm = defineStyle({
    fontSize: 'sm',
    px: '6',
    h: '8',
    borderRadius: 'md',
  })
export const textareaTheme = defineStyleConfig({
  variants: { outline },
  sizes: { sm },
})