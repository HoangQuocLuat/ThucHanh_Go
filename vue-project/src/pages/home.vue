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
  var base64Url = token.split('.')[1];
  var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
  var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function (c) {
    return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
  }).join(''));

  return JSON.parse(jsonPayload);
}
export default {
  data() {
    return {
      user : {}
    }
  },
  created() {
    let token = getCookie("token");
    if (!token) {
      this.$router.push("/login");
      return
    }

    let payload = (parseJwt(token))

    fetch(`http://127.0.0.1:8888/user/infor/${payload.id}`, {
      headers: { 'Authorization': 'Bearer ' + token }
    }).then(res => res.json()).then(
      res => {
        console.log(res)
        this.user = res
      })
  },
  methods: {
    logout() {
        setCookie("token", getCookie("token"), 0) 
        this.$router.push('/login')
    }
  }
 
};
</script>
<template>
  <div id="login-form">
    <h1 class="signin-form-header">Home</h1>
    <div id="login-body">
      <h3 style="margin-bottom: 20px;">ID: <span style="color:red;">{{ user.id }}</span></h3>
      <h3 style="margin-bottom: 20px;">Name: <span style="color:rgb(214, 90, 90);">{{ user.fullname }}</span></h3>
      
    </div>
    <button class="signin-button" @click="logout"><font-awesome-icon :icon ="['fas', 'right-from-bracket']"/></button>
  </div>
</template>
