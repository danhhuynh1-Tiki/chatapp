// Should use redux old version
import { createStore } from "redux";

let roomID = {id : undefined}

const SetRoomID = (state,action) => {
    state.id = action.room_id
    return state
}
const RoomID = createStore(SetRoomID,roomID)
// RoomID.dispatch(add)
export {RoomID}


