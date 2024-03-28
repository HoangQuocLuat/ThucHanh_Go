

<template>
  <div class="row" id="chat-room">
    <div class="chat-1 col-3">
      <div class="chat-1_head">
        <div class="chat_user"></div>
        <div class="chat_logout"></div>
      </div>
      <div class="chat_add-room">
        <button style="border: none;"><font-awesome-icon :icon="['far', 'square-plus']" />Thêm phòng</button>
      </div>
      <br>
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
.chat-1 {
  background-color: rgb(78, 144, 144);
  height: 98vh;
}
.chat-2 {
  background-color: whitesmoke;
}
</style>