
import axios from "axios";

const AddMessage = async (room_id,content) =>{
    const data = JSON.stringify({content : content})
    console.log("data send message",room_id,content)
    try{
    
        const response = await axios.post(`http://localhost:8080/api/message/${room_id}`,data,{headers: { 'Content-Type': 'application/json' }})
        // console.log(response)
        return response.data
    }catch(error){
        console.log(error)
    }
}
const GetMessage = async (room_id) => {
    if(typeof(room_id) != 'undefined'){
        // console.log("get meesage",room_id)
        // console.log(`http://localhost:8080/api/message/${room_id}`)
        try{
            const response = await axios.get(`http://localhost:8080/api/message/${room_id}`)
            // console.log("message from api",response)
            return response.data
        }catch(error){
            console.log("loi ne",error)
        }
    }
    // }else{
    //     console.log("get message undefined",room_id)
    // }
}
export { AddMessage,GetMessage}