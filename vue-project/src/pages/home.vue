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
    }

    let payload = (parseJwt(token))

    fetch(`http://127.0.0.1:8888/user/infor/${payload.id}`, {
      headers: { 'Authorization': 'Bearer ' + token }
    }).then(res => res.json()).then(
      res => {
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
  <div id="home">
    <h1 style="">Home</h1>
    <div>
      <p>ID: {{ user.id }}</p>
      <p>FullName: {{ user.fullname }}</p>
      <p>Email: {{ user.email }}</p>
    </div>
    <button style="height: 20px; width: 30px;"><font-awesome-icon :icon ="['fas', 'right-from-bracket']" @click="logout"/></button>
  </div>
</template>
