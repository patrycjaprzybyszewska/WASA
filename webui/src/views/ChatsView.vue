
<script>
export default {
  data() {
    return {
            chats: [],
            messages: [],
			loading: false,
			error: null,
            chatId: null,
            messageId: null,
			userId: localStorage.getItem("userId"),
            selectedChat: null,
            MessagetoForward: null,
            chattoforwardId: null,
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
      				  const response = await this.$axios.get('/conversation', {
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
        async getConversation(chatId){
            this.loading = true;
            this.error = null;
            try{
                const response = await this.$axios.get(`/conversation/${chatId}`, {
       				   headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` },
       		 });
                this.messages = response.data;
                this.selectedChat = chatId;
            } catch (err) {
        console.error("Error fetching messages:", err);
        this.error = `Unable to fetch messages for chatId ${chatId}. Error: ${
          err.response ? err.response.status : err.message
        }`;}
     },
     async deleteMessage(messageId){
            this.loading = true;
            this.error = null;
            try{
                const response = await this.$axios.delete(`/message/${messageId}`, {
       				   headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` },
       		 });
                this.messages = this.messages.filter((msg) => msg.messageId !== messageId);
            } catch (err) {
        console.error("Error deleteing messages:", err);
        this.error = `Unable to delete message with ID ${messageId}. Error: $${
          err.response ? err.response.status : err.message
        }`;}
        
     },
     setMessagetoForward(messageId) {
      this.MessagetoForward = { messageId };
    },
     async forwardMessage(chattoforwardId){
            this.loading = true;
            this.error = null;
            try{
                const response = await this.$axios.put(`/message/forward/${this.MessagetoForward}/${chattoforwardId}`, {
       				   headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` },
       		 });
                this.successmsg = "Message forwarded!";
            } catch (err) {
        console.error("Error deleteing messages:", err);
        this.error = `Unable to delete message with ID ${messageId}. Error: $${
          err.response ? err.response.status : err.message
        }`;}},
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
    <p>Chat ID: <button @click="getConversation(chat.chatId)">{{ chat.chatId }}</button></p>
  </li>
</ul>
  <div v-if="selectedChat">
      <h2>Messages for Chat: {{ selectedChat }}</h2>
      <ul>
        <li v-for="message in messages" :key="message.messageId">
          <p><strong>{{ message.senderId }}</strong>: {{ message.content }}</p>
          <p>deleteMessage: <button @click="deleteMessage(message.messageId)">{{ message.messageId }}</button></p>
          <p>forwardMessage: <button @click="setMessagetoForward(message.messageId)">{{ message.messageId }}Forward</button></p>
        </li>
      </ul>
    </div><div v-if="MessagetoForward">
      <h2>Select Chat</h2>
      <ul>
        <li v-for="chat in chats" :key="chat.chatId">
          <button @click="forwardMessage(chat.chatId)">
            Forward to {{ chat.chatName }} (Chat ID: {{ chat.chatId }})
          </button>
        </li>
      </ul>
      <button @click="MessagetoForward = null">Cancel</button>
    </div>
    <div v-if="successmsg" class="alert alert-success">{{ successmsg }}</div>
<div v-if="errormsg" class="alert alert-danger">{{ errormsg }}</div>

</div>
</template>