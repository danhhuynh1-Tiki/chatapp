import axios from "axios";

const CallCreateRoom = async (id) => {
    try{
        const response = await axios.get(`http://localhost:8080/api/room/${id}`)
        return response.data
    }catch(error){
        console.log(error)
    }
}
export { CallCreateRoom }