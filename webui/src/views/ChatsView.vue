
<script>
export default {
  data() {
    return {
            chats: [],
            comments: [":)", ":("],
            messages: [],
			loading: false,
			error: null,
            chatId: null,
            chatName: "",
            chatPhoto: "",
            messageId: null,
			userId: localStorage.getItem("userId"),
            selectedChat: null,
            forwardchatName: "",
            MessagetoForward: null,
            MessagetoComment: null,
            chattoforwardId: null,
            commentId: null,
            successmsg: null, 
            usertoad: null,
            errormsg: null,
            showSettings: false,
            showForward: false,
    };
  },
  created() {
		this.getConversations();
	},
	methods: {
		async getConversations(){
       
			this.loading = true;
      		this.error = null;
     			 try {
                    const response = await this.$axios.get(`/conversation`, {
       				   headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` },
       		 });
      	
			console.log("Response data:", response.data)

       		 this.chats = response.data; 
   		   } catch (err) {
     	   console.error("Error fetching conversations:", err);
          if (e.response.status === 404) {
            this.error = `  U have no conversations!`;
          }
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
                this.messages = response.data.map(item => ({
                 ...item.message,
                 comments: item.comments || [], 
                  }));

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
                this.messages = this.messages.filter((deleted) => deleted.messageId !== messageId);
            } catch (err) {
        console.error("Error deleteing messages:", err);
        this.error = `Unable to delete message with ID ${messageId}. Error: $${
          err.response ? err.response.status : err.message
        }`;}
        
     },
     setMessagetoForward(messageId) {
      this.MessagetoForward = { messageId };
      this.showForward = true;
    },

     async forwardMessage(){
            this.loading = true;
            this.error = null;
            this.successmsg = null;
            this.errormsg = null;
            try{
                const response = await this.$axios.put(`/message/forward/${this.MessagetoForward.messageId}`, {chatName: this.forwardchatName},
                {}, {
       				   headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` },
       		 })
		            this.forwardchatName = response.data.forwardchatName;
                this.successmsg = "Message forwarded!";
                this.showForward = false;
                this.MessagetoForward = null;
            } catch (err) {
        console.error("Error deleteing messages:", err);
        this.errormsg = `Unable to delete message with ID ${messageId}. Error: $${
          err.response ? err.response.status : err.message
        }`;}},
        setMessagetoComment(messageId) {
      this.MessagetoComment = { messageId };
    },
    async setComment(){
        this.loading = true;
    this.error = null;
    this.successmsg = null;
    this.errormsg = null;
    
    try {
      const commentData = {
        content: this.selectedComment, 
        messageId: this.MessagetoComment.messageId,
      };

      const response = await this.$axios.put(`/message/comment/${this.MessagetoComment.messageId}`, commentData, {
        headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` },
      });

      this.successmsg = "Comment!";
      this.setMessagetoComment = null;
      this.selectedComment = null; 
    } catch (err) {
      console.error("Error adding comment:", err);
      this.errormsg = `Unable to add comment. Error: ${err.response ? err.response.status : err.message}`;
    } finally {
      this.loading = false;
    }
  },
  input(event) {
    const file = event.target.files[0];
    if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.newphoto = e.target.result; 
        };
        reader.readAsDataURL(file);}},
	async setMyGroupName() {
		this.loading = true;
	try{

		let response = await this.$axios.put(`/groupchat/${this.chatId}/groupName`, {chatName: this.newname, chatPhoto: ""},  { headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` }});
  	localStorage.setItem("chatName", this.newname);
		this.chatName = response.data.name;
    localStorage.setItem("chatName", this.newname);
		this.errormsg = null;
    this.loading = false;
    this.successmsg = "Name set!";
	}	catch (e){ 		
      console.error("Error setting username:", e);
				if (e.response && e.response.status === 400) {
            

					}else{
            this.errormsg = e.toString();						
					}
	} finally {
        this.loading = false;
      }
    },async setMyPhoto() {
		this.loading = true;
	try{

		let response = await this.$axios.put(`/groupchat/${this.chatId}/groupPhoto`, {chatName: "", chatPhoto: this.newphoto},  { headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` }});
  	localStorage.setItem("userPhoto", this.newphoto);
		this.userPhoto = this.newphoto;
		this.errormsg = null;
    this.loading = false;
    this.successmsg = "Photo set!";
	}	catch (e){ 		
    console.error("Error setting photo:", e);
				if (e.response && e.response.status === 400) {
            

					}else{
            this.errormsg = e.toString();						
					}
	} finally {
        this.loading = false;
      }
	},
    async leaveGroup (chatId){
        this.successmsg = null;
        this.loading = true;
            this.error = null;
            try{
                const response = await this.$axios.delete(`/groupchat/${chatId}/leave/${this.userId}`,  {
       				   headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` },
       		 });
                this.successmsg = "Group left!";
                
             this.chats = this.chats.filter(chat => chat.chatId !== chatId);
            } catch (err) {
        console.error("Error leaving group:", err);
        console.error("Error leaving group:", err);
        this.error = `Unable to leave group. Error: ${err.response ? err.response.status : err.message}`;
     ;}
    },
    async addToGroup (selectedChat){
        this.successmsg = null;
        this.loading = true;
        if (!this.usertoad) {
        this.error = "User ID to add is required.";
         this.loading = false;
    return;
}
            this.error = null;
            try{
                const response = await this.$axios.put(`/groupchat/${selectedChat}/add/${this.usertoad}`, {}, {
       				   headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` },
       		 });
                this.successmsg = "User added!";
            } catch (err) {
   console.error("Error adding user to group:", err);
        this.error = `Unable to add user to group. Error: ${err.response ? err.response.status : err.message}`
     ;}
    },
    async uncommentMessage (commentId){
        this.successmsg = null;
        this.loading = true;
       this.error = null;
            try{
                const response = await this.$axios.delete(`/comment/${commentId}`,  {
       				   headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` },
       		 });
                this.messages.forEach((message) => {
                  message.comments = message.comments.filter((deleted) => deleted.commentId !== commentId);
                  });
                this.successmsg = "Comment removed!";
            } catch (err) {
   console.error("Error deleting:", err);
        this.error = ` Error: ${err.response ? err.response.status : err.message}`
     ;}
    },
    toggleSettings(chatId) {

      this.showSettings = !this.showSettings;
      if (this.showSettings) {
        this.chatId = chatId; 
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
    <p>Chat ID: <button @click="getConversation(chat.chatId)">{{ chat.chatId }}</button></p>
    <button @click="leaveGroup(chat.chatId)">Leave Group</button>
    

  </li>
</ul>
  <div v-if="selectedChat">
      <h2>Messages for Chat: {{ selectedChat }}</h2>
      <button @click="toggleSettings(selectedChat)">Settings</button>

      <label for="usertoad" class="form-label">User id: </label>
            <input
              type="text"
              id="usertoad"
              class="form-control"
              v-model="usertoad"
              placeholder="user to add"
            />
            <button @click="addToGroup(selectedChat)">Add user</button>
      <ul>
        <li v-for="message in messages" :key="message.messageId">
          <p><strong>Sender Name: {{ message.SenderName }}</strong><br>Content: {{ message.content }}<br>Date: {{ message.messageDate }}<br>Time: {{ message.messageTime }}<br> Status: 
          <span v-if="message.state === 'delivered'">☑</span>
          <span v-else>{☑☑}</span>
        </p>
          <ul>
             <li v-for="(comment, index) in message.comments" :key="index">
                 <button @click="uncommentMessage(comment.commentId)">{{ comment.content }}</button>
            </li>
          </ul>
          <p>deleteMessage: <button @click="deleteMessage(message.messageId)">{{ message.messageId }}</button></p>
          <p>forwardMessage: <button @click="setMessagetoForward(message.messageId)">Forward</button></p>
         
          <p>commmentMessage:<button @click="setMessagetoComment(message.messageId)">Comment</button></p>
        </li>
      </ul>

</div>
       <transition name="fade">
        <div v-if="showSettings" class="settings-window">
      <h2> Chat: {{  }}</h2>
      <button @click="showSettings = false">Close</button>
      <div class="user-profile">
        <div class="user-details">
          <h1 class="h2">CHAT</h1>
          <div class="mb-3">
            <label for="chatname" class="form-label">Change ChatName: </label>
            <input
              type="text"
              id="chatname"
              class="form-control"
              v-model="newname"
              placeholder="new chatname"
            />
            <button @click="setMyGroupName">OK</button>
          </div>
          <p class="h3">ChatName: {{ chatName }}</p>
        </div>
        <div v-if="successmsg" class="alert alert-success">{{ successmsg }}</div>
        <div v-if="errormsg" class="alert alert-danger">{{ errormsg }}</div>

        <div class="chat-photo">
          <img v-if="chatPhoto" :src="chatPhoto" alt="Chat Photo" style="width: 200px; height: 200px; object-fit: cover;" />
          <div v-else>
            <p>Photo</p>
          </div>
          <label for="chatPhoto" class="form-label">Change UserPhoto: </label>
          <input
            type="file"
            id="chatPhoto"
            class="form-control"
            @change="input"
          />
          <button @click="setGroupPhoto">OK</button>
        </div>
      </div>
    </div>
  </transition>

  <transition name="fade">
        <div v-if="MessagetoForward" class="settings-window">

      <div class="mb-3">
				<label for="forwardchatName" class="form-label">Name: </label>
				<input
					type="text"
					id="forwardchatName"
					class="form-control"
					v-model="forwardchatName"
					placeholder="type here name of user,chat or new chat name to forward message"
				/>
        <button @click="forwardMessage()">OK</button>
        <button @click="showForward = false">Close</button>
				</div>

      <div v-if="errormsg" class="alert alert-danger">{{ errormsg }}</div>

      <div v-if="successmsg" class="alert alert-success">{{ successmsg }}</div>


    </div>
  </transition>




</div><div v-if="MessagetoComment">
    <h2>Select Comment</h2>
        <div v-for="(comment, index) in comments" :key="index">
    <button @click="selectedComment = comment">{{ comment }}</button>
  </div>
  <button @click="setComment">Comment</button>
     
</div>
</template>


<style scoped>
.settings-window {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  width: 80%;
  max-width: 500px;
}

.settings-window h2 {
  margin-bottom: 20px;
}

.settings-window button {
  margin-top: 10px;
}

</style>