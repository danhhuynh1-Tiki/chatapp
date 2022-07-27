import Userchat from './user/Userchat';
import { Row , Col } from 'antd';
import React, {useState,useEffect} from  'react';
import { CallUsersApi } from '../services/UsersService';
import { useNavigate } from 'react-router-dom';
const Chatuser =  () => {
    const [users,setUsers] = useState([])
    let navigate = useNavigate()
    useEffect( () => {
        const fetchData = async () => {
            const response = await CallUsersApi()
            // console.log("user chat user",response)
            if (response === undefined){
                navigate("/login")
            }else{
                // console.log(response)
                setUsers(response.data)
            }
        }
        let interval = setInterval(fetchData(),5000)
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