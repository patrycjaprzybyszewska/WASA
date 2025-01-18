
<script>
export default {
  data() {
    return {
      userId: localStorage.getItem("userId"),  
      name: localStorage.getItem("name"), 
      userPhoto: localStorage.getItem("userPhoto"),
      newphoto:"",
			newname: "",
      successmsg: null,
      errormsg: null,
      loading: false,
    };
  },
methods: { 
  
  input(event) {
    const file = event.target.files[0];
    if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.newphoto = e.target.result; 
        };
        reader.readAsDataURL(file);}},
	async setMyUserName() {
		this.loading = true;
	try{

		let response = await this.$axios.put(`/session/${this.userId}/userName`, {name: this.newname, userPhoto: ""},  { headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` }});
  	localStorage.setItem("name", this.newname);
		this.name = response.data.name;
    localStorage.setItem("name", this.newname);
		this.errormsg = null;
    this.loading = false;
	}	catch (e){ 		
				if (e.response && e.response.status === 400) {
            

					}else{
            this.errormsg = e.toString();						
					}
	}
	},async setMyPhoto() {
		this.loading = true;
	try{

		let response = await this.$axios.put(`/session/${this.userId}/userPhoto`, {name: "", userPhoto: this.newphoto},  { headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` }});
  	localStorage.setItem("userPhoto", this.newphoto);
		this.userPhoto = this.newphoto;
		this.errormsg = null;
    this.loading = false;
    this.successmsg = "Photo set!";
	}	catch (e){ 		
    console.error("Error:", e); 
				if (e.response && e.response.status === 400) {
            

					}else{
            this.errormsg = e.toString();						
					}
  }
	},
},


};
</script>






<template>
  <div class="user-profile">
    <div class="user-details">
      <h1 class="h2">USER PROFILE</h1>
      <p class="h3">UserId: {{ userId }}</p>
						<div class="mb-3">
				<label for="username" class="form-label">Change UserName: </label>
				<input
					type="text"
					id="username"
					class="form-control"
					v-model="newname"
					placeholder="new username"
				/>
				<button @click="setMyUserName">
OK			</button>
			</div>
      <p class="h3">UserName: {{ name }}</p>
    </div>
    <div class="user-photo">
      <img v-if="this.userPhoto" :src="userPhoto" alt="User Photo" style="width: 200px; height: 200px; object-fit: cover;" />
      <div v-else>
        <p>Photo</p>
      </div>
      <label for="userPhoto" class="form-label">Change UserPhoto: </label>
      <input
        type="file"
        id="userPhoto"
        class="form-control"
        @change="input"
				/>
        				<button @click="setMyPhoto">
OK					</button>
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
