import Groupchat from './user/Groupchat';
import React, {useRef, useState} from 'react';
import { Row, Col, Button,Modal,Input, message} from 'antd';
import { GetMessage } from '../services/MessageService';
import { RoomID } from '../redux/RoomRedux';
import { EmailUser } from '../redux/UserRedux';
import { useInterval } from 'react-use';
import { GetGroup,CallCreateGroup } from '../services/RoomService';



const Chatgroup = () => {
    

    const Name = useRef(null)
    const Memebers = useRef(null)

    const [isModalVisible, setIsModalVisible] = useState(false);
    const [name,setName] = useState("")
    const [email,setEmail] = useState("")
    const [group,setGroup] = useState([])
    const showModal = () => {
        setIsModalVisible(true);
    };
    const getName = (e) =>{
        setName(e.target.value)
    }
    const getEmail = (e) => {
        setEmail(e.target.value)
    }
    const handleOk = async () => {
        // console.log(name)
        // console.log(email+"," + EmailUser.getState().email)
        const response = await CallCreateGroup(name,EmailUser.getState().email,email+","+EmailUser.getState().email)
        if(response === undefined){
            alert("cannot create group")
        }else{

        }
        setIsModalVisible(false);
    };
  
    const handleCancel = () => {

        setIsModalVisible(false);
    };

    const fetchData = async () => {
        // console.log(EmailUser.getState().email)
        const response = await GetGroup(EmailUser.getState().email)
        if(response === undefined){

        }else{
            setGroup(response)
        }
        // console.log(group)
    }

    useInterval(fetchData,1000)
    let listGroup
    if(group != undefined){
        listGroup = group.map((g) => <Groupchat group={g} />)
    }
    return (
        <>
            <Row style={{textAlign:'center',marginBottom : '1px'}}>
                <Col span={24}>
                    <Button onClick={showModal}>+</Button>
                </Col>
            </Row>
            <Modal title="Create Group" visible={isModalVisible} onOk={handleOk} onCancel={handleCancel}>
                <Input placeholder="Name" onChange={getName}/>
                <Input placeholder="Memebers" onChange={getEmail}/>
            </Modal>
            <Row style={{ height : '65vh', overflow:'scroll'}}>
                <Col span={24} style={{textAlign:'center'}}>
                    {listGroup}
                </Col>
            </Row>
           
        </>
    )
}
export default Chatgroup;