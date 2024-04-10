<script>
export default {
  data() {
    return {
      idroom: "",
      nameroom: "",
      inforRoom: false,
    };
  },
  methods: {
    roomform() {
      this.inforRoom = !this.inforRoom;
    },
    createRoom() {
      fetch("127.0.0.1:8888/ws/createRoom", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id: this.idroom,
          name: this.nameroom,
        }),
      })
        .then((resp) => resp.json())
        .then((resp) => {
          console.log(resp);
        });
    },
  },
};
</script>

<template>
  <div class="row" id="chat-room">
    <div class="chat-1 col-3">
      <div class="chat-1_head">
        <div class="chat_user"></div>
        <div class="chat_logout"></div>
      </div>
      <div class="chat_add-room">
        <button style="border: none" @click="roomform">
          <font-awesome-icon :icon="['far', 'square-plus']" />Thêm phòng
        </button>
      </div>
      <div v-if="inforRoom" class="inforRoom">
        <h3>Điền thông tin phòng</h3>
        <form action="" style="text-align: center;">
          <label for="">Id phòng: </label>
          <br />
          <input
            type="text"
            style="
              background-color: rgb(55, 153, 153);
              border-bottom: 1px solid black;
              border-left: none;
              border-right: none;
              border-top: none;
              margin-left: 10px;
            "
          />
          <br />
          <label for="">Tên phòng: </label>
          <br />
          <input
            type="text"
            style="
              background-color: rgb(55, 153, 153);
              border-bottom: 1px solid black;
              border-left: none;
              border-right: none;
              border-top: none;
              margin-left: 10px;
            "
          />
          <div style="margin-top: 10px; text-align: center">
            <button>Tạo phòng</button>
          </div>
        </form>
      </div>
      <br />
      <div class="chat_list-room">
        <a-collapse>
          <a-collapse-panel key="1" header="Danh sách phòng">
            <a href="">Room 1</a>
          </a-collapse-panel>
        </a-collapse>
      </div>
    </div>
    <div class="chat-2 col-9">
      <div class="chat-2_head">
        <div class="chat_name"></div>
        <div class="chat_add-member"></div>
        <div class="chat_member"></div>
      </div>
    </div>
  </div>
</template>
<style>
.chat-room {
  width: 80vw;
  height: 100%;
}
.inforRoom {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: rgb(55, 153, 153);
  border-radius: 10px;
  padding: 20px;
}

.chat-1 {
  background-color: rgb(78, 144, 144);
  height: 98vh;
}
.chat-2 {
  background-color: whitesmoke;
}
</style>