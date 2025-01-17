
<script>
export default {
  data() {
    return {
      userId: localStorage.getItem("userId"),  
      userName: localStorage.getItem("name"), 
      userPhoto: "",
      successmsg: null,
      errormsg: null,
      loading: false,
    };
  },
/*
methods: { 
	async setMyUserName() {
		this.loading = true;
	try{
		let response = await this.$axios.put("/session/:userId/userName", {userId: this.userId, name: this.name});
    localStorage.setItem("userId", response.data.userId);
    localStorage.setItem("name", response.data.name);
		this.errormsg = null;
    this.loading = false;
	}	catch (e){ 		
				if (e.response && e.response.status === 400) {
            

					}else{
            this.errormsg = e.toString();						
					}
	}
	},
},
*/

};
</script>






<template>
  <div class="user-profile">
    <div class="user-details">
      <h1 class="h2">USER PROFILE</h1>
      <p class="h3">UserId: {{ userId }}</p>
      <p class="h3">UserName: {{ userName }}</p>
    </div>
    <div class="user-photo">
      <img :src="userPhoto" alt="User Photo" class="photo" />
    </div>
  </div>
</template>


<style scoped>



.user-photo {
  flex: 0;
  margin-left: 20px;
}

.photo {
  width: 150px;
  height: 150px;
  border-radius: 50%;
  border: 2px solid #ccc;
}

h2 {
  text-align: left;
  font-size: 24px;
}

h3 {
  text-align: left;
  font-size: 20px;
}

</style>
