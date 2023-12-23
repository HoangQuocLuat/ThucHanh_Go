import { createApp } from "vue";
import App from "./App.vue";
import { createRouter, createWebHistory } from "vue-router";
import Login from "./pages/login.vue";
import Register from "./pages/register.vue";
import Home from "./pages/home.vue"
/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

/* import specific icons */
import { fas } from '@fortawesome/free-solid-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'

/* add icons to the library */
library.add(fas, fab, far)

const routes = [
  { path: "/", component: Login},
  { path: "/login", component: Login },
  { path: "/register", component: Register },
  { path: "/home", component: Home},
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

const app = createApp(App);
app.use(router);
app.component("font-awesome-icon", FontAwesomeIcon);
app.mount("#app"); 
