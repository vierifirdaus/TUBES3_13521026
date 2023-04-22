import { extendTheme } from '@chakra-ui/react'
import { buttonTheme } from './button'
export const theme = extendTheme({
    components: {
        Button: buttonTheme
    }
})