<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	data(){
		return{
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
			this.loading = true;
      		this.error = null;
     			 try {
      				  const response = await this.$axios.get('/conversations', {
       				   headers: { Authorization: `Bearer ${localStorage.getItem("userId")}` },
       		 });
       		 this.chats = response.data; 
   		   } catch (err) {
     	   console.error("Error fetching conversations:", err);
       		this.error = "Unable to fetch conversations.";
      } finally {
        this.loading = false;
      }
		},
	},
};
const userId = localStorage.getItem("userId");
</script>

<template>

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">Example App</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/login" class="nav-link">
 								 <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
 									Log
							</RouterLink>

						</li>
						<li class="nav-item">
							<RouterLink :to="`/user/${userId}`" class="nav-link">

								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout"/></svg>
								UserProfile
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/message" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
								SendMessage
							</RouterLink>
						</li>
					</ul>

					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>Secondary menu</span>
					</h6>
					<ul class="nav flex-column">
							<li v-for="chat in chats" :key="chat.id" class="nav-item">
								<RouterLink :to="'/chat/' + chat.id" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
								{{ chat.name }}
							</RouterLink>
						</li>
					</ul>
				</div>
				<div v-if="error" class="text-danger">
      			{{ error }}
    </div>
			</nav>

			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView />
			</main>
		</div>
	</div>
</template>

<style>
</style>
