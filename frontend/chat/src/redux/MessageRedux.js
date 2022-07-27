import {createStore} from 'redux'
let message = {
    room_id : "",
    messages : []    
}

const SetDataMessage = (state,action) => {
    state.room_id = action.room_id
    state.messages = action.messages
    return state
}

const DataMessage = createStore(SetDataMessage,message)

export { DataMessage }