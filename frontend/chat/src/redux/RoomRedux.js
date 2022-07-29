// Should use redux old version
import { createStore } from "redux";

let roomID = {id : undefined}
let checkGroup = {key : 1}
const SetRoomID = (state,action) => {
    state.id = action.room_id
    return state
}
const SetKey = (state,action) =>{
    state.key = action.key
    return state
}
const RoomID = createStore(SetRoomID,roomID)
const KeyRoom = createStore(SetKey,checkGroup)
// RoomID.dispatch(add)
export {RoomID,KeyRoom}


