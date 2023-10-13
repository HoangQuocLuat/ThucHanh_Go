<script>
export default {
  data() {
    return {
      fullname: "",
      name: "",
      password: "",
      email:"",
      confirmPassword: "",
      showPassword1: false,
      showPassword2: false,
    };
  },
  computed: {
    passwordsMatch() {
      return this.password = !this.confirmPassword;
    },
  },
  methods: {
    
    togglePasswordVisibility(fieldNumber) {
      if (fieldNumber === 1) {
        this.showPassword1 = !this.showPassword1;
      } else if (fieldNumber === 2) {
        this.showPassword2 = !this.showPassword2;
      }
    },
    register() {
      if (this.password !== this.confirmPassword) {
        alert("Mật khẩu và mật khẩu nhập lại không khớp.");
        return;
      }
      fetch("http://127.0.0.1:8888/user/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          fullname: this.fullname,
          name: this.name,
          hashpassword: this.password,
          email: this.email
        }),
      })
        .then((resp) => resp.json())
        .then((resp) => {
          console.log(resp);
          if (resp.result = 400) {
            alert("Tài khoản đã tồn tại");
            return
          }
          if (resp.result = 200) {
            alert("Đăng ký thành công");
            return;
          }
        });
    },
  },
};
</script>

<template>
  <div id="signin-form">
    <div class="signin-form-header">
      <span class="title-header">ĐĂNG KÝ TÀI KHOẢN MỚI MIỄN PHÍ</span>
    </div>
    <div id="signin-body">
      <div class="label-input_sig">Họ và tên</div>
      <input type="text" placeholder="Nhập họ và tên của bạn" v-model="fullname"/>

      <div class="label-input_sig">Tên đăng nhập</div>
      <input type="text" placeholder="Nhập tên đăng nhập của bạn" v-model="name"/>

      <div class="label-input_sig">Mật khẩu</div>
      <div class="input-pass" style="position: relative">
        <input
        :type="showPassword1 ? 'text' : 'password'"
          v-model="password"
          placeholder="* * * * * * * *"
        />
        <font-awesome-icon
          @click="togglePasswordVisibility(1)"
          style="position: absolute; right: 2px; top: 32%"
          :icon="showPassword1 ? 'fa-solid fa-eye' : 'fa-solid fa-eye-slash'"
        />
      </div>
      <div class="label-input_sig">Nhập lại mật khẩu</div>
      <div class="input-pass" style="position: relative">
        <input
          :type="showPassword2 ? 'text' : 'password'"
          v-model="confirmPassword"
          placeholder="* * * * * * * *"
        />
        <font-awesome-icon
          @click="togglePasswordVisibility(2)"
          style="position: absolute; right: 2px; top: 32%"
          :icon="showPassword2 ? 'fa-solid fa-eye' : 'fa-solid fa-eye-slash'"
        />
      </div>

      <div class="label-input_sig">Email</div>
      <input type="text" placeholder="Nhập email của bạn" v-model="email"/>
      <div style="display: flex; margin-top: 24px">
        <input style="margin-left: 0px; display: block" type="checkbox" />
        <label style="margin-left: 3px">
          Tôi đồng ý với các
          <router-link to="#">điều kiện và điều khoản</router-link>
        </label>
      </div>
      <button class="signin-button" @click="register">Đăng ký</button>
      <div style="border: 1px solid #EEE; margin-top: 30px; margin-bottom: 20px; position: relative;">
        <p class="signin-with" style="position: absolute; left: 50%; transform: translate(-50%, -50%); background-color: white;">Hoặc</p>
      </div>
      <div
        style="display: flex; justify-content: space-around; padding-top: 10px"
      >
        <button
          style="border-radius: 50%; height: 50px; width: 50px; border: none"
        >
          <font-awesome-icon
            style="height: 30px; width: 30px"
            icon="fa-brands fa-facebook"
          />
        </button>
        <button
          style="border-radius: 50%; height: 50px; width: 50px; border: none"
        >
          <font-awesome-icon
            style="height: 30px; width: 30px"
            icon="fa-brands fa-google-plus-g"
          />
        </button>
        <button
          style="border-radius: 50%; height: 50px; width: 50px; border: none"
        >
          <font-awesome-icon
            style="height: 30px; width: 30px"
            icon="fa-brands fa-apple"
          />
        </button>
      </div>
      <p class="footer">
        Bạn đã có tài khoản rồi?
        <router-link to="/login">Đăng nhập</router-link>
      </p>
    </div>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
}

body {
  background: #f0f2f5;
  display: flex;
  height: auto;
  justify-content: center;
  align-items: center;
}

#signin-form {
  width: 600px;
  height: auto;
  margin-top: 30px;
  margin-bottom: 63px;
  border-radius: 14px;
  background: var(--light-greyscale-greyscale-200, #fff);
  box-shadow: 0px 12px 40px 0px rgba(0, 0, 0, 0.16);
}

.signin-form-header {
  display: flex;
  justify-content: center;
  border-bottom: 1px solid #f7f7f7;
}

.title-header {
  padding-top: 16px;
  padding-bottom: 16px;
  color: var(--light-greyscale-greyscale-900, #000);
  font-family: Roboto;
  font-size: 18px;
  font-style: normal;
  font-weight: 500;
  line-height: 24px;
  /* 133.333% */
  text-transform: uppercase;
}

#signin-body {
  padding: 36.74px 80px 0px 80px;
}

.label-input_sig {
  margin-top: 24px;
  color: var(--light-greyscale-greyscale-900, #000);
  font-family: Roboto;
  font-size: 16px;
  font-style: normal;
  font-weight: 500;
  line-height: 22px;
  /* 137.5% */
}

#signin-body >input {
  display: flex;
  width: 440px;
  height: 22px;
  padding: 16px 15px;
  align-items: flex-start;
  gap: 10px;
  flex-shrink: 0;
  border-radius: 8px;
  background: #f7f7f7;
  border: 0px;
}
.input-pass >input {
  display: flex;
  width: 440px;
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
  line-height: 22px;
  /* 137.5% */
}

.signin-button {
  margin: 0 auto;
  margin-top: 40px;
  display: flex;
  width: 245px;
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
.signin-button:hover {
  background: var(--Green, #7e3d3d);
}

.signin-with {
  
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
  padding-bottom: 30px;
  margin-top: 51px;
  text-align: center;
}
</style>