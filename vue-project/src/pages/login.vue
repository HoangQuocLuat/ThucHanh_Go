<script>
function setCookie(cname, cvalue, exdays) {
  const d = new Date();
  d.setTime(d.getTime() + exdays * 24 * 60 * 60 * 1000);
  let expires = "expires=" + d.toUTCString();
  document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}
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
export default {
  data() {
    return {
      username: "",
      password: "",
      showPassword: false,
    };
  },
  methods: {
    togglePasswordVisibility() {
      this.showPassword = !this.showPassword;
    },
    login() {
      fetch("http://127.0.0.1:8888/user/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username: this.username,
          password: this.password,
        }),
      })
        .then((resp) => resp.json())
        .then((resp) => {
          console.log(resp);
          if (resp.StatusCode != 200) {
            alert("Thông tin tài khoản không chính xác ");
            return;
          }
          setCookie("token", resp.Data.token, 1);
          this.$router.push("/mess");
        });
    },
  },
  created() {
    console.log(getCookie("token"));
    let token = getCookie("token");
    if (token) {
      this.$router.push("/home");
    }
  },
};
</script>

<template>
  <div id="home">
    <div id="login-form">
      <div class="login-form-header">
        <span class="title-header">Đăng nhập</span>
        <span style="position: absolute; top: 19px; right: 19px">X</span>
      </div>
      <div id="login-body">
        <div class="label-input">Tên đăng nhập</div>
        <input type="text" v-model="username" />
        <div
          style="
            margin-top: 40px;
            display: flex;
            justify-content: space-between;
            width: 440px;
          "
        >
          <span class="label-input">Mật khẩu</span
          ><span
            ><a href="google.com" id="forgot-password">Quên mật khẩu?</a></span
          >
        </div>
        <div style="position: relative">
          <input :type="showPassword ? 'text' : 'password'" v-model="password" />
          <font-awesome-icon
            @click="togglePasswordVisibility"
            style="position: absolute; right: 2px; top: 32%"
            :icon="showPassword ? 'fa-solid fa-eye' : 'fa-solid fa-eye-slash'"
          />
        </div>
        <button class="login-button" @click="login">Đăng nhập</button>
        <p class="login-with">hoặc đăng nhập bằng</p>
        <div
          style="
            display: flex;
            justify-content: space-around;
            padding-top: 10px;
          "
        >
          <span class="otherway">
            <font-awesome-icon
              style="margin-right: 5px"
              icon="fa-brands fa-facebook-f"
            />
            Facebook
          </span>
          <span class="otherway">
            <font-awesome-icon
              style="margin-right: 5px"
              icon="fa-brands fa-google"
            />
            Google</span
          >
        </div>
        <p class="footer">
          Bạn chưa có tài khoản?
          <router-link to="/register">Đăng ký ngay!</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<style>
#home {
  background: #f0f2f5;
  display: flex;
  height: 100vh;
  justify-content: center;
  align-items: center;
}

#login-form {
  margin-top: 30px;
  width: 600px;
  height: 589px;
  border-radius: 14px;
  background: var(--light-greyscale-greyscale-200, #fff);
  box-shadow: 0px 12px 40px 0px rgba(0, 0, 0, 0.16);
}

.login-form-header {
  position: relative;
  max-height: 58.264px;
  border-bottom: 1px solid #f7f7f7;
  padding: 16px 240px 18.26px 240px;
}

.title-header {
  padding-top: 16px;
  color: var(--light-greyscale-greyscale-900, #000);
  font-family: Roboto;
  font-size: 18px;
  font-style: normal;
  font-weight: 500;
  line-height: 24px; /* 133.333% */
  text-transform: uppercase;
}

#login-body {
  padding: 36.74px 80px 0px 80px;
}

.label-input {
  color: var(--light-greyscale-greyscale-900, #000);
  font-family: Roboto;
  font-size: 16px;
  font-style: normal;
  font-weight: 500;
  line-height: 22px; /* 137.5% */
}

#login-body input {
  display: flex;
  width: 416px;
  height: 22px;
  padding: 16px 15px;
  align-items: flex-start;
  gap: 10px;
  flex-shrink: 0;
  border-radius: 8px;
  background: #f7f7f7;
  border: 0px;
}

#forgot-password {
  color: var(--dark-other-link-500, #2f80ed);
  text-align: right;
  font-family: Roboto;
  font-size: 16px;
  font-style: normal;
  font-weight: 400;
  line-height: 22px; /* 137.5% */
}

.login-button {
  margin-top: 40px;
  display: flex;
  width: 440px;
  height: 45px;
  padding: 15px 15px;
  justify-content: center;
  align-items: center;
  gap: 10px;
  flex-shrink: 0;
  border-radius: 40px;
  background: var(--Green, #00bf6f);
  box-shadow: 0px 2px 10px 0px rgba(244, 103, 0, 0.05);

  color: var(--light-text-active, #fff);
  text-align: center;
  /* GG/16px/Med/Button Text */
  font-family: Roboto;
  font-size: 16px;
  font-style: normal;
  font-weight: 500;
  line-height: normal;
}
.login-button:hover {
  background: var(--Green, #7e3d3d);
}

.login-with {
  margin-top: 30px;
  color: var(--light-transparent-greyscale-65, rgba(0, 0, 0, 0.65));
  text-align: center;
  font-family: Roboto;
  font-size: 16px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  letter-spacing: 0.192px;
}

.otherway {
  flex-shrink: 0;
  border-radius: 24px;
  background: var(--light-greyscale-greyscale-300, #e5e6ec);

  color: var(--light-greyscale-greyscale-900, #000);
  font-family: Roboto;
  font-size: 16px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  text-align: center;
  padding: 13.02px 55px 13px 55px;
}

.footer {
  margin-top: 51px;
  text-align: center;
}
</style>