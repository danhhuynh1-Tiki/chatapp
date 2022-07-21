const CallStatusApi = async () =>{
    try{
        const response = await axios.Get("")
        console.log(response.data)
        return response.data
    }catch(error){
        console.log(error)
    }
}