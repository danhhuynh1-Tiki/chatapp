import axios from "axios"

const CallUsersApi = async () => {
    try{

        const response = await axios.get("http://localhost:8080/mychat/v1/users")
        return response.data
    }catch(error){
        alert(error)
    }
}

export default CallUsersApi;