import Vue from 'vue';
import VueCookies from 'vue-cookies'
import App from './App.vue';
import router from './router';
import ElementUI from 'element-ui';
import VueI18n from 'vue-i18n';
import { messages } from './components/common/i18n';
import 'element-ui/lib/theme-chalk/index.css'; 
import './assets/css/icon.css';
import './components/common/directives';
import 'babel-polyfill';
import axios from 'axios'
import VueAxios from 'vue-axios'

if (process.env.NODE_ENV == 'development'){
    require( './mock');
}

Vue.config.productionTip = false;

Vue.use(VueI18n);
const i18n = new VueI18n({
    locale: 'zh',
    messages
});

Vue.use(ElementUI, {
    size: 'small',
    i18n: (key, value) => i18n.t(key, value)
});

axios.defaults.withCredentials=true
Vue.use(VueAxios,axios)

Vue.use(VueCookies)

//使用钩子函数对路由进行权限跳转
router.beforeEach((to, from, next) => {
    document.title = i18n.t(to.meta.title) || i18n.t('i18n.system_name');
    const role = localStorage.getItem('ms_userrole');
    let sessionId = window.$cookies.get("session_id")

    if (!sessionId && to.path !== '/login') {
        next('/login');
    } else if (to.meta.permission) {
        // 如果是管理员权限则可进入，这里只是简单的模拟管理员权限而已
        role === 'superadmin' ? next() : next('/403');
    } else {
        // 简单的判断IE10及以下不进入富文本编辑器，该组件不兼容
        if (navigator.userAgent.indexOf('MSIE') > -1 && to.path === '/editor') {
            Vue.prototype.$alert('vue-quill-editor组件不兼容IE10及以下浏览器，请使用更高版本的浏览器查看', '浏览器不兼容通知', {
                confirmButtonText: '确定'
            });
        } else {
            next();
        }
    }
});

new Vue({
    router,
    i18n,
    render: h => h(App)
}).$mount('#app');