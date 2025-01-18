
<script>
export default {
  data() {
    return {
      userId: localStorage.getItem("userId"),  
      name: localStorage.getItem("name"), 
      userPhoto: localStorage.getItem("userPhoto"),
      content:"",
      chatId: null,
      successmsg: null,
      errormsg: null,
      loading: false,
    };
  },
methods: { 

	async sendMessage() {
		this.loading = true;
	try{

		let response = await this.$axios.put(`/message`, {content: this.content, chatId: this.chatId},  { headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` }});
		this.content = response.data.content;
		this.chatId = response.data.chatId;
		this.errormsg = null;
        this.loading = false;
		this.successmsg = "Message sent!";
	}	catch (e) { console.error("Error sending message:", e);
  if (e.response) {
    console.error("Response data:", e.response.data);
    console.error("Response status:", e.response.status);
  }
  this.errormsg = "Failed to send message..";
}
	},}


};
</script>






<template>
  <div class="user-profile">
    <div class="user-details">
      <h1 class="h2">SEND MESSAGE</h1>
						<div class="mb-3">
				<label for="content" class="form-label">Content: </label>
				<input
					type="text"
					id="content"
					class="form-control"
					v-model="content"
					placeholder="type here"
				/>
				</div>
				<div class="mb-3">
				<label for="user" class="form-label">UserId: </label>
				<input
					type="text"
					id="chatId"
					class="form-control"
					v-model="chatId"
					placeholder="type here id of user or a chat"
				/>
				</div>
    </div>
				<button @click="sendMessage">
OK			</button>
<div v-if="successmsg" class="alert alert-success">{{ successmsg }}</div>
<div v-if="errormsg" class="alert alert-danger">{{ errormsg }}</div>

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
