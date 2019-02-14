import React from 'react'
import styled from 'styled-components';
import { UseInputValue } from './lib/Helpers'
import { authenticate } from './lib/Auth'

const Wrapper = styled.form`
    padding: 0.5em;
    background: white;
    display: grid;
    grid-grap: 3px;
    justify-content: center;
`

const Input = styled.input`
  font-size: 1.5em;
  text-align: center;
  color: palevioletred;
`;


const Login = () => {
    
    const email = UseInputValue('')
    const password = UseInputValue('')

    const handleSubmit = () => {
        const formData = new FormData()

        formData.append('email', email.value)
        formData.append('password', password.value)

        authenticate(formData)

    }

    return (
        <Wrapper onSubmit={e => {
            e.preventDefault()
            handleSubmit()
        }}> 
            <Input {...email} type="text" placeholder="Email"/> 
            <Input {...password} placeholder="*********" type="password"/> 

            <button type="submit"> Submit</button>
        </Wrapper>
    )
}

export default Login