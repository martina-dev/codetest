import React from 'react'
import styled from 'styled-components';

const Wrapper = styled.div`
    padding: 0.5em;
    background: white;
    display: flex;
    justify-content: space-between;
`

const Title = styled.h2`
  font-size: 1.5em;
  text-align: center;
  color: palevioletred;
`;

const Kcpe = () => {

    return (
        <Wrapper>
            <Title> KCPE </Title>
        </Wrapper>
    )
}

export default Kcpe