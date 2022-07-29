import React,{useEffect, useState} from 'react';
import { Row,Col,Input,Button } from 'antd';
import Message from './chat/Message';
import Usermessage from './chat/Usermessage';
import { DataMessage } from '../redux/MessageRedux';
import { AddMessage,GetMessage } from '../services/MessageService';
import { RoomID } from '../redux/RoomRedux';
// import { EmailUser } from '../redux/UserRedux';
import { useNavigate } from 'react-router-dom';
import { EmailUser } from '../redux/UserRedux';
import { useInterval } from 'react-use';


const message = () =>{

    const [content, setContent] = useState("")
    
    const [messages,SetMessage] = useState([])
    
    const navigate = useNavigate()
    
    const getContent = (e) => {
        setContent(e.target.value)
    }
    // Send Message with post and update in redux
    const SendMessage = async () =>{
        console.log("add mesage room_id",RoomID.getState().id)
        console.log("content",content)
        const response = await AddMessage(RoomID.getState().id,content)
        console.log("add mesasge res",response)
        // DataMessage.dispatch({type:'',room_id :response.room_id,messages : response.messages})
    }
    const SetSizeMessage =  () => {

        
    }
    const featchData = async () => {
        const response = await  GetMessage(RoomID.getState().id)
        if(response === undefined){

        } else{
            // console.log(response)
            SetMessage(response.messages)
        }
    }
    useInterval(featchData,1000)
    let listMessage = messages.map((m) => {
        
        if(m.email == EmailUser.getState().email){
            return <Usermessage message={m} />
        }else{
            return <Message message={m} />
        }
    })

    return (
        <>
                <Row style={{textAlign : 'center'}}>
                    <Col span={24}>
                        <Button onClick={SetSizeMessage} >Load More</Button>
                    </Col>
                </Row>
                <Row style={{height : '85vh',overflow:'scroll'}}>
                    <Col span={24}>
                           {listMessage}
                    </Col>
                </Row>

                
                <Row style={{width : '100%',textAlign:'center'}}>
                    <Col span={20}>
                    <Input placeholder="Message" type="text" onChange={getContent} style={{height:'50px'}}/>
                    </Col>
                    <Col span={4}>
                        <Button type="primary" style={{height:'50px'}} onClick={SendMessage}>Send</Button>
                    </Col>
                </Row>
        </>
    )
}

export default message;