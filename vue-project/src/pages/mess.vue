<script>
function getCookie(cname) {
  let name = cname + "=";
  let ca = document.cookie.split(";");
  for (let i = 0; i < ca.length; i++) {
    let c = ca[i];
    while (c.charAt(0) == " ") {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}
function setCookie(cname, cvalue, exdays) {
  const d = new Date();
  d.setTime(d.getTime() + exdays * 24 * 60 * 60 * 1000);
  let expires = "expires=" + d.toUTCString();
  document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}
function parseJwt(token) {
  var base64Url = token.split(".")[1];
  var base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
  var jsonPayload = decodeURIComponent(
    window
      .atob(base64)
      .split("")
      .map(function (c) {
        return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
      })
      .join("")
  );

  return JSON.parse(jsonPayload);
}
export default {
  data() {
    return {
      idroom: "",
      nameroom: "",
      inforRoom: false,
      user: {},
      dataRoom: [],
      ws: null,
    };
  },
  created() {
    let token = getCookie("token");
    if (!token) {
      this.$router.push("/login");
      return;
    }

    let payload = parseJwt(token);

    fetch(`http://127.0.0.1:8888/user/infor/${payload.id}`, {
      headers: { Authorization: "Bearer " + token },
    })
      .then((res) => res.json())
      .then((res) => {
        console.log(res);
        this.user = res;
      });
  },

  methods: {
    roomform() {
      this.inforRoom = !this.inforRoom;
    },
    createRoom(event) {
      event.preventDefault(); // tránh trường hợp load lại toàn bộ trang
      fetch("http://127.0.0.1:8888/ws/createRoom", {
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
          if (resp.StatusCode === 200) {
            alert(resp.Message);
            this.dataRoom.push({
              idroom: this.idroom,
              nameroom: this.nameroom,
            });
          }
          console.log(resp);
          return resp;
        });
    },

    selecRoom() {
      fetch("http://127.0.0.1:8888/ws/getRooms")
        .then((res) => res.json())
        .then((res) => {
          this.dataRoom = res.Data;
        });
    },

    join(idroom, iduser, username) {
      console.log("Bắt đầu kết nối websocket");
      this.ws = new WebSocket(
        `ws://127.0.0.1:8888/ws/joinRoom/${idroom}?userId=${iduser}&username=${username}`
      );
      this.ws.addEventListener("open", function () {
        console.log(`${idroom}, ${iduser}, ${username}`);
        console.log("Kết nối WebSocket thành công");
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
      <div class="chat_add-room" style="display: flex; justify-content: center">
        <button style="border: none; margin-top: 20px" @click="roomform">
          <font-awesome-icon :icon="['far', 'square-plus']" />
          Thêm phòng
        </button>
      </div>
      <div v-if="inforRoom" class="inforRoom" style="color: azure">
        <h3>Điền thông tin phòng</h3>
        <form action="" style="text-align: center">
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
            v-model="idroom"
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
            v-model="nameroom"
          />
          <div style="margin-top: 10px; text-align: center">
            <button @click="createRoom($event)">Tạo phòng</button>
          </div>
        </form>
      </div>
      <br />
      <div class="chat_list-room">
        <a-collapse>
          <a-collapse-panel key="1" header="DANH SÁCH PHÒNG" @click="selecRoom">
            <button
              v-for="(room, index) in dataRoom"
              @click="join(room.id, 2)"
              :key="index"
              style="
                display: block;
                border: none;
                margin: 10px;
                background-color: rgb(78, 144, 144);
                border-radius: 10px;
                width: 100%;
                font-size: 16px;
                color: azure;
              "
            >
              {{ room.id }} - {{ room.name }}
            </button>
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