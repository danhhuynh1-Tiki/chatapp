import Userchat from './user/Userchat';
import { Row , Col } from 'antd';
import React, {useState,useEffect} from  'react';
import UsersService from '../services/UsersService';

const Chatuser =  () => {
    const [users,setUsers] = useState([])

    useEffect( () => {
        const fetchData = async () => {
            const response = await UsersService()
            setUsers(response.data)
        }
        const interval = setInterval(fetchData(),1000)
        return () => clearInterval(interval);
    },[users])
    const listUserChat = users.map((user) => 
        <Userchat user={user} />
    )
    
    return (
        <>
            <Row style={{height:'70vh',overflow:'scroll'}}>
                <Col span={24} style={{textAlign:'center'}}>
                   {/* <Userchat/>
                   <Userchat/>
                   <Userchat/> */}
                   {listUserChat}
                </Col>
            </Row>
        </>
    )
}

export default Chatuser;