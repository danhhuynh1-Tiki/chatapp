import React,{useRef, useState} from 'react';
import { Row, Col, Avatar,Button,Modal,Input, List} from 'antd';
import { UserOutlined } from '@ant-design/icons';
import { RoomID } from '../../redux/RoomRedux';
import { AddMember, CallRemoveMember, GetMembers } from '../../services/RoomService';

const styleGroup ={
    marginTop : '5px',
    marginBottom : '5px',
    borderRadius : '3px 3px'
}


const Groupmessage = (props) => {
    const [listEmail,SetEmail] = useState([])
    const [email,SetEmailUser] = useState("")
    const RemoveMember = async (email) =>{
        const response = await CallRemoveMember(props.group.room_id,email)
        console.log(response)
       setIsModalVisible(false)
    }
    const getEmail = (e) =>{
        SetEmailUser(e.target.value)
    }
    const [isModalVisible, setIsModalVisible] = useState(false);
    const showModal = async () => {
        const response = await GetMembers(props.group.room_id)
        console.log(response)
        SetEmail(response)
        setIsModalVisible(true);
    };
    const handleOk = async () => {
        const response = await AddMember(props.group.room_id,email)
        if(response === undefined){
            alert("Cannot add members")
        }else{
            console.log(response)
        }
        setIsModalVisible(false);
    };
  
    const handleCancel = () => {

        setIsModalVisible(false);
    };
    const ChatGroup = () => {
        RoomID.dispatch({type:'',room_id:props.group.room_id})
        console.log(RoomID.getState().id)
    }
    if(props.group.key == 1){
    return (
            <>
            <Row>
                <Col span={20}>
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
                    <Button type="primary"></Button>
                    </Col>
                </Row>
            </Col>
            <Col span={4}>
                <Row>
                    <Col span={24}><Button type="primary" onClick={showModal}>+</Button></Col>
                </Row>
                <Modal title="Add Members" visible={isModalVisible} onOk={handleOk} onCancel={handleCancel}>
                    <Input placeholder="Email" onChange={getEmail} />
                    <div style={{height:200,overflow:'scroll'}}>
                    <List
                        itemLayout="horizontal"
                        dataSource={listEmail}
                        renderItem={(item) => (
                        <List.Item>
                            <List.Item.Meta
                            avatar={<Avatar src="https://joeschmoe.io/api/v1/random" />}
                            title={<a href="https://ant.design">{item}</a>}
                            />
                            <Button type="danger" onClick={event => RemoveMember(item)}>x</Button>
                        </List.Item>
                        )}
                    />
                    </div>
                </Modal>
            </Col>
            </Row>
            </>

    )
    }
    else{
        return (
            <>
            <Row>
                <Col span={20}>
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
                    <Button type="primary"></Button>
                    </Col>
                </Row>
            </Col>
            <Col span={4}>
                <Row>
                    <Col span={24}><Button type="primary" onClick={showModal}>I</Button></Col>
                </Row>
                <Modal title="Add Members" visible={isModalVisible} onOk={handleOk} onCancel={handleCancel}>
                    {/* <Input placeholder="Email" onChange={getEmail} /> */}
                    <div style={{height:200,overflow:'scroll'}}>
                    <List
                        itemLayout="horizontal"
                        dataSource={listEmail}
                        renderItem={(item) => (
                        <List.Item>
                            <List.Item.Meta
                            avatar={<Avatar src="https://joeschmoe.io/api/v1/random" />}
                            title={<a href="https://ant.design">{item}</a>}
                            />
                            {/* <Button type="danger" onClick={event => RemoveMember(item)}>x</Button> */}
                        </List.Item>
                        )}
                    />
                    </div>
                </Modal>
            </Col>
            </Row>
            </>

    )
    }
}

export default Groupmessage;