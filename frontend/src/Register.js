import React from 'react'
import styled from 'styled-components';
import { UseInputValue } from './lib/Helpers'
import { register } from './lib/Auth'

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


const Register = () => {
    
    const email = UseInputValue('')
    const password = UseInputValue('')
    const name = UseInputValue('')
    const role = UseInputValue('')

    const handleSubmit = () => {
        const formData = new FormData()

        formData.append('email', email.value)
        formData.append('password', password.value)
        formData.append('name', name.value)
        formData.append('role', role.value)

        register(formData)

    }

    return (
        <Wrapper onSubmit={e => {
            e.preventDefault()
            handleSubmit()
        }}>     
            <Input {...name} type="text" placeholder="Name"/> 
            <Input {...email} type="text" placeholder="Email"/> 
            <Input {...password} placeholder="*********" type="password"/> 
            <Input {...role} type="text" placeholder="Role"/> 

            <button type="submit"> Submit</button>
        </Wrapper>
    )
}

export default Register