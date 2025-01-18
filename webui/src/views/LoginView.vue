<script>
export default {
	data: function(){
		return{
			errormsg: null,
			loading: false,
			successmsg: null,
			userId: "",
			name: "",
			userPhoto: "",
		}
	},
methods: { 
	async doLogin() {
		this.loading = true;
	try{
		localStorage.removeItem('userPhoto');
		localStorage.removeItem('userId');
		localStorage.removeItem('name');
		let response = await this.$axios.post("/session", {name: this.name}, { headers: { "Content-Type": "application/json" }});
    	localStorage.setItem("userId", response.data.userId);
  		localStorage.setItem("name", response.data.name);
		this.successmsg = "You are logged!";
		this.errormsg = null;
    this.loading = false;
		
	}	catch (e){ 		
				if (e.response && e.response.status === 400) {
            this.errormsg = "Failed to login";

					}else{
            this.errormsg = e.toString();						
					}
	}
	},
},


};

</script>
<template>
	<div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">LOG INTO WASAText!</h1>
		</div>
		<form @submit.prevent="doLogin">
			      <div v-if="errormsg" class="alert alert-danger">
        {{ errormsg }}
      </div>
      <div v-if="successmsg" class="alert alert-success">
        {{ successmsg }}
      </div>
			<div class="mb-3">
				<label for="username" class="form-label">UserName: </label>
				<input
					type="text"
					id="username"
					class="form-control"
					v-model="name"
					placeholder="UserName"
				/>
			</div>
		      <button type="submit" class="btn btn-primary" :disabled="loading">
					{{ loading ? "Logging in..." : "OK" }}
      </button>
		</form>
	</div>
</template>


<style>

</style>


