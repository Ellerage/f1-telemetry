import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import { router } from "./router";
import PrimeVue from "primevue/config";
// import "primevue/resources/themes/md-light-indigo/theme.css"; // тема
// import "primevue/resources/primevue.min.css"; // базовые стили
// import "primeicons/primeicons.css";

const app = createApp(App);
app.use(router);
app.use(PrimeVue);
app.mount("#app");
