import React, { useEffect, useState } from 'react';
import {Row,Col,Avatar,Button, Space} from 'antd';
import { UserOutlined } from '@ant-design/icons';
import { LogoutUsersApi,GetUserApi} from '../../services/UsersService';
import { useNavigate } from 'react-router-dom';
const User = () =>{
    const [user,userData] = useState({})
    let navigate = useNavigate()
    useEffect( () => {
            const fetchData = async () => {
                const response = await GetUserApi()
                console.log("fetchData" , response)
                if(response === undefined){
                    navigate("/login")
                }else{
                    userData(response)
                }
            }
            fetchData()
            
    },user)

    return (                                        
        <>
            <Row className="User" style={{marginTop : '10px'}}>
                <Col span={1}></Col>
                <Col span={5}>
                {/* <Avatar src={<Image src={avatar} style={{ width: 32 }} />} /> */}
                <Avatar shape="square" size={50} icon={<UserOutlined />} />
                </Col>
                <Col span={10}>
                    <Space align="center">
                        <p style={{size:'1px'}}>{user.email}</p>
                   </Space>
                </Col>
                <Col span={8}>
                <Button type="danger" onClick={LogoutUsersApi}>Logout</Button>
                </Col>
            </Row>
        </>
    )
}

export default User;