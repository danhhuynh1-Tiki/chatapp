import React from 'react';
import {Col,Row,Avatar,Button,Badge} from 'antd';
import { UserOutlined } from '@ant-design/icons';
import { CallCreateRoom } from '../../services/RoomService';
import { RoomID } from '../../redux/RoomRedux'
import { GetMessage} from '../../services/MessageService';
import { DataMessage } from '../../redux/MessageRedux';
const styleUserChat = {
    borderRadius : '3px 3px'
}
const UserChat = (props) => {
    let status = (props.user.status === 1 ) ? true : false 
    const CreateRoom = async () => {
        // console.log("user_id createroom",props.user.id)
        const response = await CallCreateRoom(props.user.id)
        // console.log("create room_id",response.room_id)
        RoomID.dispatch({type:'',room_id : response.room_id})
        const message = await GetMessage(response.room_id)
        // console.log("message in user chat",message.messages) 
        DataMessage.dispatch({type:'',room_id:response.room_id,messages : message.messages})
    }
    return (    
        <>
             <Row className="Userchat" style={styleUserChat} onClick={CreateRoom}>
                <Col span={5}>
                {/* <Avatar src={<Image src={avatar} style={{ width: 32 }} />} /> */}
                <span>
                <Badge dot={status}>
                    <Avatar size={50} shape="square" icon={<UserOutlined />} />
                </Badge>
                </span>
                </Col>
                <Col span={10}>
                   <h7>{props.user.email}</h7>
                </Col>
                <Col span={9}>
                <Button type="danger"></Button>
                </Col>
            </Row>
        </>
    )
}
export default UserChat;