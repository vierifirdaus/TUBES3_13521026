import { extendTheme } from '@chakra-ui/react'
import { buttonTheme } from './button'
import { inputTheme } from './Input'
import { textareaTheme } from './textarea'
export const theme = extendTheme({
    components: {
        Button: buttonTheme,
        Input: inputTheme,
        Textarea: textareaTheme,
    },
})