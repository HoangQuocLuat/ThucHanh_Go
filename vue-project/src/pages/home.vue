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
export default {
  created() {
    let token = getCookie("token");
    if (!token) {
      this.$router.push("/login");
    }
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
    <h1>Home</h1>
    <font-awesome-icon :icon ="['fas', 'right-from-bracket']" @click="logout"/>
  </div>
</template>
