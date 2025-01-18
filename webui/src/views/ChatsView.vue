
<script>
export default {
  data() {
    return {
            chats: [],
			loading: false,
			error: null,
			userId: localStorage.getItem("userId")
    };
  },
  created() {
		this.getConversations();
	},
	methods: {
		async getConversations(){
            const userId = localStorage.getItem("userId");
    if (!userId) {
      this.error = "User ID not found in localStorage.";
      this.loading = false;
      return;}
			this.loading = true;
      		this.error = null;
     			 try {
      				  const response = await this.$axios.get('/conversations', {
       				   headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` },
       		 });
			console.log("Response data:", response.data)
       		 this.chats = response.data; 
   		   } catch (err) {
     	   console.error("Error fetching conversations:", err);
       		this.error = `Unable to fetch conversations.Error: ${err.response ? err.response.status : err.message}`;
      } finally {
        this.loading = false;
      }
		},
	},
};
</script>
<template>
<div>
    <h1>Chats</h1>
<div v-if="error" class="error">{{ error }}</div>
<ul v-if="!loading && !error">
  <li v-for="chat in chats" :key="chat.chatId">
    <img :src="chat.chatPhoto" alt="Chat photo" v-if="chat.chatPhoto" />
    <p><strong>{{ chat.chatName }}</strong></p>
  </li>
</ul>
</div>
</template>