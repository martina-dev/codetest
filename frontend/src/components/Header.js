import React, { useEffect, useState } from 'react'
import styled from 'styled-components';
import { Link } from 'react-router-dom';
import { getRole } from '../lib/Auth'

const Wrapper = styled.div`
    padding: 0.5em;
    background: white;
    display: flex;
    justify-content: space-between;
    border: 1px solid #DDD;
    margin: 10px;
`

const Title = styled.h2`
  font-size: 1.5em;
  text-align: center;
  color: palevioletred;
`;

const Header = () => {

    const [role, setRole] = useState()

    useEffect(() => {
        setRole(getRole())
    })
    
    if (role === "student") {
        return (
            <Wrapper>
                    <Title>
                    <Link to="/student" >Home </Link>
                    </Title>
                    <Title > 
                        <Link to="/subjects" >Subjects </Link>
                    </Title>
                    <Title>  
                        <Link to="/kcpe" >KCPE </Link>
                    </Title>
                    <Title>  
                        <Link to="/kcpe" >{role} </Link>
                    </Title>
            </Wrapper>
        )
    }
    return (
        <Wrapper>
                <Title>
                <Link to="/student" >Home </Link>
                </Title>
        </Wrapper>
    )
    
}

export default Header