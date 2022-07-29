import React from 'react';
import { Row, Col, Avatar,Button} from 'antd';
import { UserOutlined } from '@ant-design/icons';
import { RoomID } from '../../redux/RoomRedux';

const styleGroup ={
    marginTop : '5px',
    marginBottom : '5px',
    borderRadius : '3px 3px'
}


const Groupmessage = (props) => {
    const ChatGroup = () => {
        RoomID.dispatch({type:'',room_id:props.group.room_id})
        console.log(RoomID.getState().id)
    }
    return (
            <>
             <Row style={styleGroup} className="GroupChat" onClick={ChatGroup}>
                <Col span={5}>
                {/* <Avatar src={<Image src={avatar} style={{ width: 32 }} />} /> */}
                <span>
                {/* <Badge dot> */}
                    <Avatar size={50} shape="square" icon={<UserOutlined />} />
                {/* </Badge> */}
                </span>
                </Col>
                <Col span={10}>
                   <h7>{props.group.name}</h7>
                </Col>
                <Col span={9}>
                <Button type="danger"></Button>
                </Col>
            </Row>
            </>

    )
}

export default Groupmessage;