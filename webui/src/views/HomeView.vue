<template>
    <div class="login-container">
      <h1 class="text-center mb-4">Log into WasaText!</h1>
      <h2 class="text-center mb-4">
        Connect with your friends effortlessly using WASAText! Send and receive
        messages, whether one-on-one or in groups, all from the convenience of
        your PC. Enjoy seamless conversations with text or GIFs and easily stay
        in touch through your private chats or group discussions.
      </h2>
      <div class="mb-3">
        <input type="text" v-model="userId" class="form-control" placeholder="User ID">
        <input type="text" v-model="name" class="form-control" placeholder="Username">
      </div>
      <div class="mb-3">
        <button class="btn btn-success btn-block" @click="doLogin">OK</button>
      </div>
    </div>
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
  </template>
  
  <script>
  export default {
    data() {
      return {
        errormsg: null,
        name: "",
        userId: ""
      };
    },
    methods: {
      async doLogin() { 
          try {
            let response = await this.$axios.post("/session", { userId: this.userId, username: this.name });
            localStorage.setItem("token", this.userId);
            localStorage.setItem("name", this.name);
            this.$router.push({ path: '/session' });
          } catch (e) {
            if (e.response && e.response.status === 400) {
              this.errormsg =
                "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
            } else if (e.response && e.response.status === 500) {
              this.errormsg =
                "An internal error occurred. We will be notified. Please try again later.";
            } else {
              this.errormsg = e.toString();
            }
          }
      },
    },
  };
  </script>
  
  <style scoped>
  .login-container {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: flex-start; 
    justify-content: flex-start; 
    background-color: aliceblue;
    padding: 20px;
  }
  
  h1 {
    font-size: 1.5rem;
    margin-bottom: 1rem;
  }
  
  h2 {
    font-size: 1rem;
    margin-bottom: 1rem;
  }
  
  input {
    padding: 0.5rem;
    margin-bottom: 1rem;
    border: 1px solid #ccc;
    border-radius: 4px;
  }
  
  button {
    width: 300px; 
    padding: 0.5rem;
    background-color: #0d6efd;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    margin: 0 auto; 
    display: block; 
  }
  
  .error-message {
    margin-top: 1rem;
    color: #f44336;
  }
  
  </style> 