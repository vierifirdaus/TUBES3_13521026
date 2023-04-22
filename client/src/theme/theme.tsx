import { extendTheme } from '@chakra-ui/react'
import { buttonTheme } from './button'
import { inputTheme } from './Input'
export const theme = extendTheme({
    components: {
        Button: buttonTheme,
        Input: inputTheme,
    },
})