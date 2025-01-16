import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import UserView from '../views/UserView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/login', component: LoginView},
		{path: '/user/:userId', component: HomeView},
		{path: '/link2', component: HomeView},
		{path: '/some/:id/link', component: HomeView},
	]
})

export default router
