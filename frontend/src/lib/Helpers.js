import {useState} from 'react'

export const UseInputValue = (initValue) => {
    const [value, setValue] = useState(initValue)

    return {
        value,
        onChange: e => setValue(e.target.value)
    }
}

